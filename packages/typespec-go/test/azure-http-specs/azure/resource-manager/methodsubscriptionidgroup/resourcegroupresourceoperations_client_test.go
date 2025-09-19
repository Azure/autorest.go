// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package methodsubscriptionidgroup_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/stretchr/testify/require"

	"methodsubscriptionidgroup"
)

// parseTime parses an RFC3339 time string and panics if parsing fails.
func parseTime(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		log.Fatalf("failed to parse time: %v", err)
	}
	return t
}

var validResource1 = methodsubscriptionidgroup.ResourceGroupResource{
	ID:   to.Ptr(`/subscriptions/${SUBSCRIPTION_ID_EXPECTED}/providers/Azure.ResourceManager.MethodSubscriptionId/subscriptionResource1s/sub-resource-1`),
	Name: to.Ptr("sub-resource-1"),
	Type: to.Ptr("Azure.ResourceManager.MethodSubscriptionId/subscriptionResource1s"),
	Properties: &methodsubscriptionidgroup.ResourceGroupResourceProperties{
		ProvisioningState:    to.Ptr(methodsubscriptionidgroup.ResourceProvisioningState("Succeeded")),
		ResourceGroupSetting: to.Ptr(string("Valid subscription resource 1")),
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

var validResource2 = methodsubscriptionidgroup.ResourceGroupResource{
	ID:   to.Ptr(`/subscriptions/${SUBSCRIPTION_ID_EXPECTED}/providers/Azure.ResourceManager.MethodSubscriptionId/subscriptionResource1s/sub-resource-2`),
	Name: to.Ptr("sub-resource-1"),
	Type: to.Ptr("Azure.ResourceManager.MethodSubscriptionId/subscriptionResource2s"),
	Properties: &methodsubscriptionidgroup.ResourceGroupResourceProperties{
		ProvisioningState:    to.Ptr(methodsubscriptionidgroup.ResourceProvisioningState("Succeeded")),
		ResourceGroupSetting: to.Ptr(string("Valid subscription resource 1")),
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

func TestResourceGroupResourceOperationsClient_Delete(t *testing.T) {
	subscriptionID := "00000000-0000-0000-0000-000000000000"
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err)

	client, err := methodsubscriptionidgroup.NewResourceGroupResourceOperationsClient(subscriptionID, cred, nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	resourceGroupName := "test-rg"
	resourceName := "sub-resource-1"

	// Create (Put)
	delResp, err := client.Delete(context.Background(), resourceGroupName, resourceName, &methodsubscriptionidgroup.ResourceGroupResourceOperationsClientDeleteOptions{})
	require.NoError(t, err)
	require.NotNil(t, delResp)
}

func TestResourceGroupResourceOperationsClient_Put(t *testing.T) {
	subscriptionID := "00000000-0000-0000-0000-000000000000"
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err)

	client, err := methodsubscriptionidgroup.NewResourceGroupResourceOperationsClient(subscriptionID, cred, nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	resourceGroupName := "test-rg"
	resourceName := "sub-resource-1"
	resource := methodsubscriptionidgroup.ResourceGroupResource{
		// Fill with required fields for creation
	}

	// Create (Put)
	putResp, err := client.Put(context.Background(), resourceGroupName, resourceName, resource, nil)
	require.NoError(t, err)
	require.NotNil(t, putResp.ResourceGroupResource)
}
func TestResourceGroupResourceOperationsClient_CRUD(t *testing.T) {
	subscriptionID := "00000000-0000-0000-0000-000000000000"
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err)

	client, err := methodsubscriptionidgroup.NewResourceGroupResourceOperationsClient(subscriptionID, cred, nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	resourceGroupName := "test-rg"
	resourceName := "sub-resource-1"
	resource := methodsubscriptionidgroup.ResourceGroupResource{
		// Fill with required fields for creation
	}

	// Create (Put)
	putResp, err := client.Put(context.Background(), resourceGroupName, resourceName, resource, nil)
	require.NoError(t, err)
	require.NotNil(t, putResp.ResourceGroupResource)

	// Get
	getResp, err := client.Get(context.Background(), resourceGroupName, resourceName, nil)
	require.NoError(t, err)
	require.NotNil(t, getResp.ResourceGroupResource)

	// Delete
	delResp, err := client.Delete(context.Background(), resourceGroupName, resourceName, nil)
	require.NoError(t, err)
	require.NotNil(t, delResp)
}
