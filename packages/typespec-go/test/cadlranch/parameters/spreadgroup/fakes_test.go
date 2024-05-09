// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package spreadgroup_test

import (
	"context"
	"net/http"
	"spreadgroup"
	"spreadgroup/fake"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/stretchr/testify/require"
)

func TestFake_SpreadWithMultipleParameters(t *testing.T) {
	const (
		idVal    = "id"
		header   = "header"
		prop1Val = "prop1"
		prop2Val = "prop2"
		prop3Val = "prop3"
		prop4Val = "prop4"
		prop5Val = "prop5"
		prop6Val = "prop6"
	)
	server := fake.SpreadAliasServer{
		SpreadWithMultipleParameters: func(ctx context.Context, id, xMSTestHeader, prop1, prop2, prop3, prop4, prop5, prop6 string, options *spreadgroup.SpreadAliasClientSpreadWithMultipleParametersOptions) (resp azfake.Responder[spreadgroup.SpreadAliasClientSpreadWithMultipleParametersResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, idVal, id)
			require.EqualValues(t, header, xMSTestHeader)
			require.EqualValues(t, prop1Val, prop1)
			require.EqualValues(t, prop2Val, prop2)
			require.EqualValues(t, prop3Val, prop3)
			require.EqualValues(t, prop4Val, prop4)
			require.EqualValues(t, prop5Val, prop5)
			require.EqualValues(t, prop6Val, prop6)
			resp.SetResponse(http.StatusNoContent, spreadgroup.SpreadAliasClientSpreadWithMultipleParametersResponse{}, nil)
			return
		},
	}
	client, err := spreadgroup.NewSpreadClient(&azcore.ClientOptions{
		Transport: fake.NewSpreadAliasServerTransport(&server),
	})
	require.NoError(t, err)

	_, err = client.NewSpreadAliasClient().SpreadWithMultipleParameters(context.Background(), idVal, header, prop1Val, prop2Val, prop3Val, prop4Val, prop5Val, prop6Val, nil)
	require.NoError(t, err)
}
