// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func newArrayClient() complexgroup.ArrayOperations {
	return complexgroup.NewArrayClient(complexgroup.NewDefaultClient(nil))
}

func TestArrayGetEmpty(t *testing.T) {
	client := newArrayClient()
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.ArrayWrapper, &complexgroup.ArrayWrapper{
		Array: &[]string{},
	})
}

func TestArrayGetNotProvided(t *testing.T) {
	client := newArrayClient()
	result, err := client.GetNotProvided(context.Background())
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.ArrayWrapper, &complexgroup.ArrayWrapper{})
}

func TestArrayGetValid(t *testing.T) {
	client := newArrayClient()
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.ArrayWrapper, &complexgroup.ArrayWrapper{
		Array: &[]string{"1, 2, 3, 4", "", "", "&S#$(*Y", "The quick brown fox jumps over the lazy dog"},
	})
}

func TestArrayPutEmpty(t *testing.T) {
	client := newArrayClient()
	result, err := client.PutEmpty(context.Background(), complexgroup.ArrayWrapper{Array: &[]string{}})
	if err != nil {
		t.Fatalf("PutEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

/*
test is currently invalid, missing x-nullable but expects null
func TestArrayPutValid(t *testing.T) {
	client := newArrayClient()
	result, err := client.PutValid(context.Background(), complexgroup.ArrayWrapper{Array: &[]string{"1, 2, 3, 4", "", nil, "&S#$(*Y", "The quick brown fox jumps over the lazy dog"}})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}*/
