// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cluster

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/cli/cli/config/types"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/moby/buildkit/client"
	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/build/buildkit/buildkitd"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/files"
	"namespacelabs.dev/foundation/internal/fnapi"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/localexec"
	"namespacelabs.dev/foundation/internal/providers/nscloud/api"
	"namespacelabs.dev/foundation/internal/sdk/buildctl"
	"namespacelabs.dev/foundation/internal/sdk/host"
)

func NewBuildctlCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buildctl -- ...",
		Short: "Run buildctl on the target build cluster.",
	}

	cmd.RunE = fncobra.RunE(func(ctx context.Context, args []string) error {
		buildctlBin, err := buildctl.EnsureSDK(ctx, host.HostPlatform())
		if err != nil {
			return fnerrors.New("failed to download buildctl: %w", err)
		}

		p, err := runBuildProxy(ctx)
		if err != nil {
			return err
		}

		defer p.Cleanup()

		return runBuildctl(ctx, buildctlBin, p, args...)
	})

	return cmd
}

func runBuildctl(ctx context.Context, buildctlBin buildctl.Buildctl, p *buildProxy, args ...string) error {
	cmdLine := []string{"--addr", "unix://" + p.BuildkitAddr}
	cmdLine = append(cmdLine, args...)

	fmt.Fprintf(console.Debug(ctx), "buildctl %s\n", strings.Join(cmdLine, " "))

	buildctl := exec.CommandContext(ctx, string(buildctlBin), cmdLine...)
	buildctl.Env = append(buildctl.Env, fmt.Sprintf("DOCKER_CONFIG="+p.DockerConfigDir))

	return localexec.RunInteractive(ctx, buildctl)
}

type buildProxy struct {
	BuildkitAddr     string
	DockerConfigDir  string
	RegistryEndpoint string
	Cleanup          func()
}

func runBuildProxy(ctx context.Context) (*buildProxy, error) {
	existing := config.LoadDefaultConfigFile(console.Stderr(ctx))

	response, err := api.EnsureBuildCluster(ctx, api.Endpoint)
	if err != nil {
		return nil, err
	}

	if response.BuildCluster == nil || response.BuildCluster.Colocated == nil {
		return nil, fnerrors.New("cluster is not a build cluster")
	}

	if err := waitUntilReady(ctx, response); err != nil {
		return nil, err
	}

	p, err := runUnixSocketProxy(ctx, "buildkit", response.ClusterId, func(ctx context.Context) (net.Conn, error) {
		return connect(ctx, response)
	})
	if err != nil {
		return nil, err
	}

	t, err := api.RegistryCreds(ctx)
	if err != nil {
		p.Cleanup()
		return nil, err
	}

	// We don't copy over all authentication settings; only some.
	// XXX replace with custom buildctl invocation that merges auth in-memory.
	newConfig := configfile.ConfigFile{
		AuthConfigs:       existing.AuthConfigs,
		CredentialHelpers: existing.CredentialHelpers,
	}

	newConfig.AuthConfigs[response.Registry.EndpointAddress] = types.AuthConfig{
		Username: t.Username,
		Password: t.Password,
	}

	credsFile := filepath.Join(p.TempDir, config.ConfigFileName)
	if err := files.WriteJson(credsFile, newConfig, 0600); err != nil {
		p.Cleanup()
		return nil, err
	}

	return &buildProxy{p.SocketAddr, p.TempDir, response.Registry.EndpointAddress, p.Cleanup}, nil
}

func waitUntilReady(ctx context.Context, response *api.CreateClusterResult) error {
	return buildkitd.WaitReadiness(ctx, func() (*client.Client, error) {
		// We must fetch a token with our parent context, so we get a task sink etc.
		token, err := fnapi.FetchTenantToken(ctx)
		if err != nil {
			return nil, err
		}

		return client.New(ctx, response.ClusterId, client.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return api.DialPortWithToken(ctx, token, response.Cluster, int(response.BuildCluster.Colocated.TargetPort))
		}))
	})
}

func serveBuildProxy(ctx context.Context, listener net.Listener, response *api.CreateClusterResult) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			defer conn.Close()

			peerConn, err := connect(ctx, response)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
				return
			}

			defer peerConn.Close()

			go func() {
				_, _ = io.Copy(conn, peerConn)
			}()

			_, _ = io.Copy(peerConn, conn)
		}()
	}
}

func connect(ctx context.Context, response *api.CreateClusterResult) (net.Conn, error) {
	return api.DialPort(ctx, response.Cluster, int(response.BuildCluster.Colocated.TargetPort))
}

func NewBuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build an image in a build cluster.",
		Args:  cobra.MaximumNArgs(1),
	}

	dockerFile := cmd.Flags().StringP("file", "f", "", "If set, specifies what Dockerfile to build.")
	pushTarget := cmd.Flags().String("push", "", "If specified, pushes the image to the target repository.")
	pushToRepository := cmd.Flags().String("push_to_nsc_repo", "", "If specified, pushes the image to nsc's private registry, to the specified repository.")
	tags := cmd.Flags().StringArrayP("tag", "t", nil, "List of tags to attach to the image.")

	cmd.RunE = fncobra.RunE(func(ctx context.Context, specifiedArgs []string) error {
		if *pushTarget == "" && *pushToRepository == "" {
			return fnerrors.New("one of --push or --push_to_nsc_repo are required")
		}

		buildctlBin, err := buildctl.EnsureSDK(ctx, host.HostPlatform())
		if err != nil {
			return fnerrors.New("failed to download buildctl: %w", err)
		}

		p, err := runBuildProxy(ctx)
		if err != nil {
			return err
		}

		defer p.Cleanup()

		contextDir := "."
		if len(specifiedArgs) > 0 {
			contextDir = specifiedArgs[0]
		}

		var imageName string

		if *pushTarget != "" {
			imageName = *pushTarget
		}

		if *pushToRepository != "" {
			imageName = fmt.Sprintf("%s/%s", p.RegistryEndpoint, *pushToRepository)
		}

		parsed, err := name.NewTag(imageName)
		if err != nil {
			return err
		}

		imageNames := []string{imageName}
		for _, tag := range *tags {
			imageNames = append(imageNames, parsed.Tag(tag).Name())
		}

		args := []string{
			"build",
			"--frontend=dockerfile.v0",
			"--local", "context=" + contextDir,
			"--local", "dockerfile=" + contextDir,
			// buildctl parses output as csv; need to quote to pass commas to `name`.
			"--output", fmt.Sprintf("type=image,push=true,%q", "name="+strings.Join(imageNames, ",")),
		}

		if *dockerFile != "" {
			args = append(args, "--opt", "filename="+*dockerFile)
		}

		if err := runBuildctl(ctx, buildctlBin, p, args...); err != nil {
			return err
		}

		fmt.Fprintf(console.Stdout(ctx), "Pushed:\n")
		for _, imageName := range imageNames {
			fmt.Fprintf(console.Stdout(ctx), "  %s\n", imageName)
		}
		return nil
	})

	return cmd
}
