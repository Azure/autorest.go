// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package templatesgroup

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/stretchr/testify/require"
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func TestCheckNameAvailabilityClient_CheckGlobal_Success(t *testing.T) {
	reason := CheckNameAvailabilityReason("AlreadyExists")
	expected := CheckNameAvailabilityClientCheckLocalResponse{
		CheckNameAvailabilityResponse: CheckNameAvailabilityResponse{
			NameAvailable: toPtr(false),
			Reason:        &reason,
			Message:       toPtr("Name already exists"),
		},
	}
	tenantID := getEnv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		TenantID: tenantID,
	})
	require.NoError(t, err)
	client, err := NewCheckNameAvailabilityClient(getEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000"), cred, nil)
	require.NoError(t, err)
	body := CheckNameAvailabilityRequest{Name: toPtr("testName")}
	resp, err := client.CheckGlobal(context.Background(), body, nil)
	require.NoError(t, err)
	require.Equal(t, expected, resp.CheckNameAvailabilityResponse)
}

func TestCheckNameAvailabilityClient_CheckGlobal_ErrorStatus(t *testing.T) {
	tenantID := getEnv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		TenantID: tenantID,
	})
	require.NoError(t, err)
	client, err := NewCheckNameAvailabilityClient(getEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000"), cred, nil)
	require.NoError(t, err)
	body := CheckNameAvailabilityRequest{Name: toPtr("badName")}
	_, err = client.CheckGlobal(context.Background(), body, nil)
	require.Error(t, err)
}

func TestCheckNameAvailabilityClient_CheckLocal_Success(t *testing.T) {
	reason := CheckNameAvailabilityReason("AlreadyExists")

	expected := CheckNameAvailabilityResponse{
		NameAvailable: toPtr(false),
		Reason:        &reason,
		Message:       toPtr("Name is already taken"),
	}
	tenantID := getEnv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		TenantID: tenantID,
	})
	require.NoError(t, err)
	client, err := NewCheckNameAvailabilityClient(getEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000"), cred, nil)
	require.NoError(t, err)
	body := CheckNameAvailabilityRequest{Name: toPtr("existingName")}
	resp, err := client.CheckLocal(context.Background(), "eastus", body, nil)
	require.NoError(t, err)
	require.Equal(t, expected, resp.CheckNameAvailabilityResponse)
}

func TestCheckNameAvailabilityClient_CheckLocal_ErrorStatus(t *testing.T) {
	tenantID := getEnv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		TenantID: tenantID,
	})
	require.NoError(t, err)
	client, err := NewCheckNameAvailabilityClient(getEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000"), cred, nil)
	body := CheckNameAvailabilityRequest{Name: toPtr("conflictName")}
	_, err = client.CheckLocal(context.Background(), "westus", body, nil)
	require.Error(t, err)
}

func TestCheckNameAvailabilityClient_checkGlobalCreateRequest_EmptySubscriptionID(t *testing.T) {
	client := &CheckNameAvailabilityClient{}
	_, err := client.checkGlobalCreateRequest(context.Background(), CheckNameAvailabilityRequest{}, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "subscriptionID")
}

func TestCheckNameAvailabilityClient_checkLocalCreateRequest_EmptyLocation(t *testing.T) {
	client := &CheckNameAvailabilityClient{subscriptionID: "subid"}
	_, err := client.checkLocalCreateRequest(context.Background(), "", CheckNameAvailabilityRequest{}, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "location")
}

func toPtr[T any](v T) *T {
	return &v
}
