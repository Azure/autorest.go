// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package morecustombaseurigrouptest

import (
	"context"
	"generatortests/autorest/generated/morecustombaseurigroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func getMoreCustomBaseURIClient(t *testing.T) morecustombaseurigroup.PathsOperations {
	client, err := morecustombaseurigroup.NewClient(nil)
	if err != nil {
		t.Fatalf("failed to create more custom base URL client: %v", err)
	}
	// dnsSuffix string, subscriptionID string
	return client.PathsOperations(to.StringPtr(":3000"), "test12")
}

func TestGetEmpty(t *testing.T) {
	client := getMoreCustomBaseURIClient(t)
	// vault string, secret string, keyName string, options *PathsGetEmptyOptions
	result, err := client.GetEmpty(context.Background(), "http://localhost", "", "key1", &morecustombaseurigroup.PathsGetEmptyOptions{
		KeyVersion: to.StringPtr("v1"),
	})
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
