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

func getSubscriptionInCredentialsOperations(t *testing.T) azurespecialsgroup.SubscriptionInCredentialsOperations {
	client, err := azurespecialsgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client.SubscriptionInCredentialsOperations("1234-5678-9012-3456")
}

// PostMethodGlobalNotProvidedValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func TestPostMethodGlobalNotProvidedValid(t *testing.T) {
	client := getSubscriptionInCredentialsOperations(t)
	result, err := client.PostMethodGlobalNotProvidedValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// PostMethodGlobalNull - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to null, and client-side validation should prevent you from making this call
func TestPostMethodGlobalNull(t *testing.T) {
	t.Skip("invalid test, subscription ID is not x-nullable")
}

// PostMethodGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func TestPostMethodGlobalValid(t *testing.T) {
	client := getSubscriptionInCredentialsOperations(t)
	result, err := client.PostMethodGlobalValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// PostPathGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func TestPostPathGlobalValid(t *testing.T) {
	client := getSubscriptionInCredentialsOperations(t)
	result, err := client.PostPathGlobalValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// PostSwaggerGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func TestPostSwaggerGlobalValid(t *testing.T) {
	client := getSubscriptionInCredentialsOperations(t)
	result, err := client.PostSwaggerGlobalValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
