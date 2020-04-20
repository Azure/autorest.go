// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func getPolymorphicrecursiveOperations(t *testing.T) complexgroup.PolymorphicrecursiveOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.PolymorphicrecursiveOperations()
}

// GetValid - Get complex types that are polymorphic and have recursive references
func TestGetValid(t *testing.T) {
	client := getPolymorphicrecursiveOperations(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	helpers.DeepEqualOrFatal(t, result.Fish, &complexgroup.Salmon{
		Fishtype: to.StringPtr("salmon"),
		Iswild:   to.BoolPtr(true),
		Length:   to.Float32Ptr(1),
		Location: to.StringPtr("alaska"),
		Siblings: &[]complexgroup.FishType{
			&complexgroup.Shark{
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
				Fishtype: to.StringPtr("shark"),
				Length:   to.Float32Ptr(20),
				Siblings: &[]complexgroup.FishType{
					&complexgroup.Salmon{
						Fishtype: to.StringPtr("salmon"),
						Iswild:   to.BoolPtr(true),
						Length:   to.Float32Ptr(2),
						Location: to.StringPtr("atlantic"),
						Siblings: &[]complexgroup.FishType{
							&complexgroup.Shark{
								Age:      to.Int32Ptr(6),
								Birthday: &sharkBday,
								Fishtype: to.StringPtr("shark"),
								Length:   to.Float32Ptr(20),
								Species:  to.StringPtr("predator"),
							},
							&complexgroup.Sawshark{
								Age:      to.Int32Ptr(105),
								Birthday: &sawBday,
								Fishtype: to.StringPtr("sawshark"),
								Length:   to.Float32Ptr(10),
								Picture:  &[]byte{255, 255, 255, 255, 254},
								Species:  to.StringPtr("dangerous"),
							},
						},
						Species: to.StringPtr("coho"),
					},
					&complexgroup.Sawshark{
						Age:      to.Int32Ptr(105),
						Birthday: &sawBday,
						Fishtype: to.StringPtr("sawshark"),
						Length:   to.Float32Ptr(10),
						Picture:  &[]byte{255, 255, 255, 255, 254},
						Siblings: &[]complexgroup.FishType{},
						Species:  to.StringPtr("dangerous"),
					},
				},
				Species: to.StringPtr("predator"),
			},
			&complexgroup.Sawshark{
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Fishtype: to.StringPtr("sawshark"),
				Length:   to.Float32Ptr(10),
				Picture:  &[]byte{255, 255, 255, 255, 254},
				Siblings: &[]complexgroup.FishType{},
				Species:  to.StringPtr("dangerous"),
			},
		},
		Species: to.StringPtr("king"),
	})
}

// PutValid - Put complex types that are polymorphic and have recursive references
func TestPutValid(t *testing.T) {
	client := getPolymorphicrecursiveOperations(t)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	result, err := client.PutValid(context.Background(), &complexgroup.Salmon{
		Fishtype: to.StringPtr("salmon"),
		Iswild:   to.BoolPtr(true),
		Length:   to.Float32Ptr(1),
		Location: to.StringPtr("alaska"),
		Siblings: &[]complexgroup.FishType{
			&complexgroup.Shark{
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
				Fishtype: to.StringPtr("shark"),
				Length:   to.Float32Ptr(20),
				Siblings: &[]complexgroup.FishType{
					&complexgroup.Salmon{
						Fishtype: to.StringPtr("salmon"),
						Iswild:   to.BoolPtr(true),
						Length:   to.Float32Ptr(2),
						Location: to.StringPtr("atlantic"),
						Siblings: &[]complexgroup.FishType{
							&complexgroup.Shark{
								Age:      to.Int32Ptr(6),
								Birthday: &sharkBday,
								Fishtype: to.StringPtr("shark"),
								Length:   to.Float32Ptr(20),
								Species:  to.StringPtr("predator"),
							},
							&complexgroup.Sawshark{
								Age:      to.Int32Ptr(105),
								Birthday: &sawBday,
								Fishtype: to.StringPtr("sawshark"),
								Length:   to.Float32Ptr(10),
								Picture:  &[]byte{255, 255, 255, 255, 254},
								Species:  to.StringPtr("dangerous"),
							},
						},
						Species: to.StringPtr("coho"),
					},
					&complexgroup.Sawshark{
						Age:      to.Int32Ptr(105),
						Birthday: &sawBday,
						Fishtype: to.StringPtr("sawshark"),
						Length:   to.Float32Ptr(10),
						Picture:  &[]byte{255, 255, 255, 255, 254},
						Siblings: &[]complexgroup.FishType{},
						Species:  to.StringPtr("dangerous"),
					},
				},
				Species: to.StringPtr("predator"),
			},
			&complexgroup.Sawshark{
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Fishtype: to.StringPtr("sawshark"),
				Length:   to.Float32Ptr(10),
				Picture:  &[]byte{255, 255, 255, 255, 254},
				Siblings: &[]complexgroup.FishType{},
				Species:  to.StringPtr("dangerous"),
			},
		},
		Species: to.StringPtr("king"),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
