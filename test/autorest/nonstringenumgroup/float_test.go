// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonstringenumgroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newFloatClient() *FloatClient {
	return NewFloatClient(NewDefaultConnection(nil))
}

// Get - Get a float enum
func TestFloatGet(t *testing.T) {
	client := newFloatClient()
	result, err := client.Get(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Value, FloatEnumFourHundredTwentyNine1.ToPtr()); r != "" {
		t.Fatal(r)
	}
}

// Put - Put a float enum
func TestFloatPut(t *testing.T) {
	client := newFloatClient()
	result, err := client.Put(context.Background(), &FloatPutOptions{
		Input: FloatEnumTwoHundred4.ToPtr(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
