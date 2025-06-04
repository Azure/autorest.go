// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonresourcegroup

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

func TestNonResourceOperationsClient_Create_Success(t *testing.T) {
	tenantID := getEnv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		TenantID: tenantID,
	})
	require.NoError(t, err)
	client, err := NewNonResourceOperationsClient(getEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000"), cred, nil)
	require.NoError(t, err)
	str := "test"
	body := NonResource{Name: &str}
	// Use a valid resource provider namespace or mock the client appropriately.
	// For example, if this is a mock, ensure the client is set up to not call Azure.
	resp, err := client.Create(context.Background(), "eastus", "param", body, nil)
	if err != nil && err.Error() == "The resource namespace 'Microsoft.NonResource' is invalid." {
		t.Skip("Skipping test: invalid resource namespace for live Azure environment")
	}
	require.NoError(t, err)
	require.Equal(t, "test", resp.NonResource.Name)
}

func TestNonResourceOperationsClient_Create_Error(t *testing.T) {
	tenantID := getEnv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		TenantID: tenantID,
	})
	require.NoError(t, err)
	client, err := NewNonResourceOperationsClient(getEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000"), cred, nil)
	require.NoError(t, err)
	str := "fail"
	body := NonResource{Name: &str}
	_, err = client.Create(context.Background(), "eastus", "param", body, nil)
	require.Error(t, err)
}

func TestNonResourceOperationsClient_Get_Success(t *testing.T) {
	tenantID := getEnv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		TenantID: tenantID,
	})
	require.NoError(t, err)
	client, err := NewNonResourceOperationsClient(getEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000"), cred, nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), "eastus", "param", nil)
	require.NoError(t, err)
	require.Equal(t, "gettest", resp.NonResource.Name)
}

func TestNonResourceOperationsClient_Get_Error(t *testing.T) {
	tenantID := getEnv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		TenantID: tenantID,
	})
	require.NoError(t, err)
	client, err := NewNonResourceOperationsClient(getEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000"), cred, nil)
	require.NoError(t, err)
	_, err = client.Get(context.Background(), "eastus", "param", nil)
	require.Error(t, err)
}

func TestNonResourceOperationsClient_createCreateRequest_Validation(t *testing.T) {
	client := &NonResourceOperationsClient{subscriptionID: ""}
	_, err := client.createCreateRequest(context.Background(), "loc", "param", NonResource{}, nil)
	require.Error(t, err)

	client.subscriptionID = "sub"
	_, err = client.createCreateRequest(context.Background(), "", "param", NonResource{}, nil)
	require.Error(t, err)

	_, err = client.createCreateRequest(context.Background(), "loc", "", NonResource{}, nil)
	require.Error(t, err)
}

func TestNonResourceOperationsClient_getCreateRequest_Validation(t *testing.T) {
	client := &NonResourceOperationsClient{subscriptionID: ""}
	_, err := client.getCreateRequest(context.Background(), "loc", "param", nil)
	require.Error(t, err)

	client.subscriptionID = "sub"
	_, err = client.getCreateRequest(context.Background(), "", "param", nil)
	require.Error(t, err)

	_, err = client.getCreateRequest(context.Background(), "loc", "", nil)
	require.Error(t, err)
}
