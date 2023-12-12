// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package git

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

var AssumeSSHAuth = false

type EnvVars map[string]string

func NoPromptEnv() TupleList {
	// Disable password prompts as we don't handle them properly, yet.
	env := EnvVars{"GIT_TERMINAL_PROMPT": "0"}

	// Also disable prompting for passwords by the 'ssh' subprocess spawned by Git.
	//
	// See https://github.com/golang/go/blob/fad67f8a5342f4bc309f26f0ae021ce9d21724e6/src/cmd/go/internal/get/get.go#L129
	if os.Getenv("GIT_SSH") == "" && os.Getenv("GIT_SSH_COMMAND") == "" {
		env["GIT_SSH_COMMAND"] = "ssh -o ControlMaster=no -o BatchMode=yes"
	}

	// And one more source of Git prompts: the Git Credential Manager Core for Windows.
	//
	// See https://github.com/microsoft/Git-Credential-Manager-Core/blob/master/docs/environment.md#gcm_interactive.
	if os.Getenv("GCM_INTERACTIVE") == "" {
		env["GCM_INTERACTIVE"] = "never"
	}

	var overrides [][2]string

	env["GIT_CONFIG_COUNT"] = fmt.Sprintf("%d", len(overrides))
	for k, override := range overrides {
		env[fmt.Sprintf("GIT_CONFIG_KEY_%d", k)] = override[0]
		env[fmt.Sprintf("GIT_CONFIG_VALUE_%d", k)] = override[1]
	}

	return env.Deterministic()
}

func ResetConfigGlobal() (TupleList, func(), error) {
	f, err := os.CreateTemp(os.TempDir(), "empty")
	if err != nil {
		return TupleList{}, nil, err
	}

	if err := f.Close(); err != nil {
		os.Remove(f.Name())
		return TupleList{}, nil, err
	}

	env := EnvVars{"GIT_CONFIG_GLOBAL": f.Name()}

	return env.Deterministic(), func() {
		os.Remove(f.Name())
	}, nil
}

type TupleList [][2]string

func (vars EnvVars) Deterministic() TupleList {
	var t TupleList
	for k, v := range vars {
		t = append(t, [2]string{k, v})
	}
	slices.SortFunc(t, func(a, b [2]string) int {
		return strings.Compare(a[0], b[0])
	})
	return t
}

func (tl TupleList) Serialize() []string {
	var t []string
	for _, ent := range tl {
		t = append(t, fmt.Sprintf("%s=%s", ent[0], ent[1]))
	}
	return t
}
