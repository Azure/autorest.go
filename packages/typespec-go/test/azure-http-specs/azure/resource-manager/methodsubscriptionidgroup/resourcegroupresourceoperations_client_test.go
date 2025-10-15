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

var validResource = methodsubscriptionidgroup.ResourceGroupResource{
	ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/test-rg/providers/Azure.ResourceManager.MethodSubscriptionId/resourceGroupResources/rg-resource", subscriptionIdExpected)),
	Location: to.Ptr("eastus"),
	Name:     to.Ptr("rg-resource"),
	Properties: &methodsubscriptionidgroup.ResourceGroupResourceProperties{
		ProvisioningState:    to.Ptr(methodsubscriptionidgroup.ResourceProvisioningState("Succeeded")),
		ResourceGroupSetting: to.Ptr(string("test-setting")),
	},
	SystemData: &methodsubscriptionidgroup.SystemData{
		CreatedAt:          to.Ptr(parseTime("2023-01-01T00:00:00.000Z")),
		CreatedBy:          to.Ptr(string("AzureSDK")),
		CreatedByType:      to.Ptr(methodsubscriptionidgroup.CreatedByType("User")),
		LastModifiedAt:     to.Ptr(parseTime("2023-01-01T00:00:00.000Z")),
		LastModifiedBy:     to.Ptr(string("AzureSDK")),
		LastModifiedByType: to.Ptr(methodsubscriptionidgroup.CreatedByType("User")),
	},
	Type: to.Ptr("Azure.ResourceManager.MethodSubscriptionId/resourceGroupResources"),
}

func TestResourceGroupResourceOperationsClient_Delete(t *testing.T) {
	delResp, err := clientFactory.NewResourceGroupResourceOperationsClient().Delete(context.Background(), "test-rg", "sub-resource-1", nil)
	require.NoError(t, err)
	require.Zero(t, delResp)
}

func TestResourceGroupResourceOperationsClient_Put(t *testing.T) {
	var reqResource = methodsubscriptionidgroup.ResourceGroupResource{
		Location: to.Ptr("eastus"),
		Properties: &methodsubscriptionidgroup.ResourceGroupResourceProperties{
			ResourceGroupSetting: to.Ptr(string("test-setting")),
		},
	}
	putResp, err := clientFactory.NewResourceGroupResourceOperationsClient().Put(context.Background(), "test-rg", "sub-resource-1", reqResource, nil)
	require.NoError(t, err)
	require.NotNil(t, putResp)
	require.Equal(t, validResource, putResp.ResourceGroupResource)
}

func TestResourceGroupResourceOperationsClient_Get(t *testing.T) {
	getResp, err := clientFactory.NewResourceGroupResourceOperationsClient().Get(context.Background(), "test-rg", "sub-resource-1", nil)
	require.NoError(t, err)
	require.NotNil(t, getResp)
	require.Equal(t, validResource, getResp.ResourceGroupResource)
}
