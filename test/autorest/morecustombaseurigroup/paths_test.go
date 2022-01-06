// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package morecustombaseurigroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func newPathsClient() *PathsClient {
	// dnsSuffix string, subscriptionID string
	return NewPathsClient("test12", &PathsClientOptions{
		DnsSuffix: to.StringPtr(":3000"),
	})
}

func TestGetEmpty(t *testing.T) {
	client := newPathsClient()
	// vault string, secret string, keyName string, options *PathsGetEmptyOptions
	result, err := client.GetEmpty(context.Background(), "http://localhost", "", "key1", &PathsClientGetEmptyOptions{
		KeyVersion: to.StringPtr("v1"),
	})
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
