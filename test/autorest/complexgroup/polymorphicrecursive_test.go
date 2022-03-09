// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newPolymorphicrecursiveClient() *PolymorphicrecursiveClient {
	return NewPolymorphicrecursiveClient(nil)
}

// GetValid - Get complex types that are polymorphic and have recursive references
func TestGetValid(t *testing.T) {
	client := newPolymorphicrecursiveClient()
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	if r := cmp.Diff(result.FishClassification, &Salmon{
		Fishtype: to.StringPtr("salmon"),
		Length:   to.Float32Ptr(1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.StringPtr("shark"),
				Length:   to.Float32Ptr(20),
				Siblings: []FishClassification{
					&Salmon{
						Fishtype: to.StringPtr("salmon"),
						Length:   to.Float32Ptr(2),
						Siblings: []FishClassification{
							&Shark{
								Fishtype: to.StringPtr("shark"),
								Length:   to.Float32Ptr(20),
								Species:  to.StringPtr("predator"),
								Age:      to.Int32Ptr(6),
								Birthday: &sharkBday,
							},
							&Sawshark{
								Fishtype: to.StringPtr("sawshark"),
								Length:   to.Float32Ptr(10),
								Species:  to.StringPtr("dangerous"),
								Age:      to.Int32Ptr(105),
								Birthday: &sawBday,
								Picture:  []byte{255, 255, 255, 255, 254},
							},
						},
						Species:  to.StringPtr("coho"),
						Iswild:   to.BoolPtr(true),
						Location: to.StringPtr("atlantic"),
					},
					&Sawshark{
						Fishtype: to.StringPtr("sawshark"),
						Length:   to.Float32Ptr(10),
						Siblings: []FishClassification{},
						Species:  to.StringPtr("dangerous"),
						Age:      to.Int32Ptr(105),
						Birthday: &sawBday,
						Picture:  []byte{255, 255, 255, 255, 254},
					},
				},
				Species:  to.StringPtr("predator"),
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.StringPtr("sawshark"),
				Length:   to.Float32Ptr(10),
				Siblings: []FishClassification{},
				Species:  to.StringPtr("dangerous"),
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
		},
		Species:  to.StringPtr("king"),
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
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
		Fishtype: to.StringPtr("salmon"),
		Length:   to.Float32Ptr(1),
		Siblings: []FishClassification{
			&Shark{
				Fishtype: to.StringPtr("shark"),
				Length:   to.Float32Ptr(20),
				Siblings: []FishClassification{
					&Salmon{
						Fishtype: to.StringPtr("salmon"),
						Length:   to.Float32Ptr(2),
						Siblings: []FishClassification{
							&Shark{
								Fishtype: to.StringPtr("shark"),
								Length:   to.Float32Ptr(20),
								Species:  to.StringPtr("predator"),
								Age:      to.Int32Ptr(6),
								Birthday: &sharkBday,
							},
							&Sawshark{
								Fishtype: to.StringPtr("sawshark"),
								Length:   to.Float32Ptr(10),
								Species:  to.StringPtr("dangerous"),
								Age:      to.Int32Ptr(105),
								Birthday: &sawBday,
								Picture:  []byte{255, 255, 255, 255, 254},
							},
						},
						Species:  to.StringPtr("coho"),
						Iswild:   to.BoolPtr(true),
						Location: to.StringPtr("atlantic"),
					},
					&Sawshark{
						Fishtype: to.StringPtr("sawshark"),
						Length:   to.Float32Ptr(10),
						Siblings: []FishClassification{},
						Species:  to.StringPtr("dangerous"),
						Age:      to.Int32Ptr(105),
						Birthday: &sawBday,
						Picture:  []byte{255, 255, 255, 255, 254},
					},
				},
				Species:  to.StringPtr("predator"),
				Age:      to.Int32Ptr(6),
				Birthday: &sharkBday,
			},
			&Sawshark{
				Fishtype: to.StringPtr("sawshark"),
				Length:   to.Float32Ptr(10),
				Siblings: []FishClassification{},
				Species:  to.StringPtr("dangerous"),
				Age:      to.Int32Ptr(105),
				Birthday: &sawBday,
				Picture:  []byte{255, 255, 255, 255, 254},
			},
		},
		Species:  to.StringPtr("king"),
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
	}, nil)
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
