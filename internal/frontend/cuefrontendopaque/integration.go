// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package cuefrontendopaque

import (
	"context"

	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/frontend/fncue"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/pkggraph"
)

type cueIntegration struct {
	cueIntegrationDocker

	Kind string `json:"kind"`

	// Shortcuts
	Docker *cueIntegrationDocker `json:"docker"`
}

type cueIntegrationDocker struct {
	Dockerfile string `json:"dockerfile"`
	Image      string `json:"image"`
}

// Mutates "pkg""
func parseIntegration(ctx context.Context, loc pkggraph.Location, v *fncue.CueV, pkg *pkggraph.Package) error {
	var bits cueIntegration
	if err := v.Val.Decode(&bits); err != nil {
		return err
	}

	// Parsing shortcuts
	if bits.Kind == "" {
		if bits.Docker != nil {
			bits.cueIntegrationDocker = *bits.Docker
			bits.Kind = serverKindDockerfile
		}
	}

	switch bits.Kind {
	case serverKindDockerfile:
		if bits.Dockerfile != "" && bits.Image != "" {
			return fnerrors.UserError(loc, "only one of dockerfile and image can be specified")
		}

		buildPlan := &schema.ImageBuildPlan{}
		if bits.Dockerfile != "" {
			buildPlan.Dockerfile = bits.cueIntegrationDocker.Dockerfile
		}
		if bits.Image != "" {
			buildPlan.ImageId = bits.cueIntegrationDocker.Image
		}

		pkg.Binaries = append(pkg.Binaries, &schema.Binary{
			Name: pkg.Server.Name,
			BuildPlan: &schema.LayeredImageBuildPlan{
				LayerBuildPlan: []*schema.ImageBuildPlan{buildPlan},
			},
		})
		pkg.Server.Binary = &schema.Server_Binary{
			PackageRef: schema.MakePackageRef(loc.PackageName, pkg.Server.Name),
		}

		return nil
	default:
		return fnerrors.UserError(loc, "unsupported integration kind %q", bits.Kind)
	}
}