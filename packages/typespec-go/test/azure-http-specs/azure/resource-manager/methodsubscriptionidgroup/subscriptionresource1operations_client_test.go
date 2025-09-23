// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package methodsubscriptionidgroup_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"

	"methodsubscriptionidgroup"
)

// parseTime parses an RFC3339 time string and panics if parsing fails.
func parseTime(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		log.Fatalf("failed to parse time: %v", err)
	}
	return t
}

var validResource1 = methodsubscriptionidgroup.SubscriptionResource1{
	ID:   to.Ptr(fmt.Sprintf(`/subscriptions/%s/providers/Azure.ResourceManager.MethodSubscriptionId/subscriptionResource1s/sub-resource-1`, subscriptionIdExpected)),
	Name: to.Ptr("sub-resource-1"),
	Type: to.Ptr("Azure.ResourceManager.MethodSubscriptionId/subscriptionResource1s"),
	Properties: &methodsubscriptionidgroup.SubscriptionResource1Properties{
		ProvisioningState: to.Ptr(methodsubscriptionidgroup.ResourceProvisioningState("Succeeded")),
		Description:       to.Ptr(string("Valid subscription resource 1")),
	},
	SystemData: &methodsubscriptionidgroup.SystemData{
		CreatedBy:          to.Ptr(string("AzureSDK")),
		CreatedByType:      to.Ptr(methodsubscriptionidgroup.CreatedByType("User")),
		CreatedAt:          to.Ptr(parseTime("2023-01-01T00:00:00.000Z")),
		LastModifiedBy:     to.Ptr(string("AzureSDK")),
		LastModifiedAt:     to.Ptr(parseTime("2023-01-01T00:00:00.000Z")),
		LastModifiedByType: to.Ptr(methodsubscriptionidgroup.CreatedByType("User")),
	},
}

func TestSubscriptionResource1OperationsClient_Delete(t *testing.T) {
	delResp, err := clientFactory.NewSubscriptionResource1OperationsClient().Delete(context.Background(), subscriptionIdExpected, "sub-resource-1", nil)
	require.NoError(t, err)
	require.Zero(t, delResp)
}

func TestSubscriptionResource1OperationsClient_Put(t *testing.T) {
	var validResource = methodsubscriptionidgroup.SubscriptionResource1{
		Properties: &methodsubscriptionidgroup.SubscriptionResource1Properties{
			Description: to.Ptr(string("Valid subscription resource 1")),
		},
	}
	putResp, err := clientFactory.NewSubscriptionResource1OperationsClient().Put(context.Background(), subscriptionIdExpected, "sub-resource-1", validResource, nil)
	require.NoError(t, err)
	require.NotNil(t, putResp)
	require.Equal(t, validResource1, putResp.SubscriptionResource1)
}

func TestSubscriptionResource1OperationsClient_Get(t *testing.T) {
	getResp, err := clientFactory.NewSubscriptionResource1OperationsClient().Get(context.Background(), subscriptionIdExpected, "sub-resource-1", nil)
	require.NoError(t, err)
	require.NotNil(t, getResp)
	require.Equal(t, validResource1, getResp.SubscriptionResource1)
}
