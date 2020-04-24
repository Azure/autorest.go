// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgrouptest

import (
	"context"
	"generatortests/autorest/generated/azurespecialsgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getXMSClientRequestIDOperations(t *testing.T) azurespecialsgroup.XMSClientRequestIDOperations {
	client, err := azurespecialsgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client.XMSClientRequestIDOperations()
}

// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func TestGet(t *testing.T) {
	t.Skip("cannot set custom header values")
}

// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func TestParamGet(t *testing.T) {
	client := getXMSClientRequestIDOperations(t)
	result, err := client.ParamGet(context.Background(), "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0")
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
