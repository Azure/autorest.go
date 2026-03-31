// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup_test

import (
	"context"
	"testing"
	"xmlgroup"

	"github.com/stretchr/testify/require"
)

func TestXMLModelWithWrappedPrimitiveCustomItemNamesValueClient_Get(t *testing.T) {
	t.Skip("codegen bug: struct tag uses 'ItemsTags>string' but server returns 'ItemsTags>ItemName'")
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithWrappedPrimitiveCustomItemNamesValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, xmlgroup.ModelWithWrappedPrimitiveCustomItemNames{
		Tags: []string{"fiction", "classic"},
	}, resp.ModelWithWrappedPrimitiveCustomItemNames)
}

func TestXMLModelWithWrappedPrimitiveCustomItemNamesValueClient_Put(t *testing.T) {
	t.Skip("codegen bug: struct tag uses 'ItemsTags>string' but server expects 'ItemsTags>ItemName'")
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithWrappedPrimitiveCustomItemNamesValueClient().Put(context.Background(), xmlgroup.ModelWithWrappedPrimitiveCustomItemNames{
		Tags: []string{"fiction", "classic"},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
