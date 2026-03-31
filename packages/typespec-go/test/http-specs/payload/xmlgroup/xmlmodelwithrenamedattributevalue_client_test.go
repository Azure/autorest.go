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

func TestXMLModelWithRenamedAttributeValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithRenamedAttributeValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithRenamedAttribute{
		ID:     to.Ptr[int32](123),
		Title:  to.Ptr("The Great Gatsby"),
		Author: to.Ptr("F. Scott Fitzgerald"),
	}, resp.ModelWithRenamedAttribute)
}

func TestXMLModelWithRenamedAttributeValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithRenamedAttributeValueClient().Put(context.Background(), xmlgroup.ModelWithRenamedAttribute{
		ID:     to.Ptr[int32](123),
		Title:  to.Ptr("The Great Gatsby"),
		Author: to.Ptr("F. Scott Fitzgerald"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
