// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package tracing

import (
	"context"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
)

type HttpClientProvider struct {
	provider trace.TracerProvider
}

func (hp HttpClientProvider) New() *http.Client {
	return hp.Wrap(&http.Client{})
}

func (hp HttpClientProvider) Wrap(client *http.Client) *http.Client {
	client.Transport = hp.WrapTransport(client.Transport)
	return client
}

func (hp HttpClientProvider) WrapTransport(base http.RoundTripper) http.RoundTripper {
	return otelhttp.NewTransport(base, otelhttp.WithTracerProvider(hp.provider), otelhttp.WithPropagators(propagators))
}

func ProvideHttpClientProvider(ctx context.Context, _ *NoArgs, deps ExtensionDeps) (HttpClientProvider, error) {
	provider, err := getTracerProvider()
	if err != nil {
		return HttpClientProvider{}, err
	}

	return HttpClientProvider{provider}, nil
}
