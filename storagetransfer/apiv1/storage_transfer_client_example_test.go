// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package storagetransfer_test

import (
	"context"

	storagetransfer "cloud.google.com/go/storagetransfer/apiv1"
	"google.golang.org/api/iterator"
	storagetransferpb "google.golang.org/genproto/googleapis/storagetransfer/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_GetGoogleServiceAccount() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.GetGoogleServiceAccountRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#GetGoogleServiceAccountRequest.
	}
	resp, err := c.GetGoogleServiceAccount(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateTransferJob() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.CreateTransferJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#CreateTransferJobRequest.
	}
	resp, err := c.CreateTransferJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateTransferJob() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.UpdateTransferJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#UpdateTransferJobRequest.
	}
	resp, err := c.UpdateTransferJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetTransferJob() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.GetTransferJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#GetTransferJobRequest.
	}
	resp, err := c.GetTransferJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListTransferJobs() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.ListTransferJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#ListTransferJobsRequest.
	}
	it := c.ListTransferJobs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_PauseTransferOperation() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.PauseTransferOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#PauseTransferOperationRequest.
	}
	err = c.PauseTransferOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ResumeTransferOperation() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.ResumeTransferOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#ResumeTransferOperationRequest.
	}
	err = c.ResumeTransferOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_RunTransferJob() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.RunTransferJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#RunTransferJobRequest.
	}
	op, err := c.RunTransferJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_CreateAgentPool() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.CreateAgentPoolRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#CreateAgentPoolRequest.
	}
	resp, err := c.CreateAgentPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateAgentPool() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.UpdateAgentPoolRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#UpdateAgentPoolRequest.
	}
	resp, err := c.UpdateAgentPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetAgentPool() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.GetAgentPoolRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#GetAgentPoolRequest.
	}
	resp, err := c.GetAgentPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListAgentPools() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.ListAgentPoolsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#ListAgentPoolsRequest.
	}
	it := c.ListAgentPools(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_DeleteAgentPool() {
	ctx := context.Background()
	c, err := storagetransfer.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagetransferpb.DeleteAgentPoolRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/storagetransfer/v1#DeleteAgentPoolRequest.
	}
	err = c.DeleteAgentPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
