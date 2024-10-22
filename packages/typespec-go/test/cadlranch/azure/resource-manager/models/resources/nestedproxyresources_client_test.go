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
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources/top/nestedProxyResources/nested"),
		Name: to.Ptr("nested"),
		Type: to.Ptr("Azure.ResourceManager.Models.Resources/topLevelTrackedResources/top/nestedProxyResources"),
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

func TestNestedProxyResourcesClient_Get(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	nestedProxyResourcesClientGetResponse, err := clientFactory.NewNestedProxyResourcesClient().Get(
		ctx,
		"test-rg",
		"top",
		"nested",
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *validNestedResource.ID, *nestedProxyResourcesClientGetResponse.ID)
	require.Equal(t, *validNestedResource.Name, *nestedProxyResourcesClientGetResponse.Name)
	require.Equal(t, *validNestedResource.Type, *nestedProxyResourcesClientGetResponse.Type)
}

func TestNestedProxyResourcesClient_CreateOrReplase(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	nestedProxyResourcesClientCreateOrReplaceResponsePoller, err := clientFactory.NewNestedProxyResourcesClient().BeginCreateOrReplace(
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
	nestedProxyResourcesClientCreateOrReplaceResponse, err := nestedProxyResourcesClientCreateOrReplaceResponsePoller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validNestedResource.ID, *nestedProxyResourcesClientCreateOrReplaceResponse.ID)
	require.Equal(t, *validNestedResource.Name, *nestedProxyResourcesClientCreateOrReplaceResponse.Name)
	require.Equal(t, *validNestedResource.Type, *nestedProxyResourcesClientCreateOrReplaceResponse.Type)
	require.Equal(t, *validNestedResource.Properties.Description, *nestedProxyResourcesClientCreateOrReplaceResponse.Properties.Description)
}

func TestNestedProxyResourcesClient_BeginUpdate(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	nestedProxyResourcesClientUpdateResponsePoller, err := clientFactory.NewNestedProxyResourcesClient().BeginUpdate(
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
	nestedProxyResourcesClientUpdateResponse, err := nestedProxyResourcesClientUpdateResponsePoller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validNestedResource.ID, *nestedProxyResourcesClientUpdateResponse.ID)
	require.Equal(t, *validNestedResource.Name, *nestedProxyResourcesClientUpdateResponse.Name)
	require.Equal(t, *validNestedResource.Type, *nestedProxyResourcesClientUpdateResponse.Type)
	require.Equal(t, "valid2", *nestedProxyResourcesClientUpdateResponse.Properties.Description)
}

func TestNestedProxyResourcesClient_BeginDelete(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	nestedProxyResourcesClientDeleteResponsePoller, err := clientFactory.NewNestedProxyResourcesClient().BeginDelete(ctx, "test-rg", "top", "nested", nil)
	require.NoError(t, err)
	nestedProxyResourcesClientDeleteResponse, err := nestedProxyResourcesClientDeleteResponsePoller.Poll(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, nestedProxyResourcesClientDeleteResponse.StatusCode)
}

func TestNestedProxyResourcesClient_NewListByTopLevelTrackedResourcePager(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/1709")
	nestedProxyResourcesClientListByTopLevelTrackedResourceResponsePager := clientFactory.NewNestedProxyResourcesClient().NewListByTopLevelTrackedResourcePager("test-rg", "top", nil)
	require.True(t, nestedProxyResourcesClientListByTopLevelTrackedResourceResponsePager.More())
	nestedProxyResourcesClientListByTopLevelTrackedResourceResponse, err := nestedProxyResourcesClientListByTopLevelTrackedResourceResponsePager.NextPage(ctx)
	require.NoError(t, err)
	require.Len(t, nestedProxyResourcesClientListByTopLevelTrackedResourceResponse.Value, 1)
	require.Equal(t, *validNestedResource.ID, *nestedProxyResourcesClientListByTopLevelTrackedResourceResponse.Value[0].ID)
	require.Equal(t, *validNestedResource.Name, *nestedProxyResourcesClientListByTopLevelTrackedResourceResponse.Value[0].Name)
	require.Equal(t, *validNestedResource.Type, *nestedProxyResourcesClientListByTopLevelTrackedResourceResponse.Value[0].Type)
}
