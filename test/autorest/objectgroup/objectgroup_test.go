// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package objectgroup

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newObjectTypeClient() *ObjectTypeClient {
	return NewObjectTypeClient(nil)
}

func TestGet(t *testing.T) {
	client := newObjectTypeClient()
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Interface, map[string]interface{}{
		"message": "An object was successfully returned",
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPut(t *testing.T) {
	client := newObjectTypeClient()
	result, err := client.Put(context.Background(), map[string]interface{}{
		"foo": "bar",
	}, nil)
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
