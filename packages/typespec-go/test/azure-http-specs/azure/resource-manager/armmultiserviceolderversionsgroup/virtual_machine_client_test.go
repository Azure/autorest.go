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
	validVirtualMachineResource = armmultiserviceolderversionsgroup.VirtualMachine{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachinesOld/vm-old1", subscriptionIdExpected, resourceGroupExpected)),
		Name:     to.Ptr("vm-old1"),
		Type:     to.Ptr("Microsoft.Compute/virtualMachinesOld"),
		Location: to.Ptr(locationExpected),
		Properties: &armmultiserviceolderversionsgroup.VirtualMachineProperties{
			ProvisioningState: to.Ptr(armmultiserviceolderversionsgroup.ResourceProvisioningStateSucceeded),
			Size:              to.Ptr("Standard_D2s_v3"),
		},
	}
)

func TestVirtualMachineClient_Get(t *testing.T) {
	virtualMachineClientGetResponse, err := clientFactory.NewVirtualMachinesClient().Get(ctx, resourceGroupExpected, "vm-old1", nil)
	require.NoError(t, err)
	require.Equal(t, *validVirtualMachineResource.ID, *virtualMachineClientGetResponse.ID)
	require.Equal(t, *validVirtualMachineResource.Name, *virtualMachineClientGetResponse.Name)
	require.Equal(t, *validVirtualMachineResource.Type, *virtualMachineClientGetResponse.Type)
	require.Equal(t, *validVirtualMachineResource.Location, *virtualMachineClientGetResponse.Location)
	require.Equal(t, *validVirtualMachineResource.Properties.ProvisioningState, *virtualMachineClientGetResponse.Properties.ProvisioningState)
	require.Equal(t, *validVirtualMachineResource.Properties.Size, *virtualMachineClientGetResponse.Properties.Size)
}

func TestVirtualMachineClient_CreateOrUpdate(t *testing.T) {
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, resourceGroupExpected, "vm-old1", armmultiserviceolderversionsgroup.VirtualMachine{
		Location: validVirtualMachineResource.Location,
		Properties: &armmultiserviceolderversionsgroup.VirtualMachineProperties{
			Size: to.Ptr("Standard_D2s_v3"),
		},
	}, nil)
	require.NoError(t, err)

	virtualMachineClientCreateOrUpdateResponse, err := poller.PollUntilDone(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, *validVirtualMachineResource.ID, *virtualMachineClientCreateOrUpdateResponse.ID)
	require.Equal(t, *validVirtualMachineResource.Name, *virtualMachineClientCreateOrUpdateResponse.Name)
	require.Equal(t, *validVirtualMachineResource.Type, *virtualMachineClientCreateOrUpdateResponse.Type)
	require.Equal(t, *validVirtualMachineResource.Location, *virtualMachineClientCreateOrUpdateResponse.Location)
	require.Equal(t, *validVirtualMachineResource.Properties.ProvisioningState, *virtualMachineClientCreateOrUpdateResponse.Properties.ProvisioningState)
	require.Equal(t, *validVirtualMachineResource.Properties.Size, *virtualMachineClientCreateOrUpdateResponse.Properties.Size)
}
