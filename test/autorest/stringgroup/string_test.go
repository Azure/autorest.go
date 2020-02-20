// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgrouptest

import (
	"context"
	"generatortests/autorest/generated/stringgroup"
	"net/http"
	"testing"
)

func getStringClient(t *testing.T) stringgroup.StringOperations {
	client, err := stringgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create string client: %v", err)
	}
	return client.StringOperations()
}

func TestStringGetMBCS(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetMBCS(context.Background())
	if err != nil {
		t.Fatalf("GetMBCS: %v", err)
	}
	expected := &stringgroup.StringGetMBCSResponse{
		StatusCode: http.StatusOK,
		Value:      toStrPtr("啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€"),
	}
	deepEqualOrFatal(t, result, expected)
}

func TestStringPutMBCS(t *testing.T) {
	client := getStringClient(t)
	result, err := client.PutMBCS(context.Background())
	if err != nil {
		t.Fatalf("PutMBCS: %v", err)
	}
	expected := &stringgroup.StringPutMBCSResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestStringGetBase64Encoded(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetBase64Encoded(context.Background())
	if err != nil {
		t.Fatalf("GetBase64Encoded: %v", err)
	}
	val := []byte("a string that gets encoded with base64")
	expected := &stringgroup.StringGetBase64EncodedResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

// func TestStringGetBase64URLEncoded(t *testing.T) {
// 	client := getStringClient(t)
// 	result, err := client.GetBase64URLEncoded(context.Background())
// 	if err != nil {
// 		t.Fatalf("GetBase64URLEncoded: %v", err)
// 	}
// 	expected := &stringgroup.StringGetBase64URLEncodedResponse{
// 		StatusCode: http.StatusOK,
// 		Value:      []byte("a string that gets encoded with base64url"),
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

func TestStringGetEmpty(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	expected := &stringgroup.StringGetEmptyResponse{
		StatusCode: http.StatusOK,
		Value:      toStrPtr(""),
	}
	deepEqualOrFatal(t, result, expected)
}

func TestStringGetNotProvided(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetNotProvided(context.Background())
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	expected := &stringgroup.StringGetNotProvidedResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestStringGetNull(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	expected := &stringgroup.StringGetNullResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestStringGetNullBase64URLEncoded(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetNullBase64URLEncoded(context.Background())
	if err != nil {
		t.Fatalf("GetNullBase64URLEncoded: %v", err)
	}
	expected := &stringgroup.StringGetNullBase64URLEncodedResponse{
		StatusCode: http.StatusOK,
		Value:      nil,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestStringGetWhitespace(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetWhitespace(context.Background())
	if err != nil {
		t.Fatalf("GetWhitespace: %v", err)
	}
	expected := &stringgroup.StringGetWhitespaceResponse{
		StatusCode: http.StatusOK,
		Value:      toStrPtr("    Now is the time for all good men to come to the aid of their country    "),
	}
	deepEqualOrFatal(t, result, expected)
}

// func TestStringPutBase64URLEncoded(t *testing.T) {
// 	client := getStringClient(t)
// 	result, err := client.PutBase64URLEncoded(context.Background(), []byte("a string that gets encoded with base64url"))
// 	if err != nil {
// 		t.Fatalf("PutBase64URLEncoded: %v", err)
// 	}
// 	expected := &stringgroup.StringPutBase64URLEncodedResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

func TestStringPutEmpty(t *testing.T) {
	client := getStringClient(t)
	result, err := client.PutEmpty(context.Background())
	if err != nil {
		t.Fatalf("PutEmpty: %v", err)
	}
	expected := &stringgroup.StringPutEmptyResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

// func TestStringPutNull(t *testing.T) {
// 	client := getStringClient(t)
// 	result, err := client.PutNull(context.Background())
// 	if err != nil {
// 		t.Fatalf("PutNull: %v", err)
// 	}
// 	expected := &stringgroup.StringPutNullResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

func TestStringPutWhitespace(t *testing.T) {
	client := getStringClient(t)
	result, err := client.PutWhitespace(context.Background())
	if err != nil {
		t.Fatalf("PutWhitespace: %v", err)
	}
	expected := &stringgroup.StringPutWhitespaceResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func toStrPtr(s string) *string {
	return &s
}
