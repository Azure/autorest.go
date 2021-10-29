// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package custombaseurlgroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func newPathsClient() *PathsClient {
	return NewPathsClient(to.StringPtr(":3000"), nil)
}

func TestGetEmpty(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetEmpty(context.Background(), "localhost", nil)
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
