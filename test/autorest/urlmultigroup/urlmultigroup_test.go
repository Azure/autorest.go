// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlmultigrouptest

import (
	"context"
	"generatortests/autorest/generated/urlmultigroup"
	"generatortests/helpers"
	"net/http"
	"net/url"
	"testing"
)

func newQueriesClient() urlmultigroup.QueriesOperations {
	return urlmultigroup.NewQueriesClient(urlmultigroup.NewDefaultClient(nil))
}

func TestURLMultiArrayStringMultiEmpty(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringMultiEmpty(context.Background(), &urlmultigroup.QueriesArrayStringMultiEmptyOptions{
		ArrayQuery: &[]string{},
	})
	if err != nil {
		t.Fatalf("ArrayStringMultiEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestURLMultiArrayStringMultiNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringMultiNull(context.Background(), &urlmultigroup.QueriesArrayStringMultiNullOptions{
		ArrayQuery: nil,
	})
	if err != nil {
		t.Fatalf("ArrayStringMultiNull: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestURLMultiArrayStringMultiValid(t *testing.T) {
	t.Skip("Cannot set nil for string value in string slice")
	client := newQueriesClient()
	result, err := client.ArrayStringMultiValid(context.Background(), &urlmultigroup.QueriesArrayStringMultiValidOptions{
		ArrayQuery: &[]string{"ArrayQuery1", url.QueryEscape("begin!*'();:@ &=+$,/?#[]end"), "", ""},
	})
	if err != nil {
		t.Fatalf("ArrayStringMultiValid: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
