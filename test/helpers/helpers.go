// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package helpers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

// DeepEqualOrFatal fails the test if the result and expected aren't equal.
// reflect.DeepEqual is used to make the comparison.
func DeepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		// spit out in JSON format to make it easier to read
		// if marshalling fails then fall back to %+v
		resultJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			t.Fatalf("got %+v, want %+v", result, expected)
		}
		expectedJSON, err := json.MarshalIndent(expected, "", "  ")
		if err != nil {
			t.Fatalf("got %+v, want %+v", result, expected)
		}
		t.Fatalf("got %+v, want %+v", resultJSON, expectedJSON)
	}
}

// VerifyStatusCode fails the test if the response's status code doesn't match the expected status code.
func VerifyStatusCode(t *testing.T, resp *http.Response, expected int) {
	if resp == nil {
		t.Fatal("unexpected nil response")
	}
	if resp.StatusCode != expected {
		t.Fatalf("bad status code: got %d, want %d", resp.StatusCode, expected)
	}
}
