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

func newPolymorphicrecursiveClient() *PolymorphicrecursiveClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewPolymorphicrecursiveClient(pl)
}

// GetValid - Get complex types that are polymorphic and have recursive references
func TestGetValid(t *testing.T) {
	client := newPolymorphicrecursiveClient()
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	if r := cmp.Diff(result.FishClassification, &Salmon{
		Fishtype: to.Ptr("salmon"),
		Length:   to.Ptr[float32](1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.Ptr("shark"),
				Length:   to.Ptr[float32](20),
				Siblings: []FishClassification{
					&Salmon{
						Fishtype: to.Ptr("salmon"),
						Length:   to.Ptr[float32](2),
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
						},
						Species:  to.Ptr("coho"),
						Iswild:   to.Ptr(true),
						Location: to.Ptr("atlantic"),
					},
					&Sawshark{
						Fishtype: to.Ptr("sawshark"),
						Length:   to.Ptr[float32](10),
						Siblings: []FishClassification{},
						Species:  to.Ptr("dangerous"),
						Age:      to.Ptr[int32](105),
						Birthday: &sawBday,
						Picture:  []byte{255, 255, 255, 255, 254},
					},
				},
				Species:  to.Ptr("predator"),
				Age:      to.Ptr[int32](6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.Ptr("sawshark"),
				Length:   to.Ptr[float32](10),
				Siblings: []FishClassification{},
				Species:  to.Ptr("dangerous"),
				Age:      to.Ptr[int32](105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
		},
		Species:  to.Ptr("king"),
		Iswild:   to.Ptr(true),
		Location: to.Ptr("alaska"),
	}); r != "" {
		t.Fatal(r)
	}
}

// PutValid - Put complex types that are polymorphic and have recursive references
func TestPutValid(t *testing.T) {
	client := newPolymorphicrecursiveClient()
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	result, err := client.PutValid(context.Background(), &Salmon{
		Fishtype: to.Ptr("salmon"),
		Length:   to.Ptr[float32](1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.Ptr("shark"),
				Length:   to.Ptr[float32](20),
				Siblings: []FishClassification{
					&Salmon{
						Fishtype: to.Ptr("salmon"),
						Length:   to.Ptr[float32](2),
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
						},
						Species:  to.Ptr("coho"),
						Iswild:   to.Ptr(true),
						Location: to.Ptr("atlantic"),
					},
					&Sawshark{
						Fishtype: to.Ptr("sawshark"),
						Length:   to.Ptr[float32](10),
						Siblings: []FishClassification{},
						Species:  to.Ptr("dangerous"),
						Age:      to.Ptr[int32](105),
						Birthday: &sawBday,
						Picture:  []byte{255, 255, 255, 255, 254},
					},
				},
				Species:  to.Ptr("predator"),
				Age:      to.Ptr[int32](6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.Ptr("sawshark"),
				Length:   to.Ptr[float32](10),
				Siblings: []FishClassification{},
				Species:  to.Ptr("dangerous"),
				Age:      to.Ptr[int32](105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
		},
		Species:  to.Ptr("king"),
		Iswild:   to.Ptr(true),
		Location: to.Ptr("alaska"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
