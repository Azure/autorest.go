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
		optionalInt        = 1
		requiredIntList    = []int32{1, 2}
		optionalStringList = []string{"a", "b"}
	)
	server := fake.SpreadAliasServer{
		SpreadWithMultipleParameters: func(ctx context.Context, id string, xMSTestHeader string, requiredString string, optionalInt int32, requiredIntList []int32, optionalStringList []string, options *spreadgroup.SpreadAliasClientSpreadWithMultipleParametersOptions) (resp azfake.Responder[spreadgroup.SpreadAliasClientSpreadWithMultipleParametersResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, idVal, id)
			require.EqualValues(t, header, xMSTestHeader)
			require.EqualValues(t, requiredString, requiredString)
			require.EqualValues(t, optionalInt, optionalInt)
			require.EqualValues(t, requiredIntList, requiredIntList)
			require.EqualValues(t, optionalStringList, optionalStringList)
			resp.SetResponse(http.StatusNoContent, spreadgroup.SpreadAliasClientSpreadWithMultipleParametersResponse{}, nil)
			return
		},
	}
	client, err := spreadgroup.NewSpreadClient(&azcore.ClientOptions{
		Transport: fake.NewSpreadAliasServerTransport(&server),
	})
	require.NoError(t, err)

	_, err = client.NewSpreadAliasClient().SpreadWithMultipleParameters(context.Background(), idVal, header, requiredString, int32(optionalInt), requiredIntList, optionalStringList, nil)
	require.NoError(t, err)
}
