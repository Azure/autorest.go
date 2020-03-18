// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package custombaseurlgrouptest

import (
	"context"
	"generatortests/autorest/generated/custombaseurlgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getCustomBaseURLClient(t *testing.T) custombaseurlgroup.PathsOperations {
	client, err := custombaseurlgroup.NewClient("http://localhost:3000", nil)
	if err != nil {
		t.Fatalf("failed to create custom base URL client: %v", err)
	}
	return client.PathsOperations()
}

func TestGetEmpty(t *testing.T) {
	client := getCustomBaseURLClient(t)
	result, err := client.GetEmpty(context.Background(), "")
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
