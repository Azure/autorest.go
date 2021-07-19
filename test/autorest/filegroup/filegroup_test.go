// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package filegroup

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"
)

func newFilesClient() *FilesClient {
	return NewFilesClient(NewDefaultConnection(nil))
}

func TestGetEmptyFile(t *testing.T) {
	client := newFilesClient()
	result, err := client.GetEmptyFile(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if result.RawResponse.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	if result.RawResponse.ContentLength != 0 {
		t.Fatalf("expected zero ContentLength, got %d", result.RawResponse.ContentLength)
	}
	result.RawResponse.Body.Close()
}

func TestGetFile(t *testing.T) {
	client := newFilesClient()
	result, err := client.GetFile(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if result.RawResponse.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	b, err := ioutil.ReadAll(result.RawResponse.Body)
	if err != nil {
		t.Fatal(err)
	}
	result.RawResponse.Body.Close()
	if l := len(b); l != 8725 {
		t.Fatalf("unexpected byte count: want 8725, got %d", l)
	}
}

func TestGetFileLarge(t *testing.T) {
	t.Skip("test is unreliable, can fail when running on a machine with low memory")
	client := newFilesClient()
	result, err := client.GetFileLarge(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if result.RawResponse.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	b, err := ioutil.ReadAll(result.RawResponse.Body)
	if err != nil {
		t.Fatal(err)
	}
	result.RawResponse.Body.Close()
	const size = 3000 * 1024 * 1024
	if l := len(b); l != size {
		t.Fatalf("unexpected byte count: want %d, got %d", size, l)
	}
}
