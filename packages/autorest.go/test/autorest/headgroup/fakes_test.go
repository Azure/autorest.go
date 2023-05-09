// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headgroup_test

import (
	"context"
	"generatortests/headgroup"
	"generatortests/headgroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/stretchr/testify/require"
)

func TestFakeHead204(t *testing.T) {
	server := fake.HTTPSuccessServer{
		Head204: func(ctx context.Context, options *headgroup.HTTPSuccessClientHead204Options) (resp azfake.Responder[headgroup.HTTPSuccessClientHead204Response], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusNoContent, headgroup.HTTPSuccessClientHead204Response{
				Success: true,
			}, nil)
			return
		},
	}
	client, err := headgroup.NewHTTPSuccessClient(&azcore.ClientOptions{
		Transport: fake.NewHTTPSuccessServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.Head204(context.Background(), nil)
	require.NoError(t, err)
	require.True(t, resp.Success)
}
