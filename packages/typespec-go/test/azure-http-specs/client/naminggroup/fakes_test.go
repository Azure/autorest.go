// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package naminggroup_test

import (
	"context"
	"naminggroup"
	"naminggroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeNamingServer(t *testing.T) {
	calledModelServerClient := false
	calledNamingServerClient := false
	server := fake.NamingServer{
		NamingModelServer: fake.NamingModelServer{
			Client: func(ctx context.Context, body naminggroup.ClientModel, options *naminggroup.NamingModelClientClientOptions) (resp azfake.Responder[naminggroup.NamingModelClientClientResponse], errResp azfake.ErrorResponder) {
				require.NotNil(t, body.DefaultName)
				require.True(t, *body.DefaultName)
				calledModelServerClient = true
				resp.SetResponse(http.StatusNoContent, naminggroup.NamingModelClientClientResponse{}, nil)
				return
			},
		},
		Client: func(ctx context.Context, body naminggroup.ClientNameModel, options *naminggroup.NamingClientClientOptions) (resp azfake.Responder[naminggroup.NamingClientClientResponse], errResp azfake.ErrorResponder) {
			require.NotNil(t, body.ClientName)
			require.True(t, *body.ClientName)
			calledNamingServerClient = true
			resp.SetResponse(http.StatusNoContent, naminggroup.NamingClientClientResponse{}, nil)
			return
		},
	}
	client, err := naminggroup.NewNamingClient("http://localhost:3000", &azcore.ClientOptions{
		Transport: fake.NewNamingServerTransport(&server),
	})
	require.NoError(t, err)

	_, err = client.NewNamingModelClient().Client(context.Background(), naminggroup.ClientModel{
		DefaultName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)

	_, err = client.Client(context.Background(), naminggroup.ClientNameModel{
		ClientName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)

	require.True(t, calledModelServerClient)
	require.True(t, calledNamingServerClient)
}
