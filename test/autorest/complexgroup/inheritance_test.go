// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newInheritanceClient() *InheritanceClient {
	return NewInheritanceClient(nil)
}

func TestInheritanceGetValid(t *testing.T) {
	client := newInheritanceClient()
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Siamese, Siamese{
		ID:    to.Int32Ptr(2),
		Name:  to.StringPtr("Siameeee"),
		Color: to.StringPtr("green"),
		Hates: []*Dog{
			{
				ID:   to.Int32Ptr(1),
				Name: to.StringPtr("Potato"),
				Food: to.StringPtr("tomato"),
			},
			{
				ID:   to.Int32Ptr(-1),
				Name: to.StringPtr("Tomato"),
				Food: to.StringPtr("french fries"),
			},
		},
		Breed: to.StringPtr("persian"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestInheritancePutValid(t *testing.T) {
	client := newInheritanceClient()
	result, err := client.PutValid(context.Background(), Siamese{
		ID:    to.Int32Ptr(2),
		Name:  to.StringPtr("Siameeee"),
		Color: to.StringPtr("green"),
		Hates: []*Dog{
			{
				ID:   to.Int32Ptr(1),
				Name: to.StringPtr("Potato"),
				Food: to.StringPtr("tomato"),
			},
			{
				ID:   to.Int32Ptr(-1),
				Name: to.StringPtr("Tomato"),
				Food: to.StringPtr("french fries"),
			},
		},
		Breed: to.StringPtr("persian"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
