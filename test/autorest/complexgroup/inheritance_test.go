// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
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
	helpers.DeepEqualOrFatal(t, result.Siamese, &complexgroup.Siamese{
		Breed: to.StringPtr("persian"),
		Color: to.StringPtr("green"),
		Hates: &[]complexgroup.Dog{
			complexgroup.Dog{
				Food: to.StringPtr("tomato"),
				ID:   to.Int32Ptr(1),
				Name: to.StringPtr("Potato"),
			},
			complexgroup.Dog{
				Food: to.StringPtr("french fries"),
				ID:   to.Int32Ptr(-1),
				Name: to.StringPtr("Tomato"),
			},
		},
		ID:   to.Int32Ptr(2),
		Name: to.StringPtr("Siameeee"),
	})
}

func TestInheritancePutValid(t *testing.T) {
	client := getInheritanceOperations(t)
	result, err := client.PutValid(context.Background(), complexgroup.Siamese{
		Breed: to.StringPtr("persian"),
		Color: to.StringPtr("green"),
		Hates: &[]complexgroup.Dog{
			complexgroup.Dog{
				Food: to.StringPtr("tomato"),
				ID:   to.Int32Ptr(1),
				Name: to.StringPtr("Potato"),
			},
			complexgroup.Dog{
				Food: to.StringPtr("french fries"),
				ID:   to.Int32Ptr(-1),
				Name: to.StringPtr("Tomato"),
			},
		},
		ID:   to.Int32Ptr(2),
		Name: to.StringPtr("Siameeee"),
	})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
