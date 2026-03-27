// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armtest_test

import (
	"armtest/v2"
	"armtest/v2/fake"
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func TestConfigurationsClient_GetStreamingContent(t *testing.T) {
	expectedBody := []byte("Configuration file content as a streaming response")
	srv := fake.ConfigurationsServer{
		GetStreamingContent: func(ctx context.Context, resourceGroupName string, configurationName string, options *armtest.ConfigurationsClientGetStreamingContentOptions) (resp azfake.Responder[armtest.ConfigurationsClientGetStreamingContentResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, armtest.ConfigurationsClientGetStreamingContentResponse{
				Body:        io.NopCloser(bytes.NewReader(expectedBody)),
				ContentType: to.Ptr("text/powershell"),
			}, nil)
			return
		},
	}

	clientFactory, err := armtest.NewClientFactory("00000000-0000-0000-0000-000000000000", &azfake.TokenCredential{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewConfigurationsServerTransport(&srv),
		},
	})
	if err != nil {
		t.Fatalf("failed to create client factory: %v", err)
	}

	resp, err := clientFactory.NewConfigurationsClient().GetStreamingContent(context.Background(), "myResourceGroup", "myConfiguration", nil)
	if err != nil {
		t.Fatalf("failed to finish the request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	defer resp.Body.Close()

	if !bytes.Equal(body, expectedBody) {
		t.Fatalf("expected body %q, got %q", expectedBody, body)
	}
	if resp.ContentType == nil || *resp.ContentType != "text/powershell" {
		t.Fatalf("expected content type %q, got %v", "text/powershell", resp.ContentType)
	}
}
