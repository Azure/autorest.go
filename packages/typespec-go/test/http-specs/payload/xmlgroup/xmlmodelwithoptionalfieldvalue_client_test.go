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

func TestXMLModelWithOptionalFieldValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithOptionalFieldValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithOptionalField{
		Item: to.Ptr("widget"),
	}, resp.ModelWithOptionalField)
}

func TestXMLModelWithOptionalFieldValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithOptionalFieldValueClient().Put(context.Background(), xmlgroup.ModelWithOptionalField{
		Item: to.Ptr("widget"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
