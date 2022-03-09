// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package binarygroup

import (
	"bytes"
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/stretchr/testify/require"
)

func newBinaryGroupClient() *UploadClient {
	return NewUploadClient(nil)
}

func TestBinary(t *testing.T) {
	client := newBinaryGroupClient()
	resp, err := client.Binary(context.Background(), streaming.NopCloser(bytes.NewReader([]byte{0xff, 0xfe, 0xfd})), nil)
	require.NoError(t, err)
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestFile(t *testing.T) {
	client := newBinaryGroupClient()
	jsonFile := strings.NewReader(`{ "more": "cowbell" }`)
	resp, err := client.File(context.Background(), streaming.NopCloser(jsonFile), nil)
	require.NoError(t, err)
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
