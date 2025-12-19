// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armmultipleservicegroup_test

import (
	"armmultipleservicegroup"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var (
	validDiskResource = armmultipleservicegroup.Disk{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/disks/disk1", subscriptionIdExpected, resourceGroupExpected)),
		Name:     to.Ptr("disk1"),
		Type:     to.Ptr("Microsoft.Compute/disks"),
		Location: to.Ptr(locationExpected),
		Properties: &armmultipleservicegroup.DiskProperties{
			ProvisioningState: to.Ptr(armmultipleservicegroup.ResourceProvisioningStateSucceeded),
		},
	}
)

func TestDiskClient_Get(t *testing.T) {
	diskClientGetResponse, err := clientFactory.NewDisksClient().Get(ctx, resourceGroupExpected, "disk1", nil)
	require.NoError(t, err)
	require.Equal(t, *validDiskResource.ID, *diskClientGetResponse.ID)
	require.Equal(t, *validDiskResource.Name, *diskClientGetResponse.Name)
	require.Equal(t, *validDiskResource.Type, *diskClientGetResponse.Type)
	require.Equal(t, *validDiskResource.Location, *diskClientGetResponse.Location)
	require.Equal(t, *validDiskResource.Properties.ProvisioningState, *diskClientGetResponse.Properties.ProvisioningState)
}

func TestDiskClient_CreateOrUpdate(t *testing.T) {
	poller, err := clientFactory.NewDisksClient().BeginCreateOrUpdate(ctx, resourceGroupExpected, "disk1", armmultipleservicegroup.Disk{
		Location:   validDiskResource.Location,
		Properties: &armmultipleservicegroup.DiskProperties{},
	}, nil)
	require.NoError(t, err)

	diskClientCreateOrUpdateResponse, err := poller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validDiskResource.ID, *diskClientCreateOrUpdateResponse.ID)
	require.Equal(t, *validDiskResource.Name, *diskClientCreateOrUpdateResponse.Name)
	require.Equal(t, *validDiskResource.Type, *diskClientCreateOrUpdateResponse.Type)
	require.Equal(t, *validDiskResource.Location, *diskClientCreateOrUpdateResponse.Location)
	require.Equal(t, *validDiskResource.Properties.ProvisioningState, *diskClientCreateOrUpdateResponse.Properties.ProvisioningState)
}
