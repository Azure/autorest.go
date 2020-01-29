// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgrouptest

import (
	"context"
	"generatortests/autorest/generated/stringgroup"
	"net/http"
	"reflect"
	"testing"
)

func getStringClient(t *testing.T) *stringgroup.StringClient {
	client, err := stringgroup.NewStringClient(nil)
	if err != nil {
		t.Fatalf("failed to create string client: %v", err)
	}
	return client
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
}

func TestGetMBCS(t *testing.T) {
	client := getStringClient(t)
	result, err := client.GetMBCS(context.Background())
	if err != nil {
		t.Fatalf("GetMBCS: %v", err)
	}
	expected := &stringgroup.GetMBCSResponse{
		StatusCode: http.StatusOK,
		Value:      "啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€",
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutMBCS(t *testing.T) {
	client := getStringClient(t)
	result, err := client.PutMBCS(context.Background(), "啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€")
	if err != nil {
		t.Fatalf("PutMBCS: %v", err)
	}
	expected := &stringgroup.PutMBCSResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
