// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newInheritanceClient(t *testing.T) *InheritanceClient {
	client, err := NewInheritanceClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewInheritanceClient(endpoint string, options *azcore.ClientOptions) (*InheritanceClient, error) {
	client, err := azcore.NewClient("complexgroup.InheritanceClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &InheritanceClient{internal: client, endpoint: endpoint}, nil
}

func TestInheritanceGetValid(t *testing.T) {
	client := newInheritanceClient(t)
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Siamese, Siamese{
		ID:    to.Ptr[int32](2),
		Name:  to.Ptr("Siameeee"),
		Color: to.Ptr("green"),
		Hates: []*Dog{
			{
				ID:   to.Ptr[int32](1),
				Name: to.Ptr("Potato"),
				Food: to.Ptr("tomato"),
			},
			{
				ID:   to.Ptr[int32](-1),
				Name: to.Ptr("Tomato"),
				Food: to.Ptr("french fries"),
			},
		},
		Breed: to.Ptr("persian"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestInheritancePutValid(t *testing.T) {
	client := newInheritanceClient(t)
	result, err := client.PutValid(context.Background(), Siamese{
		ID:    to.Ptr[int32](2),
		Name:  to.Ptr("Siameeee"),
		Color: to.Ptr("green"),
		Hates: []*Dog{
			{
				ID:   to.Ptr[int32](1),
				Name: to.Ptr("Potato"),
				Food: to.Ptr("tomato"),
			},
			{
				ID:   to.Ptr[int32](-1),
				Name: to.Ptr("Tomato"),
				Food: to.Ptr("french fries"),
			},
		},
		Breed: to.Ptr("persian"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
