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

func TestIsUnknownDiscriminatedClient_Get(t *testing.T) {
	client, err := addlpropsgroup.NewIsUnknownDiscriminatedClient(nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, &addlpropsgroup.IsUnknownAdditionalPropertiesDiscriminatedDerived{
		Index: to.Ptr[int32](314),
		Kind:  to.Ptr("derived"),
		Name:  to.Ptr("Derived"),
		AdditionalProperties: map[string]any{
			"prop1": float64(32),
			"prop2": true,
			"prop3": "abc",
		},
		Age: to.Ptr[float32](2.71828),
	}, resp.IsUnknownAdditionalPropertiesDiscriminatedClassification)
}

func TestIsUnknownDiscriminatedClient_Put(t *testing.T) {
	client, err := addlpropsgroup.NewIsUnknownDiscriminatedClient(nil)
	require.NoError(t, err)
	resp, err := client.Put(context.Background(), &addlpropsgroup.IsUnknownAdditionalPropertiesDiscriminatedDerived{
		Index: to.Ptr[int32](314),
		Kind:  to.Ptr("derived"),
		Name:  to.Ptr("Derived"),
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
