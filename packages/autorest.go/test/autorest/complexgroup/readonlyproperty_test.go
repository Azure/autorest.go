// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newReadonlypropertyClient(t *testing.T) *ReadonlypropertyClient {
	client, err := NewReadonlypropertyClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func TestReadonlypropertyGetValid(t *testing.T) {
	client := newReadonlypropertyClient(t)
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.ReadonlyObj, ReadonlyObj{ID: to.Ptr("1234"), Size: to.Ptr[int32](2)}); r != "" {
		t.Fatal(r)
	}
}

func TestReadonlypropertyPutValid(t *testing.T) {
	client := newReadonlypropertyClient(t)
	result, err := client.PutValid(context.Background(), ReadonlyObj{Size: to.Ptr[int32](2)}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
