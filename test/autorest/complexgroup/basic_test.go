// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newBasicClient() *BasicClient {
	return NewBasicClient(nil)
}

func TestBasicGetValid(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Basic, Basic{ID: to.Int32Ptr(2), Name: to.StringPtr("abc"), Color: CMYKColorsYELLOW.ToPtr()}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicPutValid(t *testing.T) {
	client := newBasicClient()
	result, err := client.PutValid(context.Background(), Basic{
		ID:    to.Int32Ptr(2),
		Name:  to.StringPtr("abc"),
		Color: CMYKColorsMagenta.ToPtr(),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestBasicGetInvalid(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetInvalid(context.Background(), nil)
	require.Error(t, err)
	if r := cmp.Diff(result, BasicClientGetInvalidResponse{}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicGetEmpty(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Basic, Basic{}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicGetNull(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Basic, Basic{}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicGetNotProvided(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetNotProvided(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Basic, Basic{}); r != "" {
		t.Fatal(r)
	}
}
