// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreinitializationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceHeaderParamClient_WithBody_Success(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceHeaderParamClient()

	resp, err := client.WithBody(context.Background(), "sample", Input{}, nil)
	require.NoError(t, err)
	require.Equal(t, ServiceHeaderParamClientWithBodyResponse{}, resp)
}

func TestServiceHeaderParamClient_WithBody_ErrorOnRequest(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceHeaderParamClient()

	_, err = client.WithBody(context.Background(), "sample", Input{}, nil)
	require.Error(t, err)
}

func TestServiceHeaderParamClient_WithBody_Non204Status(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceHeaderParamClient()

	_, err = client.WithBody(context.Background(), "sample", Input{}, nil)
	require.Error(t, err)
}
