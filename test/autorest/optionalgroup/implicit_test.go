// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package optionalgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func newImplicitClient() ImplicitOperations {
	return NewImplicitClient(NewDefaultConnection(nil), "", "", nil)
}

func TestImplicitGetOptionalGlobalQuery(t *testing.T) {
	client := newImplicitClient()
	result, err := client.GetOptionalGlobalQuery(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetOptionalGlobalQuery: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestImplicitGetRequiredGlobalPath(t *testing.T) {
	t.Skip("Cannot set nil for string parameter so test invalid for Go")
	client := newImplicitClient()
	result, err := client.GetRequiredGlobalPath(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetRequiredGlobalPath: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestImplicitGetRequiredGlobalQuery(t *testing.T) {
	t.Skip("Cannot set nil for string parameter so test invalid for Go")
	client := newImplicitClient()
	result, err := client.GetRequiredGlobalQuery(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetRequiredGlobalQuery: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestImplicitGetRequiredPath(t *testing.T) {
	t.Skip("Cannot set nil for string parameter so test invalid for Go")
	client := newImplicitClient()
	result, err := client.GetRequiredPath(context.Background(), "", nil)
	if err != nil {
		t.Fatalf("GetRequiredPath: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestImplicitPutOptionalBody(t *testing.T) {
	client := newImplicitClient()
	result, err := client.PutOptionalBody(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutOptionalBody: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestImplicitPutOptionalHeader(t *testing.T) {
	client := newImplicitClient()
	result, err := client.PutOptionalHeader(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutOptionalHeader: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestImplicitPutOptionalQuery(t *testing.T) {
	client := newImplicitClient()
	result, err := client.PutOptionalQuery(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutOptionalQuery: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
