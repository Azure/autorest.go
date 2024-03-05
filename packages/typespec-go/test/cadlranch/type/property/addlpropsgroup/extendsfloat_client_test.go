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

func TestExtendsFloatClient_Get(t *testing.T) {
	client, err := addlpropsgroup.NewAdditionalPropertiesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewExtendsFloatClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, addlpropsgroup.ExtendsFloatAdditionalProperties{
		ID: to.Ptr[float32](43.125),
		AdditionalProperties: map[string]*float32{
			"prop": to.Ptr[float32](43.125),
		},
	}, resp.ExtendsFloatAdditionalProperties)
}

func TestExtendsFloatClient_Put(t *testing.T) {
	client, err := addlpropsgroup.NewAdditionalPropertiesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewExtendsFloatClient().Put(context.Background(), addlpropsgroup.ExtendsFloatAdditionalProperties{
		ID: to.Ptr[float32](43.125),
		AdditionalProperties: map[string]*float32{
			"prop": to.Ptr[float32](43.125),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
