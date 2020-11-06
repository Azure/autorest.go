// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headgroup

import (
	"context"
	"testing"
)

func newHTTPSuccessOperationsClient() HTTPSuccessOperations {
	return NewHTTPSuccessClient(NewDefaultConnection(nil))
}

// Head200 - Return 200 status code if successful
func TestHead200(t *testing.T) {
	client := newHTTPSuccessOperationsClient()
	resp, err := client.Head200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Success {
		t.Fatal("expected success")
	}
}

// Head204 - Return 204 status code if successful
func TestHead204(t *testing.T) {
	client := newHTTPSuccessOperationsClient()
	resp, err := client.Head204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Success {
		t.Fatal("expected success")
	}
}

// Head404 - Return 404 status code if successful
func TestHead404(t *testing.T) {
	client := newHTTPSuccessOperationsClient()
	resp, err := client.Head404(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Success {
		t.Fatal("expected non-success")
	}
}
