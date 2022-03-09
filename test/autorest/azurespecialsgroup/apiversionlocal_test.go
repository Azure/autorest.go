// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func newAPIVersionLocalClient() *APIVersionLocalClient {
	return NewAPIVersionLocalClient(nil)
}

// GetMethodLocalNull - Get method with api-version modeled in the method.  pass in api-version = null to succeed
func TestGetMethodLocalNull(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetMethodLocalNull(context.Background(), nil)
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// GetMethodLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetMethodLocalValid(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetMethodLocalValid(context.Background(), nil)
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// GetPathLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetPathLocalValid(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetPathLocalValid(context.Background(), nil)
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// GetSwaggerLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetSwaggerLocalValid(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetSwaggerLocalValid(context.Background(), nil)
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
