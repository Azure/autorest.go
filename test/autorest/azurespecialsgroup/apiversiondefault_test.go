// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgrouptest

import (
	"context"
	"generatortests/autorest/generated/azurespecialsgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

// GetMethodGlobalNotProvidedValid - GET method with api-version modeled in global settings.
func TestGetMethodGlobalNotProvidedValid(t *testing.T) {
	client := azurespecialsgroup.NewDefaultClient(nil).APIVersionDefaultOperations()
	result, err := client.GetMethodGlobalNotProvidedValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// GetMethodGlobalValid - GET method with api-version modeled in global settings.
func TestGetMethodGlobalValid(t *testing.T) {
	client := azurespecialsgroup.NewDefaultClient(nil).APIVersionDefaultOperations()
	result, err := client.GetMethodGlobalValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// GetPathGlobalValid - GET method with api-version modeled in global settings.
func TestGetPathGlobalValid(t *testing.T) {
	client := azurespecialsgroup.NewDefaultClient(nil).APIVersionDefaultOperations()
	result, err := client.GetPathGlobalValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// GetSwaggerGlobalValid - GET method with api-version modeled in global settings.
func TestGetSwaggerGlobalValid(t *testing.T) {
	client := azurespecialsgroup.NewDefaultClient(nil).APIVersionDefaultOperations()
	result, err := client.GetSwaggerGlobalValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
