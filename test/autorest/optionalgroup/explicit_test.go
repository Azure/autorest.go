// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package optionalgrouptest

import (
	"context"
	"generatortests/autorest/generated/optionalgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getExplicitClient(t *testing.T) optionalgroup.ExplicitOperations {
	client, err := optionalgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create explicit client: %v", err)
	}
	return client.ExplicitOperations()
}

func TestExplicitPostOptionalArrayHeader(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalArrayHeader(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalArrayHeader: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestExplicitPostOptionalArrayParameter(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalArrayParameter(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalArrayParameter: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestExplicitPostOptionalArrayProperty(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalArrayProperty(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalArrayProperty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestExplicitPostOptionalClassParameter(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalClassParameter(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalClassParameter: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestExplicitPostOptionalClassProperty(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalClassProperty(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalClassProperty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestExplicitPostOptionalIntegerHeader(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalIntegerHeader(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalIntegerHeader: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestExplicitPostOptionalIntegerParameter(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalIntegerParameter(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalIntegerParameter: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestExplicitPostOptionalIntegerProperty(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalIntegerProperty(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalIntegerProperty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestExplicitPostOptionalStringHeader(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalStringHeader(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalStringHeader: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestExplicitPostOptionalStringParameter(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalStringParameter(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalStringParameter: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestExplicitPostOptionalStringProperty(t *testing.T) {
	client := getExplicitClient(t)
	result, err := client.PostOptionalStringProperty(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalStringProperty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// TODO the goal of this test is to throw an exception but nils are acceptable for  []strings in go
func TestExplicitPostRequiredArrayHeader(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := getExplicitClient(t)
	result, err := client.PostRequiredArrayHeader(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO the goal of this test is to throw an exception but nils are acceptable for  []strings in go
func TestExplicitPostRequiredArrayParameter(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := getExplicitClient(t)
	result, err := client.PostRequiredArrayParameter(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

func TestExplicitPostRequiredArrayProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := getExplicitClient(t)
	result, err := client.PostRequiredArrayProperty(context.Background(), optionalgroup.ArrayWrapper{Value: nil})
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test
func TestExplicitPostRequiredClassParameter(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := getExplicitClient(t)
	result, err := client.PostRequiredClassParameter(context.Background(), optionalgroup.Product{})
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

func TestExplicitPostRequiredClassProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := getExplicitClient(t)
	result, err := client.PostRequiredClassProperty(context.Background(), optionalgroup.ClassWrapper{})
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredIntegerHeader(t *testing.T) {
	t.Skip("cannot set nil for int32 in Go")
	client := getExplicitClient(t)
	result, err := client.PostRequiredIntegerHeader(context.Background(), 0)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredIntegerParameter(t *testing.T) {
	t.Skip("cannot set nil for int32 in Go")
	client := getExplicitClient(t)
	result, err := client.PostRequiredIntegerParameter(context.Background(), 0)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredIntegerProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := getExplicitClient(t)
	result, err := client.PostRequiredIntegerProperty(context.Background(), optionalgroup.IntWrapper{})
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredStringHeader(t *testing.T) {
	t.Skip("cannot set nil for string in Go")
	client := getExplicitClient(t)
	result, err := client.PostRequiredStringHeader(context.Background(), "")
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredStringParameter(t *testing.T) {
	t.Skip("cannot set nil for string in Go")
	client := getExplicitClient(t)
	result, err := client.PostRequiredStringParameter(context.Background(), "")
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredStringProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := getExplicitClient(t)
	result, err := client.PostRequiredStringProperty(context.Background(), optionalgroup.StringWrapper{})
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}
