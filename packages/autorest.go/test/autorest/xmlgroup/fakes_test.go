// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup_test

import (
	"context"
	"generatortests"
	"generatortests/xmlgroup"
	"generatortests/xmlgroup/fake"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeGetACLs(t *testing.T) {
	value := []*xmlgroup.SignedIdentifier{
		{
			AccessPolicy: &xmlgroup.AccessPolicy{
				Expiry:     to.Ptr(time.Now().Add(time.Hour).UTC()),
				Permission: to.Ptr("rw"),
				Start:      to.Ptr(time.Now().UTC()),
			},
			ID: to.Ptr("entry 1"),
		},
		{
			AccessPolicy: &xmlgroup.AccessPolicy{
				Expiry:     to.Ptr(time.Now().Add(2 * time.Hour).UTC()),
				Permission: to.Ptr("ro"),
				Start:      to.Ptr(time.Now().UTC()),
			},
			ID: to.Ptr("entry 2"),
		},
	}
	server := fake.XMLServer{
		GetACLs: func(ctx context.Context, options *xmlgroup.XMLClientGetACLsOptions) (resp azfake.Responder[xmlgroup.XMLClientGetACLsResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, xmlgroup.XMLClientGetACLsResponse{
				SignedIdentifiers: value,
			}, nil)
			return
		},
	}
	client, err := xmlgroup.NewXMLClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewXMLServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.GetACLs(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, value, resp.SignedIdentifiers)
}
