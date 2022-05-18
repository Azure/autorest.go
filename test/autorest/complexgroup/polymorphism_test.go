// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"generatortests"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newPolymorphismClient() *PolymorphismClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewPolymorphismClient(pl)
}

// GetComplicated - Get complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
func TestPolymorphismGetComplicated(t *testing.T) {
	client := newPolymorphismClient()
	result, err := client.GetComplicated(context.Background(), nil)
	require.NoError(t, err)
	salmon, ok := result.SalmonClassification.(*SmartSalmon)
	if !ok {
		t.Fatal("fish wasn't a smart salmon")
	}
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	expectedFish := Fish{
		Fishtype: to.Ptr("smart_salmon"),
		Length:   to.Ptr[float32](1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.Ptr("shark"),
				Length:   to.Ptr[float32](20),
				Species:  to.Ptr("predator"),
				Age:      to.Ptr[int32](6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.Ptr("sawshark"),
				Length:   to.Ptr[float32](10),
				Species:  to.Ptr("dangerous"),
				Age:      to.Ptr[int32](105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.Ptr("goblin"),
				Length:   to.Ptr[float32](30),
				Species:  to.Ptr("scary"),
				Age:      to.Ptr[int32](1),
				Birthday: &goblinBday,
				Color:    to.Ptr(GoblinSharkColor("pinkish-gray")),
				Jawsize:  to.Ptr[int32](5),
			},
		},
		Species: to.Ptr("king"),
	}
	expectedSalmon := Salmon{
		Fishtype: to.Ptr("smart_salmon"),
		Length:   to.Ptr[float32](1),
		Iswild:   to.Ptr(true),
		Location: to.Ptr("alaska"),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.Ptr("shark"),
				Length:   to.Ptr[float32](20),
				Species:  to.Ptr("predator"),
				Age:      to.Ptr[int32](6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.Ptr("sawshark"),
				Length:   to.Ptr[float32](10),
				Species:  to.Ptr("dangerous"),
				Age:      to.Ptr[int32](105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.Ptr("goblin"),
				Length:   to.Ptr[float32](30),
				Species:  to.Ptr("scary"),
				Age:      to.Ptr[int32](1),
				Birthday: &goblinBday,
				Color:    to.Ptr(GoblinSharkColor("pinkish-gray")),
				Jawsize:  to.Ptr[int32](5),
			},
		},
		Species: to.Ptr("king"),
	}
	if r := cmp.Diff(salmon, &SmartSalmon{
		Fishtype: to.Ptr("smart_salmon"),
		Length:   to.Ptr[float32](1),
		Iswild:   to.Ptr(true),
		Location: to.Ptr("alaska"),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.Ptr("shark"),
				Length:   to.Ptr[float32](20),
				Species:  to.Ptr("predator"),
				Age:      to.Ptr[int32](6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.Ptr("sawshark"),
				Length:   to.Ptr[float32](10),
				Species:  to.Ptr("dangerous"),
				Age:      to.Ptr[int32](105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.Ptr("goblin"),
				Length:   to.Ptr[float32](30),
				Species:  to.Ptr("scary"),
				Age:      to.Ptr[int32](1),
				Birthday: &goblinBday,
				Color:    to.Ptr(GoblinSharkColor("pinkish-gray")),
				Jawsize:  to.Ptr[int32](5),
			},
		},
		Species: to.Ptr("king"),
		AdditionalProperties: map[string]interface{}{
			"additionalProperty1": float64(1),
			"additionalProperty2": false,
			"additionalProperty3": "hello",
			"additionalProperty4": map[string]interface{}{
				"a": float64(1),
				"b": float64(2),
			},
			"additionalProperty5": []interface{}{
				float64(1), float64(3),
			},
		},
	}); r != "" {
		t.Fatal(r)
	}
	if r := cmp.Diff(result.SalmonClassification.GetSalmon(), &expectedSalmon); r != "" {
		t.Fatal(r)
	}
	if r := cmp.Diff(result.SalmonClassification.GetFish(), &expectedFish); r != "" {
		t.Fatal(r)
	}
}

// GetComposedWithDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, with discriminator specified. Deserialization must NOT fail and use the discriminator type specified on the wire.
func TestPolymorphismGetComposedWithDiscriminator(t *testing.T) {
	client := newPolymorphismClient()
	result, err := client.GetComposedWithDiscriminator(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.DotFishMarket, DotFishMarket{
		Fishes: []DotFishClassification{
			&DotSalmon{
				FishType: to.Ptr("DotSalmon"),
				Species:  to.Ptr("king"),
				Location: to.Ptr("australia"),
				Iswild:   to.Ptr(false),
			},
			&DotSalmon{
				FishType: to.Ptr("DotSalmon"),
				Species:  to.Ptr("king"),
				Location: to.Ptr("canada"),
				Iswild:   to.Ptr(true),
			},
		},
		Salmons: []*DotSalmon{
			{
				FishType: to.Ptr("DotSalmon"),
				Species:  to.Ptr("king"),
				Location: to.Ptr("sweden"),
				Iswild:   to.Ptr(false),
			},
			{
				FishType: to.Ptr("DotSalmon"),
				Species:  to.Ptr("king"),
				Location: to.Ptr("atlantic"),
				Iswild:   to.Ptr(true),
			},
		},
		SampleFish: &DotSalmon{
			FishType: to.Ptr("DotSalmon"),
			Species:  to.Ptr("king"),
			Location: to.Ptr("australia"),
			Iswild:   to.Ptr(false),
		},
		SampleSalmon: &DotSalmon{
			FishType: to.Ptr("DotSalmon"),
			Species:  to.Ptr("king"),
			Location: to.Ptr("sweden"),
			Iswild:   to.Ptr(false),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComposedWithoutDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, without discriminator specified on wire. Deserialization must NOT fail and use the explicit type of the property.
func TestPolymorphismGetComposedWithoutDiscriminator(t *testing.T) {
	client := newPolymorphismClient()
	result, err := client.GetComposedWithoutDiscriminator(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.DotFishMarket, DotFishMarket{
		Fishes: []DotFishClassification{
			&DotFish{
				Species: to.Ptr("king"),
			},
			&DotFish{
				Species: to.Ptr("king"),
			},
		},
		Salmons: []*DotSalmon{
			{
				Species:  to.Ptr("king"),
				Location: to.Ptr("sweden"),
				Iswild:   to.Ptr(false),
			},
			{
				Species:  to.Ptr("king"),
				Location: to.Ptr("atlantic"),
				Iswild:   to.Ptr(true),
			},
		},
		SampleFish: &DotFish{
			Species: to.Ptr("king"),
		},
		SampleSalmon: &DotSalmon{
			Species:  to.Ptr("king"),
			Location: to.Ptr("sweden"),
			Iswild:   to.Ptr(false),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDotSyntax - Get complex types that are polymorphic, JSON key contains a dot
func TestPolymorphismGetDotSyntax(t *testing.T) {
	client := newPolymorphismClient()
	result, err := client.GetDotSyntax(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.DotFishClassification, &DotSalmon{
		FishType: to.Ptr("DotSalmon"),
		Species:  to.Ptr("king"),
		Location: to.Ptr("sweden"),
		Iswild:   to.Ptr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetValid - Get complex types that are polymorphic
func TestPolymorphismGetValid(t *testing.T) {
	client := newPolymorphismClient()
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	salmon, ok := result.FishClassification.(*Salmon)
	if !ok {
		t.Fatal("fish wasn't a salmon")
	}
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	if r := cmp.Diff(salmon, &Salmon{
		Fishtype: to.Ptr("salmon"),
		Length:   to.Ptr[float32](1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.Ptr("shark"),
				Length:   to.Ptr[float32](20),
				Species:  to.Ptr("predator"),
				Age:      to.Ptr[int32](6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.Ptr("sawshark"),
				Length:   to.Ptr[float32](10),
				Species:  to.Ptr("dangerous"),
				Age:      to.Ptr[int32](105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.Ptr("goblin"),
				Length:   to.Ptr[float32](30),
				Species:  to.Ptr("scary"),
				Age:      to.Ptr[int32](1),
				Birthday: &goblinBday,
				Color:    to.Ptr(GoblinSharkColor("pinkish-gray")),
				Jawsize:  to.Ptr[int32](5),
			},
		},
		Species:  to.Ptr("king"),
		Iswild:   to.Ptr(true),
		Location: to.Ptr("alaska"),
	}); r != "" {
		t.Fatal(r)
	}
}

// PutComplicated - Put complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
func TestPolymorphismPutComplicated(t *testing.T) {
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	client := newPolymorphismClient()
	result, err := client.PutComplicated(context.Background(), &SmartSalmon{
		Fishtype: to.Ptr("smart_salmon"),
		Length:   to.Ptr[float32](1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.Ptr("shark"),
				Length:   to.Ptr[float32](20),
				Species:  to.Ptr("predator"),
				Age:      to.Ptr[int32](6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.Ptr("sawshark"),
				Length:   to.Ptr[float32](10),
				Species:  to.Ptr("dangerous"),
				Age:      to.Ptr[int32](105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.Ptr("goblin"),
				Length:   to.Ptr[float32](30),
				Species:  to.Ptr("scary"),
				Age:      to.Ptr[int32](1),
				Birthday: &goblinBday,
				Color:    to.Ptr(GoblinSharkColor("pinkish-gray")),
				Jawsize:  to.Ptr[int32](5),
			},
		},
		Species:  to.Ptr("king"),
		Iswild:   to.Ptr(true),
		Location: to.Ptr("alaska"),
		AdditionalProperties: map[string]interface{}{
			"additionalProperty1": float64(1),
			"additionalProperty2": false,
			"additionalProperty3": "hello",
			"additionalProperty4": map[string]interface{}{
				"a": float64(1),
				"b": float64(2),
			},
			"additionalProperty5": []interface{}{
				float64(1), float64(3),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PutMissingDiscriminator - Put complex types that are polymorphic, omitting the discriminator
func TestPolymorphismPutMissingDiscriminator(t *testing.T) {
	client := newPolymorphismClient()
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	result, err := client.PutMissingDiscriminator(context.Background(), &Salmon{
		Length: to.Ptr[float32](1),
		Siblings: []FishClassification{
			&Shark{
				Length:   to.Ptr[float32](20),
				Species:  to.Ptr("predator"),
				Age:      to.Ptr[int32](6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Length:   to.Ptr[float32](10),
				Species:  to.Ptr("dangerous"),
				Age:      to.Ptr[int32](105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Length:   to.Ptr[float32](30),
				Species:  to.Ptr("scary"),
				Age:      to.Ptr[int32](1),
				Birthday: &goblinBday,
				Color:    to.Ptr(GoblinSharkColor("pinkish-gray")),
				Jawsize:  to.Ptr[int32](5),
			},
		},
		Species:  to.Ptr("king"),
		Iswild:   to.Ptr(true),
		Location: to.Ptr("alaska"),
	}, nil)
	require.NoError(t, err)
	expectedSalmon := &Salmon{
		Length: to.Ptr[float32](1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.Ptr("shark"),
				Length:   to.Ptr[float32](20),
				Species:  to.Ptr("predator"),
				Age:      to.Ptr[int32](6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.Ptr("sawshark"),
				Length:   to.Ptr[float32](10),
				Species:  to.Ptr("dangerous"),
				Age:      to.Ptr[int32](105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Fishtype: to.Ptr("goblin"),
				Length:   to.Ptr[float32](30),
				Species:  to.Ptr("scary"),
				Age:      to.Ptr[int32](1),
				Birthday: &goblinBday,
				Color:    to.Ptr(GoblinSharkColor("pinkish-gray")),
				Jawsize:  to.Ptr[int32](5),
			},
		},
		Species:  to.Ptr("king"),
		Iswild:   to.Ptr(true),
		Location: to.Ptr("alaska"),
	}
	if r := cmp.Diff(expectedSalmon, result.SalmonClassification); r != "" {
		t.Fatal(r)
	}
}

// PutValid - Put complex types that are polymorphic
func TestPolymorphismPutValid(t *testing.T) {
	client := newPolymorphismClient()
	goblinBday := time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	result, err := client.PutValid(context.Background(), &Salmon{
		Length: to.Ptr[float32](1),
		Siblings: []FishClassification{
			&Shark{
				Length:   to.Ptr[float32](20),
				Species:  to.Ptr("predator"),
				Age:      to.Ptr[int32](6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Length:   to.Ptr[float32](10),
				Species:  to.Ptr("dangerous"),
				Age:      to.Ptr[int32](105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
			&Goblinshark{
				Length:   to.Ptr[float32](30),
				Species:  to.Ptr("scary"),
				Age:      to.Ptr[int32](1),
				Birthday: &goblinBday,
				Color:    to.Ptr(GoblinSharkColor("pinkish-gray")),
				Jawsize:  to.Ptr[int32](5),
			},
		},
		Species:  to.Ptr("king"),
		Iswild:   to.Ptr(true),
		Location: to.Ptr("alaska"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PutValidMissingRequired - Put complex types that are polymorphic, attempting to omit required 'birthday' field - the request should not be allowed from the client
func TestPolymorphismPutValidMissingRequired(t *testing.T) {
	t.Skip("client side validation not applicable to track 2")
}
