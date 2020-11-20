// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package morecustombaseurigroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newPathsClient() PathsClient {
	client := NewConnection(to.StringPtr(":3000"), nil)
	// dnsSuffix string, subscriptionID string
	return NewPathsClient(client, "test12")
}

func TestGetEmpty(t *testing.T) {
	client := newPathsClient()
	// vault string, secret string, keyName string, options *PathsGetEmptyOptions
	result, err := client.GetEmpty(context.Background(), "http://localhost", "", "key1", &PathsGetEmptyOptions{
		KeyVersion: to.StringPtr("v1"),
	})
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
