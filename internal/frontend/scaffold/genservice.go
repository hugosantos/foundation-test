// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package scaffold

import (
	"context"
	"text/template"

	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/schema"
)

const (
	serviceFileName = "service.cue"
)

type GenServiceOpts struct {
	ExportedServiceName string
	Framework           schema.Framework
	HttpBackendPkg      string
}

func CreateServiceScaffold(ctx context.Context, fsfs fnfs.ReadWriteFS, loc fnfs.Location, opts GenServiceOpts) error {
	return generateCueSource(ctx, fsfs, loc.Rel(serviceFileName), serviceTmpl, opts)
}

var serviceTmpl = template.Must(template.New(serviceFileName).Parse(`
import (
	"namespacelabs.dev/foundation/std/fn"
	"namespacelabs.dev/foundation/std/fn:inputs"
	{{if .HttpBackendPkg -}}
	"namespacelabs.dev/foundation/std/web/http"
	{{end -}}
)

{{if .ExportedServiceName}}
// Load the protobuf definition so its contents are available to Namespace.
$proto: inputs.#Proto & {
	source: "service.proto"
}
{{end}}

service: fn.#Service & {
	framework: "{{.Framework}}"

	{{if .ExportedServiceName}}
	// Export a grpc-based API, defined within service.proto.
	exportService: $proto.services.{{.ExportedServiceName}}

	// Make this service available to the public Internet.
	ingress:              "INTERNET_FACING"
	{{end}}

	{{if .HttpBackendPkg}}
	instantiate: {
		// Wire the API backend's configuration (e.g. public address) automatically.
		apiBackend: http.#Exports.Backend & {
			endpointOwner: "{{.HttpBackendPkg}}"
			manager:       "namespacelabs.dev/foundation/std/grpc/httptranscoding"
		}
	}
	{{end}}
}
`))
