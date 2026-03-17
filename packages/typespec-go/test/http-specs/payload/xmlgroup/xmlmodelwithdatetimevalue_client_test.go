// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup_test

import (
	"context"
	"testing"
	"time"
	"xmlgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestXMLModelWithDatetimeValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithDatetimeValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	rfc3339 := time.Date(2022, 8, 26, 18, 38, 0, 0, time.UTC)
	rfc7231 := time.Date(2022, 8, 26, 14, 38, 0, 0, time.UTC)
	require.NotNil(t, resp.RFC3339)
	require.True(t, rfc3339.Equal(*resp.RFC3339))
	require.NotNil(t, resp.RFC7231)
	require.True(t, rfc7231.Equal(*resp.RFC7231))
}

func TestXMLModelWithDatetimeValueClient_Put(t *testing.T) {
	t.Skip("Go's datetime.RFC1123 marshals as 'UTC' but the mock server expects 'GMT'")
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	rfc3339 := time.Date(2022, 8, 26, 18, 38, 0, 0, time.UTC)
	rfc7231 := time.Date(2022, 8, 26, 14, 38, 0, 0, time.UTC)
	resp, err := client.NewXMLModelWithDatetimeValueClient().Put(context.Background(), xmlgroup.ModelWithDatetime{
		RFC3339: to.Ptr(rfc3339),
		RFC7231: to.Ptr(rfc7231),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
