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

func TestInheritanceGetValid(t *testing.T) {
	client := getInheritanceOperations(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	d1ID, d1Name, d1Food, d2ID, d2Name, d2Food := int32(1), "Potato", "tomato", int32(-1), "Tomato", "french fries"
	id, name, color, breed, hates := int32(2), "Siameeee", "green", "persian", []complexgroup.Dog{complexgroup.Dog{ID: &d1ID, Name: &d1Name, Food: &d1Food}, complexgroup.Dog{ID: &d2ID, Name: &d2Name, Food: &d2Food}}
	expected := &complexgroup.InheritanceGetValidResponse{
		StatusCode: http.StatusOK,
		Siamese: &complexgroup.Siamese{
			Breed: &breed,
			Color: &color,
			Hates: hates,
			ID:    &id,
			Name:  &name,
		},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestInheritancePutValid(t *testing.T) {
	client := getInheritanceOperations(t)
	d1ID, d1Name, d1Food, d2ID, d2Name, d2Food := int32(1), "Potato", "tomato", int32(-1), "Tomato", "french fries"
	id, name, color, breed, hates := int32(2), "Siameeee", "green", "persian", []complexgroup.Dog{complexgroup.Dog{ID: &d1ID, Name: &d1Name, Food: &d1Food}, complexgroup.Dog{ID: &d2ID, Name: &d2Name, Food: &d2Food}}
	cat := complexgroup.Siamese{
		Breed: &breed,
		Color: &color,
		Hates: hates,
		ID:    &id,
		Name:  &name,
	}
	result, err := client.PutValid(context.Background(), cat)
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	expected := &complexgroup.InheritancePutValidResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
