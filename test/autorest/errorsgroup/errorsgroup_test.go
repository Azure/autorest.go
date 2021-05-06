// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package errorsgroup

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
	if r := cmp.Diff(result.PetAction, &PetAction{}); r != "" {
		t.Fatal(r)
	}
}

func TestDoSomethingError1(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "jump", nil)
	var sadErr *PetSadError
	if !errors.As(err, &sadErr) {
		t.Fatalf("expected PetSadError: %v", err)
	}
	if r := cmp.Diff(sadErr, &PetSadError{
		PetActionError: PetActionError{
			ErrorMessage: to.StringPtr("casper aint happy"),
			ErrorType:    to.StringPtr("PetSadError"),
		},
		Reason: to.StringPtr("need more treats"),
	}, cmpopts.IgnoreUnexported(PetActionError{})); r != "" {
		t.Fatal(r)
	}
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
	if r := cmp.Diff(hungrErr, &PetHungryOrThirstyError{
		PetSadError: PetSadError{
			PetActionError: PetActionError{
				ErrorMessage: to.StringPtr("scooby is low"),
				ErrorType:    to.StringPtr("PetHungryOrThirstyError"),
			},
			Reason: to.StringPtr("need more everything"),
		},
		HungryOrThirsty: to.StringPtr("hungry and thirsty"),
	}, cmpopts.IgnoreUnexported(PetActionError{})); r != "" {
		t.Fatal(r)
	}
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
	if r := cmp.Diff(actErr, &PetActionError{}, cmpopts.IgnoreUnexported(PetActionError{})); r != "" {
		t.Fatal(r)
	}
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
	if r := cmp.Diff(result.Pet, &Pet{
		Animal: Animal{
			AniType: to.StringPtr("Dog"),
		},
		Name: to.StringPtr("Tommy Tomson"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestGetPetByIDSuccess2(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "django", nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusAccepted {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestGetPetByIDError1(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "coyoteUgly", nil)
	var anfe *AnimalNotFound
	if !errors.As(err, &anfe) {
		t.Fatal("expected AnimalNotFoundError")
	}
	if r := cmp.Diff(anfe, &AnimalNotFound{
		NotFoundErrorBase: NotFoundErrorBase{
			BaseError: BaseError{
				SomeBaseProp: to.StringPtr("problem finding animal"),
			},
			Reason:       to.StringPtr("the type of animal requested is not available"),
			WhatNotFound: to.StringPtr("AnimalNotFound"),
		},
		Name: to.StringPtr("coyote"),
	}, cmpopts.IgnoreUnexported(NotFoundErrorBase{})); r != "" {
		t.Fatal(r)
	}
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
	if r := cmp.Diff(lnfe, &LinkNotFound{
		NotFoundErrorBase: NotFoundErrorBase{
			BaseError: BaseError{
				SomeBaseProp: to.StringPtr("problem finding pet"),
			},
			Reason:       to.StringPtr("link to pet not found"),
			WhatNotFound: to.StringPtr("InvalidResourceLink"),
		},
		WhatSubAddress: to.StringPtr("pet/yourpet was not found"),
	}, cmpopts.IgnoreUnexported(NotFoundErrorBase{})); r != "" {
		t.Fatal(r)
	}
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
	if r := cmp.Diff(err.Error(), "That's all folks!!"); r != "" {
		t.Fatal(r)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}
