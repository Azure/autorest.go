// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package optionalgroup

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func newExplicitClient() *ExplicitClient {
	return NewExplicitClient(nil)
}

func TestExplicitPostOptionalArrayHeader(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalArrayHeader(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalArrayHeader: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestExplicitPostOptionalArrayParameter(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalArrayParameter(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalArrayParameter: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestExplicitPostOptionalArrayProperty(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalArrayProperty(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalArrayProperty: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestExplicitPostOptionalClassParameter(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalClassParameter(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalClassParameter: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestExplicitPostOptionalClassProperty(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalClassProperty(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalClassProperty: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestExplicitPostOptionalIntegerHeader(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalIntegerHeader(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalIntegerHeader: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestExplicitPostOptionalIntegerParameter(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalIntegerParameter(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalIntegerParameter: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestExplicitPostOptionalIntegerProperty(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalIntegerProperty(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalIntegerProperty: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestExplicitPostOptionalStringHeader(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalStringHeader(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalStringHeader: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestExplicitPostOptionalStringParameter(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalStringParameter(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalStringParameter: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestExplicitPostOptionalStringProperty(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalStringProperty(context.Background(), nil)
	if err != nil {
		t.Fatalf("PostOptionalStringProperty: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// TODO the goal of this test is to throw an exception but nils are acceptable for  []strings in go
func TestExplicitPostRequiredArrayHeader(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredArrayHeader(context.Background(), nil, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO the goal of this test is to throw an exception but nils are acceptable for  []strings in go
func TestExplicitPostRequiredArrayParameter(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredArrayParameter(context.Background(), nil, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

func TestExplicitPostRequiredArrayProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredArrayProperty(context.Background(), ArrayWrapper{Value: nil}, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test
func TestExplicitPostRequiredClassParameter(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredClassParameter(context.Background(), Product{}, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

func TestExplicitPostRequiredClassProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredClassProperty(context.Background(), ClassWrapper{}, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredIntegerHeader(t *testing.T) {
	t.Skip("cannot set nil for int32 in Go")
	client := newExplicitClient()
	result, err := client.PostRequiredIntegerHeader(context.Background(), 0, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredIntegerParameter(t *testing.T) {
	t.Skip("cannot set nil for int32 in Go")
	client := newExplicitClient()
	result, err := client.PostRequiredIntegerParameter(context.Background(), 0, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredIntegerProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredIntegerProperty(context.Background(), IntWrapper{}, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredStringHeader(t *testing.T) {
	t.Skip("cannot set nil for string in Go")
	client := newExplicitClient()
	result, err := client.PostRequiredStringHeader(context.Background(), "", nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredStringParameter(t *testing.T) {
	t.Skip("cannot set nil for string in Go")
	client := newExplicitClient()
	result, err := client.PostRequiredStringParameter(context.Background(), "", nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredStringProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredStringProperty(context.Background(), StringWrapper{}, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result but received informaiton in result")
	}
}
