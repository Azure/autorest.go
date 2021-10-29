// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package binarygroup

import (
	"bytes"
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
)

func newBinaryGroupClient() *UploadClient {
	return NewUploadClient(nil)
}

func TestBinary(t *testing.T) {
	client := newBinaryGroupClient()
	resp, err := client.Binary(context.Background(), streaming.NopCloser(bytes.NewReader([]byte{0xff, 0xfe, 0xfd})), nil)
	if err != nil {
		t.Fatal(err)
	}
	if sc := resp.RawResponse.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestFile(t *testing.T) {
	client := newBinaryGroupClient()
	jsonFile := strings.NewReader(`{ "more": "cowbell" }`)
	resp, err := client.File(context.Background(), streaming.NopCloser(jsonFile), nil)
	if err != nil {
		t.Fatal(err)
	}
	if sc := resp.RawResponse.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}
