//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armdataboxedge_test

import (
	"armdataboxedge/v2"
	"armdataboxedge/v2/fake"
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeAddonsClientBeginCreateOrUpdate(t *testing.T) {
	const theID = "the_id"
	server := fake.AddonsServer{
		BeginCreateOrUpdate: func(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, addon armdataboxedge.AddonClassification, options *armdataboxedge.AddonsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdataboxedge.AddonsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder) {
			result := armdataboxedge.AddonsClientCreateOrUpdateResponse{
				AddonClassification: &armdataboxedge.ArcAddon{
					ID: to.Ptr(theID),
				},
			}
			resp.SetTerminalResponse(http.StatusOK, result, nil)
			return
		},
	}
	client, err := armdataboxedge.NewAddonsClient("subID", azfake.NewTokenCredential(), &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewAddonsServerTransport(&server),
		},
	})
	require.NoError(t, err)
	poller, err := client.BeginCreateOrUpdate(context.Background(), "device", "role", "addon", "rg", &armdataboxedge.Addon{}, nil)
	require.NoError(t, err)
	resp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: time.Second,
	})
	require.NoError(t, err)
	addon, ok := resp.AddonClassification.(*armdataboxedge.ArcAddon)
	require.True(t, ok)
	require.NotNil(t, addon.ID)
	require.EqualValues(t, theID, *addon.ID)
}
