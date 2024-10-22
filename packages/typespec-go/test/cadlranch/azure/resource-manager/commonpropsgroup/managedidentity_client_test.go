// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package commonpropsgroup_test

import (
	"commonpropsgroup"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var (
	validSystemAssignedManagedIdentityResource = commonpropsgroup.ManagedIdentityTrackedResource{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Azure.ResourceManager.CommonProperties/managedIdentityTrackedResources/identity", subscriptionIdExpected, resourceGroupExpected)),
		Location: to.Ptr(locationExpected),
		Tags: map[string]*string{
			"tagKey1": to.Ptr("tagValue1"),
		},
		Identity: &commonpropsgroup.ManagedServiceIdentity{
			Type:        to.Ptr(commonpropsgroup.ManagedServiceIdentityType(identityTypeSystemAssigendExpected)),
			PrincipalID: to.Ptr(principalIdExpected),
			TenantID:    to.Ptr(tenantIdExpected),
		},
		Properties: &commonpropsgroup.ManagedIdentityTrackedResourceProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}

	validUserAssignedAndSystemAssignedManagedIdentityResource = commonpropsgroup.ManagedIdentityTrackedResource{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Azure.ResourceManager.CommonProperties/managedIdentityTrackedResources/identity", subscriptionIdExpected, resourceGroupExpected)),
		Location: to.Ptr(locationExpected),
		Tags: map[string]*string{
			"tagKey1": to.Ptr("tagValue1"),
		},
		Identity: &commonpropsgroup.ManagedServiceIdentity{
			Type: to.Ptr(commonpropsgroup.ManagedServiceIdentityType(identityTypeSystemUserAssignedExpected)),
			UserAssignedIdentities: map[string]*commonpropsgroup.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {
					PrincipalID: to.Ptr(principalIdExpected),
					ClientID:    to.Ptr(clientIdExpected),
				},
			},
			PrincipalID: to.Ptr(principalIdExpected),
			TenantID:    to.Ptr(tenantIdExpected),
		},
		Properties: &commonpropsgroup.ManagedIdentityTrackedResourceProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}
)

func TestManagedIdentityTrackedResourcesClient_Get(t *testing.T) {
	managedIdentityTrackedResourcesClientGetResponse, err := clientFactory.NewManagedIdentityClient().Get(ctx, resourceGroupExpected, "identity", nil)
	require.NoError(t, err)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.ID, *managedIdentityTrackedResourcesClientGetResponse.ID)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.Location, *managedIdentityTrackedResourcesClientGetResponse.Location)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.Identity.Type, *managedIdentityTrackedResourcesClientGetResponse.Identity.Type)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.Identity.PrincipalID, *managedIdentityTrackedResourcesClientGetResponse.Identity.PrincipalID)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.Identity.TenantID, *managedIdentityTrackedResourcesClientGetResponse.Identity.TenantID)
	require.Equal(t, validSystemAssignedManagedIdentityResource.Tags, managedIdentityTrackedResourcesClientGetResponse.Tags)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.Properties.ProvisioningState, *managedIdentityTrackedResourcesClientGetResponse.Properties.ProvisioningState)
}

func TestManagedIdentityTrackedResourcesClient_CreateWithSystemAssigned(t *testing.T) {
	managedIdentityTrackedResourcesClientCreateWithSystemAssignedResponse, err := clientFactory.NewManagedIdentityClient().CreateWithSystemAssigned(
		ctx,
		resourceGroupExpected,
		"identity",
		commonpropsgroup.ManagedIdentityTrackedResource{
			Location: to.Ptr(locationExpected),
			Identity: &commonpropsgroup.ManagedServiceIdentity{
				Type: to.Ptr(commonpropsgroup.ManagedServiceIdentityTypeSystemAssigned),
			},
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.ID, *managedIdentityTrackedResourcesClientCreateWithSystemAssignedResponse.ID)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.Location, *managedIdentityTrackedResourcesClientCreateWithSystemAssignedResponse.Location)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.Identity.Type, *managedIdentityTrackedResourcesClientCreateWithSystemAssignedResponse.Identity.Type)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.Identity.PrincipalID, *managedIdentityTrackedResourcesClientCreateWithSystemAssignedResponse.Identity.PrincipalID)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.Identity.TenantID, *managedIdentityTrackedResourcesClientCreateWithSystemAssignedResponse.Identity.TenantID)
	require.Equal(t, validSystemAssignedManagedIdentityResource.Tags, managedIdentityTrackedResourcesClientCreateWithSystemAssignedResponse.Tags)
	require.Equal(t, *validSystemAssignedManagedIdentityResource.Properties, *managedIdentityTrackedResourcesClientCreateWithSystemAssignedResponse.Properties)
}

func TestManagedIdentityTrackedResourcesClient_UpdateWithUserAssignedAndSystemAssigned(t *testing.T) {
	managedIdentityTrackedResourcesClientUpdateWithUserAssignedAndSystemAssignedResponse, err := clientFactory.NewManagedIdentityClient().UpdateWithUserAssignedAndSystemAssigned(
		ctx,
		resourceGroupExpected,
		"identity",
		commonpropsgroup.ManagedIdentityTrackedResource{
			Location: to.Ptr(locationExpected),
			Identity: &commonpropsgroup.ManagedServiceIdentity{
				Type: to.Ptr(commonpropsgroup.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*commonpropsgroup.UserAssignedIdentity{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
				},
			},
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *validUserAssignedAndSystemAssignedManagedIdentityResource.ID, *managedIdentityTrackedResourcesClientUpdateWithUserAssignedAndSystemAssignedResponse.ID)
	require.Equal(t, *validUserAssignedAndSystemAssignedManagedIdentityResource.Location, *managedIdentityTrackedResourcesClientUpdateWithUserAssignedAndSystemAssignedResponse.Location)
	require.Equal(t, *validUserAssignedAndSystemAssignedManagedIdentityResource.Identity, *managedIdentityTrackedResourcesClientUpdateWithUserAssignedAndSystemAssignedResponse.Identity)
	require.Equal(t, validUserAssignedAndSystemAssignedManagedIdentityResource.Tags, managedIdentityTrackedResourcesClientUpdateWithUserAssignedAndSystemAssignedResponse.Tags)
	require.Equal(t, *validUserAssignedAndSystemAssignedManagedIdentityResource.Properties, *managedIdentityTrackedResourcesClientUpdateWithUserAssignedAndSystemAssignedResponse.Properties)
}
