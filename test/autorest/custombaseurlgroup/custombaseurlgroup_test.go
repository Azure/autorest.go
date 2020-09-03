// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package custombaseurlgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newPathsClient() PathsOperations {
	client := NewClient(to.StringPtr(":3000"), nil)
	return NewPathsClient(client)
}

func TestGetEmpty(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetEmpty(context.Background(), "localhost")
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
