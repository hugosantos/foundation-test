// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package kubeops

import (
	"context"
	"encoding/json"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/runtime/kubernetes/client"
	"namespacelabs.dev/foundation/runtime/kubernetes/kubedef"
	"namespacelabs.dev/foundation/runtime/kubernetes/kubeparser"
	fnschema "namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace/tasks"
)

func registerCreate() {
	ops.RegisterFuncs(ops.Funcs[*kubedef.OpCreate]{
		Handle: func(ctx context.Context, d *fnschema.SerializedInvocation, create *kubedef.OpCreate) (*ops.HandleResult, error) {
			if create.BodyJson == "" {
				return nil, fnerrors.InternalError("%s: apply.Body is required", d.Description)
			}

			var obj kubeparser.ObjHeader
			if err := json.Unmarshal([]byte(create.BodyJson), &obj); err != nil {
				return nil, fnerrors.BadInputError("%s: kubernetes.create: failed to parse resource: %w", d.Description, err)
			}

			if obj.Kind == "" {
				// XXX support older prebuilts
				if create.Resource == "secrets" {
					obj.SetGroupVersionKind(schema.GroupVersionKind{Version: "v1", Kind: "Secret"})
				} else {
					return nil, fnerrors.InternalError("%s: object Kind is required", d.Description)
				}
			}

			if obj.Name == "" {
				return nil, fnerrors.InternalError("%s: create.Name is required", d.Description)
			}

			cluster, err := kubedef.InjectedKubeCluster(ctx)
			if err != nil {
				return nil, err
			}

			resource, err := resolveResource(ctx, cluster, obj.GetObjectKind().GroupVersionKind())
			if err != nil {
				return nil, err
			}

			if create.SkipIfAlreadyExists || create.UpdateIfExisting {
				obj, err := fetchResource(ctx, cluster, d.Description, *resource, obj.Name,
					obj.Namespace, fnschema.PackageNames(d.Scope...))
				if err != nil {
					return nil, fnerrors.New("failed to fetch resource: %w", err)
				}

				if obj != nil {
					if create.SkipIfAlreadyExists {
						return nil, nil // Nothing to do.
					}

					if create.UpdateIfExisting {
						msg := &unstructured.Unstructured{Object: map[string]interface{}{}}
						if err := msg.UnmarshalJSON([]byte(create.BodyJson)); err != nil {
							return nil, fnerrors.New("failed to parse create body: %w", err)
						}

						// This is not advised. Overwriting without reading.
						msg.SetResourceVersion(obj.GetResourceVersion())

						return nil, updateResource(ctx, d, *resource, msg, cluster.RESTConfig())
					}
				}
			}

			if err := tasks.Action("kubernetes.create").Scope(fnschema.PackageNames(d.Scope...)...).
				HumanReadablef(d.Description).
				Arg("resource", resource.Resource).
				Arg("name", obj.Name).
				Arg("namespace", obj.Namespace).Run(ctx, func(ctx context.Context) error {
				client, err := client.MakeGroupVersionBasedClient(ctx, cluster.RESTConfig(), resource.GroupVersion())
				if err != nil {
					return err
				}

				req := client.Post()
				opts := metav1.CreateOptions{
					FieldManager: kubedef.K8sFieldManager,
				}

				if obj.Namespace != "" {
					req.Namespace(obj.Namespace)
				}

				r := req.Resource(resource.Resource).
					VersionedParams(&opts, metav1.ParameterCodec).
					Body([]byte(create.BodyJson))

				if OutputKubeApiURLs {
					fmt.Fprintf(console.Debug(ctx), "kubernetes: api post call %q\n", r.URL())
				}

				if err := r.Do(ctx).Error(); err != nil {
					return err
				}

				return nil
			}); err != nil {
				if !errors.IsNotFound(err) {
					return nil, fnerrors.InvocationError("%s: failed to create: %w", d.Description, err)
				}
			}

			return nil, nil
		},

		PlanOrder: func(create *kubedef.OpCreate) (*fnschema.ScheduleOrder, error) {
			if create.BodyJson == "" {
				return nil, fnerrors.InternalError("create.Body is required")
			}

			// XXX handle old versions of the secrets prebuilt which don't return a Kind within the body.
			if create.Resource == "secrets" {
				return kubedef.PlanOrder(schema.GroupVersionKind{Version: "v1", Kind: "Secret"}), nil
			}

			var parsed unstructured.Unstructured
			if err := json.Unmarshal([]byte(create.BodyJson), &parsed); err != nil {
				return nil, fnerrors.BadInputError("kubernetes.apply: failed to parse resource: %w", err)
			}

			return kubedef.PlanOrder(parsed.GroupVersionKind()), nil
		},
	})
}

func updateResource(ctx context.Context, d *fnschema.SerializedInvocation, resource schema.GroupVersionResource, body *unstructured.Unstructured, restcfg *rest.Config) error {
	return tasks.Action("kubernetes.update").Scope(fnschema.PackageNames(d.Scope...)...).
		HumanReadablef(d.Description).
		Arg("resource", resource.Resource).
		Arg("name", body.GetName()).
		Arg("namespace", body.GetNamespace()).Run(ctx, func(ctx context.Context) error {
		client, err := client.MakeGroupVersionBasedClient(ctx, restcfg, resource.GroupVersion())
		if err != nil {
			return err
		}

		req := client.Put()
		opts := metav1.UpdateOptions{
			FieldManager: kubedef.K8sFieldManager,
		}

		if body.GetNamespace() != "" {
			req.Namespace(body.GetNamespace())
		}

		r := req.Resource(resource.Resource).
			Name(body.GetName()).
			VersionedParams(&opts, metav1.ParameterCodec).
			Body(body)

		if OutputKubeApiURLs {
			fmt.Fprintf(console.Debug(ctx), "kubernetes: api put call %q\n", r.URL())
		}

		return r.Do(ctx).Error()
	})
}
