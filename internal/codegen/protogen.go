// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package codegen

import (
	"context"
	"fmt"
	"io/fs"

	"golang.org/x/exp/slices"
	"namespacelabs.dev/foundation/framework/rpcerrors/multierr"
	"namespacelabs.dev/foundation/internal/bytestream"
	srcprotos "namespacelabs.dev/foundation/internal/codegen/protos"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/parsing"
	"namespacelabs.dev/foundation/internal/protos"
	"namespacelabs.dev/foundation/internal/sdk/buf"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/cfg"
	"namespacelabs.dev/foundation/std/execution"
	"namespacelabs.dev/foundation/std/pkggraph"
)

func RegisterGraphHandlers() {
	execution.RegisterHandlerFunc(func(ctx context.Context, _ *schema.SerializedInvocation, op *OpMultiProtoGen) (*execution.HandleResult, error) {
		request := map[schema.Framework]*srcprotos.FileDescriptorSetAndDeps{}
		var errs []error
		for _, entry := range op.Protos {
			var err error
			request[entry.Framework], err = srcprotos.Merge(entry.Protos...)
			errs = append(errs, err)
		}

		if mergeErr := multierr.New(errs...); mergeErr != nil {
			return nil, mergeErr
		}

		module, err := execution.Get(ctx, pkggraph.MutableModuleInjection)
		if err != nil {
			return nil, err
		}

		config, err := execution.Get(ctx, execution.ConfigurationInjection)
		if err != nil {
			return nil, err
		}

		if err := generateProtoSrcs(ctx, config, request, module.ReadWriteFS()); err != nil {
			return nil, err
		}

		return nil, nil
	})

	execution.Compile[*OpProtoGen](func(ctx context.Context, inputs []*schema.SerializedInvocation) ([]*schema.SerializedInvocation, error) {
		module, err := execution.Get(ctx, pkggraph.MutableModuleInjection)
		if err != nil {
			return nil, err
		}

		loader, err := execution.Get(ctx, pkggraph.PackageLoaderInjection)
		if err != nil {
			return nil, err
		}

		requests := map[schema.Framework][]*srcprotos.FileDescriptorSetAndDeps{}
		for _, input := range inputs {
			msg := &OpProtoGen{}
			if err := input.Impl.UnmarshalTo(msg); err != nil {
				return nil, err
			}

			loc, err := loader.Resolve(ctx, schema.PackageName(msg.PackageName))
			if err != nil {
				return nil, err
			}

			if loc.Module.ModuleName() != module.ModuleName() {
				return nil, fnerrors.BadInputError("%s: can't perform codegen for packages in %q", module.ModuleName(), loc.Module.ModuleName())
			}

			requests[msg.Framework] = append(requests[msg.Framework], msg.Protos)
		}

		multi := &OpMultiProtoGen{}
		for fmwk, protos := range requests {
			multi.Protos = append(multi.Protos, &OpMultiProtoGen_ProtosByFramework{
				Framework: fmwk,
				Protos:    protos,
			})
		}
		slices.SortFunc(multi.Protos, func(a, b *OpMultiProtoGen_ProtosByFramework) int {
			return int(a.Framework) - int(b.Framework)
		})

		return []*schema.SerializedInvocation{
			{
				Impl: protos.WrapAnyOrDie(multi),
			},
		}, nil
	})
}

func generateProtoSrcs(ctx context.Context, env cfg.Configuration, request map[schema.Framework]*srcprotos.FileDescriptorSetAndDeps, out fnfs.ReadWriteFS) error {
	protogen, err := buf.MakeProtoSrcs(ctx, env, request)
	if err != nil {
		return err
	}

	merged, err := compute.GetValue(ctx, protogen)
	if err != nil {
		return err
	}

	if console.DebugToConsole {
		d := console.Debug(ctx)

		fmt.Fprintln(d, "Codegen results:")
		_ = fnfs.VisitFiles(ctx, merged, func(path string, _ bytestream.ByteStream, _ fs.DirEntry) error {
			fmt.Fprintf(d, "  %s\n", path)
			return nil
		})
	}

	if err := fnfs.WriteFSToWorkspace(ctx, console.Stdout(ctx), out, merged); err != nil {
		return err
	}

	return nil
}

func GenProtosAtPaths(ctx context.Context, env cfg.Context, fmwk schema.Framework, fsys fs.FS, paths []string, out fnfs.ReadWriteFS) error {
	opts, err := parsing.MakeProtoParseOpts(ctx, parsing.NewPackageLoader(env), env.Workspace().Proto())
	if err != nil {
		return err
	}

	parsed, err := opts.Parse(fsys, paths)
	if err != nil {
		return err
	}

	return generateProtoSrcs(ctx, env.Configuration(), map[schema.Framework]*srcprotos.FileDescriptorSetAndDeps{
		fmwk: parsed,
	}, out)
}
