// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package orchestration

import (
	"context"
	"fmt"

	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/provision/deploy"
	"namespacelabs.dev/foundation/runtime"
	"namespacelabs.dev/foundation/schema"
	orchpb "namespacelabs.dev/foundation/schema/orchestration"
	"namespacelabs.dev/foundation/std/execution"
	"namespacelabs.dev/foundation/std/planning"
	"namespacelabs.dev/foundation/workspace/tasks"
)

func Deploy(ctx context.Context, env planning.Context, cluster runtime.ClusterNamespace, plan *schema.DeployPlan, wait, outputProgress bool) error {
	if !UseOrchestrator {
		if !wait {
			return fnerrors.BadInputError("waiting is mandatory without the orchestrator")
		}

		p := execution.NewPlan(plan.Program.Invocation...)

		// Make sure that the cluster is accessible to a serialized invocation implementation.
		return execution.Execute(ctx, env, "deployment.execute", p,
			deploy.MaybeRenderBlock(env, cluster, outputProgress),
			runtime.InjectCluster(cluster)...)
	}

	return tasks.Action("orchestrator.deploy").Scope(schema.PackageNames(plan.FocusServer...)...).
		Run(ctx, func(ctx context.Context) error {
			debug := console.Debug(ctx)
			fmt.Fprintf(debug, "deploying program:\n")
			for k, inv := range plan.GetProgram().GetInvocation() {
				fmt.Fprintf(debug, " #%d %q --> cats:%v after:%v\n", k, inv.Description,
					inv.GetOrder().GetSchedCategory(),
					inv.GetOrder().GetSchedAfterCategory())
			}

			raw, err := cluster.Cluster().EnsureState(ctx, orchestratorStateKey)
			if err != nil {
				return err
			}

			conn, err := raw.(*RemoteOrchestrator).Connect(ctx)
			if err != nil {
				return err
			}

			defer conn.Close()

			id, err := CallDeploy(ctx, env, conn, plan)
			if err != nil {
				return err
			}

			if wait {
				var ch chan *orchpb.Event
				var cleanup func(ctx context.Context) error

				if outputProgress {
					ch, cleanup = deploy.MaybeRenderBlock(env, cluster, true)(ctx)
				}

				err := WireDeploymentStatus(ctx, conn, id, ch)
				if cleanup != nil {
					cleanupErr := cleanup(ctx)
					if err == nil {
						return cleanupErr
					}
				}

				return err
			}

			return nil
		})
}
