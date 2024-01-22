//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package addlpropsgroup_test

import (
	"addlpropsgroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestIsUnknownClient_Get(t *testing.T) {
	client, err := addlpropsgroup.NewIsUnknownClient(nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, addlpropsgroup.IsUnknownAdditionalProperties{
		Name: to.Ptr("IsUnknownAdditionalProperties"),
		AdditionalProperties: map[string]any{
			"prop1": float64(32),
			"prop2": true,
			"prop3": "abc",
		},
	}, resp.IsUnknownAdditionalProperties)
}

func TestIsUnknownClient_Put(t *testing.T) {
	client, err := addlpropsgroup.NewIsUnknownClient(nil)
	require.NoError(t, err)
	resp, err := client.Put(context.Background(), addlpropsgroup.IsUnknownAdditionalProperties{
		Name: to.Ptr("IsUnknownAdditionalProperties"),
		AdditionalProperties: map[string]any{
			"prop1": float64(32),
			"prop2": true,
			"prop3": "abc",
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
