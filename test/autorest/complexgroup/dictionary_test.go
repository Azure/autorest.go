// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func newDictionaryClient() DictionaryOperations {
	return NewDictionaryClient(NewDefaultClient(nil))
}

func TestDictionaryGetEmpty(t *testing.T) {
	client := newDictionaryClient()
	result, err := client.GetEmpty(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.DictionaryWrapper, &DictionaryWrapper{DefaultProgram: &map[string]string{}})
}

func TestDictionaryGetNotProvided(t *testing.T) {
	client := newDictionaryClient()
	result, err := client.GetNotProvided(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.DictionaryWrapper, &DictionaryWrapper{})
}

func TestDictionaryGetNull(t *testing.T) {
	client := newDictionaryClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.DictionaryWrapper, &DictionaryWrapper{})
}

/*
test is invalid, expects null values but missing x-nullable
func TestDictionaryGetValid(t *testing.T) {
	client := newDictionaryClient()
	result, err := client.GetValid(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	s1, s2, s3, s4 := "notepad", "mspaint", "excel", ""
	val := DictionaryWrapper{DefaultProgram: &map[string]string{"txt": s1, "bmp": s2, "xls": s3, "exe": s4, "": nil}}
	helpers.DeepEqualOrFatal(t, result.DictionaryWrapper, &val)
}*/

func TestDictionaryPutEmpty(t *testing.T) {
	client := newDictionaryClient()
	result, err := client.PutEmpty(context.Background(), DictionaryWrapper{DefaultProgram: &map[string]string{}}, nil)
	if err != nil {
		t.Fatalf("PutEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

/*
test is invalid, expects null values but missing x-nullable
func TestDictionaryPutValid(t *testing.T) {
	client := newDictionaryClient()
	s1, s2, s3, s4 := "notepad", "mspaint", "excel", ""
	result, err := client.PutValid(context.Background(), DictionaryWrapper{DefaultProgram: &map[string]string{"txt": s1, "bmp": s2, "xls": s3, "exe": s4, "": nil}})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}*/
