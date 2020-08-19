// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexmodelgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexmodelgroup"
	"testing"
)

func TestCreate(t *testing.T) {
	client := complexmodelgroup.NewDefaultClient(nil).ComplexModelClientOperations()
	_, err := client.Create(context.Background(), "sub", "rg", complexmodelgroup.CatalogDictionaryOfArray{})
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestList(t *testing.T) {
	client := complexmodelgroup.NewDefaultClient(nil).ComplexModelClientOperations()
	_, err := client.List(context.Background(), "")
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestUpdate(t *testing.T) {
	client := complexmodelgroup.NewDefaultClient(nil).ComplexModelClientOperations()
	_, err := client.Update(context.Background(), "", "", complexmodelgroup.CatalogArrayOfDictionary{})
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}
