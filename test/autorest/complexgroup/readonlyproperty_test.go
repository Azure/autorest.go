// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/google/go-cmp/cmp"
)

func newReadonlypropertyClient() *ReadonlypropertyClient {
	return NewReadonlypropertyClient(NewDefaultConnection(nil))
}

func TestReadonlypropertyGetValid(t *testing.T) {
	client := newReadonlypropertyClient()
	result, err := client.GetValid(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	if r := cmp.Diff(result.ReadonlyObj, ReadonlyObj{ID: to.StringPtr("1234"), Size: to.Int32Ptr(2)}); r != "" {
		t.Fatal(r)
	}
}

func TestReadonlypropertyPutValid(t *testing.T) {
	client := newReadonlypropertyClient()
	id, size := "1234", int32(2)
	result, err := client.PutValid(context.Background(), ReadonlyObj{ID: &id, Size: &size}, nil)
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
