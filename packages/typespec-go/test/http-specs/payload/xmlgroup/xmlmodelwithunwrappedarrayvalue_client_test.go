// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup_test

import (
	"context"
	"testing"
	"xmlgroup"

	"github.com/stretchr/testify/require"
)

func TestXMLModelWithUnwrappedArrayValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithUnwrappedArrayValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithUnwrappedArray{
		Colors: []string{"red", "green", "blue"},
		Counts: []int32{1, 2},
	}, resp.ModelWithUnwrappedArray)
}

func TestXMLModelWithUnwrappedArrayValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithUnwrappedArrayValueClient().Put(context.Background(), xmlgroup.ModelWithUnwrappedArray{
		Colors: []string{"red", "green", "blue"},
		Counts: []int32{1, 2},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
