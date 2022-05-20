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

func newReadonlypropertyClient() *ReadonlypropertyClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewReadonlypropertyClient(pl)
}

func TestReadonlypropertyGetValid(t *testing.T) {
	client := newReadonlypropertyClient()
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.ReadonlyObj, ReadonlyObj{ID: to.Ptr("1234"), Size: to.Ptr[int32](2)}); r != "" {
		t.Fatal(r)
	}
}

func TestReadonlypropertyPutValid(t *testing.T) {
	client := newReadonlypropertyClient()
	id, size := "1234", int32(2)
	result, err := client.PutValid(context.Background(), ReadonlyObj{ID: &id, Size: &size}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
