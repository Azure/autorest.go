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

func TestExtendsModelClient_Get(t *testing.T) {
	client, err := addlpropsgroup.NewExtendsModelClient(nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, addlpropsgroup.ExtendsModelAdditionalProperties{
		AdditionalProperties: map[string]*addlpropsgroup.ModelForRecord{
			"prop": {
				State: to.Ptr("ok"),
			},
		},
	}, resp.ExtendsModelAdditionalProperties)
}

func TestExtendsModelClient_Put(t *testing.T) {
	client, err := addlpropsgroup.NewExtendsModelClient(nil)
	require.NoError(t, err)
	resp, err := client.Put(context.Background(), addlpropsgroup.ExtendsModelAdditionalProperties{
		AdditionalProperties: map[string]*addlpropsgroup.ModelForRecord{
			"prop": {
				State: to.Ptr("ok"),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}