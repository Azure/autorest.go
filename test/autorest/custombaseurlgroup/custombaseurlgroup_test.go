// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package custombaseurlgroup

import (
	"context"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func newPathsClient() *PathsClient {
	return NewPathsClient(&PathsClientOptions{
		Host: to.StringPtr(":3000"),
	})
}

func TestGetEmpty(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetEmpty(context.Background(), "localhost", nil)
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
