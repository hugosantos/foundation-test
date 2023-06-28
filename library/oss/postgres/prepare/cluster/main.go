// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"fmt"
	"log"
	"net"

	"namespacelabs.dev/foundation/framework/resources"
	"namespacelabs.dev/foundation/framework/resources/provider"
	postgresclass "namespacelabs.dev/foundation/library/database/postgres"
	"namespacelabs.dev/foundation/library/oss/postgres"
)

const (
	providerPkg = "namespacelabs.dev/foundation/library/oss/postgres"
	user        = "postgres"
)

func main() {
	_, p := provider.MustPrepare[*postgres.ClusterIntent]()

	endpoint, err := resources.LookupServerEndpoint(p.Resources, fmt.Sprintf("%s:server", providerPkg), "postgres")
	if err != nil {
		log.Fatalf("failed to get Postgres server endpoint: %v", err)
	}

	password, err := resources.ReadSecret(p.Resources, fmt.Sprintf("%s:password", providerPkg))
	if err != nil {
		log.Fatalf("failed to read Postgres password: %v", err)
	}

	host, port, err := net.SplitHostPort(endpoint)
	if err != nil {
		log.Fatalf("invalid Postgres endpoint %q: %v", endpoint, err)
	}

	instance := &postgresclass.ClusterInstance{
		Address:  endpoint,
		User:     user,
		Password: string(password),
		Host:     host,
		Port:     port,
	}

	p.EmitResult(instance)
}
