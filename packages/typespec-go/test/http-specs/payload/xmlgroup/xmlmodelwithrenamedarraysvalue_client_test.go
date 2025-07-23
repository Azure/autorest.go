// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup_test

import (
	"context"
	"testing"
	"xmlgroup"

	"github.com/stretchr/testify/require"
)

func TestXMLModelWithRenamedArraysValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithRenamedArraysValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, xmlgroup.ModelWithRenamedArrays{
		Colors: []string{"red", "green", "blue"},
		Counts: []int32{1, 2},
	}, resp.ModelWithRenamedArrays)
}

func TestXMLModelWithRenamedArraysValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithRenamedArraysValueClient().Put(context.Background(), xmlgroup.ModelWithRenamedArrays{
		Colors: []string{"red", "green", "blue"},
		Counts: []int32{1, 2},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
