// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newPolymorphicrecursiveClient() PolymorphicrecursiveOperations {
	return NewPolymorphicrecursiveClient(NewDefaultConnection(nil))
}

// GetValid - Get complex types that are polymorphic and have recursive references
func TestGetValid(t *testing.T) {
	client := newPolymorphicrecursiveClient()
	result, err := client.GetValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	helpers.DeepEqualOrFatal(t, result.Fish, &Salmon{
		Fish: Fish{
			Fishtype: to.StringPtr("salmon"),
			Length:   to.Float32Ptr(1),
			Siblings: &[]FishClassification{
				&Shark{
					Fish: Fish{
						Fishtype: to.StringPtr("shark"),
						Length:   to.Float32Ptr(20),
						Siblings: &[]FishClassification{
							&Salmon{
								Fish: Fish{
									Fishtype: to.StringPtr("salmon"),
									Length:   to.Float32Ptr(2),
									Siblings: &[]FishClassification{
										&Shark{
											Fish: Fish{
												Fishtype: to.StringPtr("shark"),
												Length:   to.Float32Ptr(20),
												Species:  to.StringPtr("predator"),
											},
											Age:      to.Int32Ptr(6),
											Birthday: &sharkBday,
										},
										&Sawshark{
											Shark: Shark{
												Fish: Fish{
													Fishtype: to.StringPtr("sawshark"),
													Length:   to.Float32Ptr(10),
													Species:  to.StringPtr("dangerous"),
												},
												Age:      to.Int32Ptr(105),
												Birthday: &sawBday,
											},
											Picture: &[]byte{255, 255, 255, 255, 254},
										},
									},
									Species: to.StringPtr("coho"),
								},
								Iswild:   to.BoolPtr(true),
								Location: to.StringPtr("atlantic"),
							},
							&Sawshark{
								Shark: Shark{
									Fish: Fish{
										Fishtype: to.StringPtr("sawshark"),
										Length:   to.Float32Ptr(10),
										Siblings: &[]FishClassification{},
										Species:  to.StringPtr("dangerous"),
									},
									Age:      to.Int32Ptr(105),
									Birthday: &sawBday,
								},
								Picture: &[]byte{255, 255, 255, 255, 254},
							},
						},
						Species: to.StringPtr("predator"),
					},
					Age:      to.Int32Ptr(6),
					Birthday: &sharkBday,
				},
				&Sawshark{
					Shark: Shark{
						Fish: Fish{
							Fishtype: to.StringPtr("sawshark"),
							Length:   to.Float32Ptr(10),
							Siblings: &[]FishClassification{},
							Species:  to.StringPtr("dangerous"),
						},
						Age:      to.Int32Ptr(105),
						Birthday: &sawBday,
					},
					Picture: &[]byte{255, 255, 255, 255, 254},
				},
			},
			Species: to.StringPtr("king"),
		},
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
	})
}

// PutValid - Put complex types that are polymorphic and have recursive references
func TestPutValid(t *testing.T) {
	client := newPolymorphicrecursiveClient()
	sawBday := time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)
	sharkBday := time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)
	result, err := client.PutValid(context.Background(), &Salmon{
		Fish: Fish{
			Fishtype: to.StringPtr("salmon"),
			Length:   to.Float32Ptr(1),
			Siblings: &[]FishClassification{
				&Shark{
					Fish: Fish{
						Fishtype: to.StringPtr("shark"),
						Length:   to.Float32Ptr(20),
						Siblings: &[]FishClassification{
							&Salmon{
								Fish: Fish{
									Fishtype: to.StringPtr("salmon"),
									Length:   to.Float32Ptr(2),
									Siblings: &[]FishClassification{
										&Shark{
											Fish: Fish{
												Fishtype: to.StringPtr("shark"),
												Length:   to.Float32Ptr(20),
												Species:  to.StringPtr("predator"),
											},
											Age:      to.Int32Ptr(6),
											Birthday: &sharkBday,
										},
										&Sawshark{
											Shark: Shark{
												Fish: Fish{
													Fishtype: to.StringPtr("sawshark"),
													Length:   to.Float32Ptr(10),
													Species:  to.StringPtr("dangerous"),
												},
												Age:      to.Int32Ptr(105),
												Birthday: &sawBday,
											},
											Picture: &[]byte{255, 255, 255, 255, 254},
										},
									},
									Species: to.StringPtr("coho"),
								},
								Iswild:   to.BoolPtr(true),
								Location: to.StringPtr("atlantic"),
							},
							&Sawshark{
								Shark: Shark{
									Fish: Fish{
										Fishtype: to.StringPtr("sawshark"),
										Length:   to.Float32Ptr(10),
										Siblings: &[]FishClassification{},
										Species:  to.StringPtr("dangerous"),
									},
									Age:      to.Int32Ptr(105),
									Birthday: &sawBday,
								},
								Picture: &[]byte{255, 255, 255, 255, 254},
							},
						},
						Species: to.StringPtr("predator"),
					},
					Age:      to.Int32Ptr(6),
					Birthday: &sharkBday,
				},
				&Sawshark{
					Shark: Shark{
						Fish: Fish{
							Fishtype: to.StringPtr("sawshark"),
							Length:   to.Float32Ptr(10),
							Siblings: &[]FishClassification{},
							Species:  to.StringPtr("dangerous"),
						},
						Age:      to.Int32Ptr(105),
						Birthday: &sawBday,
					},
					Picture: &[]byte{255, 255, 255, 255, 254},
				},
			},
			Species: to.StringPtr("king"),
		},
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
