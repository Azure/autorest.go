// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

func newXMSClientRequestIDClient() *XMSClientRequestIDClient {
	return NewXMSClientRequestIDClient(NewDefaultConnection(nil))
}

// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func TestGet(t *testing.T) {
	client := newXMSClientRequestIDClient()
	result, err := client.Get(policy.WithHTTPHeader(context.Background(), http.Header{
		"x-ms-client-request-id": []string{"9C4D50EE-2D56-4CD3-8152-34347DC9F2B0"},
	}), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func TestParamGet(t *testing.T) {
	client := newXMSClientRequestIDClient()
	result, err := client.ParamGet(context.Background(), "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0", nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
