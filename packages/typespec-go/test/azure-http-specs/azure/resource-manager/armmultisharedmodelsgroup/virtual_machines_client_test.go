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
	validVirtualMachineResource = armmultisharedmodelsgroup.VirtualMachine{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachinesShared/vm-shared1", subscriptionIDExpected, resourceGroupExpected)),
		Name:     to.Ptr("vm-shared1"),
		Type:     to.Ptr("Microsoft.Compute/virtualMachinesShared"),
		Location: to.Ptr("eastus"),
		Properties: &armmultisharedmodelsgroup.VirtualMachineProperties{
			ProvisioningState: to.Ptr(armmultisharedmodelsgroup.ResourceProvisioningStateSucceeded),
		},
	}
)

func TestVirtualMachinesClient_Get(t *testing.T) {
	resp, err := clientFactory.NewVirtualMachinesClient().Get(ctx, resourceGroupExpected, "vm-shared1", nil)
	require.NoError(t, err)
	require.Equal(t, *validVirtualMachineResource.ID, *resp.ID)
	require.Equal(t, *validVirtualMachineResource.Name, *resp.Name)
	require.Equal(t, *validVirtualMachineResource.Type, *resp.Type)
	require.Equal(t, *validVirtualMachineResource.Location, *resp.Location)
	require.Equal(t, *validVirtualMachineResource.Properties.ProvisioningState, *resp.Properties.ProvisioningState)
}

func TestVirtualMachinesClient_CreateOrUpdate(t *testing.T) {
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, resourceGroupExpected, "vm-shared1", armmultisharedmodelsgroup.VirtualMachine{
		Location: to.Ptr("eastus"),
		Properties: &armmultisharedmodelsgroup.VirtualMachineProperties{
			Metadata: &armmultisharedmodelsgroup.SharedMetadata{
				CreatedBy: to.Ptr("user@example.com"),
				Tags:      map[string]*string{"environment": to.Ptr("production")},
			},
		},
	}, nil)
	require.NoError(t, err)

	resp, err := poller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validVirtualMachineResource.ID, *resp.ID)
	require.Equal(t, *validVirtualMachineResource.Name, *resp.Name)
	require.Equal(t, *validVirtualMachineResource.Type, *resp.Type)
	require.Equal(t, *validVirtualMachineResource.Location, *resp.Location)
	require.Equal(t, *validVirtualMachineResource.Properties.ProvisioningState, *resp.Properties.ProvisioningState)
}
