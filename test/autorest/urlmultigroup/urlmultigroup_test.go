// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlmultigroup

import (
	"context"
	"net/url"
	"reflect"
	"testing"
)

func newQueriesClient() *QueriesClient {
	return NewQueriesClient(nil)
}

func TestURLMultiArrayStringMultiEmpty(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringMultiEmpty(context.Background(), &QueriesClientArrayStringMultiEmptyOptions{
		ArrayQuery: []string{},
	})
	if err != nil {
		t.Fatalf("ArrayStringMultiEmpty: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestURLMultiArrayStringMultiNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringMultiNull(context.Background(), &QueriesClientArrayStringMultiNullOptions{
		ArrayQuery: nil,
	})
	if err != nil {
		t.Fatalf("ArrayStringMultiNull: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestURLMultiArrayStringMultiValid(t *testing.T) {
	t.Skip("Cannot set nil for string value in string slice")
	client := newQueriesClient()
	result, err := client.ArrayStringMultiValid(context.Background(), &QueriesClientArrayStringMultiValidOptions{
		ArrayQuery: []string{
			"ArrayQuery1",
			url.QueryEscape("begin!*'();:@ &=+$,/?#[]end"),
			"",
			""},
	})
	if err != nil {
		t.Fatalf("ArrayStringMultiValid: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
