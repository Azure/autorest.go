// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package examplebasicgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicServiceOperationGroupClient_Basic_Success(t *testing.T) {
	client, err := NewBasicServiceOperationGroupClient(nil)
	require.NoError(t, err)
	expectedResp := ActionResponse{
		// Fill with expected fields
	}
	reqBody := ActionRequest{
		// Fill with test request fields
	}
	resp, err := client.Basic(context.Background(), "queryValue", "headerValue", reqBody, nil)
	require.NoError(t, err)
	require.Equal(t, expectedResp, resp.ActionResponse)
}

func TestBasicServiceOperationGroupClient_Basic_ErrorStatus(t *testing.T) {
	client, err := NewBasicServiceOperationGroupClient(nil)
	require.NoError(t, err)
	reqBody := ActionRequest{}
	_, err = client.Basic(context.Background(), "q", "h", reqBody, nil)
	require.Error(t, err)
}

func TestBasicServiceOperationGroupClient_Basic_TransportError(t *testing.T) {
	client, err := NewBasicServiceOperationGroupClient(nil)
	require.NoError(t, err)
	reqBody := ActionRequest{}
	_, err = client.Basic(context.Background(), "q", "h", reqBody, nil)
	require.Error(t, err)
}
