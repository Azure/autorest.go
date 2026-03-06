// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armmultiserviceolderversionsgroup_test

import (
	"armmultiserviceolderversionsgroup"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var (
	validDiskResource = armmultiserviceolderversionsgroup.Disk{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/disksOld/disk-old1", subscriptionIdExpected, resourceGroupExpected)),
		Name:     to.Ptr("disk-old1"),
		Type:     to.Ptr("Microsoft.Compute/disksOld"),
		Location: to.Ptr(locationExpected),
		Properties: &armmultiserviceolderversionsgroup.DiskProperties{
			ProvisioningState: to.Ptr(armmultiserviceolderversionsgroup.ResourceProvisioningStateSucceeded),
			DiskSizeGB:        to.Ptr(int32(128)),
		},
	}
)

func TestDiskClient_Get(t *testing.T) {
	diskClientGetResponse, err := clientFactory.NewDisksClient().Get(ctx, resourceGroupExpected, "disk-old1", nil)
	require.NoError(t, err)
	require.Equal(t, *validDiskResource.ID, *diskClientGetResponse.ID)
	require.Equal(t, *validDiskResource.Name, *diskClientGetResponse.Name)
	require.Equal(t, *validDiskResource.Type, *diskClientGetResponse.Type)
	require.Equal(t, *validDiskResource.Location, *diskClientGetResponse.Location)
	require.Equal(t, *validDiskResource.Properties.ProvisioningState, *diskClientGetResponse.Properties.ProvisioningState)
	require.Equal(t, *validDiskResource.Properties.DiskSizeGB, *diskClientGetResponse.Properties.DiskSizeGB)
}

func TestDiskClient_CreateOrUpdate(t *testing.T) {
	poller, err := clientFactory.NewDisksClient().BeginCreateOrUpdate(ctx, resourceGroupExpected, "disk-old1", armmultiserviceolderversionsgroup.Disk{
		Location: validDiskResource.Location,
		Properties: &armmultiserviceolderversionsgroup.DiskProperties{
			DiskSizeGB: to.Ptr(int32(128)),
		},
	}, nil)
	require.NoError(t, err)

	diskClientCreateOrUpdateResponse, err := poller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validDiskResource.ID, *diskClientCreateOrUpdateResponse.ID)
	require.Equal(t, *validDiskResource.Name, *diskClientCreateOrUpdateResponse.Name)
	require.Equal(t, *validDiskResource.Type, *diskClientCreateOrUpdateResponse.Type)
	require.Equal(t, *validDiskResource.Location, *diskClientCreateOrUpdateResponse.Location)
	require.Equal(t, *validDiskResource.Properties.ProvisioningState, *diskClientCreateOrUpdateResponse.Properties.ProvisioningState)
	require.Equal(t, *validDiskResource.Properties.DiskSizeGB, *diskClientCreateOrUpdateResponse.Properties.DiskSizeGB)
}
