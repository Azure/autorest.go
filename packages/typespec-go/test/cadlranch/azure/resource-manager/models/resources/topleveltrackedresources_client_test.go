// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package resources_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"resources"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var (
	validTopLevelResource = resources.TopLevelTrackedResource{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources/top", subscriptionIdExpected, resourceGroupExpected)),
		Name:     to.Ptr("top"),
		Type:     to.Ptr("Azure.ResourceManager.Models.Resources/topLevelTrackedResources"),
		Location: to.Ptr("eastus"),
		Properties: &resources.TopLevelTrackedResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		SystemData: &resources.SystemData{
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			CreatedAt:          to.Ptr(time.Now()),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedAt:     to.Ptr(time.Now()),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
	}
)

func TestTopLevelTrackedResourcesClient_Get(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	topLevelTrackedResourcesClientGetResponse, err := clientFactory.NewTopLevelTrackedResourcesClient().Get(
		ctx,
		"test-rg",
		"top",
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *validTopLevelResource.ID, *topLevelTrackedResourcesClientGetResponse.ID)
	require.Equal(t, *validTopLevelResource.Name, *topLevelTrackedResourcesClientGetResponse.Name)
	require.Equal(t, *validTopLevelResource.Type, *topLevelTrackedResourcesClientGetResponse.Type)
}

func TestTopLevelTrackedResourcesClient_CreateOrReplace(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	topLevelTrackedResourcesClientCreateOrReplaceResponsePoller, err := clientFactory.NewTopLevelTrackedResourcesClient().BeginCreateOrReplace(
		ctx,
		"test-rg",
		"top",
		resources.TopLevelTrackedResource{
			Location: to.Ptr("eastus"),
			Properties: &resources.TopLevelTrackedResourceProperties{
				Description: to.Ptr("valid"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	topLevelTrackedResourcesClientCreateOrReplaceResponse, err := topLevelTrackedResourcesClientCreateOrReplaceResponsePoller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validTopLevelResource.ID, *topLevelTrackedResourcesClientCreateOrReplaceResponse.ID)
	require.Equal(t, *validTopLevelResource.Name, *topLevelTrackedResourcesClientCreateOrReplaceResponse.Name)
	require.Equal(t, *validTopLevelResource.Type, *topLevelTrackedResourcesClientCreateOrReplaceResponse.Type)
	require.Equal(t, *validTopLevelResource.Location, *topLevelTrackedResourcesClientCreateOrReplaceResponse.Location)
	require.Equal(t, *validTopLevelResource.Properties.Description, *topLevelTrackedResourcesClientCreateOrReplaceResponse.Properties.Description)
}

func TestTopLevelTrackedResourcesClient_BeginUpdate(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	topLevelTrackedResourcesClientUpdateResponsePoller, err := clientFactory.NewTopLevelTrackedResourcesClient().BeginUpdate(
		ctx,
		"test-rg",
		"top",
		resources.TopLevelTrackedResource{
			Properties: &resources.TopLevelTrackedResourceProperties{
				Description: to.Ptr("valid2"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	topLevelTrackedResourcesClientUpdateResponse, err := topLevelTrackedResourcesClientUpdateResponsePoller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validTopLevelResource.ID, *topLevelTrackedResourcesClientUpdateResponse.ID)
	require.Equal(t, *validTopLevelResource.Name, *topLevelTrackedResourcesClientUpdateResponse.Name)
	require.Equal(t, *validTopLevelResource.Type, *topLevelTrackedResourcesClientUpdateResponse.Type)
	require.Equal(t, *validTopLevelResource.Location, *topLevelTrackedResourcesClientUpdateResponse.Location)
	require.Equal(t, "valid2", *topLevelTrackedResourcesClientUpdateResponse.Properties.Description)
}

func TestTopLevelTrackedResourcesClient_BeginDelete(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	topLevelTrackedResourcesClientDeleteResponsePoller, err := clientFactory.NewTopLevelTrackedResourcesClient().BeginDelete(
		ctx,
		"test-rg",
		"top",
		nil,
	)
	require.NoError(t, err)
	topLevelTrackedResourcesClientDeleteResponse, err := topLevelTrackedResourcesClientDeleteResponsePoller.Poll(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, topLevelTrackedResourcesClientDeleteResponse.StatusCode)
}

func TestTopLevelTrackedResourcesClient_NewListByResourceGroupPager(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	topLevelTrackedResourcesClientListByResourceGroupResponsePager := clientFactory.NewTopLevelTrackedResourcesClient().NewListByResourceGroupPager("test-rg", nil)
	require.True(t, topLevelTrackedResourcesClientListByResourceGroupResponsePager.More())
	topLevelTrackedResourcesClientListByResourceGroupResponse, err := topLevelTrackedResourcesClientListByResourceGroupResponsePager.NextPage(ctx)
	require.NoError(t, err)
	require.Len(t, topLevelTrackedResourcesClientListByResourceGroupResponse.Value, 1)
	require.Equal(t, *validTopLevelResource.ID, *topLevelTrackedResourcesClientListByResourceGroupResponse.Value[0].ID)
	require.Equal(t, *validTopLevelResource.Name, *topLevelTrackedResourcesClientListByResourceGroupResponse.Value[0].Name)
	require.Equal(t, *validTopLevelResource.Type, *topLevelTrackedResourcesClientListByResourceGroupResponse.Value[0].Type)
	require.Equal(t, *validTopLevelResource.Location, *topLevelTrackedResourcesClientListByResourceGroupResponse.Value[0].Location)
}

func TestTopLevelTrackedResourcesClient_NewListBySubscriptionPager(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	TopLevelTrackedResourcesClientListBySubscriptionResponsePager := clientFactory.NewTopLevelTrackedResourcesClient().NewListBySubscriptionPager(nil)
	require.True(t, TopLevelTrackedResourcesClientListBySubscriptionResponsePager.More())
	TopLevelTrackedResourcesClientListBySubscriptionResponse, err := TopLevelTrackedResourcesClientListBySubscriptionResponsePager.NextPage(ctx)
	require.NoError(t, err)
	require.Len(t, TopLevelTrackedResourcesClientListBySubscriptionResponse.Value, 1)
	require.Equal(t, *validTopLevelResource.ID, *TopLevelTrackedResourcesClientListBySubscriptionResponse.Value[0].ID)
	require.Equal(t, *validTopLevelResource.Name, *TopLevelTrackedResourcesClientListBySubscriptionResponse.Value[0].Name)
	require.Equal(t, *validTopLevelResource.Type, *TopLevelTrackedResourcesClientListBySubscriptionResponse.Value[0].Type)
	require.Equal(t, *validTopLevelResource.Location, *TopLevelTrackedResourcesClientListBySubscriptionResponse.Value[0].Location)
}

func TestTopLevelTrackedResourcesClient_ActionSync(t *testing.T) {
	_, err := clientFactory.NewTopLevelTrackedResourcesClient().ActionSync(
		ctx,
		resourceGroupExpected,
		"top",
		resources.NotificationDetails{
			Message: to.Ptr("Resource action at top level."),
			Urgent:  to.Ptr(true),
		},
		nil,
	)
	require.NoError(t, err)
}
