// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package extenumsgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newPetClient() *PetClient {
	return NewPetClient(nil)
}

func TestAddPet(t *testing.T) {
	client := newPetClient()
	result, err := client.AddPet(context.Background(), &PetClientAddPetOptions{
		PetParam: &Pet{
			Name: to.StringPtr("Retriever"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Pet, Pet{
		Name: to.StringPtr("Retriever"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestGetByPetIDExpected(t *testing.T) {
	client := newPetClient()
	result, err := client.GetByPetID(context.Background(), "tommy", nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Pet, Pet{
		DaysOfWeek: DaysOfWeekExtensibleEnumMonday.ToPtr(),
		IntEnum:    IntEnumOne.ToPtr(),
		Name:       to.StringPtr("Tommy Tomson"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestGetByPetIDUnexpected(t *testing.T) {
	client := newPetClient()
	result, err := client.GetByPetID(context.Background(), "casper", nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Pet, Pet{
		DaysOfWeek: (*DaysOfWeekExtensibleEnum)(to.StringPtr("Weekend")),
		IntEnum:    IntEnumTwo.ToPtr(),
		Name:       to.StringPtr("Casper Ghosty"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestGetByPetIDAllowed(t *testing.T) {
	client := newPetClient()
	result, err := client.GetByPetID(context.Background(), "scooby", nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Pet, Pet{
		DaysOfWeek: DaysOfWeekExtensibleEnumThursday.ToPtr(),
		IntEnum:    (*IntEnum)(to.StringPtr("2.1")),
		Name:       to.StringPtr("Scooby Scarface"),
	}); r != "" {
		t.Fatal(r)
	}
}
