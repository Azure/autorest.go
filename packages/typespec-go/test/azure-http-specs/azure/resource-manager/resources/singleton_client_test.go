// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package resources_test

import (
	"resources"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var (
	validSingletonResource = resources.SingletonTrackedResource{
		ID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/singletonTrackedResources/default"),
		Name:     to.Ptr("default"),
		Type:     to.Ptr("Azure.ResourceManager.Resources/singletonTrackedResources"),
		Location: to.Ptr("eastus"),
		Properties: &resources.SingletonTrackedResourceProperties{
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

func TestSingletonClient_GetByResourceGroup(t *testing.T) {
	singletonClientGetByResourceGroupResponse, err := clientFactory.NewSingletonClient(subscriptionIdExpected).GetByResourceGroup(
		ctx,
		"test-rg",
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *validSingletonResource.ID, *singletonClientGetByResourceGroupResponse.ID)
	require.Equal(t, *validSingletonResource.Name, *singletonClientGetByResourceGroupResponse.Name)
	require.Equal(t, *validSingletonResource.Type, *singletonClientGetByResourceGroupResponse.Type)
	require.Equal(t, *validSingletonResource.Location, *singletonClientGetByResourceGroupResponse.Location)
	require.Equal(t, *validSingletonResource.Properties.Description, *singletonClientGetByResourceGroupResponse.Properties.Description)
	require.Equal(t, *validSingletonResource.Properties.ProvisioningState, *singletonClientGetByResourceGroupResponse.Properties.ProvisioningState)
}

func TestSingletonClient_BeginCreateOrUpdate(t *testing.T) {
	singletonClientBeginCreateOrUpdatePoller, err := clientFactory.NewSingletonClient(subscriptionIdExpected).BeginCreateOrUpdate(
		ctx,
		"test-rg",
		resources.SingletonTrackedResource{
			Location: to.Ptr("eastus"),
			Properties: &resources.SingletonTrackedResourceProperties{
				Description: to.Ptr("valid"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	singletonClientBeginCreateOrUpdateResponse, err := singletonClientBeginCreateOrUpdatePoller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validSingletonResource.ID, *singletonClientBeginCreateOrUpdateResponse.ID)
	require.Equal(t, *validSingletonResource.Name, *singletonClientBeginCreateOrUpdateResponse.Name)
	require.Equal(t, *validSingletonResource.Type, *singletonClientBeginCreateOrUpdateResponse.Type)
	require.Equal(t, *validSingletonResource.Location, *singletonClientBeginCreateOrUpdateResponse.Location)
	require.Equal(t, *validSingletonResource.Properties.Description, *singletonClientBeginCreateOrUpdateResponse.Properties.Description)
	require.Equal(t, *validSingletonResource.Properties.ProvisioningState, *singletonClientBeginCreateOrUpdateResponse.Properties.ProvisioningState)
}

func TestSingletonClient_Update(t *testing.T) {
	singletonClientUpdateResponse, err := clientFactory.NewSingletonClient(subscriptionIdExpected).Update(
		ctx,
		"test-rg",
		resources.SingletonTrackedResource{
			Properties: &resources.SingletonTrackedResourceProperties{
				Description: to.Ptr("valid2"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *validSingletonResource.ID, *singletonClientUpdateResponse.ID)
	require.Equal(t, *validSingletonResource.Name, *singletonClientUpdateResponse.Name)
	require.Equal(t, *validSingletonResource.Type, *singletonClientUpdateResponse.Type)
	require.Equal(t, *validSingletonResource.Location, *singletonClientUpdateResponse.Location)
	require.Equal(t, "valid2", *singletonClientUpdateResponse.Properties.Description)
	require.Equal(t, *validSingletonResource.Properties.ProvisioningState, *singletonClientUpdateResponse.Properties.ProvisioningState)
}

func TestSingletonClient_NewListByResourceGroupPager(t *testing.T) {
	singletonClientListByResourceGroupResponsePager := clientFactory.NewSingletonClient(subscriptionIdExpected).NewListByResourceGroupPager("test-rg", nil)
	require.True(t, singletonClientListByResourceGroupResponsePager.More())
	singletonClientListByResourceGroupResponse, err := singletonClientListByResourceGroupResponsePager.NextPage(ctx)
	require.NoError(t, err)
	require.Len(t, singletonClientListByResourceGroupResponse.Value, 1)
	require.Equal(t, *validSingletonResource.ID, *singletonClientListByResourceGroupResponse.Value[0].ID)
	require.Equal(t, *validSingletonResource.Name, *singletonClientListByResourceGroupResponse.Value[0].Name)
	require.Equal(t, *validSingletonResource.Type, *singletonClientListByResourceGroupResponse.Value[0].Type)
	require.Equal(t, *validSingletonResource.Location, *singletonClientListByResourceGroupResponse.Value[0].Location)
	require.Equal(t, *validSingletonResource.Properties.Description, *singletonClientListByResourceGroupResponse.Value[0].Properties.Description)
	require.Equal(t, *validSingletonResource.Properties.ProvisioningState, *singletonClientListByResourceGroupResponse.Value[0].Properties.ProvisioningState)
}
