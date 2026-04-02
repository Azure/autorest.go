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

func TestXMLModelWithRenamedPropertyValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithRenamedPropertyValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithRenamedProperty{
		Title:  to.Ptr("foo"),
		Author: to.Ptr("bar"),
	}, resp.ModelWithRenamedProperty)
}

func TestXMLModelWithRenamedPropertyValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithRenamedPropertyValueClient().Put(context.Background(), xmlgroup.ModelWithRenamedProperty{
		Title:  to.Ptr("foo"),
		Author: to.Ptr("bar"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
