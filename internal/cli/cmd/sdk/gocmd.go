// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package sdk

import (
	"context"

	"github.com/spf13/cobra"
	"k8s.io/utils/ptr"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/integrations/golang/rungo"
	"namespacelabs.dev/foundation/internal/parsing"
	golangsdk "namespacelabs.dev/foundation/internal/sdk/golang"
	"namespacelabs.dev/foundation/internal/sdk/host"
	"namespacelabs.dev/foundation/schema"
)

func newGoCmd(goVersion string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                "go -- ...",
		Short:              "Run Go.",
		Hidden:             true,
		DisableFlagParsing: true,
	}

	env := fncobra.EnvFromValue(cmd, ptr.To("dev"))
	cmd.RunE = fncobra.RunE(func(ctx context.Context, args []string) error {
		pl := parsing.NewPackageLoader(*env)
		loc, err := pl.Resolve(ctx, schema.MakePackageName((*env).Workspace().ModuleName()))
		if err != nil {
			return err
		}

		sdk, err := golangsdk.MatchSDK(goVersion, host.HostPlatform())
		if err != nil {
			return fnerrors.AttachLocation(loc, err)
		}

		localSDK, err := compute.GetValue(ctx, sdk)
		if err != nil {
			return fnerrors.AttachLocation(loc, err)
		}

		return rungo.RunGo(ctx, loc, localSDK, args...)
	})
	return cmd
}
