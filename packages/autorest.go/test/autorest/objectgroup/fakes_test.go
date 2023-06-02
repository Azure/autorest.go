// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package objectgroup_test

import (
	"context"
	"generatortests/objectgroup"
	"generatortests/objectgroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/stretchr/testify/require"
)

func TestFakeGet(t *testing.T) {
	raw := []byte(`{"foo": "bar"}`)
	server := fake.ObjectTypeServer{
		Get: func(ctx context.Context, options *objectgroup.ObjectTypeClientGetOptions) (resp azfake.Responder[objectgroup.ObjectTypeClientGetResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, objectgroup.ObjectTypeClientGetResponse{RawJSON: raw}, nil)
			return
		},
	}
	client, err := objectgroup.NewObjectTypeClient(&azcore.ClientOptions{
		Transport: fake.NewObjectTypeServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, raw, resp.RawJSON)
}

func TestFakePut(t *testing.T) {
	raw := []byte(`{"foo": "bar"}`)
	server := fake.ObjectTypeServer{
		Put: func(ctx context.Context, putObject []byte, options *objectgroup.ObjectTypeClientPutOptions) (resp azfake.Responder[objectgroup.ObjectTypeClientPutResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, raw, putObject)
			resp.SetResponse(http.StatusOK, objectgroup.ObjectTypeClientPutResponse{}, nil)
			return
		},
	}
	client, err := objectgroup.NewObjectTypeClient(&azcore.ClientOptions{
		Transport: fake.NewObjectTypeServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.Put(context.Background(), raw, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
