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
	expectedFish := complexgroup.Fish{
		Fishtype: to.StringPtr("smart_salmon"),
		Length:   to.Float32Ptr(1),
		Siblings: &[]complexgroup.FishClassification{
			&complexgroup.Shark{
				Fish: complexgroup.Fish{
					Fishtype: to.StringPtr("shark"),
					Length:   to.Float32Ptr(20),
					Species:  to.StringPtr("predator")},
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&complexgroup.Sawshark{
				Shark: complexgroup.Shark{
					Fish: complexgroup.Fish{
						Fishtype: to.StringPtr("sawshark"),
						Length:   to.Float32Ptr(10),
						Species:  to.StringPtr("dangerous"),
					},
					Age:      to.Int32Ptr(105),
					Birthday: &sawBday,
				},
				Picture: &[]byte{255, 255, 255, 255, 254},
			},
			&complexgroup.Goblinshark{
				Shark: complexgroup.Shark{
					Fish: complexgroup.Fish{
						Fishtype: to.StringPtr("goblin"),
						Length:   to.Float32Ptr(30),
						Species:  to.StringPtr("scary"),
					},
					Age:      to.Int32Ptr(1),
					Birthday: &goblinBday,
				},
				Color:   complexgroup.GoblinSharkColor("pinkish-gray").ToPtr(),
				Jawsize: to.Int32Ptr(5),
			},
		},
		Species: to.StringPtr("king"),
	}
	expectedSalmon := complexgroup.Salmon{
		Fish:     expectedFish,
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
	}
	helpers.DeepEqualOrFatal(t, salmon, &complexgroup.SmartSalmon{
		Salmon: expectedSalmon,
	})
	helpers.DeepEqualOrFatal(t, result.Salmon.GetSalmon(), &expectedSalmon)
	helpers.DeepEqualOrFatal(t, result.Salmon.GetFish(), &expectedFish)
}

// GetComposedWithDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, with discriminator specified. Deserialization must NOT fail and use the discriminator type specified on the wire.
func TestPolymorphismGetComposedWithDiscriminator(t *testing.T) {
	client := getPolymorphismOperations(t)
	result, err := client.GetComposedWithDiscriminator(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.DotFishMarket, &complexgroup.DotFishMarket{
		Fishes: &[]complexgroup.DotFishClassification{
			&complexgroup.DotSalmon{
				DotFish: complexgroup.DotFish{
					FishType: to.StringPtr("DotSalmon"),
					Species:  to.StringPtr("king"),
				},
				Location: to.StringPtr("australia"),
				Iswild:   to.BoolPtr(false),
			},
			&complexgroup.DotSalmon{
				DotFish: complexgroup.DotFish{
					FishType: to.StringPtr("DotSalmon"),
					Species:  to.StringPtr("king"),
				},
				Location: to.StringPtr("canada"),
				Iswild:   to.BoolPtr(true),
			},
		},
		Salmons: &[]complexgroup.DotSalmon{
			{
				DotFish: complexgroup.DotFish{
					FishType: to.StringPtr("DotSalmon"),
					Species:  to.StringPtr("king"),
				},
				Location: to.StringPtr("sweden"),
				Iswild:   to.BoolPtr(false),
			},
			{
				DotFish: complexgroup.DotFish{
					FishType: to.StringPtr("DotSalmon"),
					Species:  to.StringPtr("king"),
				},
				Location: to.StringPtr("atlantic"),
				Iswild:   to.BoolPtr(true),
			},
		},
		SampleFish: &complexgroup.DotSalmon{
			DotFish: complexgroup.DotFish{
				FishType: to.StringPtr("DotSalmon"),
				Species:  to.StringPtr("king"),
			},
			Location: to.StringPtr("australia"),
			Iswild:   to.BoolPtr(false),
		},
		SampleSalmon: &complexgroup.DotSalmon{
			DotFish: complexgroup.DotFish{
				FishType: to.StringPtr("DotSalmon"),
				Species:  to.StringPtr("king"),
			},
			Location: to.StringPtr("sweden"),
			Iswild:   to.BoolPtr(false),
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
		Fishes: &[]complexgroup.DotFishClassification{
			&complexgroup.DotFish{
				Species: to.StringPtr("king"),
			},
			&complexgroup.DotFish{
				Species: to.StringPtr("king"),
			},
		},
		Salmons: &[]complexgroup.DotSalmon{
			{
				DotFish: complexgroup.DotFish{
					Species: to.StringPtr("king"),
				},
				Location: to.StringPtr("sweden"),
				Iswild:   to.BoolPtr(false),
			},
			{
				DotFish: complexgroup.DotFish{
					Species: to.StringPtr("king"),
				},
				Location: to.StringPtr("atlantic"),
				Iswild:   to.BoolPtr(true),
			},
		},
		SampleFish: &complexgroup.DotFish{
			Species: to.StringPtr("king"),
		},
		SampleSalmon: &complexgroup.DotSalmon{
			DotFish: complexgroup.DotFish{
				Species: to.StringPtr("king"),
			},
			Location: to.StringPtr("sweden"),
			Iswild:   to.BoolPtr(false),
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
		DotFish: complexgroup.DotFish{
			FishType: to.StringPtr("DotSalmon"),
			Species:  to.StringPtr("king"),
		},
		Location: to.StringPtr("sweden"),
		Iswild:   to.BoolPtr(true),
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
		Fish: complexgroup.Fish{
			Fishtype: to.StringPtr("salmon"),
			Length:   to.Float32Ptr(1),
			Siblings: &[]complexgroup.FishClassification{
				&complexgroup.Shark{
					Fish: complexgroup.Fish{
						Fishtype: to.StringPtr("shark"),
						Length:   to.Float32Ptr(20),
						Species:  to.StringPtr("predator"),
					},
					Age:      to.Int32Ptr(6),
					Birthday: &sharkBday,
				},
				&complexgroup.Sawshark{
					Shark: complexgroup.Shark{
						Fish: complexgroup.Fish{
							Fishtype: to.StringPtr("sawshark"),
							Length:   to.Float32Ptr(10),
							Species:  to.StringPtr("dangerous"),
						},
						Age:      to.Int32Ptr(105),
						Birthday: &sawBday,
					},
					Picture: &[]byte{255, 255, 255, 255, 254},
				},
				&complexgroup.Goblinshark{
					Shark: complexgroup.Shark{
						Fish: complexgroup.Fish{
							Fishtype: to.StringPtr("goblin"),
							Length:   to.Float32Ptr(30),
							Species:  to.StringPtr("scary"),
						},
						Age:      to.Int32Ptr(1),
						Birthday: &goblinBday,
					},
					Color:   complexgroup.GoblinSharkColor("pinkish-gray").ToPtr(),
					Jawsize: to.Int32Ptr(5),
				},
			},
			Species: to.StringPtr("king"),
		},
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
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
		Fish: complexgroup.Fish{
			Length: to.Float32Ptr(1),
			Siblings: &[]complexgroup.FishClassification{
				&complexgroup.Shark{
					Fish: complexgroup.Fish{
						Length:  to.Float32Ptr(20),
						Species: to.StringPtr("predator"),
					},
					Age:      to.Int32Ptr(6),
					Birthday: &sharkBday,
				},
				&complexgroup.Sawshark{
					Shark: complexgroup.Shark{
						Fish: complexgroup.Fish{
							Length:  to.Float32Ptr(10),
							Species: to.StringPtr("dangerous"),
						},
						Age:      to.Int32Ptr(105),
						Birthday: &sawBday,
					},
					Picture: &[]byte{255, 255, 255, 255, 254},
				},
				&complexgroup.Goblinshark{
					Shark: complexgroup.Shark{
						Fish: complexgroup.Fish{
							Length:  to.Float32Ptr(30),
							Species: to.StringPtr("scary"),
						},
						Age:      to.Int32Ptr(1),
						Birthday: &goblinBday,
					},
					Color:   complexgroup.GoblinSharkColor("pinkish-gray").ToPtr(),
					Jawsize: to.Int32Ptr(5),
				},
			},
			Species: to.StringPtr("king"),
		},
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
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
		Fish: complexgroup.Fish{
			Length: to.Float32Ptr(1),
			Siblings: &[]complexgroup.FishClassification{
				&complexgroup.Shark{
					Fish: complexgroup.Fish{
						Length:  to.Float32Ptr(20),
						Species: to.StringPtr("predator"),
					},
					Age:      to.Int32Ptr(6),
					Birthday: &sharkBday,
				},
				&complexgroup.Sawshark{
					Shark: complexgroup.Shark{
						Fish: complexgroup.Fish{
							Length:  to.Float32Ptr(10),
							Species: to.StringPtr("dangerous"),
						},
						Age:      to.Int32Ptr(105),
						Birthday: &sawBday,
					},
					Picture: &[]byte{255, 255, 255, 255, 254},
				},
				&complexgroup.Goblinshark{
					Shark: complexgroup.Shark{
						Fish: complexgroup.Fish{
							Length:  to.Float32Ptr(30),
							Species: to.StringPtr("scary"),
						},
						Age:      to.Int32Ptr(1),
						Birthday: &goblinBday,
					},
					Color:   complexgroup.GoblinSharkColor("pinkish-gray").ToPtr(),
					Jawsize: to.Int32Ptr(5),
				},
			},
			Species: to.StringPtr("king"),
		},
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
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
