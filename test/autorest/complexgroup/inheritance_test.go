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

func TestInheritanceGetValid(t *testing.T) {
	client := complexgroup.NewDefaultClient(nil).InheritanceOperations()
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.Siamese, &complexgroup.Siamese{
		Cat: complexgroup.Cat{
			Pet: complexgroup.Pet{
				ID:   to.Int32Ptr(2),
				Name: to.StringPtr("Siameeee"),
			},
			Color: to.StringPtr("green"),
			Hates: &[]complexgroup.Dog{
				{
					Pet: complexgroup.Pet{
						ID:   to.Int32Ptr(1),
						Name: to.StringPtr("Potato"),
					},
					Food: to.StringPtr("tomato"),
				},
				{
					Pet: complexgroup.Pet{
						ID:   to.Int32Ptr(-1),
						Name: to.StringPtr("Tomato"),
					},
					Food: to.StringPtr("french fries"),
				},
			},
		},
		Breed: to.StringPtr("persian"),
	})
}

func TestInheritancePutValid(t *testing.T) {
	client := complexgroup.NewDefaultClient(nil).InheritanceOperations()
	result, err := client.PutValid(context.Background(), complexgroup.Siamese{
		Cat: complexgroup.Cat{
			Pet: complexgroup.Pet{
				ID:   to.Int32Ptr(2),
				Name: to.StringPtr("Siameeee"),
			},
			Color: to.StringPtr("green"),
			Hates: &[]complexgroup.Dog{
				{
					Pet: complexgroup.Pet{
						ID:   to.Int32Ptr(1),
						Name: to.StringPtr("Potato"),
					},
					Food: to.StringPtr("tomato"),
				},
				{
					Pet: complexgroup.Pet{
						ID:   to.Int32Ptr(-1),
						Name: to.StringPtr("Tomato"),
					},
					Food: to.StringPtr("french fries"),
				},
			},
		},
		Breed: to.StringPtr("persian"),
	})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
