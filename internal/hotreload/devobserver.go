// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package hotreload

import (
	"context"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"namespacelabs.dev/foundation/framework/rpcerrors/multierr"
	"namespacelabs.dev/foundation/internal/cli/fncobra/name"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs/workspace/wsremote"
	hrconstants "namespacelabs.dev/foundation/internal/hotreload/constants"
	"namespacelabs.dev/foundation/internal/integrations"
	"namespacelabs.dev/foundation/internal/planning"
	"namespacelabs.dev/foundation/internal/runtime"
	"namespacelabs.dev/foundation/internal/wscontents"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/tasks"
)

type FileSyncDevObserver struct {
	log          io.Writer
	server       runtime.Deployable
	cluster      runtime.ClusterNamespace
	fileSyncPort int32

	mu   sync.Mutex
	conn *grpc.ClientConn
}

func ConfigureFileSyncDevObserver(ctx context.Context, cluster runtime.ClusterNamespace, srv planning.Server) (context.Context, integrations.DevObserver, error) {
	if wsremote.Ctx(ctx) != nil {
		return nil, nil, fnerrors.NewWithLocation(srv.Location, "`%s dev` on multiple web/nodejs servers not supported", name.CmdName)
	}

	devObserver := newFileSyncDevObserver(ctx, cluster, srv, hrconstants.FileSyncPort)

	newCtx, _ := wsremote.BufferAndSinkTo(ctx, devObserver.Deposit)

	return newCtx, devObserver, nil
}

func newFileSyncDevObserver(ctx context.Context, cluster runtime.ClusterNamespace, srv planning.Server, fileSyncPort int32) *FileSyncDevObserver {
	return &FileSyncDevObserver{
		log:          console.TypedOutput(ctx, "hot reload", console.CatOutputUs),
		server:       srv.Proto(),
		cluster:      cluster,
		fileSyncPort: fileSyncPort,
	}
}

func (do *FileSyncDevObserver) Close() error {
	do.mu.Lock()
	defer do.mu.Unlock()
	return do.cleanup()
}

func (do *FileSyncDevObserver) cleanup() error {
	var errs []error

	if do.conn != nil {
		if err := do.conn.Close(); err != nil {
			errs = append(errs, err)
		}
		do.conn = nil
	}

	return multierr.New(errs...)
}

func (do *FileSyncDevObserver) OnDeployment(ctx context.Context) {
	do.mu.Lock()
	defer do.mu.Unlock()

	err := do.cleanup()
	if err != nil {
		fmt.Fprintln(do.log, "failed to port forwarding cleanup", err)
	}

	orch := compute.On(ctx)
	sink := tasks.SinkFrom(ctx)

	conn, err := grpc.NewClient("passthrough:///filesync-"+do.server.GetName(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			patchedContext := compute.AttachOrch(tasks.WithSink(ctx, sink), orch)

			return do.cluster.DialServer(patchedContext, do.server, &schema.Endpoint_Port{
				ContainerPort: do.fileSyncPort,
			})
		}),
	)
	if err != nil {
		fmt.Fprintln(do.log, "failed to connect to filesync", err)
		return
	}

	do.conn = conn
}

func (do *FileSyncDevObserver) Deposit(ctx context.Context, s *wsremote.Signature, events []*wscontents.FileEvent) (bool, error) {
	do.mu.Lock()
	defer do.mu.Unlock()

	if do.conn == nil {
		return false, nil
	}

	var labels []string
	for _, ev := range events {
		labels = append(labels, fmt.Sprintf("%s %s", ev.Event, ev.Path))
	}

	fmt.Fprintf(do.log, "FileSync event: %s\n", strings.Join(labels, ", "))

	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if _, err := wsremote.NewFileSyncServiceClient(do.conn).Push(newCtx, &wsremote.PushRequest{
		Signature: s,
		FileEvent: events,
	}); err != nil {
		return false, err
	}

	return true, nil
}
