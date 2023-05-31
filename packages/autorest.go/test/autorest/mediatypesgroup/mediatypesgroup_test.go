// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package mediatypesgroup

import (
	"bytes"
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newMediaTypesClient(t *testing.T) *MediaTypesClient {
	client, err := NewMediaTypesClient(&azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewMediaTypesClient(options *azcore.ClientOptions) (*MediaTypesClient, error) {
	client, err := azcore.NewClient("mediatypesgroup.MediaTypesClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &MediaTypesClient{internal: client}, nil
}

func TestAnalyzeBody(t *testing.T) {
	client := newMediaTypesClient(t)
	body := streaming.NopCloser(bytes.NewReader([]byte("PDF")))
	result, err := client.AnalyzeBody(context.Background(), ContentTypeApplicationPDF, &MediaTypesClientAnalyzeBodyOptions{
		Input: body,
	})
	require.NoError(t, err)
	if *result.Value != "Nice job with PDF" {
		t.Fatalf("unexpected result %s", *result.Value)
	}
}

func TestAnalyzeBodyWithJSON(t *testing.T) {
	client := newMediaTypesClient(t)
	body := SourcePath{
		Source: to.Ptr("test"),
	}
	result, err := client.AnalyzeBodyWithJSON(context.Background(), &MediaTypesClientAnalyzeBodyWithJSONOptions{Input: &body})
	require.NoError(t, err)
	if *result.Value != "Nice job with JSON" {
		t.Fatalf("unexpected result %s", *result.Value)
	}
}

func TestContentTypeWithEncoding(t *testing.T) {
	client := newMediaTypesClient(t)
	result, err := client.ContentTypeWithEncoding(context.Background(), &MediaTypesClientContentTypeWithEncodingOptions{
		Input: to.Ptr("foo"),
	})
	require.NoError(t, err)
	if *result.Value != "Nice job sending content type with encoding" {
		t.Fatalf("unexpected result %s", *result.Value)
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
