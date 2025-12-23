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

func TestIsModelClient_Get(t *testing.T) {
	client, err := addlpropsgroup.NewAdditionalPropertiesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewAdditionalPropertiesIsModelClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, addlpropsgroup.IsModelAdditionalProperties{
		AdditionalProperties: map[string]*addlpropsgroup.ModelForRecord{
			"knownProp": {
				State: to.Ptr("ok"),
			},
			"prop": {
				State: to.Ptr("ok"),
			},
		},
	}, resp.IsModelAdditionalProperties)
}

func TestIsModelClient_Put(t *testing.T) {
	client, err := addlpropsgroup.NewAdditionalPropertiesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewAdditionalPropertiesIsModelClient().Put(context.Background(), addlpropsgroup.IsModelAdditionalProperties{
		AdditionalProperties: map[string]*addlpropsgroup.ModelForRecord{
			"knownProp": {
				State: to.Ptr("ok"),
			},
			"prop": {
				State: to.Ptr("ok"),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
