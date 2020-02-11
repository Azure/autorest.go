// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgrouptest

import (
	"context"
	"generatortests/autorest/generated/stringgroup"
	"net/http"
	"testing"
)

func getEnumClient(t *testing.T) stringgroup.EnumOperations {
	client, err := stringgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create enum client: %v", err)
	}
	return client.EnumOperations()
}

func TestEnumGetNotExpandable(t *testing.T) {
	client := getEnumClient(t)
	result, err := client.GetNotExpandable(context.Background())
	if err != nil {
		t.Fatalf("GetNotExpandable: %v", err)
	}
	color := stringgroup.ColorsRedColor
	expected := &stringgroup.EnumGetNotExpandableResponse{
		StatusCode: http.StatusOK,
		Value:      &color,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestEnumGetReferenced(t *testing.T) {
	client := getEnumClient(t)
	result, err := client.GetReferenced(context.Background())
	if err != nil {
		t.Fatalf("GetReferenced: %v", err)
	}
	color := stringgroup.ColorsRedColor
	expected := &stringgroup.EnumGetReferencedResponse{
		StatusCode: http.StatusOK,
		Value:      &color,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestEnumGetReferencedConstant(t *testing.T) {
	client := getEnumClient(t)
	result, err := client.GetReferencedConstant(context.Background())
	if err != nil {
		t.Fatalf("GetReferencedConstant: %v", err)
	}
	val := "Sample String"
	expected := &stringgroup.EnumGetReferencedConstantResponse{
		StatusCode:       http.StatusOK,
		RefColorConstant: &stringgroup.RefColorConstant{Field1: &val},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestEnumPutNotExpandable(t *testing.T) {
	client := getEnumClient(t)
	result, err := client.PutNotExpandable(context.Background(), stringgroup.ColorsRedColor)
	if err != nil {
		t.Fatalf("PutNotExpandable: %v", err)
	}
	expected := &stringgroup.EnumPutNotExpandableResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestEnumPutReferenced(t *testing.T) {
	client := getEnumClient(t)
	result, err := client.PutReferenced(context.Background(), stringgroup.ColorsRedColor)
	if err != nil {
		t.Fatalf("PutReferenced: %v", err)
	}
	expected := &stringgroup.EnumPutReferencedResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestEnumPutReferencedConstant(t *testing.T) {
	client := getEnumClient(t)
	val := string(stringgroup.ColorsGreenColor)
	result, err := client.PutReferencedConstant(context.Background(), stringgroup.RefColorConstant{ColorConstant: &val})
	if err != nil {
		t.Fatalf("PutReferencedConstant: %v", err)
	}
	expected := &stringgroup.EnumPutReferencedConstantResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
