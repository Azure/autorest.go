// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armmultisharedmodelsgroup_test

import (
	"armmultisharedmodelsgroup"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var (
	validStorageAccountResource = armmultisharedmodelsgroup.StorageAccount{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/account1", subscriptionIDExpected, resourceGroupExpected)),
		Name:     to.Ptr("account1"),
		Type:     to.Ptr("Microsoft.Storage/storageAccounts"),
		Location: to.Ptr("westus"),
		Properties: &armmultisharedmodelsgroup.StorageAccountProperties{
			ProvisioningState: to.Ptr(armmultisharedmodelsgroup.ResourceProvisioningStateSucceeded),
		},
	}
)

func TestStorageAccountsClient_Get(t *testing.T) {
	resp, err := clientFactory.NewStorageAccountsClient().Get(ctx, resourceGroupExpected, "account1", nil)
	require.NoError(t, err)
	require.Equal(t, *validStorageAccountResource.ID, *resp.ID)
	require.Equal(t, *validStorageAccountResource.Name, *resp.Name)
	require.Equal(t, *validStorageAccountResource.Type, *resp.Type)
	require.Equal(t, *validStorageAccountResource.Location, *resp.Location)
	require.Equal(t, *validStorageAccountResource.Properties.ProvisioningState, *resp.Properties.ProvisioningState)
	require.NotNil(t, resp.Properties.Metadata)
	require.EqualValues(t, "admin@example.com", *resp.Properties.Metadata.CreatedBy)
	require.EqualValues(t, "engineering", *resp.Properties.Metadata.Tags["department"])
}

func TestStorageAccountsClient_CreateOrUpdate(t *testing.T) {
	poller, err := clientFactory.NewStorageAccountsClient().BeginCreateOrUpdate(ctx, resourceGroupExpected, "account1", armmultisharedmodelsgroup.StorageAccount{
		Location: to.Ptr("westus"),
		Properties: &armmultisharedmodelsgroup.StorageAccountProperties{
			Metadata: &armmultisharedmodelsgroup.SharedMetadata{
				CreatedBy: to.Ptr("admin@example.com"),
				Tags:      map[string]*string{"department": to.Ptr("engineering")},
			},
		},
	}, nil)
	require.NoError(t, err)

	resp, err := poller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validStorageAccountResource.ID, *resp.ID)
	require.Equal(t, *validStorageAccountResource.Name, *resp.Name)
	require.Equal(t, *validStorageAccountResource.Type, *resp.Type)
	require.Equal(t, *validStorageAccountResource.Location, *resp.Location)
	require.Equal(t, *validStorageAccountResource.Properties.ProvisioningState, *resp.Properties.ProvisioningState)
}
