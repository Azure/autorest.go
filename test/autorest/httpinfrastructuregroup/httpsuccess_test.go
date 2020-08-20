// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregrouptest

import (
	"context"
	"generatortests/autorest/generated/httpinfrastructuregroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func TestHTTPSuccessDelete200(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Delete200(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPSuccessDelete202(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Delete202(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusAccepted)
}

func TestHTTPSuccessDelete204(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Delete204(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusNoContent)
}

func TestHTTPSuccessGet200(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Get200(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPSuccessHead200(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Head200(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPSuccessHead204(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Head204(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusNoContent)
}

func TestHTTPSuccessHead404(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Head404(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusNotFound)
}

func TestHTTPSuccessOptions200(t *testing.T) {
	t.Skip("options method not enabled by test server")
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Options200(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPSuccessPatch200(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Patch200(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPSuccessPatch202(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Patch202(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusAccepted)
}

func TestHTTPSuccessPatch204(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Patch204(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusNoContent)
}

func TestHTTPSuccessPost200(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Post200(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPSuccessPost201(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Post201(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestHTTPSuccessPost202(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Post202(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusAccepted)
}

func TestHTTPSuccessPost204(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Post204(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusNoContent)
}

func TestHTTPSuccessPut200(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Put200(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPSuccessPut201(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Put201(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestHTTPSuccessPut202(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Put202(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusAccepted)
}

func TestHTTPSuccessPut204(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPSuccessOperations()
	result, err := client.Put204(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusNoContent)
}
