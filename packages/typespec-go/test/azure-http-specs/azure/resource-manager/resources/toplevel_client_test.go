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
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top", subscriptionIdExpected, resourceGroupExpected)),
		Name:     to.Ptr("top"),
		Type:     to.Ptr("Azure.ResourceManager.Resources/topLevelTrackedResources"),
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

func TestTopLevelClient_Get(t *testing.T) {
	topLevelClientGetResponse, err := clientFactory.NewTopLevelClient().Get(
		ctx,
		"test-rg",
		"top",
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *validTopLevelResource.ID, *topLevelClientGetResponse.ID)
	require.Equal(t, *validTopLevelResource.Name, *topLevelClientGetResponse.Name)
	require.Equal(t, *validTopLevelResource.Type, *topLevelClientGetResponse.Type)
	require.Equal(t, *validTopLevelResource.Location, *topLevelClientGetResponse.Location)
	require.Equal(t, *validTopLevelResource.Properties.Description, *topLevelClientGetResponse.Properties.Description)
	require.Equal(t, *validTopLevelResource.Properties.ProvisioningState, *topLevelClientGetResponse.Properties.ProvisioningState)
}

func TestTopLevelClient_CreateOrReplace(t *testing.T) {
	topLevelClientCreateOrReplaceResponsePoller, err := clientFactory.NewTopLevelClient().BeginCreateOrReplace(
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
	topLevelClientCreateOrReplaceResponse, err := topLevelClientCreateOrReplaceResponsePoller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validTopLevelResource.ID, *topLevelClientCreateOrReplaceResponse.ID)
	require.Equal(t, *validTopLevelResource.Name, *topLevelClientCreateOrReplaceResponse.Name)
	require.Equal(t, *validTopLevelResource.Type, *topLevelClientCreateOrReplaceResponse.Type)
	require.Equal(t, *validTopLevelResource.Location, *topLevelClientCreateOrReplaceResponse.Location)
	require.Equal(t, *validTopLevelResource.Properties.Description, *topLevelClientCreateOrReplaceResponse.Properties.Description)
	require.Equal(t, *validTopLevelResource.Properties.ProvisioningState, *topLevelClientCreateOrReplaceResponse.Properties.ProvisioningState)
}

func TestTopLevelClient_BeginUpdate(t *testing.T) {
	topLevelClientUpdateResponsePoller, err := clientFactory.NewTopLevelClient().BeginUpdate(
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
	topLevelClientUpdateResponse, err := topLevelClientUpdateResponsePoller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validTopLevelResource.ID, *topLevelClientUpdateResponse.ID)
	require.Equal(t, *validTopLevelResource.Name, *topLevelClientUpdateResponse.Name)
	require.Equal(t, *validTopLevelResource.Type, *topLevelClientUpdateResponse.Type)
	require.Equal(t, *validTopLevelResource.Location, *topLevelClientUpdateResponse.Location)
	require.Equal(t, "valid2", *topLevelClientUpdateResponse.Properties.Description)
	require.Equal(t, *validTopLevelResource.Properties.ProvisioningState, *topLevelClientUpdateResponse.Properties.ProvisioningState)
}

func TestTopLevelClient_BeginDelete(t *testing.T) {
	topLevelClientDeleteResponsePoller, err := clientFactory.NewTopLevelClient().BeginDelete(
		ctx,
		"test-rg",
		"top",
		nil,
	)
	require.NoError(t, err)
	topLevelClientDeleteResponse, err := topLevelClientDeleteResponsePoller.Poll(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, topLevelClientDeleteResponse.StatusCode)
}

func TestTopLevelClient_NewListByResourceGroupPager(t *testing.T) {
	topLevelClientListByResourceGroupResponsePager := clientFactory.NewTopLevelClient().NewListByResourceGroupPager("test-rg", nil)
	require.True(t, topLevelClientListByResourceGroupResponsePager.More())
	topLevelClientListByResourceGroupResponse, err := topLevelClientListByResourceGroupResponsePager.NextPage(ctx)
	require.NoError(t, err)
	require.Len(t, topLevelClientListByResourceGroupResponse.Value, 1)
	require.Equal(t, *validTopLevelResource.ID, *topLevelClientListByResourceGroupResponse.Value[0].ID)
	require.Equal(t, *validTopLevelResource.Name, *topLevelClientListByResourceGroupResponse.Value[0].Name)
	require.Equal(t, *validTopLevelResource.Type, *topLevelClientListByResourceGroupResponse.Value[0].Type)
	require.Equal(t, *validTopLevelResource.Location, *topLevelClientListByResourceGroupResponse.Value[0].Location)
	require.Equal(t, *validTopLevelResource.Properties.Description, *topLevelClientListByResourceGroupResponse.Value[0].Properties.Description)
	require.Equal(t, *validTopLevelResource.Properties.ProvisioningState, *topLevelClientListByResourceGroupResponse.Value[0].Properties.ProvisioningState)
}

func TestTopLevelClient_NewListBySubscriptionPager(t *testing.T) {
	TopLevelClientListBySubscriptionResponsePager := clientFactory.NewTopLevelClient().NewListBySubscriptionPager(nil)
	require.True(t, TopLevelClientListBySubscriptionResponsePager.More())
	TopLevelClientListBySubscriptionResponse, err := TopLevelClientListBySubscriptionResponsePager.NextPage(ctx)
	require.NoError(t, err)
	require.Len(t, TopLevelClientListBySubscriptionResponse.Value, 1)
	require.Equal(t, *validTopLevelResource.ID, *TopLevelClientListBySubscriptionResponse.Value[0].ID)
	require.Equal(t, *validTopLevelResource.Name, *TopLevelClientListBySubscriptionResponse.Value[0].Name)
	require.Equal(t, *validTopLevelResource.Type, *TopLevelClientListBySubscriptionResponse.Value[0].Type)
	require.Equal(t, *validTopLevelResource.Location, *TopLevelClientListBySubscriptionResponse.Value[0].Location)
	require.Equal(t, *validTopLevelResource.Properties.Description, *TopLevelClientListBySubscriptionResponse.Value[0].Properties.Description)
	require.Equal(t, *validTopLevelResource.Properties.ProvisioningState, *TopLevelClientListBySubscriptionResponse.Value[0].Properties.ProvisioningState)
}

func TestTopLevelClient_ActionSync(t *testing.T) {
	_, err := clientFactory.NewTopLevelClient().ActionSync(
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
