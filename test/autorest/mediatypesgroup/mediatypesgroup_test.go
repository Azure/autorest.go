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

func newMediaTypesClient() *MediaTypesClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewMediaTypesClient(pl)
}

func TestAnalyzeBody(t *testing.T) {
	client := newMediaTypesClient()
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
	client := newMediaTypesClient()
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
	client := newMediaTypesClient()
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
