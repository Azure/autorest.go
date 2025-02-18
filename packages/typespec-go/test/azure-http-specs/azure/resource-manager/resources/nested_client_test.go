// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package resources_test

import (
	"net/http"
	"resources"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var (
	validNestedResource = resources.NestedProxyResource{
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top/nestedProxyResources/nested"),
		Name: to.Ptr("nested"),
		Type: to.Ptr("Azure.ResourceManager.Resources/topLevelTrackedResources/top/nestedProxyResources"),
		Properties: &resources.NestedProxyResourceProperties{
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

func TestNestedClient_Get(t *testing.T) {
	nestedClientGetResponse, err := clientFactory.NewNestedClient(subscriptionIdExpected).Get(
		ctx,
		"test-rg",
		"top",
		"nested",
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *validNestedResource.ID, *nestedClientGetResponse.ID)
	require.Equal(t, *validNestedResource.Name, *nestedClientGetResponse.Name)
	require.Equal(t, *validNestedResource.Type, *nestedClientGetResponse.Type)
	require.Equal(t, *validNestedResource.Properties.Description, *nestedClientGetResponse.Properties.Description)
	require.Equal(t, *validNestedResource.Properties.ProvisioningState, *nestedClientGetResponse.Properties.ProvisioningState)
}

func TestNestedClient_BeginCreateOrReplace(t *testing.T) {
	nestedClientCreateOrReplaceResponsePoller, err := clientFactory.NewNestedClient(subscriptionIdExpected).BeginCreateOrReplace(
		ctx,
		"test-rg",
		"top",
		"nested",
		resources.NestedProxyResource{
			Properties: &resources.NestedProxyResourceProperties{
				Description: to.Ptr("valid"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	nestedClientCreateOrReplaceResponse, err := nestedClientCreateOrReplaceResponsePoller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validNestedResource.ID, *nestedClientCreateOrReplaceResponse.ID)
	require.Equal(t, *validNestedResource.Name, *nestedClientCreateOrReplaceResponse.Name)
	require.Equal(t, *validNestedResource.Type, *nestedClientCreateOrReplaceResponse.Type)
	require.Equal(t, *validNestedResource.Properties.Description, *nestedClientCreateOrReplaceResponse.Properties.Description)
	require.Equal(t, *validNestedResource.Properties.ProvisioningState, *nestedClientCreateOrReplaceResponse.Properties.ProvisioningState)
}

func TestNestedClient_BeginUpdate(t *testing.T) {
	nestedClientUpdateResponsePoller, err := clientFactory.NewNestedClient(subscriptionIdExpected).BeginUpdate(
		ctx,
		"test-rg",
		"top",
		"nested",
		resources.NestedProxyResource{
			Properties: &resources.NestedProxyResourceProperties{
				Description: to.Ptr("valid2"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	nestedClientUpdateResponse, err := nestedClientUpdateResponsePoller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validNestedResource.ID, *nestedClientUpdateResponse.ID)
	require.Equal(t, *validNestedResource.Name, *nestedClientUpdateResponse.Name)
	require.Equal(t, *validNestedResource.Type, *nestedClientUpdateResponse.Type)
	require.Equal(t, "valid2", *nestedClientUpdateResponse.Properties.Description)
	require.Equal(t, *validNestedResource.Properties.ProvisioningState, *nestedClientUpdateResponse.Properties.ProvisioningState)
}

func TestNestedClient_BeginDelete(t *testing.T) {
	nestedClientDeleteResponsePoller, err := clientFactory.NewNestedClient(subscriptionIdExpected).BeginDelete(ctx, "test-rg", "top", "nested", nil)
	require.NoError(t, err)
	nestedClientDeleteResponse, err := nestedClientDeleteResponsePoller.Poll(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, nestedClientDeleteResponse.StatusCode)
}

func TestNestedClient_NewListByTopLevelTrackedResourcePager(t *testing.T) {
	nestedClientListByTopLevelTrackedResourceResponsePager := clientFactory.NewNestedClient(subscriptionIdExpected).NewListByTopLevelTrackedResourcePager("test-rg", "top", nil)
	require.True(t, nestedClientListByTopLevelTrackedResourceResponsePager.More())
	nestedClientListByTopLevelTrackedResourceResponse, err := nestedClientListByTopLevelTrackedResourceResponsePager.NextPage(ctx)
	require.NoError(t, err)
	require.Len(t, nestedClientListByTopLevelTrackedResourceResponse.Value, 1)
	require.Equal(t, *validNestedResource.ID, *nestedClientListByTopLevelTrackedResourceResponse.Value[0].ID)
	require.Equal(t, *validNestedResource.Name, *nestedClientListByTopLevelTrackedResourceResponse.Value[0].Name)
	require.Equal(t, *validNestedResource.Type, *nestedClientListByTopLevelTrackedResourceResponse.Value[0].Type)
	require.Equal(t, *validNestedResource.Properties.Description, *nestedClientListByTopLevelTrackedResourceResponse.Value[0].Properties.Description)
	require.Equal(t, *validNestedResource.Properties.ProvisioningState, *nestedClientListByTopLevelTrackedResourceResponse.Value[0].Properties.ProvisioningState)
}
