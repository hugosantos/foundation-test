// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"
	"fmt"

	"namespacelabs.dev/foundation/framework/testing"
	"namespacelabs.dev/foundation/internal/testdata/service/proto"
)

func main() {
	testing.Do(func(ctx context.Context, t testing.Test) error {
		endpoint := t.MustEndpoint("namespacelabs.dev/foundation/internal/testdata/service/localstacks3", "localstacks3")

		conn, err := t.NewClient(endpoint)
		if err != nil {
			return err
		}

		cli := proto.NewFileServiceClient(conn)

		// This file is not present in the bucket.
		if _, err = cli.Get(ctx, &proto.GetFileRequest{
			Filename: "file which is not present",
		}); err == nil {
			return fmt.Errorf("calling on unexistent bucket should fail")
		}

		if _, err = cli.Add(ctx, &proto.AddFileRequest{
			Filename: "foo",
			Contents: "bar",
		}); err != nil {
			return fmt.Errorf("add failed with %v", err)
		}

		resp, err := cli.Get(ctx, &proto.GetFileRequest{
			Filename: "foo",
		})
		if err != nil {
			return fmt.Errorf("get failed with %v", err)
		}

		if resp.Contents != "bar" {
			return fmt.Errorf("expected contents to be foocontent, got %s", resp.Contents)
		}

		return nil
	})
}
