// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package redis

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel/attribute"
	"namespacelabs.dev/foundation/std/go/core"
)

var redisServerEndpoint = flag.String("redis_endpoint", "", "Redis endpoint address.")

func ProvideRedis(ctx context.Context, args *RedisArgs, deps ExtensionDeps) (*redis.Client, error) {
	if *redisServerEndpoint == "" {
		return nil, errors.New("redis_endpoint is required")
	}

	client := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     *redisServerEndpoint,
		Password: os.Getenv("REDIS_ROOT_PASSWORD"),
		DB:       int(args.Database),
	})

	tp, err := deps.OpenTelemetry.GetTracerProvider()
	if err != nil {
		return nil, err
	}

	client.AddHook(redisotel.NewTracingHook(
		redisotel.WithAttributes(attribute.Int("redis.db", int(args.Database))),
		redisotel.WithTracerProvider(tp),
	))

	// Asynchronously wait until a database connection is ready.
	deps.ReadinessCheck.Register(fmt.Sprintf("redis/%s", core.InstantiationPathFromContext(ctx)),
		core.CheckAtStartupFunc(func(ctx context.Context) error {
			return client.Ping(ctx).Err()
		}))

	return client, nil
}
