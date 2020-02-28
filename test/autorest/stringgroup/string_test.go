// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgrouptest

import (
	"context"
	"generatortests/autorest/generated/stringgroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
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
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€"))
}

func TestStringPutMBCS(t *testing.T) {
	client := getStringClient(t)
	result, err := client.PutMBCS(context.Background())
	if err != nil {
		t.Fatalf("PutMBCS: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestStringGetBase64Encoded(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetBase64Encoded(context.Background())
	if err != nil {
		t.Fatalf("GetBase64Encoded: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val := []byte("a string that gets encoded with base64")
	helpers.DeepEqualOrFatal(t, result.Value, &val)
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
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr(""))
}

func TestStringGetNotProvided(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetNotProvided(context.Background())
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestStringGetNull(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestStringGetNullBase64URLEncoded(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetNullBase64URLEncoded(context.Background())
	if err != nil {
		t.Fatalf("GetNullBase64URLEncoded: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, (*[]byte)(nil))
}

func TestStringGetWhitespace(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetWhitespace(context.Background())
	if err != nil {
		t.Fatalf("GetWhitespace: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("    Now is the time for all good men to come to the aid of their country    "))
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
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
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
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}
