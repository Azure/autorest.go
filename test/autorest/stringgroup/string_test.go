// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newStringClient() StringOperations {
	return NewStringClient(NewDefaultConnection(nil))
}

func TestStringGetMBCS(t *testing.T) {
	client := newStringClient()
	result, err := client.GetMBCS(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetMBCS: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€"))
}

func TestStringPutMBCS(t *testing.T) {
	client := newStringClient()
	result, err := client.PutMBCS(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutMBCS: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestStringGetBase64Encoded(t *testing.T) {
	client := newStringClient()
	result, err := client.GetBase64Encoded(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBase64Encoded: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val := []byte("a string that gets encoded with base64")
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestStringGetBase64URLEncoded(t *testing.T) {
	client := newStringClient()
	result, err := client.GetBase64URLEncoded(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBase64URLEncoded: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, *result.Value, []byte("a string that gets encoded with base64url"))
}

func TestStringGetEmpty(t *testing.T) {
	client := newStringClient()
	result, err := client.GetEmpty(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr(""))
}

func TestStringGetNotProvided(t *testing.T) {
	client := newStringClient()
	result, err := client.GetNotProvided(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestStringGetNull(t *testing.T) {
	client := newStringClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestStringGetNullBase64URLEncoded(t *testing.T) {
	client := newStringClient()
	result, err := client.GetNullBase64URLEncoded(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNullBase64URLEncoded: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, (*[]byte)(nil))
}

func TestStringGetWhitespace(t *testing.T) {
	client := newStringClient()
	result, err := client.GetWhitespace(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetWhitespace: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("    Now is the time for all good men to come to the aid of their country    "))
}

func TestStringPutBase64URLEncoded(t *testing.T) {
	client := newStringClient()
	result, err := client.PutBase64URLEncoded(context.Background(), []byte("a string that gets encoded with base64url"), nil)
	if err != nil {
		t.Fatalf("PutBase64URLEncoded: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestStringPutEmpty(t *testing.T) {
	client := newStringClient()
	result, err := client.PutEmpty(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestStringPutNull(t *testing.T) {
	client := newStringClient()
	result, err := client.PutNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutNull: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestStringPutWhitespace(t *testing.T) {
	client := newStringClient()
	result, err := client.PutWhitespace(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutWhitespace: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
