// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup_test

import (
	"context"
	"testing"
	"xmlgroup"

	"github.com/stretchr/testify/require"
)

func TestXMLModelWithSimpleArraysValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithSimpleArraysValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithSimpleArrays{
		Colors: []string{"red", "green", "blue"},
		Counts: []int32{1, 2},
	}, resp.ModelWithSimpleArrays)
}

func TestXMLModelWithSimpleArraysValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithSimpleArraysValueClient().Put(context.Background(), xmlgroup.ModelWithSimpleArrays{
		Colors: []string{"red", "green", "blue"},
		Counts: []int32{1, 2},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
