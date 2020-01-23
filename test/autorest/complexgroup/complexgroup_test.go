// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"net/http"
	"reflect"
	"testing"
)

func getBasicClient(t *testing.T) *complexgroup.BasicClient {
	client, err := complexgroup.NewBasicClient(complexgroup.DefaultEndpoint, nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client
}

func getPrimitiveClient(t *testing.T) *complexgroup.PrimitiveClient {
	client, err := complexgroup.NewPrimitiveClient(complexgroup.DefaultEndpoint, nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
}

func TestGetValid(t *testing.T) {
	client := getBasicClient(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	var v complexgroup.ColorType
	colors := complexgroup.PossibleColorValues()
	for _, c := range colors {
		if string(c) == "YELLOW" {
			v = c
			break
		}
	}
	i, s := int(2), "abc"
	expected := &complexgroup.GetValidResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{ID: &i, Name: &s, Color: &v},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutValid(t *testing.T) {
	client := getBasicClient(t)
	var v complexgroup.ColorType
	colors := complexgroup.PossibleColorValues()
	for _, c := range colors {
		if string(c) == "Magenta" {
			v = c
			break
		}
	}
	i, s := int(2), "abc"
	result, err := client.PutValid(context.Background(), complexgroup.Basic{ID: &i, Name: &s, Color: &v})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	expected := &complexgroup.PutValidResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

// TODO check this
func TestGetInvalid(t *testing.T) {
	client := getBasicClient(t)
	_, err := client.GetInvalid(context.Background())
	if err == nil {
		t.Fatalf("GetInvalid expected an error")
	}
	// i, s := int(1), "abc"
	// expected := &complexgroup.GetInvalidResponse{
	// 	StatusCode: http.StatusOK,
	// 	Basic:      &complexgroup.Basic{ID: &i, Name: &s},
	// }
	// deepEqualOrFatal(t, result, expected)
}

func TestGetEmpty(t *testing.T) {
	client := getBasicClient(t)
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	expected := &complexgroup.GetEmptyResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{},
	}
	deepEqualOrFatal(t, result, expected)
}

// TODO nil or null?
func TestGetNull(t *testing.T) {
	client := getBasicClient(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	expected := &complexgroup.GetNullResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetNotProvided(t *testing.T) {
	client := getBasicClient(t)
	result, err := client.GetNotProvided(context.Background())
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	expected := &complexgroup.GetNotProvidedResponse{
		StatusCode: http.StatusOK,
		Basic:      nil,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetInt(t *testing.T) {
	client := getPrimitiveClient(t)
	result, err := client.GetInt(context.Background())
	if err != nil {
		t.Fatalf("GetInt: %v", err)
	}
	a, b := int32(-1), int32(2)
	expected := &complexgroup.GetIntResponse{
		StatusCode: http.StatusOK,
		IntWrapper: &complexgroup.IntWrapper{Field1: &a, Field2: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutInt(t *testing.T) {
	client := getPrimitiveClient(t)
	a, b := int32(-1), int32(2)
	result, err := client.PutInt(context.Background(), complexgroup.IntWrapper{Field1: &a, Field2: &b})
	if err != nil {
		t.Fatalf("PutInt: %v", err)
	}
	expected := &complexgroup.PutIntResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
