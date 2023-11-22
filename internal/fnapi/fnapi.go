// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package fnapi

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/pflag"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"namespacelabs.dev/foundation/internal/auth"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/versions"
	"namespacelabs.dev/go-ids"
)

var (
	EndpointAddress             = "https://api.namespacelabs.net"
	AdminMode                   = false
	ExchangeGithubToTenantToken = false

	UserAgent = "ns/unknown"
)

func SetupFlags(flags *pflag.FlagSet) {
	flags.StringVar(&EndpointAddress, "fnapi_endpoint", EndpointAddress, "The fnapi endpoint address.")
	_ = flags.MarkHidden("fnapi_endpoint")
	flags.BoolVar(&AdminMode, "fnapi_admin", AdminMode, "Whether to enable admin mode.")
	_ = flags.MarkHidden("fnapi_admin")
}

func ResolveStaticEndpoint(endpoint string) func(context.Context, Token) (string, error) {
	return func(context.Context, Token) (string, error) {
		return endpoint, nil
	}
}

// A nil handle indicates that the caller wants to discard the response.
func AnonymousCall(ctx context.Context, endpoint string, method string, req interface{}, handle func(io.Reader) error) error {
	return Call[any]{
		Method:     method,
		FetchToken: nil, // Callers of this API do not assume that credentials are injected.
	}.Do(ctx, req, ResolveStaticEndpoint(endpoint), handle)
}

func AuthenticatedCall(ctx context.Context, endpoint string, method string, req interface{}, handle func(io.Reader) error) error {
	return Call[any]{
		Method:     method,
		FetchToken: FetchToken,
	}.Do(ctx, req, ResolveStaticEndpoint(endpoint), handle)
}

type Token interface {
	Claims(context.Context) (*auth.TokenClaims, error)
	IssueToken(context.Context, time.Duration, func(context.Context, string, time.Duration) (string, error)) (string, error)
}

func FetchSessionToken(ctx context.Context, sessionToken string, duration time.Duration) (string, error) {
	req := IssueTenantTokenFromSessionRequest{
		SessionToken:      sessionToken,
		TokenDurationSecs: int64(duration.Seconds()),
	}

	var resp IssueTenantTokenFromSessionResponse
	if err := AnonymousCall(ctx, EndpointAddress, "nsl.signin.SigninService/IssueTenantTokenFromSession", req, DecodeJSONResponse(&resp)); err != nil {
		return "", err
	}

	return resp.TenantToken, nil
}

func BearerToken(ctx context.Context, t Token) (string, error) {
	raw, err := t.IssueToken(ctx, 5*time.Minute, FetchSessionToken)
	if err != nil {
		return "", err
	}

	return "Bearer " + raw, nil
}

type Call[RequestT any] struct {
	Method     string
	FetchToken func(context.Context) (Token, error)
}

func DecodeJSONResponse(resp any) func(io.Reader) error {
	return func(body io.Reader) error {
		return json.NewDecoder(body).Decode(resp)
	}
}

func AddNamespaceHeaders(ctx context.Context, headers *http.Header) {
	headers.Add("NS-Internal-Version", fmt.Sprintf("%d", versions.Builtin().APIVersion))
	headers.Add("User-Agent", UserAgent)

	if AdminMode {
		headers.Add("NS-API-Mode", "admin")
	}
}

