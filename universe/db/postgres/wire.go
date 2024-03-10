// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package postgres

import (
	"context"

	"namespacelabs.dev/foundation/framework/resources"
	"namespacelabs.dev/foundation/internal/fnerrors"
)

func ProvideDatabase(ctx context.Context, db *DatabaseArgs, deps ExtensionDeps) (*DB, error) {
	if db.ResourceRef == "" {
		return nil, fnerrors.New("resourceRef is required")
	}

	res, err := resources.LoadResources()
	if err != nil {
		return nil, err
	}

	tp, err := deps.OpenTelemetry.GetTracerProvider()
	if err != nil {
		return nil, err
	}

	return ConnectToResource(ctx, res, db.ResourceRef, tp, &ConfigOverrides{
		MaxConns: db.MaxConns,
	})
}
