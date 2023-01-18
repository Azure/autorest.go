// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package booleangroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newBoolClient(t *testing.T) *BoolClient {
	client, err := NewBoolClient(nil)
	require.NoError(t, err)
	return client
}

func NewBoolClient(options *azcore.ClientOptions) (*BoolClient, error) {
	client, err := azcore.NewClient("booleangroup.BoolClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &BoolClient{internal: client}, nil
}

func TestGetTrue(t *testing.T) {
	client := newBoolClient(t)
	result, err := client.GetTrue(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr(true)); r != "" {
		t.Fatal(r)
	}
}

func TestGetFalse(t *testing.T) {
	client := newBoolClient(t)
	result, err := client.GetFalse(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr(false)); r != "" {
		t.Fatal(r)
	}
}

func TestGetNull(t *testing.T) {
	client := newBoolClient(t)
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, (*bool)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestGetInvalid(t *testing.T) {
	client := newBoolClient(t)
	result, err := client.GetInvalid(context.Background(), nil)
	// TODO: verify error response is clear and actionable
	require.Error(t, err)
	require.Zero(t, result)
}

func TestPutTrue(t *testing.T) {
	client := newBoolClient(t)
	result, err := client.PutTrue(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutFalse(t *testing.T) {
	client := newBoolClient(t)
	result, err := client.PutFalse(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
