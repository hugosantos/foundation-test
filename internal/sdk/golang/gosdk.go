// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package golang

import (
	"embed"
	"fmt"
	"path/filepath"

	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"golang.org/x/mod/semver"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/sdk/host"
)

var (
	//go:embed gosdk.json
	lib embed.FS

	v = &host.ParsedVersions{Name: "go", FS: lib, Filename: "gosdk.json"}
)

type LocalSDK = host.LocalSDK

func vv(version string) string {
	v := v.Get()

	for ver, minor := range v.Versions {
		// If we store `"1.22": "1.22.2"` then "1.22.1" should also map to "1.22".
		if semver.Compare("v"+minor, "v"+version) >= 0 {
			version = ver
		}
	}

	return version
}

func MatchLatestVersion(version string) string {
	version = vv(version)

	v := v.Get()

	actualVer, has := v.Versions[version]
	if !has {
		return version
	}

	return actualVer
}

func MatchSDK(version string, platform specs.Platform) (compute.Computable[LocalSDK], error) {
	return SDK(vv(version), platform)
}

func SDK(version string, platform specs.Platform) (compute.Computable[LocalSDK], error) {
	return v.SDK(version, platform, func(ver string, platform specs.Platform) (string, string) {
		return fmt.Sprintf("https://go.dev/dl/go%s.%s-%s.tar.gz", ver, platform.OS, platform.Architecture), "go/bin/go"
	})
}

func GoRoot(sdk LocalSDK) string { return filepath.Join(sdk.Path, "go") }
func GoBin(sdk LocalSDK) string  { return sdk.Binary }

func GoRootEnv(sdk LocalSDK) string {
	return fmt.Sprintf("GOROOT=%s", GoRoot(sdk))
}
