// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package provision

import (
	"context"
	"fmt"
	"strings"

	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/frontend"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace"
)

// Represents a server bound to an environment.
type Server struct {
	Location workspace.Location
	Package  *workspace.Package

	Provisioning frontend.PreparedProvisionPlan // A provisioning plan that is attached to the server itself.

	env   ServerEnv            // The environment this server instance is bound to.
	entry *schema.Stack_Entry  // The stack entry, i.e. all of the server's dependencies.
	deps  []*workspace.Package // List of parsed deps.
}

type ServerEnv interface {
	workspace.WorkspaceEnvironment
	workspace.SealedPackages
}

func (t Server) Module() *workspace.Module       { return t.Location.Module }
func (t Server) Env() ServerEnv                  { return t.env }
func (t Server) PackageName() schema.PackageName { return t.Location.PackageName }
func (t Server) StackEntry() *schema.Stack_Entry { return t.entry }
func (t Server) Proto() *schema.Server           { return t.entry.Server }
func (t Server) Name() string                    { return t.entry.Server.Name }
func (t Server) Framework() schema.Framework     { return t.entry.Server.Framework }
func (t Server) IsStateful() bool                { return t.entry.Server.IsStateful }
func (t Server) Deps() []*workspace.Package      { return t.deps }

func (t Server) GetDep(pkg schema.PackageName) *workspace.Package {
	for _, d := range t.deps {
		if d.PackageName() == pkg {
			return d
		}
	}
	return nil
}

func makeServer(ctx context.Context, loader workspace.Packages, env *schema.Environment, pkgname schema.PackageName, bind func() ServerEnv) (Server, error) {
	sealed, err := workspace.Seal(ctx, loader, pkgname, &workspace.SealHelper{
		AdditionalServerDeps: func(fmwk schema.Framework) ([]schema.PackageName, error) {
			var pkgs schema.PackageList
			pkgs.Add(schema.PackageName(fmt.Sprintf("namespacelabs.dev/foundation/std/runtime/%s", strings.ToLower(env.Runtime))))
			if handler, ok := workspace.FrameworkHandlers[fmwk]; ok {
				pkgs.AddMultiple(handler.DevelopmentPackages()...)
			}
			return pkgs.PackageNames(), nil
		},
	})
	if err != nil {
		return Server{}, err
	}

	if sealed.ParsedPackage == nil || sealed.ParsedPackage.Server == nil {
		return Server{}, fnerrors.UserError(pkgname, "not a server")
	}

	t := Server{
		Location: sealed.ParsedPackage.Location,
		env:      bind(),
	}

	for _, req := range sealed.ParsedPackage.Server.GetEnvironmentRequirement() {
		for _, r := range req.GetEnvironmentHasLabel() {
			if !env.HasLabel(r) {
				return Server{}, IncompatibleEnvironmentErr{
					Env:              t.env.Proto(),
					Server:           sealed.ParsedPackage.Server,
					RequirementOwner: schema.PackageName(req.Package),
					RequiredLabel:    r,
				}
			}
		}

		for _, r := range req.GetEnvironmentDoesNotHaveLabel() {
			if env.HasLabel(r) {
				return Server{}, IncompatibleEnvironmentErr{
					Env:               t.env.Proto(),
					Server:            sealed.ParsedPackage.Server,
					RequirementOwner:  schema.PackageName(req.Package),
					IncompatibleLabel: r,
				}
			}
		}
	}

	t.Package = sealed.ParsedPackage
	t.entry = sealed.Proto
	t.deps = sealed.Deps

	pdata, err := t.Package.Parsed.EvalProvision(ctx, t.Env(), frontend.ProvisionInputs{
		Workspace:      t.Module().Workspace,
		ServerLocation: t.Location,
	})
	if err != nil {
		return Server{}, fnerrors.Wrap(t.Location, err)
	}

	if len(pdata.DeclaredStack) > 0 {
		return Server{}, fnerrors.UserError(t.Location, "servers can't add servers to the stack")
	}

	t.Provisioning = pdata.PreparedProvisionPlan
	t.entry.ServerNaming = pdata.Naming

	return t, nil
}
