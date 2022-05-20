// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package extenumsgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newPetClient() *PetClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewPetClient(pl)
}

func TestAddPet(t *testing.T) {
	client := newPetClient()
	result, err := client.AddPet(context.Background(), &PetClientAddPetOptions{
		PetParam: &Pet{
			Name: to.Ptr("Retriever"),
		},
	})
	require.NoError(t, err)
	if r := cmp.Diff(result.Pet, Pet{
		Name: to.Ptr("Retriever"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestGetByPetIDExpected(t *testing.T) {
	client := newPetClient()
	result, err := client.GetByPetID(context.Background(), "tommy", nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Pet, Pet{
		DaysOfWeek: to.Ptr(DaysOfWeekExtensibleEnumMonday),
		IntEnum:    to.Ptr(IntEnumOne),
		Name:       to.Ptr("Tommy Tomson"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestGetByPetIDUnexpected(t *testing.T) {
	client := newPetClient()
	result, err := client.GetByPetID(context.Background(), "casper", nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Pet, Pet{
		DaysOfWeek: (*DaysOfWeekExtensibleEnum)(to.Ptr("Weekend")),
		IntEnum:    to.Ptr(IntEnumTwo),
		Name:       to.Ptr("Casper Ghosty"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestGetByPetIDAllowed(t *testing.T) {
	client := newPetClient()
	result, err := client.GetByPetID(context.Background(), "scooby", nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Pet, Pet{
		DaysOfWeek: to.Ptr(DaysOfWeekExtensibleEnumThursday),
		IntEnum:    (*IntEnum)(to.Ptr("2.1")),
		Name:       to.Ptr("Scooby Scarface"),
	}); r != "" {
		t.Fatal(r)
	}
}
