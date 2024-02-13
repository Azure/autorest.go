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

func TestExtendsUnknownClient_Get(t *testing.T) {
	client, err := addlpropsgroup.NewAdditionalPropertiesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewExtendsUnknownClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, addlpropsgroup.ExtendsUnknownAdditionalProperties{
		Name: to.Ptr("ExtendsUnknownAdditionalProperties"),
		AdditionalProperties: map[string]any{
			"prop1": float64(32),
			"prop2": true,
			"prop3": "abc",
		},
	}, resp.ExtendsUnknownAdditionalProperties)
}

func TestExtendsUnknownClient_Put(t *testing.T) {
	client, err := addlpropsgroup.NewAdditionalPropertiesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewExtendsUnknownClient().Put(context.Background(), addlpropsgroup.ExtendsUnknownAdditionalProperties{
		Name: to.Ptr("ExtendsUnknownAdditionalProperties"),
		AdditionalProperties: map[string]any{
			"prop1": float64(32),
			"prop2": true,
			"prop3": "abc",
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
