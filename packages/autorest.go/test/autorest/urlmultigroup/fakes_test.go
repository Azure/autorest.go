// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlmultigroup_test

import (
	"context"
	"generatortests/urlmultigroup"
	"generatortests/urlmultigroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/stretchr/testify/require"
)

func TestFakeArrayStringMultiValid(t *testing.T) {
	content := []string{"one", "two", "three"}
	server := fake.QueriesServer{
		ArrayStringMultiValid: func(ctx context.Context, options *urlmultigroup.QueriesClientArrayStringMultiValidOptions) (resp azfake.Responder[urlmultigroup.QueriesClientArrayStringMultiValidResponse], errResp azfake.ErrorResponder) {
			require.NotNil(t, options)
			require.EqualValues(t, content, options.ArrayQuery)
			resp.SetResponse(http.StatusOK, urlmultigroup.QueriesClientArrayStringMultiValidResponse{}, nil)
			return
		},
	}
	client, err := urlmultigroup.NewQueriesClient(&azcore.ClientOptions{
		Transport: fake.NewQueriesServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.ArrayStringMultiValid(context.Background(), &urlmultigroup.QueriesClientArrayStringMultiValidOptions{
		ArrayQuery: content,
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}
