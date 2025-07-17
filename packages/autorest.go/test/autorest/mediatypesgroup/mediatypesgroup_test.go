// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package mediatypesgroup

import (
	"bytes"
	"context"
	"generatortests"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newMediaTypesClient(t *testing.T) *MediaTypesClient {
	client, err := NewMediaTypesClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
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

func TestBodyThreeTypes(t *testing.T) {
	client := newMediaTypesClient(t)
	result, err := client.BodyThreeTypes(context.Background(), streaming.NopCloser(bytes.NewReader([]byte("foo"))), nil)
	require.NoError(t, err)
	require.NotNil(t, result.Value)
	require.Empty(t, *result.Value)
}

func TestBodyThreeTypesWithJSON(t *testing.T) {
	client := newMediaTypesClient(t)
	type hello struct {
		Hello string `json:"hello"`
	}
	result, err := client.BodyThreeTypesWithJSON(context.Background(), hello{
		Hello: "world",
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, result.Value)
	require.Empty(t, *result.Value)
}

func TestBodyThreeTypesWithText(t *testing.T) {
	client := newMediaTypesClient(t)
	result, err := client.BodyThreeTypesWithText(context.Background(), "hello, world", nil)
	require.NoError(t, err)
	require.NotNil(t, result.Value)
	require.Empty(t, *result.Value)
}

func TestBinaryBodyWithThreeContentTypes(t *testing.T) {
	client := newMediaTypesClient(t)
	result, err := client.BinaryBodyWithThreeContentTypes(context.Background(), ContentType2ApplicationJSON, streaming.NopCloser(strings.NewReader(`{"hello":"world"}`)), nil)
	require.NoError(t, err)
	require.NotNil(t, result.Value)
	require.Empty(t, *result.Value)

	result, err = client.BinaryBodyWithThreeContentTypes(context.Background(), ContentType2ApplicationOctetStream, streaming.NopCloser(bytes.NewReader([]byte("foo"))), nil)
	require.NoError(t, err)
	require.NotNil(t, result.Value)
	require.Empty(t, *result.Value)

	result, err = client.BinaryBodyWithThreeContentTypes(context.Background(), ContentType2TextPlain, streaming.NopCloser(strings.NewReader("hello, world")), nil)
	require.NoError(t, err)
	require.NotNil(t, result.Value)
	require.Empty(t, *result.Value)
}

func TestBinaryBodyWithTwoContentTypes(t *testing.T) {
	client := newMediaTypesClient(t)
	result, err := client.BinaryBodyWithTwoContentTypes(context.Background(), ContentType1ApplicationJSON, streaming.NopCloser(strings.NewReader(`{"hello":"world"}`)), nil)
	require.NoError(t, err)
	require.NotNil(t, result.Value)
	require.Empty(t, *result.Value)

	result, err = client.BinaryBodyWithTwoContentTypes(context.Background(), ContentType1ApplicationOctetStream, streaming.NopCloser(bytes.NewReader([]byte("foo"))), nil)
	require.NoError(t, err)
	require.NotNil(t, result.Value)
	require.Empty(t, *result.Value)
}

func TestPutTextAndJSONBody(t *testing.T) {
	t.Skip("no route")
}
