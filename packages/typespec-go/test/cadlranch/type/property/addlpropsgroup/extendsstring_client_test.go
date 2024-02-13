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

func TestExtendsStringClient_Get(t *testing.T) {
	client, err := addlpropsgroup.NewAdditionalPropertiesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewExtendsStringClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, addlpropsgroup.ExtendsStringAdditionalProperties{
		Name: to.Ptr("ExtendsStringAdditionalProperties"),
		AdditionalProperties: map[string]*string{
			"prop": to.Ptr("abc"),
		},
	}, resp.ExtendsStringAdditionalProperties)
}

func TestExtendsStringClient_Put(t *testing.T) {
	client, err := addlpropsgroup.NewAdditionalPropertiesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewExtendsStringClient().Put(context.Background(), addlpropsgroup.ExtendsStringAdditionalProperties{
		Name: to.Ptr("ExtendsStringAdditionalProperties"),
		AdditionalProperties: map[string]*string{
			"prop": to.Ptr("abc"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
