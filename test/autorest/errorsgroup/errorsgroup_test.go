// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package errorsgrouptest

import (
	"context"
	"generatortests/autorest/generated/errorsgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newPetClient() errorsgroup.PetOperations {
	options := errorsgroup.DefaultClientOptions()
	options.Retry.MaxRetryDelay = 20 * time.Millisecond
	client := errorsgroup.NewClient("http://localhost:3000", &options)
	return errorsgroup.NewPetClient(client)
}

// DoSomething - Asks pet to do something
func TestDoSomethingSuccess(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "stay")
	if err != nil {
		t.Fatal(err)
	}
	// bug in test server, route returns wrong JSON model so PetAction is empty
	helpers.DeepEqualOrFatal(t, result.PetAction, &errorsgroup.PetAction{})
}

func TestDoSomethingError1(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "jump")
	sadErr, ok := err.(*errorsgroup.PetSadError)
	if !ok {
		t.Fatalf("expected PetSadError: %v", err)
	}
	helpers.DeepEqualOrFatal(t, sadErr, &errorsgroup.PetSadError{
		PetActionError: errorsgroup.PetActionError{
			ErrorMessage: to.StringPtr("casper aint happy"),
			ErrorType:    to.StringPtr("PetSadError"),
		},
		Reason: to.StringPtr("need more treats"),
	})
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestDoSomethingError2(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "fetch")
	hungrErr, ok := err.(*errorsgroup.PetHungryOrThirstyError)
	if !ok {
		t.Fatal("expected PetHungryOrThirstyError")
	}
	helpers.DeepEqualOrFatal(t, hungrErr, &errorsgroup.PetHungryOrThirstyError{
		PetSadError: errorsgroup.PetSadError{
			PetActionError: errorsgroup.PetActionError{
				ErrorMessage: to.StringPtr("scooby is low"),
				ErrorType:    to.StringPtr("PetHungryOrThirstyError"),
			},
			Reason: to.StringPtr("need more everything"),
		},
		HungryOrThirsty: to.StringPtr("hungry and thirsty"),
	})
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestDoSomethingError3(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "unknown")
	actErr, ok := err.(*errorsgroup.PetActionError)
	if !ok {
		t.Fatal("expected PetActionError")
	}
	helpers.DeepEqualOrFatal(t, actErr, &errorsgroup.PetActionError{})
	if result != nil {
		t.Fatal("expected nil result")
	}
}

// GetPetByID - Gets pets by id.
func TestGetPetByIDSuccess1(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "tommy")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &errorsgroup.Pet{
		Animal: errorsgroup.Animal{
			AniType: to.StringPtr("Dog"),
		},
		Name: to.StringPtr("Tommy Tomson"),
	})
}

func TestGetPetByIDSuccess2(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "django")
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusAccepted)
}

func TestGetPetByIDError1(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "coyoteUgly")
	anfe, ok := err.(*errorsgroup.AnimalNotFound)
	if !ok {
		t.Fatal("expected AnimalNotFoundError")
	}
	helpers.DeepEqualOrFatal(t, anfe, &errorsgroup.AnimalNotFound{
		NotFoundErrorBase: errorsgroup.NotFoundErrorBase{
			BaseError: errorsgroup.BaseError{
				SomeBaseProp: to.StringPtr("problem finding animal"),
			},
			Reason:       to.StringPtr("the type of animal requested is not available"),
			WhatNotFound: to.StringPtr("AnimalNotFound"),
		},
		Name: to.StringPtr("coyote"),
	})
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestGetPetByIDError2(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "weirdAlYankovic")
	lnfe, ok := err.(*errorsgroup.LinkNotFound)
	if !ok {
		t.Fatal("expected LinkNotFoundError")
	}
	helpers.DeepEqualOrFatal(t, lnfe, &errorsgroup.LinkNotFound{
		NotFoundErrorBase: errorsgroup.NotFoundErrorBase{
			BaseError: errorsgroup.BaseError{
				SomeBaseProp: to.StringPtr("problem finding pet"),
			},
			Reason:       to.StringPtr("link to pet not found"),
			WhatNotFound: to.StringPtr("InvalidResourceLink"),
		},
		WhatSubAddress: to.StringPtr("pet/yourpet was not found"),
	})
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestGetPetByIDError3(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "ringo")
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	const expected = "ringo is missing"
	if e := err.Error(); e != expected {
		t.Fatalf("expected %s, got %s", expected, e)
	}
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestGetPetByIDError4(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "alien123")
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	expected := "123"
	if e := err.Error(); e != expected {
		t.Fatalf("expected %s, got %s", expected, e)
	}
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestGetPetByIDError5(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "unknown")
	// default generic error (no schema)
	helpers.DeepEqualOrFatal(t, err.Error(), "That's all folks!!")
	if result != nil {
		t.Fatal("expected nil result")
	}
}
