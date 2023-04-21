// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"generatortests"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newXMSClientRequestIDClient(t *testing.T) *XMSClientRequestIDClient {
	client, err := NewXMSClientRequestIDClient(nil)
	require.NoError(t, err)
	return client
}

func NewXMSClientRequestIDClient(options *azcore.ClientOptions) (*XMSClientRequestIDClient, error) {
	client, err := azcore.NewClient("azurespecialsgroup.XMSClientRequestIDClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &XMSClientRequestIDClient{internal: client}, nil
}

// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func TestGet(t *testing.T) {
	client := newXMSClientRequestIDClient(t)
	result, err := client.Get(runtime.WithHTTPHeader(context.Background(), http.Header{
		"x-ms-client-request-id": []string{"9C4D50EE-2D56-4CD3-8152-34347DC9F2B0"},
	}), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func TestParamGet(t *testing.T) {
	client := newXMSClientRequestIDClient(t)
	result, err := client.ParamGet(context.Background(), "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
