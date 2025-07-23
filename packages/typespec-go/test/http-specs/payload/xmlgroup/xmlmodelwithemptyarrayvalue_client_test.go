// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup_test

import (
	"context"
	"testing"
	"xmlgroup"

	"github.com/stretchr/testify/require"
)

func TestXMLModelWithEmptyArrayValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithEmptyArrayValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithEmptyArray{}, resp.ModelWithEmptyArray)
}

func TestXMLModelWithEmptyArrayValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithEmptyArrayValueClient().Put(context.Background(), xmlgroup.ModelWithEmptyArray{
		Items: []xmlgroup.SimpleModel{},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
