// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package grpclistener

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"namespacelabs.dev/foundation/internal/testdata/service/proto"
	"namespacelabs.dev/foundation/std/go/grpc/servercore"
	"namespacelabs.dev/foundation/std/go/server"
)

func init() {
	servercore.SetGrpcListenerConfiguration("mtls", conf{})
}

type Service struct {
}

func WireService(ctx context.Context, srv server.Registrar, deps ServiceDeps) {
	proto.RegisterEmptyServiceServer(srv, &Service{})
}

type conf struct {
	servercore.DefaultConfiguration
}

func (conf) TransportCredentials(string) credentials.TransportCredentials {
	return nil
}

func (conf) ServerOpts(string) []grpc.ServerOption { return nil }
