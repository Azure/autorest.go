// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexmodelgroup

import (
	"context"
	"testing"
)

func newComplexModelClient() ComplexModelClient {
	return NewComplexModelClient(NewDefaultConnection(nil))
}

func TestCreate(t *testing.T) {
	client := newComplexModelClient()
	_, err := client.Create(context.Background(), "sub", "rg", CatalogDictionaryOfArray{}, nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestList(t *testing.T) {
	client := newComplexModelClient()
	_, err := client.List(context.Background(), "", nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestUpdate(t *testing.T) {
	client := newComplexModelClient()
	_, err := client.Update(context.Background(), "", "", CatalogArrayOfDictionary{}, nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}
