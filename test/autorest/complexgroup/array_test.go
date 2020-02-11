// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"net/http"
	"testing"
)

func getArrayOperations(t *testing.T) complexgroup.ArrayOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.ArrayOperations()
}

func TestArrayGetEmpty(t *testing.T) {
	client := getArrayOperations(t)
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	val := complexgroup.ArrayWrapper{Array: []string{}}
	expected := &complexgroup.ArrayGetEmptyResponse{
		StatusCode:   http.StatusOK,
		ArrayWrapper: &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestArrayGetNotProvided(t *testing.T) {
	client := getArrayOperations(t)
	result, err := client.GetNotProvided(context.Background())
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	val := complexgroup.ArrayWrapper{}
	expected := &complexgroup.ArrayGetNotProvidedResponse{
		StatusCode:   http.StatusOK,
		ArrayWrapper: &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestArrayGetValid(t *testing.T) {
	client := getArrayOperations(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	val := complexgroup.ArrayWrapper{Array: []string{"1, 2, 3, 4", "", "", "&S#$(*Y", "The quick brown fox jumps over the lazy dog"}}
	expected := &complexgroup.ArrayGetValidResponse{
		StatusCode:   http.StatusOK,
		ArrayWrapper: &val,
	}
	deepEqualOrFatal(t, result, expected)
}

// TODO this works if the Array field in ArrayWrapper is of type []*string without the omitempty JSON tag
// func TestArrayPutEmpty(t *testing.T) {
// 	client := getArrayOperations(t)
// 	result, err := client.PutEmpty(context.Background(), complexgroup.ArrayWrapper{Array: []*string{}})
// 	if err != nil {
// 		t.Fatalf("PutEmpty: %v", err)
// 	}
// 	expected := &complexgroup.ArrayPutEmptyResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

// TODO this only works if the Array field on ArrayWrapper is of type []*string
// func TestArrayPutValid(t *testing.T) {
// 	client := getArrayOperations(t)
// 	result, err := client.PutValid(context.Background(), complexgroup.ArrayWrapper{Array: []*string{toStrPtr("1, 2, 3, 4"), toStrPtr(""), nil, toStrPtr("&S#$(*Y"), toStrPtr("The quick brown fox jumps over the lazy dog")}})
// 	if err != nil {
// 		t.Fatalf("PutValid: %v", err)
// 	}
// 	expected := &complexgroup.ArrayPutValidResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

// func toStrPtr(s string) *string {
// 	return &s
// }
