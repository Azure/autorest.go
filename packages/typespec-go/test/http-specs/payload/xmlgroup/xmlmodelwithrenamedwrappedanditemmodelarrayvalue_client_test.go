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

func TestXMLModelWithRenamedWrappedAndItemModelArrayValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithRenamedWrappedAndItemModelArrayValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithRenamedWrappedAndItemModelArray{
		Books: []xmlgroup.Book{
			{
				Title: to.Ptr("The Great Gatsby"),
			},
			{
				Title: to.Ptr("Les Miserables"),
			},
		},
	}, resp.ModelWithRenamedWrappedAndItemModelArray)
}

func TestXMLModelWithRenamedWrappedAndItemModelArrayValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithRenamedWrappedAndItemModelArrayValueClient().Put(context.Background(), xmlgroup.ModelWithRenamedWrappedAndItemModelArray{
		Books: []xmlgroup.Book{
			{
				Title: to.Ptr("The Great Gatsby"),
			},
			{
				Title: to.Ptr("Les Miserables"),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
