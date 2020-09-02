// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonstringenumgrouptest

import (
	"context"
	"generatortests/autorest/generated/nonstringenumgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func newFloatClient() nonstringenumgroup.FloatOperations {
	return nonstringenumgroup.NewFloatClient(nonstringenumgroup.NewDefaultClient(nil))
}

// Get - Get a float enum
func TestFloatGet(t *testing.T) {
	client := newFloatClient()
	result, err := client.Get(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, nonstringenumgroup.FloatEnumFourHundredTwentyNine1.ToPtr())
}

// Put - Put a float enum
func TestFloatPut(t *testing.T) {
	client := newFloatClient()
	result, err := client.Put(context.Background(), &nonstringenumgroup.FloatPutOptions{
		Input: nonstringenumgroup.FloatEnumTwoHundred4.ToPtr(),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}
