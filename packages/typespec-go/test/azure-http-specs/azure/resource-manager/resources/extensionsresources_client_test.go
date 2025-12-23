// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package resources_test

import (
	"context"
	"resources"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestExtensionsResourcesClient_BeginCreateOrUpdate(t *testing.T) {
	poller, err := clientFactory.NewExtensionsResourcesClient().BeginCreateOrUpdate(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg", "extension", resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description: to.Ptr("valid"),
		},
	}, nil)
	require.NoError(t, err)
	resp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: 1 * time.Second,
	})
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)


	poller2, err := clientFactory.NewExtensionsResourcesClient().BeginCreateOrUpdate(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000", "extension", resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description: to.Ptr("valid"),
		},
	}, nil)
	require.NoError(t, err)
	resp, err = poller2.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: 1 * time.Second,
	})
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)


	poller3, err := clientFactory.NewExtensionsResourcesClient().BeginCreateOrUpdate(context.Background(), "/", "extension", resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description: to.Ptr("valid"),
		},
	}, nil)
	require.NoError(t, err)
	resp, err = poller3.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: 1 * time.Second,
	})
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)

	poller4, err := clientFactory.NewExtensionsResourcesClient().BeginCreateOrUpdate(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top", "extension", resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description: to.Ptr("valid"),
		},
	}, nil)
	require.NoError(t, err)
	resp, err = poller4.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: 1 * time.Second,
	})
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)

}

func TestExtensionsResourcesClient_Delete(t *testing.T) {
	resp, err := clientFactory.NewExtensionsResourcesClient().Delete(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg", "extension", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
	resp, err = clientFactory.NewExtensionsResourcesClient().Delete(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000", "extension", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
	resp, err = clientFactory.NewExtensionsResourcesClient().Delete(context.Background(), "/", "extension", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
	resp, err = clientFactory.NewExtensionsResourcesClient().Delete(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top", "extension", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestExtensionsResourcesClient_Get(t *testing.T) {
	resp, err := clientFactory.NewExtensionsResourcesClient().Get(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg", "extension", nil)
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)

	resp, err = clientFactory.NewExtensionsResourcesClient().Get(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000", "extension", nil)
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)

	resp, err = clientFactory.NewExtensionsResourcesClient().Get(context.Background(), "/", "extension", nil)
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)

	resp, err = clientFactory.NewExtensionsResourcesClient().Get(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top", "extension", nil)
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)
}

func TestExtensionsResourcesClient_NewListByScopePager(t *testing.T) {
	pager := clientFactory.NewExtensionsResourcesClient().NewListByScopePager("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg", nil)
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Value, 1)
		require.Equal(t, &resources.ExtensionsResource{
			Properties: &resources.ExtensionsResourceProperties{
				Description:       to.Ptr("valid"),
				ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
			},
			Name: to.Ptr("extension"),
			ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
			SystemData: &resources.SystemData{
				CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
				CreatedBy:          to.Ptr("AzureSDK"),
				CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
				LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
				LastModifiedBy:     to.Ptr("AzureSDK"),
				LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
			},
			Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
		}, page.Value[0])
		pageCount++
	}
	require.EqualValues(t, 1, pageCount)


	pager = clientFactory.NewExtensionsResourcesClient().NewListByScopePager("/subscriptions/00000000-0000-0000-0000-000000000000", nil)
	pageCount = 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Value, 1)
		require.Equal(t, &resources.ExtensionsResource{
			Properties: &resources.ExtensionsResourceProperties{
				Description:       to.Ptr("valid"),
				ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
			},
			Name: to.Ptr("extension"),
			ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
			SystemData: &resources.SystemData{
				CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
				CreatedBy:          to.Ptr("AzureSDK"),
				CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
				LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
				LastModifiedBy:     to.Ptr("AzureSDK"),
				LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
			},
			Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
		}, page.Value[0])
		pageCount++
	}
	require.EqualValues(t, 1, pageCount)

	pager = clientFactory.NewExtensionsResourcesClient().NewListByScopePager("/", nil)
	pageCount = 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Value, 1)
		require.Equal(t, &resources.ExtensionsResource{
			Properties: &resources.ExtensionsResourceProperties{
				Description:       to.Ptr("valid"),
				ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
			},
			Name: to.Ptr("extension"),
			ID:   to.Ptr("/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
			SystemData: &resources.SystemData{
				CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
				CreatedBy:          to.Ptr("AzureSDK"),
				CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
				LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
				LastModifiedBy:     to.Ptr("AzureSDK"),
				LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
			},
			Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
		}, page.Value[0])
		pageCount++
	}
	require.EqualValues(t, 1, pageCount)

	
	pager = clientFactory.NewExtensionsResourcesClient().NewListByScopePager("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top", nil)
	pageCount = 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Value, 1)
		require.Equal(t, &resources.ExtensionsResource{
			Properties: &resources.ExtensionsResourceProperties{
				Description:       to.Ptr("valid"),
				ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
			},
			Name: to.Ptr("extension"),
			ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
			SystemData: &resources.SystemData{
				CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
				CreatedBy:          to.Ptr("AzureSDK"),
				CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
				LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
				LastModifiedBy:     to.Ptr("AzureSDK"),
				LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
			},
			Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
		}, page.Value[0])
		pageCount++
	}
	require.EqualValues(t, 1, pageCount)
}

func TestExtensionsResourcesClient_Update(t *testing.T) {
	resp, err := clientFactory.NewExtensionsResourcesClient().Update(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg", "extension", resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description: to.Ptr("valid2"),
		},
	}, nil)
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid2"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)

	resp, err = clientFactory.NewExtensionsResourcesClient().Update(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000", "extension", resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description: to.Ptr("valid2"),
		},
	}, nil)
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid2"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)

	resp, err = clientFactory.NewExtensionsResourcesClient().Update(context.Background(), "/", "extension", resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description: to.Ptr("valid2"),
		},
	}, nil)
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid2"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)

	resp, err = clientFactory.NewExtensionsResourcesClient().Update(context.Background(), "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top", "extension", resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description: to.Ptr("valid2"),
		},
	}, nil)
	require.NoError(t, err)
	require.Equal(t, resources.ExtensionsResource{
		Properties: &resources.ExtensionsResourceProperties{
			Description:       to.Ptr("valid2"),
			ProvisioningState: to.Ptr(resources.ProvisioningStateSucceeded),
		},
		Name: to.Ptr("extension"),
		ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.Resources/topLevelTrackedResources/top/providers/Azure.ResourceManager.Resources/extensionsResources/extension"),
		SystemData: &resources.SystemData{
			CreatedAt:          to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      to.Ptr(resources.CreatedByTypeUser),
			LastModifiedAt:     to.Ptr(time.Date(2024, time.October, 4, 0, 56, 7, 442000000, time.UTC)),
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedByType: to.Ptr(resources.CreatedByTypeUser),
		},
		Type: to.Ptr("Azure.ResourceManager.Resources/extensionsResources"),
	}, resp.ExtensionsResource)
}
