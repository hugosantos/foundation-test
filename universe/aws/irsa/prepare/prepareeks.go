// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/providers/aws/eks"
	"namespacelabs.dev/foundation/provision/configure"
	"namespacelabs.dev/foundation/provision/tool/protocol"
	"namespacelabs.dev/foundation/runtime/kubernetes/kubedef"
	"namespacelabs.dev/foundation/runtime/kubernetes/kubetool"
)

func main() {
	if err := configure.RunServer(context.Background(), func(sr grpc.ServiceRegistrar) {
		h := configure.NewHandlers()
		h.Any().HandleStack(provisionHook{})

		protocol.RegisterInvocationServiceServer(sr, h.ServiceHandler())
	}); err != nil {
		log.Fatal(err)
	}
}

type provisionHook struct{}

func (provisionHook) Apply(ctx context.Context, r configure.StackRequest, out *configure.ApplyOutput) error {
	if r.Env.Runtime != "kubernetes" {
		return fnerrors.BadInputError("universe/aws/irsa only supports kubernetes")
	}

	// Don't attempt to allocate IAM resources when the environment is ephemeral (e.g. a test).
	if r.Env.Ephemeral {
		return nil
	}

	serviceAccount := &kubedef.ServiceAccountDetails{}
	if err := r.UnpackInput(serviceAccount); err != nil {
		return err
	}

	eksCluster := &eks.EKSCluster{}
	if ok, err := r.CheckUnpackInput(eksCluster); err != nil {
		return err
	} else if !ok {
		return nil
	}

	eksServerDetails := &eks.EKSServerDetails{}
	if err := r.UnpackInput(eksServerDetails); err != nil {
		return err
	}

	namespace := kubetool.FromRequest(r).Namespace

	result, err := eks.PrepareIrsa(eksCluster, eksServerDetails.ComputedIamRoleName, namespace, serviceAccount.ServiceAccountName, r.Focus.Server)
	if err != nil {
		return err
	}

	out.Invocations = append(out.Invocations, result.Invocations...)
	out.Extensions = append(out.Extensions, result.Extensions...)

	return nil
}

func (provisionHook) Delete(ctx context.Context, r configure.StackRequest, out *configure.DeleteOutput) error {
	if r.Env.Runtime != "kubernetes" {
		return fnerrors.BadInputError("universe/aws/irsa only supports kubernetes")
	}

	return nil
}
