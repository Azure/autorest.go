// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup_test

import (
	"context"
	"testing"
	"xmlgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestXMLModelWithEncodedNamesValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithEncodedNamesValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, xmlgroup.ModelWithEncodedNames{
		Colors: []string{"red", "green", "blue"},
		ModelData: &xmlgroup.SimpleModel{
			Age:  to.Ptr(int32(123)),
			Name: to.Ptr("foo"),
		},
	}, resp.ModelWithEncodedNames)
}

func TestXMLModelWithEncodedNamesValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithEncodedNamesValueClient().Put(context.Background(), xmlgroup.ModelWithEncodedNames{
		Colors: []string{"red", "green", "blue"},
		ModelData: &xmlgroup.SimpleModel{
			Age:  to.Ptr(int32(123)),
			Name: to.Ptr("foo"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
