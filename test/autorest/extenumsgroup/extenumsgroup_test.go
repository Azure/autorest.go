// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package extenumsgrouptest

import (
	"context"
	"generatortests/autorest/generated/extenumsgroup"
	"generatortests/helpers"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func TestAddPet(t *testing.T) {
	client := extenumsgroup.NewDefaultClient(nil).PetOperations()
	result, err := client.AddPet(context.Background(), &extenumsgroup.PetAddPetOptions{
		PetParam: &extenumsgroup.Pet{
			Name: to.StringPtr("Retriever"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &extenumsgroup.Pet{
		Name: to.StringPtr("Retriever"),
	})
}

func TestGetByPetIDExpected(t *testing.T) {
	client := extenumsgroup.NewDefaultClient(nil).PetOperations()
	result, err := client.GetByPetID(context.Background(), "tommy")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &extenumsgroup.Pet{
		DaysOfWeek: extenumsgroup.DaysOfWeekExtensibleEnumMonday.ToPtr(),
		IntEnum:    extenumsgroup.IntEnumOne.ToPtr(),
		Name:       to.StringPtr("Tommy Tomson"),
	})
}

func TestGetByPetIDUnexpected(t *testing.T) {
	client := extenumsgroup.NewDefaultClient(nil).PetOperations()
	result, err := client.GetByPetID(context.Background(), "casper")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &extenumsgroup.Pet{
		DaysOfWeek: (*extenumsgroup.DaysOfWeekExtensibleEnum)(to.StringPtr("Weekend")),
		IntEnum:    extenumsgroup.IntEnumTwo.ToPtr(),
		Name:       to.StringPtr("Casper Ghosty"),
	})
}

func TestGetByPetIDAllowed(t *testing.T) {
	client := extenumsgroup.NewDefaultClient(nil).PetOperations()
	result, err := client.GetByPetID(context.Background(), "scooby")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &extenumsgroup.Pet{
		DaysOfWeek: extenumsgroup.DaysOfWeekExtensibleEnumThursday.ToPtr(),
		IntEnum:    (*extenumsgroup.IntEnum)(to.StringPtr("2.1")),
		Name:       to.StringPtr("Scooby Scarface"),
	})
}
