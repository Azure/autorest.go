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

var validResource2 = methodsubscriptionidgroup.SubscriptionResource2{
	ID:   to.Ptr(fmt.Sprintf(`/subscriptions/%s/providers/Azure.ResourceManager.MethodSubscriptionId/subscriptionResource2s/sub-resource-2`, subscriptionIdExpected)),
	Name: to.Ptr("sub-resource-2"),
	Type: to.Ptr("Azure.ResourceManager.MethodSubscriptionId/subscriptionResource2s"),
	Properties: &methodsubscriptionidgroup.SubscriptionResource2Properties{
		ProvisioningState: to.Ptr(methodsubscriptionidgroup.ResourceProvisioningState("Succeeded")),
		ConfigValue:       to.Ptr(string("test-config")),
	},
	SystemData: &methodsubscriptionidgroup.SystemData{
		CreatedBy:          to.Ptr(string("AzureSDK")),
		CreatedByType:      to.Ptr(methodsubscriptionidgroup.CreatedByType("User")),
		CreatedAt:          to.Ptr(parseTime("2023-01-01T00:00:00.000Z")),
		LastModifiedBy:     to.Ptr(string("AzureSDK")),
		LastModifiedAt:     to.Ptr(parseTime("2023-01-01T00:00:00.000Z")),
		LastModifiedByType: to.Ptr(methodsubscriptionidgroup.CreatedByType("User")),
	},
}

func TestSubscriptionResource2OperationsClient_Delete(t *testing.T) {
	delResp, err := clientFactory.NewSubscriptionResource2OperationsClient().Delete(context.Background(), subscriptionIdExpected, "sub-resource-1", nil)
	require.NoError(t, err)
	require.Zero(t, delResp)
}

func TestSubscriptionResource2OperationsClient_Put(t *testing.T) {
	var validResource = methodsubscriptionidgroup.SubscriptionResource2{
		Properties: &methodsubscriptionidgroup.SubscriptionResource2Properties{
			ConfigValue: to.Ptr(string("test-config")),
		},
	}
	putResp, err := clientFactory.NewSubscriptionResource2OperationsClient().Put(context.Background(), subscriptionIdExpected, "sub-resource-1", validResource, nil)
	require.NoError(t, err)
	require.NotNil(t, putResp)
	require.Equal(t, validResource2, putResp.SubscriptionResource2)
}

func TestSubscriptionResource2OperationsClient_Get(t *testing.T) {
	getResp, err := clientFactory.NewSubscriptionResource2OperationsClient().Get(context.Background(), subscriptionIdExpected, "sub-resource-1", nil)
	require.NoError(t, err)
	require.NotNil(t, getResp)
	require.Equal(t, validResource2, getResp.SubscriptionResource2)
}
