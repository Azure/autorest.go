// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package jsongroup_test

import (
	"context"
	"jsongroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestJSONClientGet(t *testing.T) {
	client, err := jsongroup.NewJSONClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewJSONPropertyClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.True(t, *resp.DefaultName)
}

func TestJSONClientSend(t *testing.T) {
	client, err := jsongroup.NewJSONClient("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewJSONPropertyClient().Send(context.Background(), jsongroup.JSONEncodedNameModel{DefaultName: to.Ptr(true)}, nil)
	require.NoError(t, err)
}
