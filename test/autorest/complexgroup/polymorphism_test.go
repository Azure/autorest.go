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

func getPolymorphismOperations(t *testing.T) complexgroup.PolymorphismOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.PolymorphismOperations()
}

// GetComplicated - Get complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
func TestPolymorphismGetComplicated(t *testing.T) {
	client := getPolymorphismOperations(t)
	result, err := client.GetComplicated(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	salmon, ok := result.Salmon.(*complexgroup.SmartSalmon)
	if !ok {
		t.Fatal("fish wasn't a smart salmon")
	}
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	helpers.DeepEqualOrFatal(t, salmon, &complexgroup.SmartSalmon{
		Fishtype: to.StringPtr("smart_salmon"),
		Iswild:   to.BoolPtr(true),
		Length:   to.Float32Ptr(1),
		Location: to.StringPtr("alaska"),
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
			&complexgroup.Goblinshark{
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    complexgroup.GoblinSharkColor("pinkish-gray").ToPtr(),
				Fishtype: to.StringPtr("goblin"),
				Jawsize:  to.Int32Ptr(5),
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
			},
		},
		Species: to.StringPtr("king"),
	})
}

// GetComposedWithDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, with discriminator specified. Deserialization must NOT fail and use the discriminator type specified on the wire.
func TestPolymorphismGetComposedWithDiscriminator(t *testing.T) {
	client := getPolymorphismOperations(t)
	result, err := client.GetComposedWithDiscriminator(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.DotFishMarket, &complexgroup.DotFishMarket{
		Fishes: &[]complexgroup.DotFishType{
			&complexgroup.DotSalmon{
				FishType: to.StringPtr("DotSalmon"),
				Location: to.StringPtr("australia"),
				Iswild:   to.BoolPtr(false),
				Species:  to.StringPtr("king"),
			},
			&complexgroup.DotSalmon{
				FishType: to.StringPtr("DotSalmon"),
				Location: to.StringPtr("canada"),
				Iswild:   to.BoolPtr(true),
				Species:  to.StringPtr("king"),
			},
		},
		Salmons: &[]complexgroup.DotSalmon{
			{
				FishType: to.StringPtr("DotSalmon"),
				Location: to.StringPtr("sweden"),
				Iswild:   to.BoolPtr(false),
				Species:  to.StringPtr("king"),
			},
			{
				FishType: to.StringPtr("DotSalmon"),
				Location: to.StringPtr("atlantic"),
				Iswild:   to.BoolPtr(true),
				Species:  to.StringPtr("king"),
			},
		},
		SampleFish: &complexgroup.DotSalmon{
			FishType: to.StringPtr("DotSalmon"),
			Location: to.StringPtr("australia"),
			Iswild:   to.BoolPtr(false),
			Species:  to.StringPtr("king"),
		},
		SampleSalmon: &complexgroup.DotSalmon{
			FishType: to.StringPtr("DotSalmon"),
			Location: to.StringPtr("sweden"),
			Iswild:   to.BoolPtr(false),
			Species:  to.StringPtr("king"),
		},
	})
}

// GetComposedWithoutDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, without discriminator specified on wire. Deserialization must NOT fail and use the explicit type of the property.
func TestPolymorphismGetComposedWithoutDiscriminator(t *testing.T) {
	client := getPolymorphismOperations(t)
	result, err := client.GetComposedWithoutDiscriminator(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.DotFishMarket, &complexgroup.DotFishMarket{
		Fishes: &[]complexgroup.DotFishType{
			&complexgroup.DotFish{
				Species: to.StringPtr("king"),
			},
			&complexgroup.DotFish{
				Species: to.StringPtr("king"),
			},
		},
		Salmons: &[]complexgroup.DotSalmon{
			{
				Location: to.StringPtr("sweden"),
				Iswild:   to.BoolPtr(false),
				Species:  to.StringPtr("king"),
			},
			{
				Location: to.StringPtr("atlantic"),
				Iswild:   to.BoolPtr(true),
				Species:  to.StringPtr("king"),
			},
		},
		SampleFish: &complexgroup.DotFish{
			Species: to.StringPtr("king"),
		},
		SampleSalmon: &complexgroup.DotSalmon{
			Location: to.StringPtr("sweden"),
			Iswild:   to.BoolPtr(false),
			Species:  to.StringPtr("king"),
		},
	})
}

