// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package objectgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newObjectTypeClient(t *testing.T) *ObjectTypeClient {
	client, err := NewObjectTypeClient(&azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewObjectTypeClient(options *azcore.ClientOptions) (*ObjectTypeClient, error) {
	client, err := azcore.NewClient("objectgroup.ObjectTypeClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ObjectTypeClient{internal: client}, nil
}

func TestGet(t *testing.T) {
	client := newObjectTypeClient(t)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(string(resp.RawJSON), `{ "message": "An object was successfully returned" }`); r != "" {
		t.Fatal(r)
	}
}

func TestPut(t *testing.T) {
	client := newObjectTypeClient(t)
	result, err := client.Put(context.Background(), []byte(`{ "foo": "bar" }`), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
