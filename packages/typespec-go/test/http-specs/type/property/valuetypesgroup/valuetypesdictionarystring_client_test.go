// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestValueTypesDictionaryStringClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesDictionaryStringClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*string{
		"k1": to.Ptr("hello"),
		"k2": to.Ptr("world"),
	}, resp.Property)
}

func TestValueTypesDictionaryStringClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesDictionaryStringClient().Put(context.Background(), valuetypesgroup.DictionaryStringProperty{
		Property: map[string]*string{
			"k1": to.Ptr("hello"),
			"k2": to.Ptr("world"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
