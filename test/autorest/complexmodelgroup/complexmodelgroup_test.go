// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexmodelgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexmodelgroup"
	"testing"
)

func getOperations(t *testing.T) complexmodelgroup.ComplexModelClientOperations {
	client, err := complexmodelgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.ComplexModelClientOperations()
}

func TestCreate(t *testing.T) {
	client := getOperations(t)
	_, err := client.Create(context.Background(), "sub", "rg", complexmodelgroup.CatalogDictionaryOfArray{})
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestList(t *testing.T) {
	client := getOperations(t)
	_, err := client.List(context.Background(), "")
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestUpdate(t *testing.T) {
	client := getOperations(t)
	_, err := client.Update(context.Background(), "", "", complexmodelgroup.CatalogArrayOfDictionary{})
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}
