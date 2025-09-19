// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package methodsubscriptionidgroup_test

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/stretchr/testify/require"

	"methodsubscriptionidgroup"
)

func TestSubscriptionResource2OperationsClient_CRUD(t *testing.T) {
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	require.NotEmpty(t, subscriptionID, "AZURE_SUBSCRIPTION_ID must be set")

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err)

	client, err := methodsubscriptionidgroup.NewSubscriptionResource2OperationsClient(subscriptionID, cred, nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	resourceName := "test-resource2"
	resource := methodsubscriptionidgroup.SubscriptionResource2{
		// TODO: Fill with required fields for creation according to Azure best practices
	}

	// Create (Put)
	putResp, err := client.Put(context.Background(), resourceName, resource, nil)
	require.NoError(t, err)
	require.NotNil(t, putResp.SubscriptionResource2)

	// Get
	getResp, err := client.Get(context.Background(), resourceName, nil)
	require.NoError(t, err)
	require.NotNil(t, getResp.SubscriptionResource2)

	// Delete
	delResp, err := client.Delete(context.Background(), resourceName, nil)
	require.NoError(t, err)
	require.NotNil(t, delResp)
}
