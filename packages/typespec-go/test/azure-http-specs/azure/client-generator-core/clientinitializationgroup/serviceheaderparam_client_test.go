// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitializationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceHeaderParamClient_WithBody_Success(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceHeaderParamClient()
	name := "test-name"
	resp, err := client.WithBody(context.Background(), "test-name-value", Input{Name: &name}, nil)
	require.NoError(t, err)
	require.Equal(t, ServiceHeaderParamClientWithBodyResponse{}, resp)
}

func TestServiceHeaderParamClient_WithQuery(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceHeaderParamClient()

	_, err = client.WithQuery(context.Background(), "test-name-value", "test-id", &ServiceHeaderParamClientWithQueryOptions{})
	require.NoError(t, err)
}
