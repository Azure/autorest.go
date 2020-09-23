// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

func newXMSClientRequestIDClient() XMSClientRequestIDOperations {
	return NewXMSClientRequestIDClient(NewDefaultClient(nil))
}

// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func TestGet(t *testing.T) {
	client := newXMSClientRequestIDClient()
	result, err := client.Get(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	result, err = client.Get(azcore.WithHTTPHeader(context.Background(), http.Header{
		"x-ms-client-request-id": []string{"9C4D50EE-2D56-4CD3-8152-34347DC9F2B0"},
	}), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func TestParamGet(t *testing.T) {
	client := newXMSClientRequestIDClient()
	result, err := client.ParamGet(context.Background(), "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0", nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