// GetDotSyntax - Get complex types that are polymorphic, JSON key contains a dot
func TestPolymorphismGetDotSyntax(t *testing.T) {
	client := getPolymorphismOperations(t)
	result, err := client.GetDotSyntax(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.DotFish, &complexgroup.DotSalmon{
		FishType: to.StringPtr("DotSalmon"),
		Location: to.StringPtr("sweden"),
		Iswild:   to.BoolPtr(true),
		Species:  to.StringPtr("king"),
	})
}

// GetValid - Get complex types that are polymorphic
func TestPolymorphismGetValid(t *testing.T) {
	client := getPolymorphismOperations(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	salmon, ok := result.Fish.(*complexgroup.Salmon)
	if !ok {
		t.Fatal("fish wasn't a salmon")
	}
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	helpers.DeepEqualOrFatal(t, salmon, &complexgroup.Salmon{
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
			&complexgroup.Goblinshark{
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    complexgroup.GoblinSharkColor("pinkish-gray").ToPtr(),
				Fishtype: to.StringPtr("goblin"),
				Jawsize:  to.Int32Ptr(5),
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
			},
		},
		Species: to.StringPtr("king"),
	})
}

// PutComplicated - Put complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
func TestPolymorphismPutComplicated(t *testing.T) {
	t.Skip("additional properties NYI")
}

// PutMissingDiscriminator - Put complex types that are polymorphic, omitting the discriminator
func TestPolymorphismPutMissingDiscriminator(t *testing.T) {
	client := getPolymorphismOperations(t)
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	result, err := client.PutMissingDiscriminator(context.Background(), &complexgroup.Salmon{
		Iswild:   to.BoolPtr(true),
		Length:   to.Float32Ptr(1),
		Location: to.StringPtr("alaska"),
		Siblings: &[]complexgroup.FishType{
			&complexgroup.Shark{
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
				Length:   to.Float32Ptr(20),
				Species:  to.StringPtr("predator"),
			},
			&complexgroup.Sawshark{
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Length:   to.Float32Ptr(10),
				Picture:  &[]byte{255, 255, 255, 255, 254},
				Species:  to.StringPtr("dangerous"),
			},
			&complexgroup.Goblinshark{
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    complexgroup.GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize:  to.Int32Ptr(5),
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
			},
		},
		Species: to.StringPtr("king"),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

// PutValid - Put complex types that are polymorphic
func TestPolymorphismPutValid(t *testing.T) {
	client := getPolymorphismOperations(t)
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	resp, err := client.PutValid(context.Background(), &complexgroup.Salmon{
		Iswild:   to.BoolPtr(true),
		Length:   to.Float32Ptr(1),
		Location: to.StringPtr("alaska"),
		Siblings: &[]complexgroup.FishType{
			&complexgroup.Shark{
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
				Length:   to.Float32Ptr(20),
				Species:  to.StringPtr("predator"),
			},
			&complexgroup.Sawshark{
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Length:   to.Float32Ptr(10),
				Picture:  &[]byte{255, 255, 255, 255, 254},
				Species:  to.StringPtr("dangerous"),
			},
			&complexgroup.Goblinshark{
				Age:      to.Int32Ptr(1),
				Birthday: &goblinBday,
				Color:    complexgroup.GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize:  to.Int32Ptr(5),
				Length:   to.Float32Ptr(30),
				Species:  to.StringPtr("scary"),
			},
		},
		Species: to.StringPtr("king"),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutValidMissingRequired - Put complex types that are polymorphic, attempting to omit required 'birthday' field - the request should not be allowed from the client
func TestPolymorphismPutValidMissingRequired(t *testing.T) {
	t.Skip("client side validation not applicable to track 2")
}
