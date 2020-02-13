// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"net/http"
	"testing"
)

func getInheritanceOperations(t *testing.T) complexgroup.InheritanceOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.InheritanceOperations()
}

// TODO this test passes but the endpoint is returning more fields than what we're currently checking for
func TestInheritanceGetValid(t *testing.T) {
	client := getInheritanceOperations(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	id := "persian"
	expected := &complexgroup.InheritanceGetValidResponse{
		StatusCode: http.StatusOK,
		Siamese:    &complexgroup.Siamese{Breed: &id},
	}
	deepEqualOrFatal(t, result, expected)
}

// // TODO the endpoint is expecting more fields than those that currently exist in the Siamese type because we're still missing the discriminated type functionality
// func TestInheritancePutValid(t *testing.T) {
// 	client := getInheritanceOperations(t)
// 	x := "persian"
// 	cat := complexgroup.Siamese{
// 		Breed: &x,
// 	}
// 	result, err := client.PutValid(context.Background(), cat)
// 	if err != nil {
// 		t.Fatalf("PutValid: %v", err)
// 	}
// 	expected := &complexgroup.InheritancePutValidResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }
