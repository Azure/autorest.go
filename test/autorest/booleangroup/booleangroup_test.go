// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package booleangroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newBoolClient() *BoolClient {
	return NewBoolClient(nil)
}

func TestGetTrue(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetTrue(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.BoolPtr(true)); r != "" {
		t.Fatal(r)
	}
}

func TestGetFalse(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetFalse(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.BoolPtr(false)); r != "" {
		t.Fatal(r)
	}
}

func TestGetNull(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, (*bool)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestGetInvalid(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetInvalid(context.Background(), nil)
	// TODO: verify error response is clear and actionable
	require.Error(t, err)
	require.Zero(t, result)
}

func TestPutTrue(t *testing.T) {
	client := newBoolClient()
	result, err := client.PutTrue(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutFalse(t *testing.T) {
	client := newBoolClient()
	result, err := client.PutFalse(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
