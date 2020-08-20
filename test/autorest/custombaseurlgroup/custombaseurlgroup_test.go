// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package custombaseurlgrouptest

import (
	"context"
	"generatortests/autorest/generated/custombaseurlgroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func getCustomBaseURLClient() custombaseurlgroup.PathsOperations {
	client := custombaseurlgroup.NewClient(to.StringPtr(":3000"), nil)
	return client.PathsOperations()
}

func TestGetEmpty(t *testing.T) {
	client := getCustomBaseURLClient()
	result, err := client.GetEmpty(context.Background(), "localhost")
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
