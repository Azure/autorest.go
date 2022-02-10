// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newEnumClient() *EnumClient {
	return NewEnumClient(nil)
}

func TestEnumGetNotExpandable(t *testing.T) {
	client := newEnumClient()
	result, err := client.GetNotExpandable(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNotExpandable: %v", err)
	}
	if r := cmp.Diff(result.Value, ColorsRedColor.ToPtr()); r != "" {
		t.Fatal(r)
	}
}

func TestEnumGetReferenced(t *testing.T) {
	client := newEnumClient()
	result, err := client.GetReferenced(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetReferenced: %v", err)
	}
	if r := cmp.Diff(result.Value, ColorsRedColor.ToPtr()); r != "" {
		t.Fatal(r)
	}
}

func TestEnumGetReferencedConstant(t *testing.T) {
	client := newEnumClient()
	result, err := client.GetReferencedConstant(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetReferencedConstant: %v", err)
	}
	val := "Sample String"
	if r := cmp.Diff(result.RefColorConstant, RefColorConstant{Field1: &val}); r != "" {
		t.Fatal(r)
	}
}

func TestEnumPutNotExpandable(t *testing.T) {
	client := newEnumClient()
	result, err := client.PutNotExpandable(context.Background(), ColorsRedColor, nil)
	if err != nil {
		t.Fatalf("PutNotExpandable: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestEnumPutReferenced(t *testing.T) {
	client := newEnumClient()
	result, err := client.PutReferenced(context.Background(), ColorsRedColor, nil)
	if err != nil {
		t.Fatalf("PutReferenced: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestEnumPutReferencedConstant(t *testing.T) {
	client := newEnumClient()
	val := string(ColorsGreenColor)
	result, err := client.PutReferencedConstant(context.Background(), RefColorConstant{ColorConstant: &val}, nil)
	if err != nil {
		t.Fatalf("PutReferencedConstant: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}

}
