// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package resources_test

import (
	"net/http"
	"testing"
	"time"

	"resources"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestTopLevelTrackedResources(t *testing.T) {
	// TopLevelTrackedResources Get
	expectedTopLevelTrackedResourcesClientGetResponse := resources.TopLevelTrackedResourcesClientGetResponse{
		TopLevelTrackedResource: resources.TopLevelTrackedResource{
			ID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources/top"),
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
		},
	}
	topLevelTrackedResourcesClientGetResponse, err := clientFactory.NewTopLevelTrackedResourcesClient().Get(
		ctx,
		"test-rg",
		"top",
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientGetResponse.ID, *topLevelTrackedResourcesClientGetResponse.ID)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientGetResponse.Name, *topLevelTrackedResourcesClientGetResponse.Name)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientGetResponse.Type, *topLevelTrackedResourcesClientGetResponse.Type)

	// TopLevelTrackedResources CreateOrReplace
	expectedTopLevelTrackedResourcesClientCreateOrReplaceResponse := resources.TopLevelTrackedResourcesClientCreateOrReplaceResponse{
		TopLevelTrackedResource: resources.TopLevelTrackedResource{
			ID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources/top"),
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
		},
	}
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
	require.Equal(t, *expectedTopLevelTrackedResourcesClientCreateOrReplaceResponse.ID, *topLevelTrackedResourcesClientCreateOrReplaceResponse.ID)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientCreateOrReplaceResponse.Name, *topLevelTrackedResourcesClientCreateOrReplaceResponse.Name)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientCreateOrReplaceResponse.Type, *topLevelTrackedResourcesClientCreateOrReplaceResponse.Type)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientCreateOrReplaceResponse.Location, *topLevelTrackedResourcesClientCreateOrReplaceResponse.Location)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientCreateOrReplaceResponse.Properties.Description, *topLevelTrackedResourcesClientCreateOrReplaceResponse.Properties.Description)

	// TopLevelTrackedResources Update
	expectedTopLevelTrackedResourcesClientUpdateResponse := resources.TopLevelTrackedResourcesClientUpdateResponse{
		TopLevelTrackedResource: resources.TopLevelTrackedResource{
			ID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources/top"),
			Name:     to.Ptr("top"),
			Type:     to.Ptr("Azure.ResourceManager.Models.Resources/topLevelTrackedResources"),
			Location: to.Ptr("eastus"),
			Properties: &resources.TopLevelTrackedResourceProperties{
				Description:       to.Ptr("valid2"),
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
		},
	}
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
	require.Equal(t, *expectedTopLevelTrackedResourcesClientUpdateResponse.ID, *topLevelTrackedResourcesClientUpdateResponse.ID)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientUpdateResponse.Name, *topLevelTrackedResourcesClientUpdateResponse.Name)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientUpdateResponse.Type, *topLevelTrackedResourcesClientUpdateResponse.Type)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientUpdateResponse.Location, *topLevelTrackedResourcesClientUpdateResponse.Location)
	require.Equal(t, *expectedTopLevelTrackedResourcesClientUpdateResponse.Properties.Description, *topLevelTrackedResourcesClientUpdateResponse.Properties.Description)

	// TopLevelTrackedResources Delete
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

	// TopLevelTrackedResources ListByResourceGroup
	topLevelTrackedResourcesClientListByResourceGroupResponsePager := clientFactory.NewTopLevelTrackedResourcesClient().NewListByResourceGroupPager("test-rg", nil)
	require.True(t, topLevelTrackedResourcesClientListByResourceGroupResponsePager.More())
	topLevelTrackedResourcesClientListByResourceGroupResponse, err := topLevelTrackedResourcesClientListByResourceGroupResponsePager.NextPage(ctx)
	require.NoError(t, err)
	require.Len(t, topLevelTrackedResourcesClientListByResourceGroupResponse.Value, 1)

	// TopLevelTrackedResources List
	TopLevelTrackedResourcesClientListBySubscriptionResponsePager := clientFactory.NewTopLevelTrackedResourcesClient().NewListBySubscriptionPager(nil)
	require.True(t, TopLevelTrackedResourcesClientListBySubscriptionResponsePager.More())
	TopLevelTrackedResourcesClientListBySubscriptionResponse, err := TopLevelTrackedResourcesClientListBySubscriptionResponsePager.NextPage(ctx)
	require.NoError(t, err)
	require.Len(t, TopLevelTrackedResourcesClientListBySubscriptionResponse.Value, 1)
}
