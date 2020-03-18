// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package filegrouptest

import (
	"context"
	"generatortests/autorest/generated/filegroup"
	"generatortests/helpers"
	"io/ioutil"
	"net/http"
	"testing"
)

func getFileGroupClient(t *testing.T) filegroup.FilesOperations {
	client, err := filegroup.NewClient("http://localhost:3000", nil)
	if err != nil {
		t.Fatalf("failed to create more custom base URL client: %v", err)
	}
	return client.FilesOperations()
}

func TestGetEmptyFile(t *testing.T) {
	client := getFileGroupClient(t)
	result, err := client.GetEmptyFile(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
	if result.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	if result.ContentLength != 0 {
		t.Fatalf("expected zero ContentLength, got %d", result.ContentLength)
	}
	result.Body.Close()
}

func TestGetFile(t *testing.T) {
	client := getFileGroupClient(t)
	result, err := client.GetFile(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
	if result.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	b, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	result.Body.Close()
	if l := len(b); l != 8725 {
		t.Fatalf("unexpected byte count: want 8725, got %d", l)
	}
}

func TestGetFileLarge(t *testing.T) {
	client := getFileGroupClient(t)
	result, err := client.GetFileLarge(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
	if result.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	b, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	result.Body.Close()
	const size = 3000 * 1024 * 1024
	if l := len(b); l != size {
		t.Fatalf("unexpected byte count: want %d, got %d", size, l)
	}
}
