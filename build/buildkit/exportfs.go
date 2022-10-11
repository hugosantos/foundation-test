// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package buildkit

import (
	"context"
	"io/fs"
	"os"

	"github.com/moby/buildkit/client"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/workspace/dirs"
	"namespacelabs.dev/foundation/workspace/tasks"
)

func exportToFS() exporter[fs.FS] { return &exportFS{} }

type exportFS struct {
	outputDir string
}

func (e *exportFS) Prepare(ctx context.Context) error {
	dir, err := dirs.CreateUserTempDir("buildkit", "fs")
	if err != nil {
		return err
	}

	e.outputDir = dir

	compute.On(ctx).Cleanup(tasks.Action("buildkit.build-fs.cleanup"), func(ctx context.Context) error {
		return os.RemoveAll(dir)
	})

	return nil
}

func (e *exportFS) Exports() []client.ExportEntry {
	return []client.ExportEntry{{
		Type:      client.ExporterLocal,
		OutputDir: e.outputDir,
	}}
}

func (e *exportFS) Provide(context.Context, *client.SolveResponse) (fs.FS, error) {
	return fnfs.Local(e.outputDir), nil
}