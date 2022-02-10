// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup

import (
	"context"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newStringClient() *StringClient {
	return NewStringClient(nil)
}

func TestStringGetMBCS(t *testing.T) {
	client := newStringClient()
	result, err := client.GetMBCS(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetMBCS: %v", err)
	}
	if r := cmp.Diff(result.Value, to.StringPtr("啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€")); r != "" {
		t.Fatal(r)
	}
}

func TestStringPutMBCS(t *testing.T) {
	client := newStringClient()
	result, err := client.PutMBCS(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutMBCS: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestStringGetBase64Encoded(t *testing.T) {
	client := newStringClient()
	result, err := client.GetBase64Encoded(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBase64Encoded: %v", err)
	}
	val := []byte("a string that gets encoded with base64")
	if r := cmp.Diff(result.Value, val); r != "" {
		t.Fatal(r)
	}
}

func TestStringGetBase64URLEncoded(t *testing.T) {
	client := newStringClient()
	result, err := client.GetBase64URLEncoded(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBase64URLEncoded: %v", err)
	}
	if r := cmp.Diff(result.Value, []byte("a string that gets encoded with base64url")); r != "" {
		t.Fatal(r)
	}
}

func TestStringGetEmpty(t *testing.T) {
	client := newStringClient()
	result, err := client.GetEmpty(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	if r := cmp.Diff(result.Value, to.StringPtr("")); r != "" {
		t.Fatal(r)
	}
}

func TestStringGetNotProvided(t *testing.T) {
	client := newStringClient()
	result, err := client.GetNotProvided(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestStringGetNull(t *testing.T) {
	client := newStringClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestStringGetNullBase64URLEncoded(t *testing.T) {
	client := newStringClient()
	result, err := client.GetNullBase64URLEncoded(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNullBase64URLEncoded: %v", err)
	}
	if r := cmp.Diff(result.Value, ([]byte)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestStringGetWhitespace(t *testing.T) {
	client := newStringClient()
	result, err := client.GetWhitespace(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetWhitespace: %v", err)
	}
	if r := cmp.Diff(result.Value, to.StringPtr("    Now is the time for all good men to come to the aid of their country    ")); r != "" {
		t.Fatal(r)
	}
}

func TestStringPutBase64URLEncoded(t *testing.T) {
	client := newStringClient()
	result, err := client.PutBase64URLEncoded(context.Background(), []byte("a string that gets encoded with base64url"), nil)
	if err != nil {
		t.Fatalf("PutBase64URLEncoded: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestStringPutEmpty(t *testing.T) {
	client := newStringClient()
	result, err := client.PutEmpty(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutEmpty: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestStringPutNull(t *testing.T) {
	client := newStringClient()
	result, err := client.PutNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutNull: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestStringPutWhitespace(t *testing.T) {
	client := newStringClient()
	result, err := client.PutWhitespace(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutWhitespace: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
