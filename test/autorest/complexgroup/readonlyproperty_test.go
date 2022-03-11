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

func newReadonlypropertyClient() *ReadonlypropertyClient {
	return NewReadonlypropertyClient(nil)
}

func TestReadonlypropertyGetValid(t *testing.T) {
	client := newReadonlypropertyClient()
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.ReadonlyObj, ReadonlyObj{ID: to.StringPtr("1234"), Size: to.Int32Ptr(2)}); r != "" {
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
