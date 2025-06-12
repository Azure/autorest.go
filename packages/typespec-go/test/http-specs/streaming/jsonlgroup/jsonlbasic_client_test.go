// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package jsonlgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJsonlBasicClient_Receive(t *testing.T) {
	client, err := NewJsonlGroupClient(nil)
	require.NoError(t, err)
	_, err = client.Receive(context.Background(), &JsonlBasicClientReceiveOptions{})
	require.Error(t, err)
}

func TestJsonlBasicClient_Send(t *testing.T) {
	client, err := NewJsonlGroupClient(nil)
	require.NoError(t, err)
	_, err = client.Send(context.Background(), nil)
	require.Error(t, err)
}
