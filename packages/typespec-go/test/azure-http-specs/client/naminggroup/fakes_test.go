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
	calledPropertyServerClient := false
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
		NamingPropertyServer: fake.NamingPropertyServer{
			Client: func(ctx context.Context, body naminggroup.ClientNameModel, options *naminggroup.NamingPropertyClientClientOptions) (resp azfake.Responder[naminggroup.NamingPropertyClientClientResponse], errResp azfake.ErrorResponder) {
				require.NotNil(t, body.ClientName)
				require.True(t, *body.ClientName)
				calledPropertyServerClient = true
				resp.SetResponse(http.StatusNoContent, naminggroup.NamingPropertyClientClientResponse{}, nil)
				return
			},
		},
	}
	client, err := naminggroup.NewNamingClientWithNoCredential("http://localhost:3000", &naminggroup.NamingClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewNamingServerTransport(&server),
		},
	})
	require.NoError(t, err)

	_, err = client.NewNamingModelClient().Client(context.Background(), naminggroup.ClientModel{
		DefaultName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)

	_, err = client.NewNamingPropertyClient().Client(context.Background(), naminggroup.ClientNameModel{
		ClientName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)

	require.True(t, calledModelServerClient)
	require.True(t, calledPropertyServerClient)
}
