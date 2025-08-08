// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package hierarchygroup_test

import (
	"context"
	"hierarchygroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestHierarchyBuildingClient_UpdatePet(t *testing.T) {
	client, err := hierarchygroup.NewHierarchyBuildingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.UpdatePet(context.Background(), &hierarchygroup.Pet{
		Kind:    to.Ptr("pet"),
		Name:    to.Ptr("Buddy"),
		Trained: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Equal(t, &hierarchygroup.Pet{
		Kind:    to.Ptr("pet"),
		Name:    to.Ptr("Buddy"),
		Trained: to.Ptr(true),
	}, resp.AnimalClassification)
}

func TestHierarchyBuildingClient_UpdateDog(t *testing.T) {
	client, err := hierarchygroup.NewHierarchyBuildingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.UpdateDog(context.Background(), &hierarchygroup.Dog{
		Kind:    to.Ptr("dog"),
		Name:    to.Ptr("Rex"),
		Trained: to.Ptr(true),
		Breed:   to.Ptr("German Shepherd"),
	}, nil)
	require.NoError(t, err)
	require.Equal(t, &hierarchygroup.Dog{
		Kind:    to.Ptr("dog"),
		Name:    to.Ptr("Rex"),
		Trained: to.Ptr(true),
		Breed:   to.Ptr("German Shepherd"),
	}, resp.AnimalClassification)
}
