// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreinitializationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceMultipleParamsClient_WithBody(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceMultipleParamsClient()

	resp, err := client.WithBody(context.Background(), "name1", "region1", Input{}, nil)
	require.NoError(t, err)
	require.Equal(t, ServiceMultipleParamsClientWithBodyResponse{}, resp)
}

func TestServiceMultipleParamsClient_WithQuery(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceMultipleParamsClient()

	resp, err := client.WithQuery(context.Background(), "name1", "region1", "id1", nil)
	require.NoError(t, err)
	require.Equal(t, ServiceMultipleParamsClientWithQueryResponse{}, resp)
}
