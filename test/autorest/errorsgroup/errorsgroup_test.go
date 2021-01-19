// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package errorsgroup

import (
	"context"
	"errors"
	"generatortests/helpers"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newPetClient() *PetClient {
	options := ConnectionOptions{}
	options.Retry.MaxRetryDelay = 20 * time.Millisecond
	client := NewConnection("http://localhost:3000", &options)
	return NewPetClient(client)
}

// DoSomething - Asks pet to do something
func TestDoSomethingSuccess(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "stay", nil)
	if err != nil {
		t.Fatal(err)
	}
	// bug in test server, route returns wrong JSON model so PetAction is empty
	helpers.DeepEqualOrFatal(t, result.PetAction, &PetAction{})
}

func TestDoSomethingError1(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "jump", nil)
	var sadErr *PetSadError
	if !errors.As(err, &sadErr) {
		t.Fatalf("expected PetSadError: %v", err)
	}
	helpers.DeepEqualOrFatal(t, sadErr, &PetSadError{
		PetActionError: PetActionError{
			ErrorMessage: to.StringPtr("casper aint happy"),
			ErrorType:    to.StringPtr("PetSadError"),
		},
		Reason: to.StringPtr("need more treats"),
	})
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestDoSomethingError2(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "fetch", nil)
	var hungrErr *PetHungryOrThirstyError
	if !errors.As(err, &hungrErr) {
		t.Fatal("expected PetHungryOrThirstyError")
	}
	helpers.DeepEqualOrFatal(t, hungrErr, &PetHungryOrThirstyError{
		PetSadError: PetSadError{
			PetActionError: PetActionError{
				ErrorMessage: to.StringPtr("scooby is low"),
				ErrorType:    to.StringPtr("PetHungryOrThirstyError"),
			},
			Reason: to.StringPtr("need more everything"),
		},
		HungryOrThirsty: to.StringPtr("hungry and thirsty"),
	})
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestDoSomethingError3(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "unknown", nil)
	var actErr *PetActionError
	if !errors.As(err, &actErr) {
		t.Fatal("expected PetActionError")
	}
	helpers.DeepEqualOrFatal(t, actErr, &PetActionError{})
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetPetByID - Gets pets by id.
func TestGetPetByIDSuccess1(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "tommy", nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Pet, &Pet{
		Animal: Animal{
			AniType: to.StringPtr("Dog"),
		},
		Name: to.StringPtr("Tommy Tomson"),
	})
}

func TestGetPetByIDSuccess2(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "django", nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusAccepted)
}

func TestGetPetByIDError1(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "coyoteUgly", nil)
	var anfe *AnimalNotFound
	if !errors.As(err, &anfe) {
		t.Fatal("expected AnimalNotFoundError")
	}
	helpers.DeepEqualOrFatal(t, anfe, &AnimalNotFound{
		NotFoundErrorBase: NotFoundErrorBase{
			BaseError: BaseError{
				SomeBaseProp: to.StringPtr("problem finding animal"),
			},
			Reason:       to.StringPtr("the type of animal requested is not available"),
			WhatNotFound: to.StringPtr("AnimalNotFound"),
		},
		Name: to.StringPtr("coyote"),
	})
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestGetPetByIDError2(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "weirdAlYankovic", nil)
	var lnfe *LinkNotFound
	if !errors.As(err, &lnfe) {
		t.Fatal("expected LinkNotFoundError")
	}
	helpers.DeepEqualOrFatal(t, lnfe, &LinkNotFound{
		NotFoundErrorBase: NotFoundErrorBase{
			BaseError: BaseError{
				SomeBaseProp: to.StringPtr("problem finding pet"),
			},
			Reason:       to.StringPtr("link to pet not found"),
			WhatNotFound: to.StringPtr("InvalidResourceLink"),
		},
		WhatSubAddress: to.StringPtr("pet/yourpet was not found"),
	})
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestGetPetByIDError3(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "ringo", nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	const expected = "ringo is missing"
	if e := err.Error(); e != expected {
		t.Fatalf("expected %s, got %s", expected, e)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestGetPetByIDError4(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "alien123", nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	expected := "123"
	if e := err.Error(); e != expected {
		t.Fatalf("expected %s, got %s", expected, e)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestGetPetByIDError5(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "unknown", nil)
	// default generic error (no schema)
	helpers.DeepEqualOrFatal(t, err.Error(), "That's all folks!!")
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}
