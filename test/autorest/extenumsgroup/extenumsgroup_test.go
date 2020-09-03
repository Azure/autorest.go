// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package extenumsgroup

import (
	"context"
	"generatortests/helpers"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newPetClient() PetOperations {
	return NewPetClient(NewDefaultClient(nil))
}

func TestAddPet(t *testing.T) {
	client := newPetClient()
	result, err := client.AddPet(context.Background(), &PetAddPetOptions{
		PetParam: &Pet{
			Name: to.StringPtr("Retriever"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &Pet{
		Name: to.StringPtr("Retriever"),
	})
}

func TestGetByPetIDExpected(t *testing.T) {
	client := newPetClient()
	result, err := client.GetByPetID(context.Background(), "tommy")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &Pet{
		DaysOfWeek: DaysOfWeekExtensibleEnumMonday.ToPtr(),
		IntEnum:    IntEnumOne.ToPtr(),
		Name:       to.StringPtr("Tommy Tomson"),
	})
}

func TestGetByPetIDUnexpected(t *testing.T) {
	client := newPetClient()
	result, err := client.GetByPetID(context.Background(), "casper")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &Pet{
		DaysOfWeek: (*DaysOfWeekExtensibleEnum)(to.StringPtr("Weekend")),
		IntEnum:    IntEnumTwo.ToPtr(),
		Name:       to.StringPtr("Casper Ghosty"),
	})
}

func TestGetByPetIDAllowed(t *testing.T) {
	client := newPetClient()
	result, err := client.GetByPetID(context.Background(), "scooby")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &Pet{
		DaysOfWeek: DaysOfWeekExtensibleEnumThursday.ToPtr(),
		IntEnum:    (*IntEnum)(to.StringPtr("2.1")),
		Name:       to.StringPtr("Scooby Scarface"),
	})
}
