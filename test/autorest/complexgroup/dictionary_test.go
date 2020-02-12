// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"net/http"
	"testing"
)

func getDictionaryOperations(t *testing.T) complexgroup.DictionaryOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.DictionaryOperations()
}

func TestDictionaryGetEmpty(t *testing.T) {
	client := getDictionaryOperations(t)
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	val := complexgroup.DictionaryWrapper{DefaultProgram: map[string]*string{}}
	expected := &complexgroup.DictionaryGetEmptyResponse{
		StatusCode:        http.StatusOK,
		DictionaryWrapper: &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestDictionaryGetNotProvided(t *testing.T) {
	client := getDictionaryOperations(t)
	result, err := client.GetNotProvided(context.Background())
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	val := complexgroup.DictionaryWrapper{}
	expected := &complexgroup.DictionaryGetNotProvidedResponse{
		StatusCode:        http.StatusOK,
		DictionaryWrapper: &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestDictionaryGetNull(t *testing.T) {
	client := getDictionaryOperations(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	expected := &complexgroup.DictionaryGetNullResponse{
		StatusCode:        http.StatusOK,
		DictionaryWrapper: &complexgroup.DictionaryWrapper{},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestDictionaryGetValid(t *testing.T) {
	client := getDictionaryOperations(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	s1, s2, s3, s4 := "notepad", "mspaint", "excel", ""
	val := complexgroup.DictionaryWrapper{DefaultProgram: map[string]*string{"txt": &s1, "bmp": &s2, "xls": &s3, "exe": &s4, "": nil}}
	expected := &complexgroup.DictionaryGetValidResponse{
		StatusCode:        http.StatusOK,
		DictionaryWrapper: &val,
	}
	deepEqualOrFatal(t, result, expected)
}

// // TODO this works if the DefaultProgram field in DictionaryWrapper is of type map[string]*string without the omitempty JSON tag
// func TestDictionaryPutEmpty(t *testing.T) {
// 	client := getDictionaryOperations(t)
// 	result, err := client.PutEmpty(context.Background(), complexgroup.DictionaryWrapper{DefaultProgram: map[string]*string{}})
// 	if err != nil {
// 		t.Fatalf("PutEmpty: %v", err)
// 	}
// 	expected := &complexgroup.DictionaryPutEmptyResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

func TestDictionaryPutValid(t *testing.T) {
	client := getDictionaryOperations(t)
	s1, s2, s3, s4 := "notepad", "mspaint", "excel", ""
	result, err := client.PutValid(context.Background(), complexgroup.DictionaryWrapper{DefaultProgram: map[string]*string{"txt": &s1, "bmp": &s2, "xls": &s3, "exe": &s4, "": nil}})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	expected := &complexgroup.DictionaryPutValidResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func toStrPtr(s string) *string {
	return &s
}
