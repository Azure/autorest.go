// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonstringenumgroup

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newIntClient() *IntClient {
	return NewIntClient(nil)
}

// Get - Get an int enum
func TestIntGet(t *testing.T) {
	client := newIntClient()
	result, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, IntEnumFourHundredTwentyNine.ToPtr()); r != "" {
		t.Fatal(r)
	}
}

// Put - Put an int enum
func TestIntPut(t *testing.T) {
	client := newIntClient()
	result, err := client.Put(context.Background(), &IntClientPutOptions{
		Input: IntEnumTwoHundred.ToPtr(),
	})
	require.NoError(t, err)
	if *result.Value != "Nice job posting an int enum" {
		t.Fatalf("unexpected value %s", *result.Value)
	}
}
