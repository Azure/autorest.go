// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package methodsubscriptionidgroup_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"

	"methodsubscriptionidgroup"
)

var validMixedSubscriptionResource = methodsubscriptionidgroup.SubscriptionResource{
	ID:   to.Ptr(fmt.Sprintf("/subscriptions/%s/providers/Azure.ResourceManager.MethodSubscriptionId/subscriptionResources/sub-resource", subscriptionIdExpected)),
	Name: to.Ptr("sub-resource"),
	Properties: &methodsubscriptionidgroup.SubscriptionResourceProperties{
		ProvisioningState:   to.Ptr(methodsubscriptionidgroup.ResourceProvisioningState("Succeeded")),
		SubscriptionSetting: to.Ptr(string("test-sub-setting")),
	},
	SystemData: &methodsubscriptionidgroup.SystemData{
		CreatedAt:          to.Ptr(parseTime("2023-01-01T00:00:00.000Z")),
		CreatedBy:          to.Ptr(string("AzureSDK")),
		CreatedByType:      to.Ptr(methodsubscriptionidgroup.CreatedByType("User")),
		LastModifiedAt:     to.Ptr(parseTime("2023-01-01T00:00:00.000Z")),
		LastModifiedBy:     to.Ptr(string("AzureSDK")),
		LastModifiedByType: to.Ptr(methodsubscriptionidgroup.CreatedByType("User")),
	},
	Type: to.Ptr("Azure.ResourceManager.MethodSubscriptionId/subscriptionResources"),
}

func TestSubscriptionResourceOperationsClient_Delete(t *testing.T) {
	delResp, err := clientFactory.NewSubscriptionResourceOperationsClient().Delete(context.Background(), subscriptionIdExpected, "sub-resource", nil)
	require.NoError(t, err)
	require.Zero(t, delResp)
}

func TestSubscriptionResourceOperationsClient_Put(t *testing.T) {
	var validResource = methodsubscriptionidgroup.SubscriptionResource{
		Properties: &methodsubscriptionidgroup.SubscriptionResourceProperties{
			SubscriptionSetting: to.Ptr(string("test-sub-setting")),
		},
	}
	putResp, err := clientFactory.NewSubscriptionResourceOperationsClient().Put(context.Background(), subscriptionIdExpected, "sub-resource", validResource, nil)
	require.NoError(t, err)
	require.NotNil(t, putResp)
	require.Equal(t, validMixedSubscriptionResource, putResp.SubscriptionResource)
}

func TestSubscriptionResourceOperationsClient_Get(t *testing.T) {
	getResp, err := clientFactory.NewSubscriptionResourceOperationsClient().Get(context.Background(), subscriptionIdExpected, "sub-resource", nil)
	require.NoError(t, err)
	require.NotNil(t, getResp)
	require.Equal(t, validMixedSubscriptionResource, getResp.SubscriptionResource)
}
