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

func getComplexClient(t *testing.T) *complexgroup.ComplexClient {
	client, err := complexgroup.NewComplexClient(complexgroup.DefaultEndpoint, nil)
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
	client := getComplexClient(t)
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
	client := getComplexClient(t)
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
