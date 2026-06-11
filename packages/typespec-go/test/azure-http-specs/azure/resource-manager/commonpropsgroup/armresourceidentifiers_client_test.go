// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package commonpropsgroup_test

import (
	"commonpropsgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var (
	expectedArmResourceIdentifierResource = commonpropsgroup.ArmResourceIdentifierResource{
		ID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Azure.ResourceManager.CommonProperties/armResourceIdentifierResources/armId"),
		Location: to.Ptr(locationExpected),
		Name:     to.Ptr("armId"),
		Type:     to.Ptr("Azure.ResourceManager.CommonProperties/armResourceIdentifierResources"),
		Properties: &commonpropsgroup.ArmResourceIdentifierResourceProperties{
			ProvisioningState: to.Ptr(commonpropsgroup.ResourceProvisioningStateSucceeded),
			SimpleArmID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/myVnet"),
			ArmIDWithType:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/myVnet"),
			ArmIDWithTypeAndScope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/myVnet"),
			ArmIDWithAllScopes:    to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/myVm"),
		},
	}
)

func TestArmResourceIdentifiersClient_Get(t *testing.T) {
	resp, err := clientFactory.NewArmResourceIdentifiersClient().Get(ctx, resourceGroupExpected, "armId", nil)
	require.NoError(t, err)
	require.Equal(t, *expectedArmResourceIdentifierResource.ID, *resp.ID)
	require.Equal(t, *expectedArmResourceIdentifierResource.Location, *resp.Location)
	require.Equal(t, *expectedArmResourceIdentifierResource.Name, *resp.Name)
	require.Equal(t, *expectedArmResourceIdentifierResource.Type, *resp.Type)
	require.Equal(t, *expectedArmResourceIdentifierResource.Properties.ProvisioningState, *resp.Properties.ProvisioningState)
	require.Equal(t, *expectedArmResourceIdentifierResource.Properties.SimpleArmID, *resp.Properties.SimpleArmID)
	require.Equal(t, *expectedArmResourceIdentifierResource.Properties.ArmIDWithType, *resp.Properties.ArmIDWithType)
	require.Equal(t, *expectedArmResourceIdentifierResource.Properties.ArmIDWithTypeAndScope, *resp.Properties.ArmIDWithTypeAndScope)
	require.Equal(t, *expectedArmResourceIdentifierResource.Properties.ArmIDWithAllScopes, *resp.Properties.ArmIDWithAllScopes)
}

func TestArmResourceIdentifiersClient_CreateOrReplace(t *testing.T) {
	resp, err := clientFactory.NewArmResourceIdentifiersClient().CreateOrReplace(
		ctx,
		resourceGroupExpected,
		"armId",
		commonpropsgroup.ArmResourceIdentifierResource{
			Location: to.Ptr(locationExpected),
			Properties: &commonpropsgroup.ArmResourceIdentifierResourceProperties{
				SimpleArmID:           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/myVnet"),
				ArmIDWithType:         to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/myVnet"),
				ArmIDWithTypeAndScope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/myVnet"),
				ArmIDWithAllScopes:    to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/myVm"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *expectedArmResourceIdentifierResource.ID, *resp.ID)
	require.Equal(t, *expectedArmResourceIdentifierResource.Location, *resp.Location)
	require.Equal(t, *expectedArmResourceIdentifierResource.Name, *resp.Name)
	require.Equal(t, *expectedArmResourceIdentifierResource.Type, *resp.Type)
	require.Equal(t, *expectedArmResourceIdentifierResource.Properties.ProvisioningState, *resp.Properties.ProvisioningState)
	require.Equal(t, *expectedArmResourceIdentifierResource.Properties.SimpleArmID, *resp.Properties.SimpleArmID)
	require.Equal(t, *expectedArmResourceIdentifierResource.Properties.ArmIDWithType, *resp.Properties.ArmIDWithType)
	require.Equal(t, *expectedArmResourceIdentifierResource.Properties.ArmIDWithTypeAndScope, *resp.Properties.ArmIDWithTypeAndScope)
	require.Equal(t, *expectedArmResourceIdentifierResource.Properties.ArmIDWithAllScopes, *resp.Properties.ArmIDWithAllScopes)
}
