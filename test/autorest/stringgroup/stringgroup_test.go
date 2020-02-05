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

func getEnumClient(t *testing.T) stringgroup.EnumOperations {
	client, err := stringgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create string client: %v", err)
	}
	return client.EnumOperations()
}

func getStringClient(t *testing.T) stringgroup.StringOperations {
	client, err := stringgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create string client: %v", err)
	}
	return client.StringOperations()
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
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

func toStrPtr(s string) *string {
	return &s
}
