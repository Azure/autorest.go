// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup_test

import (
	"context"
	"generatortests/urlgroup"
	"generatortests/urlgroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeGetGlobalAndLocalQueryNull(t *testing.T) {
	pathParam1 := "pathParam1"
	pathParam2 := "pathParam2"
	opts := urlgroup.PathItemsClientGetGlobalAndLocalQueryNullOptions{
		LocalStringQuery:    to.Ptr("string1"),
		PathItemStringQuery: to.Ptr("string2"),
	}
	server := fake.PathItemsServer{
		GetGlobalAndLocalQueryNull: func(ctx context.Context, pathItemStringPath, localStringPath string, options *urlgroup.PathItemsClientGetGlobalAndLocalQueryNullOptions) (resp azfake.Responder[urlgroup.PathItemsClientGetGlobalAndLocalQueryNullResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, pathParam1, pathItemStringPath)
			require.EqualValues(t, pathParam2, localStringPath)
			require.EqualValues(t, &opts, options)
			resp.SetResponse(http.StatusOK, urlgroup.PathItemsClientGetGlobalAndLocalQueryNullResponse{}, nil)
			return
		},
	}
	client, err := urlgroup.NewPathItemsClient("clientPathParam", nil, &azcore.ClientOptions{
		Transport: fake.NewPathItemsServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.GetGlobalAndLocalQueryNull(context.Background(), pathParam1, pathParam2, &opts)
	require.NoError(t, err)
	require.Zero(t, resp)
}
