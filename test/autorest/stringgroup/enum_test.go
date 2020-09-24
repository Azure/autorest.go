// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func newEnumClient() EnumOperations {
	return NewEnumClient(NewDefaultClient(nil))
}

func TestEnumGetNotExpandable(t *testing.T) {
	client := newEnumClient()
	result, err := client.GetNotExpandable(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNotExpandable: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, ColorsRedColor.ToPtr())
}

func TestEnumGetReferenced(t *testing.T) {
	client := newEnumClient()
	result, err := client.GetReferenced(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetReferenced: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, ColorsRedColor.ToPtr())
}

func TestEnumGetReferencedConstant(t *testing.T) {
	client := newEnumClient()
	result, err := client.GetReferencedConstant(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetReferencedConstant: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val := "Sample String"
	helpers.DeepEqualOrFatal(t, result.RefColorConstant, &RefColorConstant{Field1: &val})
}

func TestEnumPutNotExpandable(t *testing.T) {
	client := newEnumClient()
	result, err := client.PutNotExpandable(context.Background(), ColorsRedColor, nil)
	if err != nil {
		t.Fatalf("PutNotExpandable: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestEnumPutReferenced(t *testing.T) {
	client := newEnumClient()
	result, err := client.PutReferenced(context.Background(), ColorsRedColor, nil)
	if err != nil {
		t.Fatalf("PutReferenced: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestEnumPutReferencedConstant(t *testing.T) {
	client := newEnumClient()
	val := string(ColorsGreenColor)
	result, err := client.PutReferencedConstant(context.Background(), RefColorConstant{ColorConstant: &val}, nil)
	if err != nil {
		t.Fatalf("PutReferencedConstant: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)

}
