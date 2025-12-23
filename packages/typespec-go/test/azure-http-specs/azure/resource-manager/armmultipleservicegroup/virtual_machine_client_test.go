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
	validVirtualMachineResource = armmultipleservicegroup.VirtualMachine{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/vm1", subscriptionIdExpected, resourceGroupExpected)),
		Name:     to.Ptr("vm1"),
		Type:     to.Ptr("Microsoft.Compute/virtualMachines"),
		Location: to.Ptr(locationExpected),
		Properties: &armmultipleservicegroup.VirtualMachineProperties{
			ProvisioningState: to.Ptr(armmultipleservicegroup.ResourceProvisioningStateSucceeded),
		},
	}
)

func TestVirtualMachineClient_Get(t *testing.T) {
	virtualMachineClientGetResponse, err := clientFactory.NewVirtualMachinesClient().Get(ctx, resourceGroupExpected, "vm1", nil)
	require.NoError(t, err)
	require.Equal(t, *validVirtualMachineResource.ID, *virtualMachineClientGetResponse.ID)
	require.Equal(t, *validVirtualMachineResource.Name, *virtualMachineClientGetResponse.Name)
	require.Equal(t, *validVirtualMachineResource.Type, *virtualMachineClientGetResponse.Type)
	require.Equal(t, *validVirtualMachineResource.Location, *virtualMachineClientGetResponse.Location)
	require.Equal(t, *validVirtualMachineResource.Properties.ProvisioningState, *virtualMachineClientGetResponse.Properties.ProvisioningState)
}

func TestVirtualMachineClient_CreateOrUpdate(t *testing.T) {
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, resourceGroupExpected, "vm1", armmultipleservicegroup.VirtualMachine{
		Location:   validVirtualMachineResource.Location,
		Properties: &armmultipleservicegroup.VirtualMachineProperties{},
	}, nil)
	require.NoError(t, err)

	virtualMachineClientCreateOrUpdateResponse, err := poller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validVirtualMachineResource.ID, *virtualMachineClientCreateOrUpdateResponse.ID)
	require.Equal(t, *validVirtualMachineResource.Name, *virtualMachineClientCreateOrUpdateResponse.Name)
	require.Equal(t, *validVirtualMachineResource.Type, *virtualMachineClientCreateOrUpdateResponse.Type)
	require.Equal(t, *validVirtualMachineResource.Location, *virtualMachineClientCreateOrUpdateResponse.Location)
	require.Equal(t, *validVirtualMachineResource.Properties.ProvisioningState, *virtualMachineClientCreateOrUpdateResponse.Properties.ProvisioningState)
}
