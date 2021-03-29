// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package formdatagroup

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

func newFormdataClient() *FormdataClient {
	return NewFormdataClient(NewDefaultConnection(nil))
}

func TestUploadFile(t *testing.T) {
	client := newFormdataClient()
	s := strings.NewReader("the data")
	resp, err := client.UploadFile(context.Background(), azcore.NopCloser(s), "sample", nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "the data" {
		t.Fatalf("unexpected result %s", string(b))
	}
}

func TestUploadFileViaBody(t *testing.T) {
	client := newFormdataClient()
	s := strings.NewReader("the data")
	resp, err := client.UploadFileViaBody(context.Background(), azcore.NopCloser(s), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "the data" {
		t.Fatalf("unexpected result %s", string(b))
	}
}

func TestUploadFiles(t *testing.T) {
	t.Skip("missing route in test server")
	client := newFormdataClient()
	s1 := strings.NewReader("the data")
	s2 := strings.NewReader(" to be uploaded")
	resp, err := client.UploadFiles(context.Background(), []azcore.ReadSeekCloser{
		azcore.NopCloser(s1),
		azcore.NopCloser(s2),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	// TODO: verify response body
}
