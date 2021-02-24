// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonstringenumgroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newIntClient() *IntClient {
	return NewIntClient(NewDefaultConnection(nil))
}

// Get - Get an int enum
func TestIntGet(t *testing.T) {
	client := newIntClient()
	result, err := client.Get(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Value, IntEnumFourHundredTwentyNine.ToPtr()); r != "" {
		t.Fatal(r)
	}
}

// Put - Put an int enum
func TestIntPut(t *testing.T) {
	client := newIntClient()
	result, err := client.Put(context.Background(), &IntPutOptions{
		Input: IntEnumTwoHundred.ToPtr(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
