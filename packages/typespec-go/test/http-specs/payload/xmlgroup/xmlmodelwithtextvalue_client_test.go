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

func TestXMLModelWithTextValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithTextValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithText{
		Language: to.Ptr("foo"),
		Content:  to.Ptr("\n  This is some text.\n"),
	}, resp.ModelWithText)
}

func TestXMLModelWithTextValueClient_Put(t *testing.T) {
	t.Skip("Go's encoding/xml encodes newlines as &#xA; in chardata, which fast-xml-parser doesn't decode back to newlines")
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithTextValueClient().Put(context.Background(), xmlgroup.ModelWithText{
		Language: to.Ptr("foo"),
		Content:  to.Ptr("\n  This is some text.\n"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
