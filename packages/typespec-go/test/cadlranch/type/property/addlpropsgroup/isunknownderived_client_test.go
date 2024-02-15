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

func TestIsUnknownDerivedClient_Get(t *testing.T) {
	client, err := addlpropsgroup.NewAdditionalPropertiesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewIsUnknownDerivedClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, addlpropsgroup.IsUnknownAdditionalPropertiesDerived{
		Index: to.Ptr[int32](314),
		Name:  to.Ptr("IsUnknownAdditionalProperties"),
		AdditionalProperties: map[string]any{
			"prop1": float64(32),
			"prop2": true,
			"prop3": "abc",
		},
		Age: to.Ptr[float32](2.71828),
	}, resp.IsUnknownAdditionalPropertiesDerived)
}

func TestIsUnknownDerivedClient_Put(t *testing.T) {
	client, err := addlpropsgroup.NewAdditionalPropertiesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewIsUnknownDerivedClient().Put(context.Background(), addlpropsgroup.IsUnknownAdditionalPropertiesDerived{
		Index: to.Ptr[int32](314),
		Name:  to.Ptr("IsUnknownAdditionalProperties"),
		AdditionalProperties: map[string]any{
			"prop1": float64(32),
			"prop2": true,
			"prop3": "abc",
		},
		Age: to.Ptr[float32](2.71828),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
