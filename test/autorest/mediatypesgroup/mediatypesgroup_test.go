// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package mediatypesgroup

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func newMediaTypesClient() *MediaTypesClient {
	return NewMediaTypesClient(nil)
}

func TestAnalyzeBody(t *testing.T) {
	client := newMediaTypesClient()
	body := streaming.NopCloser(bytes.NewReader([]byte("PDF")))
	result, err := client.AnalyzeBody(context.Background(), ContentTypeApplicationPDF, &MediaTypesClientAnalyzeBodyOptions{
		Input: body,
	})
	if err != nil {
		t.Fatalf("AnalyzeBody: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestAnalyzeBodyWithJSON(t *testing.T) {
	client := newMediaTypesClient()
	body := SourcePath{
		Source: to.StringPtr("test"),
	}
	result, err := client.AnalyzeBodyWithJSON(context.Background(), &MediaTypesClientAnalyzeBodyWithJSONOptions{Input: &body})
	if err != nil {
		t.Fatalf("AnalyzeBodyWithSourcePath: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestContentTypeWithEncoding(t *testing.T) {
	client := newMediaTypesClient()
	result, err := client.ContentTypeWithEncoding(context.Background(), &MediaTypesClientContentTypeWithEncodingOptions{
		Input: to.StringPtr("foo"),
	})
	if err != nil {
		t.Fatalf("ContentTypeWithEncoding: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestBinaryBodyWithThreeContentTypes(t *testing.T) {
	t.Skip("no route")
}

func TestBinaryBodyWithThreeContentTypesWithText(t *testing.T) {
	t.Skip("no route")
}

func TestBinaryBodyWithTwoContentTypes(t *testing.T) {
	t.Skip("no route")
}

func TestPutTextAndJSONBodyWithJSON(t *testing.T) {
	t.Skip("no route")
}
