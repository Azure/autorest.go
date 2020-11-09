// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package mediatypesgroup

import (
	"bytes"
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newMediaTypesClient() MediaTypesClientOperations {
	return NewMediaTypesClient(NewDefaultConnection(nil))
}

func TestAnalyzeBody(t *testing.T) {
	client := newMediaTypesClient()
	body := azcore.NopCloser(bytes.NewReader([]byte("PDF")))
	result, err := client.AnalyzeBody(context.Background(), ContentTypeApplicationPDF, body, nil)
	if err != nil {
		t.Fatalf("AnalyzeBody: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestAnalyzeBodyWithSourcePath(t *testing.T) {
	client := newMediaTypesClient()
	body := SourcePath{
		Source: to.StringPtr("test"),
	}
	result, err := client.AnalyzeBodyWithSourcePath(context.Background(), &MediaTypesClientAnalyzeBodyWithSourcePathOptions{Input: &body})
	if err != nil {
		t.Fatalf("AnalyzeBodyWithSourcePath: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestContentTypeWithEncoding(t *testing.T) {
	client := newMediaTypesClient()
	result, err := client.ContentTypeWithEncoding(context.Background(), "foo")
	if err != nil {
		t.Fatalf("ContentTypeWithEncoding: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}
