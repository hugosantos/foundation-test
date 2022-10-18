// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/cmd"
	"namespacelabs.dev/foundation/internal/cli/cmd/cluster"
	"namespacelabs.dev/foundation/internal/cli/cmd/debug"
	"namespacelabs.dev/foundation/internal/cli/cmd/eks"
	"namespacelabs.dev/foundation/internal/cli/cmd/source"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
)

func main() {
	fncobra.DoMain("nsdev", func(root *cobra.Command) {
		cmd.RegisterCommands(root)
		root.AddCommand(debug.NewDebugCmd())
		root.AddCommand(debug.NewFnServicesCmd())
		root.AddCommand(eks.NewEksCmd())
		root.AddCommand(cmd.NewImagesCmd())
		root.AddCommand(cluster.NewClusterCmd())
		root.AddCommand(cmd.NewLintCmd())
		root.AddCommand(source.NewSourceCmd())
		root.AddCommand(cmd.NewUseCmd())
	})
}
