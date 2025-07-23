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
	var (
		idVal              = "id"
		header             = "header"
		requiredString     = "prop1"
		optionalInt        = int32(1)
		requiredIntList    = []int32{1, 2}
		optionalStringList = []string{"a", "b"}
	)
	server := fake.SpreadAliasServer{
		SpreadWithMultipleParameters: func(ctx context.Context, id string, xMSTestHeader string, requiredString string, requiredIntList []int32, options *spreadgroup.SpreadAliasClientSpreadWithMultipleParametersOptions) (resp azfake.Responder[spreadgroup.SpreadAliasClientSpreadWithMultipleParametersResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, idVal, id)
			require.EqualValues(t, header, xMSTestHeader)
			require.EqualValues(t, requiredString, requiredString)
			require.EqualValues(t, requiredIntList, requiredIntList)
			require.NotNil(t, options)
			require.NotNil(t, options.OptionalInt)
			require.EqualValues(t, optionalInt, *options.OptionalInt)
			require.EqualValues(t, optionalStringList, options.OptionalStringList)
			resp.SetResponse(http.StatusNoContent, spreadgroup.SpreadAliasClientSpreadWithMultipleParametersResponse{}, nil)
			return
		},
	}
	client, err := spreadgroup.NewSpreadClient("http://localhost:3000", &azcore.ClientOptions{
		Transport: fake.NewSpreadAliasServerTransport(&server),
	})
	require.NoError(t, err)

	_, err = client.NewSpreadAliasClient().SpreadWithMultipleParameters(context.Background(), idVal, header, requiredString, requiredIntList, &spreadgroup.SpreadAliasClientSpreadWithMultipleParametersOptions{
		OptionalInt:        &optionalInt,
		OptionalStringList: optionalStringList,
	})
	require.NoError(t, err)
}
