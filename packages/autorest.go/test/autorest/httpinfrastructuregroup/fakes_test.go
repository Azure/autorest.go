// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup_test

import (
	"context"
	"generatortests"
	"generatortests/httpinfrastructuregroup"
	"generatortests/httpinfrastructuregroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeGet200Model201ModelDefaultError200Valid(t *testing.T) {
	server := fake.MultipleResponsesServer{
		Get200Model201ModelDefaultError200Valid: func(ctx context.Context, options *httpinfrastructuregroup.MultipleResponsesClientGet200Model201ModelDefaultError200ValidOptions) (resp azfake.Responder[httpinfrastructuregroup.MultipleResponsesClientGet200Model201ModelDefaultError200ValidResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusCreated, httpinfrastructuregroup.MultipleResponsesClientGet200Model201ModelDefaultError200ValidResponse{
				Value: httpinfrastructuregroup.B{
					StatusCode: to.Ptr("status_code"),
				},
			}, nil)
			return
		},
	}
	client, err := httpinfrastructuregroup.NewMultipleResponsesClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewMultipleResponsesServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.Get200Model201ModelDefaultError200Valid(context.Background(), nil)
	require.NoError(t, err)
	b, ok := resp.Value.(httpinfrastructuregroup.B)
	require.True(t, ok)
	require.NotNil(t, b.StatusCode)
	require.EqualValues(t, "status_code", *b.StatusCode)
}
