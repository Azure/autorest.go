// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexmodelgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newComplexModelClient() *ComplexModelClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewComplexModelClient(pl)
}

func TestCreate(t *testing.T) {
	client := newComplexModelClient()
	_, err := client.Create(context.Background(), "sub", "rg", CatalogDictionaryOfArray{}, nil)
	require.Error(t, err)
}

func TestList(t *testing.T) {
	client := newComplexModelClient()
	_, err := client.List(context.Background(), "", nil)
	require.Error(t, err)
}

func TestUpdate(t *testing.T) {
	client := newComplexModelClient()
	_, err := client.Update(context.Background(), "", "", CatalogArrayOfDictionary{}, nil)
	require.Error(t, err)
}
