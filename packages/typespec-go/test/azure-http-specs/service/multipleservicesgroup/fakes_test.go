// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipleservicesgroup_test

import (
	"context"
	"multipleservicesgroup"
	"multipleservicesgroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/stretchr/testify/require"
)

func TestFakeServiceAServer(t *testing.T) {
	calledOpA := false
	calledSubOpA := false
	server := fake.ServiceAServer{
		ServiceAOperationsServer: fake.ServiceAOperationsServer{
			OpA: func(ctx context.Context, options *multipleservicesgroup.ServiceAOperationsClientOpAOptions) (resp azfake.Responder[multipleservicesgroup.ServiceAOperationsClientOpAResponse], errResp azfake.ErrorResponder) {
				calledOpA = true
				resp.SetResponse(http.StatusNoContent, multipleservicesgroup.ServiceAOperationsClientOpAResponse{}, nil)
				return
			},
		},
		ServiceASubNamespaceServer: fake.ServiceASubNamespaceServer{
			SubOpA: func(ctx context.Context, options *multipleservicesgroup.ServiceASubNamespaceClientSubOpAOptions) (resp azfake.Responder[multipleservicesgroup.ServiceASubNamespaceClientSubOpAResponse], errResp azfake.ErrorResponder) {
				calledSubOpA = true
				resp.SetResponse(http.StatusNoContent, multipleservicesgroup.ServiceASubNamespaceClientSubOpAResponse{}, nil)
				return
			},
		},
	}
	client, err := multipleservicesgroup.NewServiceAClientWithNoCredential("http://localhost:3000", &multipleservicesgroup.ServiceAClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewServiceAServerTransport(&server),
		},
	})
	require.NoError(t, err)

	_, err = client.NewServiceAOperationsClient().OpA(context.Background(), nil)
	require.NoError(t, err)

	_, err = client.NewServiceASubNamespaceClient().SubOpA(context.Background(), nil)
	require.NoError(t, err)

	require.True(t, calledOpA)
	require.True(t, calledSubOpA)
}

func TestFakeServiceBServer(t *testing.T) {
	calledOpB := false
	calledSubOpB := false
	server := fake.ServiceBServer{
		ServiceBOperationsServer: fake.ServiceBOperationsServer{
			OpB: func(ctx context.Context, options *multipleservicesgroup.ServiceBOperationsClientOpBOptions) (resp azfake.Responder[multipleservicesgroup.ServiceBOperationsClientOpBResponse], errResp azfake.ErrorResponder) {
				calledOpB = true
				resp.SetResponse(http.StatusNoContent, multipleservicesgroup.ServiceBOperationsClientOpBResponse{}, nil)
				return
			},
		},
		ServiceBSubNamespaceServer: fake.ServiceBSubNamespaceServer{
			SubOpB: func(ctx context.Context, options *multipleservicesgroup.ServiceBSubNamespaceClientSubOpBOptions) (resp azfake.Responder[multipleservicesgroup.ServiceBSubNamespaceClientSubOpBResponse], errResp azfake.ErrorResponder) {
				calledSubOpB = true
				resp.SetResponse(http.StatusNoContent, multipleservicesgroup.ServiceBSubNamespaceClientSubOpBResponse{}, nil)
				return
			},
		},
	}
	client, err := multipleservicesgroup.NewServiceBClientWithNoCredential("http://localhost:3000", &multipleservicesgroup.ServiceBClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewServiceBServerTransport(&server),
		},
	})
	require.NoError(t, err)

	_, err = client.NewServiceBOperationsClient().OpB(context.Background(), nil)
	require.NoError(t, err)

	_, err = client.NewServiceBSubNamespaceClient().SubOpB(context.Background(), nil)
	require.NoError(t, err)

	require.True(t, calledOpB)
	require.True(t, calledSubOpB)
}