func (c Call[RequestT]) Do(ctx context.Context, request RequestT, resolveEndpoint func(context.Context, Token) (string, error), handle func(io.Reader) error) error {
	headers := http.Header{}

	var resolvedToken Token
	if c.FetchToken != nil {
		tok, err := c.FetchToken(ctx)
		if err != nil {
			return err
		}

		resolvedToken = tok

		bearerToken, err := BearerToken(ctx, tok)
		if err != nil {
			return err
		}

		headers.Add("Authorization", bearerToken)
	}

	AddNamespaceHeaders(ctx, &headers)

	reqBytes, err := json.Marshal(request)
	if err != nil {
		return fnerrors.InternalError("failed to marshal request: %w", err)
	}

	endpoint, err := resolveEndpoint(ctx, resolvedToken)
	if err != nil {
		return err
	}

	tid := ids.NewRandomBase32ID(4)
	fmt.Fprintf(console.Debug(ctx), "[%s] RPC: %v (endpoint: %v)\n", tid, c.Method, endpoint)
	fmt.Fprintf(console.Debug(ctx), "[%s] Body: %s\n", tid, reqBytes)

	t := time.Now()
	url := endpoint + "/" + c.Method
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBytes))
	if err != nil {
		return fnerrors.InternalError("failed to construct request: %w", err)
	}

	for k, v := range headers {
		httpReq.Header[k] = append(httpReq.Header[k], v...)
	}

	client := &http.Client{
		Transport: http.DefaultTransport,
	}

	response, err := client.Do(httpReq)
	if err != nil {
		return fnerrors.InvocationError("namespace api", "http call failed: %w", err)
	}

	defer response.Body.Close()

	fmt.Fprintf(console.Debug(ctx), "[%s] RPC: %v: status %s took %v\n", tid, c.Method, response.Status, time.Since(t))

	if response.StatusCode == http.StatusOK {
		if handle == nil {
			return nil
		}

		return handle(response.Body)
	}

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fnerrors.InvocationError("namespace api", "reading response body: %w", err)
	}

	st := &spb.Status{}
	if err := json.Unmarshal(respBody, st); err == nil {
		return handleGrpcStatus(url, st)
	}

	fmt.Fprintf(console.Debug(ctx), "Error body response: %s\n", string(respBody))

	if grpcDetails := response.Header[http.CanonicalHeaderKey("grpc-status-details-bin")]; len(grpcDetails) > 0 {
		data, err := base64.RawStdEncoding.DecodeString(grpcDetails[0])
		if err != nil {
			return fnerrors.InternalError("failed to decode grpc details: %w", err)
		}

		if err := proto.Unmarshal(data, st); err != nil {
			return fnerrors.InternalError("failed to unmarshal grpc details: %w", err)
		}

		return handleGrpcStatus(url, st)
	}

	grpcMessage := response.Header[http.CanonicalHeaderKey("grpc-message")]
	grpcStatus := response.Header[http.CanonicalHeaderKey("grpc-status")]

	if len(grpcMessage) > 0 && len(grpcStatus) > 0 {
		intVar, err := strconv.Atoi(grpcStatus[0])
		if err == nil {
			st.Code = int32(intVar)
			st.Message = grpcMessage[0]

			return handleGrpcStatus(url, st)
		}
	}

	switch response.StatusCode {
	case http.StatusInternalServerError:
		return fnerrors.InternalError("namespace api: internal server error: %s", string(respBody))
	case http.StatusUnauthorized:
		return fnerrors.ReauthError("%s requires authentication", url)
	case http.StatusForbidden:
		return fnerrors.PermissionDeniedError("%s denied access", url)
	case http.StatusNotFound:
		return fnerrors.InternalError("%s not found: %s", url, string(respBody))
	default:
		return fnerrors.InvocationError("namespace api", "unexpected %d error reaching %q: %s", response.StatusCode, url, response.Status)
	}
}

func handleGrpcStatus(url string, st *spb.Status) error {
	switch st.Code {
	case int32(codes.Unauthenticated):
		return fnerrors.ReauthError("%s requires authentication: %w", url, status.ErrorProto(st))

	case int32(codes.PermissionDenied):
		return fnerrors.PermissionDeniedError("%s denied access: %w", url, status.ErrorProto(st))

	case int32(codes.FailedPrecondition):
		// Failed precondition is not retryable so we should not suggest that it is transient (e.g. invocation error suggests this).
		return fnerrors.New("failed to call %s: %w", url, status.ErrorProto(st))

	case int32(codes.Internal):
		return fnerrors.InternalError("failed to call %s: %w", url, status.ErrorProto(st))

	default:
		return fnerrors.InvocationError("namespace api", "failed to call %s: %w", url, status.ErrorProto(st))
	}
}
