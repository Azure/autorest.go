// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexmodelgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func newComplexModelClient(t *testing.T) *ComplexModelClient {
	client, err := NewComplexModelClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func TestCreate(t *testing.T) {
	client := newComplexModelClient(t)
	_, err := client.Create(context.Background(), "sub", "rg", CatalogDictionaryOfArray{}, nil)
	require.Error(t, err)
}

func TestList(t *testing.T) {
	client := newComplexModelClient(t)
	_, err := client.List(context.Background(), "", nil)
	require.Error(t, err)
}

func TestUpdate(t *testing.T) {
	client := newComplexModelClient(t)
	_, err := client.Update(context.Background(), "", "", CatalogArrayOfDictionary{}, nil)
	require.Error(t, err)
}
