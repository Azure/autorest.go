// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newDictionaryClient() *DictionaryClient {
	return NewDictionaryClient(nil)
}

func TestDictionaryGetEmpty(t *testing.T) {
	client := newDictionaryClient()
	result, err := client.GetEmpty(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	if r := cmp.Diff(result.DictionaryWrapper, DictionaryWrapper{DefaultProgram: map[string]*string{}}); r != "" {
		t.Fatal(r)
	}
}

func TestDictionaryGetNotProvided(t *testing.T) {
	client := newDictionaryClient()
	result, err := client.GetNotProvided(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	if r := cmp.Diff(result.DictionaryWrapper, DictionaryWrapper{}); r != "" {
		t.Fatal(r)
	}
}

func TestDictionaryGetNull(t *testing.T) {
	client := newDictionaryClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	if r := cmp.Diff(result.DictionaryWrapper, DictionaryWrapper{}); r != "" {
		t.Fatal(r)
	}
}

func TestDictionaryGetValid(t *testing.T) {
	client := newDictionaryClient()
	result, err := client.GetValid(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	s1, s2, s3, s4 := "notepad", "mspaint", "excel", ""
	val := DictionaryWrapper{DefaultProgram: map[string]*string{"txt": &s1, "bmp": &s2, "xls": &s3, "exe": &s4, "": nil}}
	if r := cmp.Diff(result.DictionaryWrapper, val); r != "" {
		t.Fatal(r)
	}
}

func TestDictionaryPutEmpty(t *testing.T) {
	client := newDictionaryClient()
	result, err := client.PutEmpty(context.Background(), DictionaryWrapper{DefaultProgram: map[string]*string{}}, nil)
	if err != nil {
		t.Fatalf("PutEmpty: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestDictionaryPutValid(t *testing.T) {
	client := newDictionaryClient()
	s1, s2, s3, s4 := "notepad", "mspaint", "excel", ""
	result, err := client.PutValid(context.Background(), DictionaryWrapper{DefaultProgram: map[string]*string{"txt": &s1, "bmp": &s2, "xls": &s3, "exe": &s4, "": nil}}, nil)
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
