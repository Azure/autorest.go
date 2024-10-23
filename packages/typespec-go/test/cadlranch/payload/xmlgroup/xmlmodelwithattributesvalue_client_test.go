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

func TestXMLModelWithAttributesValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithAttributesValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithAttributes{
		Id1:     to.Ptr[int32](123),
		Id2:     to.Ptr("foo"),
		Enabled: to.Ptr(true),
	}, resp.ModelWithAttributes)
}

func TestXMLModelWithAttributesValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithAttributesValueClient().Put(context.Background(), xmlgroup.ModelWithAttributes{
		Id1:     to.Ptr[int32](123),
		Id2:     to.Ptr("foo"),
		Enabled: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
