// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package jsonlgroup_test

import (
	"context"
	"jsonlgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApiKeyClient_Invalid(t *testing.T) {
	client, err := jsonlgroup.NewJsonlGroupClient(nil)
	require.NoError(t, err)
	resp, err := client.NewJsonlBasicClient().Receive(context.Background(), &jsonlgroup.JsonlBasicClientReceiveOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestApiKeyClient_OutputToInputOutput(t *testing.T) {
	client, err := jsonlgroup.NewJsonlGroupClient(nil)
	require.NoError(t, err)
	resp, err := client.NewJsonlBasicClient().Send(context.Background(), &jsonlgroup.JsonlBasicClientSendOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}
