// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package managed_identity_test

import (
	"fmt"
	"testing"

	"managed_identity"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var (
	validSystemAssignedManagedIdentityResource = managed_identity.ManagedIdentityTrackedResource{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Azure.ResourceManager.Models.CommonTypes.ManagedIdentity/managedIdentityTrackedResources/identity", subscriptionIdExpected, resourceGroupExpected)),
		Location: to.Ptr(locationExpected),
		Tags: map[string]*string{
			"tagKey1": to.Ptr("tagValue1"),
		},
		Identity: &managed_identity.ManagedServiceIdentity{
			Type:        to.Ptr(managed_identity.ManagedServiceIdentityType(identityTypeSystemAssigendExpected)),
			PrincipalID: to.Ptr(principalIdExpected),
			TenantID:    to.Ptr(tenantIdExpected),
		},
		Properties: &managed_identity.ManagedIdentityTrackedResourceProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}

	validUserAssignedAndSystemAssignedManagedIdentityResource = managed_identity.ManagedIdentityTrackedResource{
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Azure.ResourceManager.Models.CommonTypes.ManagedIdentity/managedIdentityTrackedResources/identity", subscriptionIdExpected, resourceGroupExpected)),
		Location: to.Ptr(locationExpected),
		Tags: map[string]*string{
			"tagKey1": to.Ptr("tagValue1"),
		},
		Identity: &managed_identity.ManagedServiceIdentity{
			Type: to.Ptr(managed_identity.ManagedServiceIdentityType(identityTypeSystemUserAssignedExpected)),
			UserAssignedIdentities: map[string]*managed_identity.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {
					PrincipalID: to.Ptr(principalIdExpected),
					ClientID:    to.Ptr(clientIdExpected),
				},
			},
			PrincipalID: to.Ptr(principalIdExpected),
			TenantID:    to.Ptr(tenantIdExpected),
		},
		Properties: &managed_identity.ManagedIdentityTrackedResourceProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}
)

func TestManagedIdentityTrackedResourcesClient_Get(t *testing.T) {
	managedIdentityTrackedResourcesClientGetResponse, err := clientFactory.NewManagedIdentityTrackedResourcesClient().Get(ctx, resourceGroupExpected, "identity", nil)
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
	managedIdentityTrackedResourcesClientCreateWithSystemAssignedResponse, err := clientFactory.NewManagedIdentityTrackedResourcesClient().CreateWithSystemAssigned(
		ctx,
		resourceGroupExpected,
		"identity",
		managed_identity.ManagedIdentityTrackedResource{
			Location: to.Ptr(locationExpected),
			Identity: &managed_identity.ManagedServiceIdentity{
				Type: to.Ptr(managed_identity.ManagedServiceIdentityTypeSystemAssigned),
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
	managedIdentityTrackedResourcesClientUpdateWithUserAssignedAndSystemAssignedResponse, err := clientFactory.NewManagedIdentityTrackedResourcesClient().UpdateWithUserAssignedAndSystemAssigned(
		ctx,
		resourceGroupExpected,
		"identity",
		managed_identity.ManagedIdentityTrackedResource{
			Location: to.Ptr(locationExpected),
			Identity: &managed_identity.ManagedServiceIdentity{
				Type: to.Ptr(managed_identity.ManagedServiceIdentityTypeSystemAndUserAssignedV3),
				UserAssignedIdentities: map[string]*managed_identity.UserAssignedIdentity{
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
