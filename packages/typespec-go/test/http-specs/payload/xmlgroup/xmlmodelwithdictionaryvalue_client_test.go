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

func TestXMLModelWithDictionaryValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithDictionaryValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	// TODO: xml helper for maps lower-cases keys :(
	require.EqualValues(t, xmlgroup.ModelWithDictionary{
		Metadata: map[string]*string{
			"color":   to.Ptr("blue"),
			"count":   to.Ptr("123"),
			"enabled": to.Ptr("false"),
		},
	}, resp.ModelWithDictionary)
}

func TestXMLModelWithDictionaryValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithDictionaryValueClient().Put(context.Background(), xmlgroup.ModelWithDictionary{
		Metadata: map[string]*string{
			"Color":   to.Ptr("blue"),
			"Count":   to.Ptr("123"),
			"Enabled": to.Ptr("false"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
