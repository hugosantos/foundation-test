// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package auth

import (
	"context"
	"fmt"

	"github.com/morikuni/aec"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"

	"namespacelabs.dev/foundation/internal/auth"
	"namespacelabs.dev/foundation/internal/clerk"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnapi"
	"namespacelabs.dev/foundation/internal/fnerrors"
)

func NewLoginCmd() *cobra.Command {
	var kind string
	var openBrowser bool

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to Namespace to use Ephemeral Clusters, Remote Builders, etc.",
		Args:  cobra.NoArgs,

		RunE: fncobra.RunE(func(ctx context.Context, args []string) error {
			res, err := fnapi.StartLogin(ctx, kind, auth.Workspace)
			if err != nil {
				return err
			}

			stdout := console.Stdout(ctx)
			fmt.Fprintf(stdout, "%s\n", aec.Bold.Apply("Login to Namespace"))

			if openBrowser && openURL(res.LoginUrl) {
				fmt.Fprintf(stdout, "Please complete the login flow in your browser.\n\n  %s\n", res.LoginUrl)
			} else {
				fmt.Fprintf(stdout, "In order to login, open the following URL in your browser:\n\n  %s\n", res.LoginUrl)
			}

			tenant, err := completeLogin(ctx, res.LoginId, res.Kind)
			if err != nil {
				return err
			}

			if err := auth.StoreTenantToken(tenant.token); err != nil {
				return err
			}

			if tenant.name != "" {
				fmt.Fprintf(stdout, "\nYou are now logged into workspace %q, have a nice day.\n", tenant.name)
			} else {
				fmt.Fprintf(stdout, "\nYou are now logged in, have a nice day.\n")
			}

			return nil
		}),
	}

	cmd.Flags().StringVar(&kind, "kind", "", "Internal kind.")
	_ = cmd.Flags().MarkHidden("kind")

	cmd.Flags().BoolVar(&openBrowser, "browser", true, "Open a browser to login.")

	return cmd
}

func openURL(url string) bool {
	err := browser.OpenURL(url)
	return err == nil
}

type tenant struct {
	token string
	name  string
}

func completeLogin(ctx context.Context, id, kind string) (tenant, error) {
	if kind == "tenant" {
		res, err := fnapi.CompleteTenantLogin(ctx, id)
		if err != nil {
			return tenant{}, err
		}

		return tenant{
			token: res.TenantToken,
			name:  res.TenantName,
		}, nil
	}

	// TODO remove old login path
	userAuth, err := getUserAuth(ctx, id, kind)
	if err != nil {
		return tenant{}, err
	}

	if _, err := auth.StoreUser(ctx, userAuth); err != nil {
		return tenant{}, err
	}

	userToken, err := auth.GenerateTokenFromUserAuth(ctx, userAuth)
	if err != nil {
		return tenant{}, err
	}

	tt, err := fnapi.ExchangeUserToken(ctx, userToken)
	if err != nil {
		return tenant{}, err
	}

	return tenant{token: tt.TenantToken}, nil
}

func getUserAuth(ctx context.Context, id, kind string) (*auth.UserAuth, error) {
	switch kind {
	case "tenant":
		return nil, fnerrors.InternalError("tenant login cannot produce user auth")

	case "clerk":
		t, err := fnapi.CompleteClerkLogin(ctx, id)
		if err != nil {
			return nil, err
		}

		n, err := clerk.Login(ctx, t.Ticket)
		if err != nil {
			return nil, err
		}

		return &auth.UserAuth{
			Username: n.Email,
			Clerk:    n,
		}, nil

	default:
		return fnapi.CompleteLogin(ctx, id)
	}
}
