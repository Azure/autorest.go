// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"net/http"
	"testing"
)

func getBasicOperations(t *testing.T) complexgroup.BasicOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.BasicOperations()
}

func TestBasicGetValid(t *testing.T) {
	client := getBasicOperations(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	var v complexgroup.CMYKColors
	colors := complexgroup.PossibleCMYKColorsValues()
	for _, c := range colors {
		if string(c) == "YELLOW" {
			v = c
			break
		}
	}
	i, s := int32(2), "abc"
	expected := &complexgroup.BasicGetValidResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{ID: &i, Name: &s, Color: &v},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestBasicPutValid(t *testing.T) {
	client := getBasicOperations(t)
	var v complexgroup.CMYKColors
	colors := complexgroup.PossibleCMYKColorsValues()
	for _, c := range colors {
		if string(c) == "Magenta" {
			v = c
			break
		}
	}
	i, s := int32(2), "abc"
	result, err := client.PutValid(context.Background(), complexgroup.Basic{ID: &i, Name: &s, Color: &v})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	expected := &complexgroup.BasicPutValidResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

// TODO check this
func TestBasicGetInvalid(t *testing.T) {
	client := getBasicOperations(t)
	result, err := client.GetInvalid(context.Background())
	if err == nil {
		t.Fatalf("GetInvalid expected an error")
	}
	var expected *complexgroup.BasicGetInvalidResponse
	expected = nil
	deepEqualOrFatal(t, result, expected)
}

func TestBasicGetEmpty(t *testing.T) {
	client := getBasicOperations(t)
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	expected := &complexgroup.BasicGetEmptyResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestBasicGetNull(t *testing.T) {
	client := getBasicOperations(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	expected := &complexgroup.BasicGetNullResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestBasicGetNotProvided(t *testing.T) {
	client := getBasicOperations(t)
	result, err := client.GetNotProvided(context.Background())
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	expected := &complexgroup.BasicGetNotProvidedResponse{
		StatusCode: http.StatusOK,
		Basic:      nil,
	}
	deepEqualOrFatal(t, result, expected)
}
