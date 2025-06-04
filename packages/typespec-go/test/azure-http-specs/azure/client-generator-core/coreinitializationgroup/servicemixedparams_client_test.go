// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreinitializationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceMixedParamsClient_WithBody(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceMixedParamsClient()

	body := WithBodyRequest{}
	resp, err := client.WithBody(context.Background(), "name1", "region1", body, nil)
	require.NoError(t, err)
	require.Equal(t, ServiceMixedParamsClientWithBodyResponse{}, resp)
}

func TestServiceMixedParamsClient_WithQuery(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceMixedParamsClient()

	resp, err := client.WithQuery(context.Background(), "name1", "region1", "id1", nil)
	require.NoError(t, err)
	require.Equal(t, ServiceMixedParamsClientWithQueryResponse{}, resp)
}
