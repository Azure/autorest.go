// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package resources_test

import (
	"context"
	"resources"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestLocationResourcesClient_CreateOrUpdate(t *testing.T) {
	resp, err := clientFactory.NewLocationResourcesClient().CreateOrUpdate(context.Background(), "eastus", "resource", resources.LocationResource{
		Properties: &resources.LocationResourceProperties{
			Description: to.Ptr("valid"),
		},
	}, nil)
	require.NoError(t, err)
	require.Equal(t, resources.LocationResource{
		Properties: &resources.LocationResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("resource"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Azure.ResourceManager.Resources/locations/eastus/locationResources/resource"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/locationResources"),
	}, resp.LocationResource)
}

func TestLocationResourcesClient_Delete(t *testing.T) {
	resp, err := clientFactory.NewLocationResourcesClient().Delete(context.Background(), "eastus", "resource", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestLocationResourcesClient_Get(t *testing.T) {
	resp, err := clientFactory.NewLocationResourcesClient().Get(context.Background(), "eastus", "resource", nil)
	require.NoError(t, err)
	require.Equal(t, resources.LocationResource{
		Properties: &resources.LocationResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("resource"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Azure.ResourceManager.Resources/locations/eastus/locationResources/resource"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/locationResources"),
	}, resp.LocationResource)
}

func TestLocationResourcesClient_NewListByScopePager(t *testing.T) {
	pager := clientFactory.NewLocationResourcesClient().NewListByLocationPager("eastus", nil)
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Value, 1)
		require.Equal(t, &resources.LocationResource{
			Properties: &resources.LocationResourceProperties{
				Description:       to.Ptr("valid"),
				ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
			},
			Name: to.Ptr("resource"),
			ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Azure.ResourceManager.Resources/locations/eastus/locationResources/resource"),
			SystemData: &resources.SystemData{
				CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
				CreatedBy:          to.Ptr("AzureSDK"),
				CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
				LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
				LastModifiedBy:     to.Ptr("AzureSDK"),
				LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
			},
			Type: to.Ptr("Azure.ResourceManager.Resources/locationResources"),
		}, page.Value[0])
		pageCount++
	}
	require.EqualValues(t, 1, pageCount)
}

func TestLocationResourcesClient_Update(t *testing.T) {
	resp, err := clientFactory.NewLocationResourcesClient().Update(context.Background(), "eastus", "resource", resources.LocationResource{
		Properties: &resources.LocationResourceProperties{
			Description: to.Ptr("valid2"),
		},
	}, nil)
	require.NoError(t, err)
	require.Equal(t, resources.LocationResource{
		Properties: &resources.LocationResourceProperties{
			Description:       to.Ptr("valid2"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("resource"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Azure.ResourceManager.Resources/locations/eastus/locationResources/resource"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/locationResources"),
	}, resp.LocationResource)
}
