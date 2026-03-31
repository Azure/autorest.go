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

func TestXMLModelWithNestedModelValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithNestedModelValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithNestedModel{
		Nested: &xmlgroup.SimpleModel{
			Name: to.Ptr("foo"),
			Age:  to.Ptr[int32](123),
		},
	}, resp.ModelWithNestedModel)
}

func TestXMLModelWithNestedModelValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithNestedModelValueClient().Put(context.Background(), xmlgroup.ModelWithNestedModel{
		Nested: &xmlgroup.SimpleModel{
			Name: to.Ptr("foo"),
			Age:  to.Ptr[int32](123),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
