// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package arraygroup_test

import (
	"context"
	"generatortests"
	"generatortests/arraygroup"
	"generatortests/arraygroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestFakeGetArrayValid(t *testing.T) {
	server := fake.ArrayServer{
		GetArrayValid: func(ctx context.Context, options *arraygroup.ArrayClientGetArrayValidOptions) (resp azfake.Responder[arraygroup.ArrayClientGetArrayValidResponse], errResp azfake.ErrorResponder) {
			require.Nil(t, options)
			resp.SetResponse(http.StatusOK, arraygroup.ArrayClientGetArrayValidResponse{
				StringArrayArray: [][]*string{
					to.SliceOfPtrs("7", "8", "9"),
					to.SliceOfPtrs("4", "5", "6"),
					to.SliceOfPtrs("1", "2", "3"),
				},
			}, nil)
			return
		},
	}
	client, err := arraygroup.NewArrayClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewArrayServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.GetArrayValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.StringArrayArray, [][]*string{
		to.SliceOfPtrs("7", "8", "9"),
		to.SliceOfPtrs("4", "5", "6"),
		to.SliceOfPtrs("1", "2", "3"),
	}); r != "" {
		t.Fatal(r)
	}
}
