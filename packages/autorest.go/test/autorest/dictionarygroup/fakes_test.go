// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package dictionarygroup_test

import (
	"context"
	"generatortests"
	"generatortests/dictionarygroup"
	"generatortests/dictionarygroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeGetArrayValid(t *testing.T) {
	server := fake.DictionaryServer{
		GetArrayValid: func(ctx context.Context, options *dictionarygroup.DictionaryClientGetArrayValidOptions) (resp azfake.Responder[dictionarygroup.DictionaryClientGetArrayValidResponse], errResp azfake.ErrorResponder) {
			require.Nil(t, options)
			resp.SetResponse(http.StatusOK, dictionarygroup.DictionaryClientGetArrayValidResponse{
				Value: map[string][]*string{
					"key": {
						to.Ptr("one"),
						to.Ptr("two"),
					},
				},
			}, nil)
			return
		},
	}
	client, err := dictionarygroup.NewDictionaryClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewDictionaryServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.GetArrayValid(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string][]*string{
		"key": {
			to.Ptr("one"),
			to.Ptr("two"),
		},
	}, resp.Value)
}

func TestFakePutDictionaryValid(t *testing.T) {
	server := fake.DictionaryServer{
		PutDictionaryValid: func(ctx context.Context, arrayBody map[string]map[string]*string, options *dictionarygroup.DictionaryClientPutDictionaryValidOptions) (resp azfake.Responder[dictionarygroup.DictionaryClientPutDictionaryValidResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, map[string]map[string]*string{
				"outer": {
					"inner": to.Ptr("foo"),
				},
			}, arrayBody)
			resp.SetResponse(http.StatusOK, dictionarygroup.DictionaryClientPutDictionaryValidResponse{}, nil)
			return
		},
	}
	client, err := dictionarygroup.NewDictionaryClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewDictionaryServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PutDictionaryValid(context.Background(), map[string]map[string]*string{
		"outer": {
			"inner": to.Ptr("foo"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
