// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"reflect"
	"testing"
)

func newSkipURLEncodingClient() *SkipURLEncodingClient {
	return NewSkipURLEncodingClient(nil)
}

// GetMethodPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
func TestGetMethodPathValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetMethodPathValid(context.Background(), "path1/path2/path3", nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// GetMethodQueryNull - Get method with unencoded query parameter with value null
func TestGetMethodQueryNull(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetMethodQueryNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// GetMethodQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
func TestGetMethodQueryValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetMethodQueryValid(context.Background(), "value1&q2=value2&q3=value3", nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// GetPathQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
func TestGetPathQueryValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetPathQueryValid(context.Background(), "value1&q2=value2&q3=value3", nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// GetPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
func TestGetPathValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetPathValid(context.Background(), "path1/path2/path3", nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// GetSwaggerPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
func TestGetSwaggerPathValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetSwaggerPathValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// GetSwaggerQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
func TestGetSwaggerQueryValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetSwaggerQueryValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
