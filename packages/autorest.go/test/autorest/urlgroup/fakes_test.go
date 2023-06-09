// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup_test

import (
	"context"
	"generatortests/urlgroup"
	"generatortests/urlgroup/fake"
	"net/http"
	"testing"
	"time"

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

func TestFakeArrayCSVInPath(t *testing.T) {
	intputArrayPath := []string{
		"/encoded;",
		"notencoded",
		"?alsoencoded@",
	}
	server := fake.PathsServer{
		ArrayCSVInPath: func(ctx context.Context, arrayPath []string, options *urlgroup.PathsClientArrayCSVInPathOptions) (resp azfake.Responder[urlgroup.PathsClientArrayCSVInPathResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, intputArrayPath, arrayPath)
			resp.SetResponse(http.StatusOK, urlgroup.PathsClientArrayCSVInPathResponse{}, nil)
			return
		},
	}
	client, err := urlgroup.NewPathsClient(&azcore.ClientOptions{
		Transport: fake.NewPathsServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.ArrayCSVInPath(context.Background(), intputArrayPath, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFakeArrayStringCSVEmpty(t *testing.T) {
	inputQuery := []string{
		"/encoded;",
		"notencoded",
		"?alsoencoded@",
	}
	server := fake.QueriesServer{
		ArrayStringCSVEmpty: func(ctx context.Context, options *urlgroup.QueriesClientArrayStringCSVEmptyOptions) (resp azfake.Responder[urlgroup.QueriesClientArrayStringCSVEmptyResponse], errResp azfake.ErrorResponder) {
			require.NotNil(t, options)
			require.EqualValues(t, inputQuery, options.ArrayQuery)
			resp.SetResponse(http.StatusOK, urlgroup.QueriesClientArrayStringCSVEmptyResponse{}, nil)
			return
		},
	}
	client, err := urlgroup.NewQueriesClient(&azcore.ClientOptions{
		Transport: fake.NewQueriesServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.ArrayStringCSVEmpty(context.Background(), &urlgroup.QueriesClientArrayStringCSVEmptyOptions{
		ArrayQuery: inputQuery,
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFakeDateTimeNull(t *testing.T) {
	now := time.Now().UTC()
	server := fake.PathsServer{
		DateTimeNull: func(ctx context.Context, dateTimePath time.Time, options *urlgroup.PathsClientDateTimeNullOptions) (resp azfake.Responder[urlgroup.PathsClientDateTimeNullResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, now, dateTimePath)
			resp.SetResponse(http.StatusBadRequest, urlgroup.PathsClientDateTimeNullResponse{}, nil)
			return
		},
	}
	client, err := urlgroup.NewPathsClient(&azcore.ClientOptions{
		Transport: fake.NewPathsServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.DateTimeNull(context.Background(), now, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
