// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package errorsgroup

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newPetClient() *PetClient {
	options := azcore.ClientOptions{}
	options.Retry.MaxRetryDelay = 20 * time.Millisecond
	return NewPetClient(&options)
}

// DoSomething - Asks pet to do something
func TestDoSomethingSuccess(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "stay", nil)
	if err != nil {
		t.Fatal(err)
	}
	// bug in test server, route returns wrong JSON model so PetAction is empty
	if r := cmp.Diff(result.PetAction, PetAction{}); r != "" {
		t.Fatal(r)
	}
}

func TestDoSomethingError1(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "jump", nil)
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatalf("expected azcore.ResponseError: %v", err)
	}
	const want = `POST http://localhost:3000/errorStatusCodes/Pets/doSomething/jump
--------------------------------------------------------------------------------
RESPONSE 500: 500 Internal Server Error
ERROR CODE UNAVAILABLE
--------------------------------------------------------------------------------
{
  "actionResponse": "grrrr",
  "errorType": "PetSadError",
  "errorMessage": "casper aint happy",
  "reason": "need more treats"
}
--------------------------------------------------------------------------------
`
	if got := respErr.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestDoSomethingError2(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "fetch", nil)
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatalf("expected azcore.ResponseError: %v", err)
	}
	const want = `POST http://localhost:3000/errorStatusCodes/Pets/doSomething/fetch
--------------------------------------------------------------------------------
RESPONSE 404: 404 Not Found
ERROR CODE UNAVAILABLE
--------------------------------------------------------------------------------
{
  "actionResponse": "howl",
  "errorType": "PetHungryOrThirstyError",
  "errorMessage": "scooby is low",
  "reason": "need more everything",
  "hungryOrThirsty": "hungry and thirsty"
}
--------------------------------------------------------------------------------
`
	if got := respErr.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestDoSomethingError3(t *testing.T) {
	client := newPetClient()
	result, err := client.DoSomething(context.Background(), "unknown", nil)
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatalf("expected azcore.ResponseError: %v", err)
	}
	const want = `POST http://localhost:3000/errorStatusCodes/Pets/doSomething/unknown
--------------------------------------------------------------------------------
RESPONSE 400: 400 Bad Request
ERROR CODE UNAVAILABLE
--------------------------------------------------------------------------------
{
  "message": "Action cannot be performed unknown",
  "status": 400
}
--------------------------------------------------------------------------------
`
	if got := respErr.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
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
	if r := cmp.Diff(result.Pet, Pet{
		AniType: to.StringPtr("Dog"),
		Name:    to.StringPtr("Tommy Tomson"),
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
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestGetPetByIDError1(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "coyoteUgly", nil)
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatalf("expected azcore.ResponseError: %v", err)
	}
	const want = `GET http://localhost:3000/errorStatusCodes/Pets/coyoteUgly/GetPet
--------------------------------------------------------------------------------
RESPONSE 404: 404 Not Found
ERROR CODE UNAVAILABLE
--------------------------------------------------------------------------------
{
  "someBaseProp": "problem finding animal",
  "reason": "the type of animal requested is not available",
  "name": "coyote",
  "whatNotFound": "AnimalNotFound"
}
--------------------------------------------------------------------------------
`
	if got := respErr.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestGetPetByIDError2(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "weirdAlYankovic", nil)
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatalf("expected azcore.ResponseError: %v", err)
	}
	const want = `GET http://localhost:3000/errorStatusCodes/Pets/weirdAlYankovic/GetPet
--------------------------------------------------------------------------------
RESPONSE 404: 404 Not Found
ERROR CODE UNAVAILABLE
--------------------------------------------------------------------------------
{
  "someBaseProp": "problem finding pet",
  "reason": "link to pet not found",
  "whatSubAddress": "pet/yourpet was not found",
  "whatNotFound": "InvalidResourceLink"
}
--------------------------------------------------------------------------------
`
	if got := respErr.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
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
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatalf("expected azcore.ResponseError: %v", err)
	}
	const want = `GET http://localhost:3000/errorStatusCodes/Pets/ringo/GetPet
--------------------------------------------------------------------------------
RESPONSE 400: 400 Bad Request
ERROR CODE UNAVAILABLE
--------------------------------------------------------------------------------
"ringo is missing"
--------------------------------------------------------------------------------
`
	if got := respErr.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
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
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatalf("expected azcore.ResponseError: %v", err)
	}
	const want = `GET http://localhost:3000/errorStatusCodes/Pets/alien123/GetPet
--------------------------------------------------------------------------------
RESPONSE 501: 501 Not Implemented
ERROR CODE UNAVAILABLE
--------------------------------------------------------------------------------
123
--------------------------------------------------------------------------------
`
	if got := respErr.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestGetPetByIDError5(t *testing.T) {
	client := newPetClient()
	result, err := client.GetPetByID(context.Background(), "unknown", nil)
	// default generic error (no schema)
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatalf("expected azcore.ResponseError: %v", err)
	}
	const want = `GET http://localhost:3000/errorStatusCodes/Pets/unknown/GetPet
--------------------------------------------------------------------------------
RESPONSE 402: 402 Payment Required
ERROR CODE UNAVAILABLE
--------------------------------------------------------------------------------
That's all folks!!
--------------------------------------------------------------------------------
`
	if got := respErr.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}
