// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package templatesgroup

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

const RESOURCE_GROUP_EXPECTED = "test-rg"

func TestNewLroClient_BeginCreateOrReplace(t *testing.T) {
	createdByType := CreatedByType("User")
	lastModifiedByType := CreatedByType("User")
	createdAt, err := time.Parse(time.RFC3339Nano, "2024-10-04T00:56:07.442Z")
	require.NoError(t, err)
	lastModifiedAt, err := time.Parse(time.RFC3339Nano, "2024-10-04T00:56:07.442Z")
	require.NoError(t, err)

	poller, err := clientFactory.NewLroClient().BeginCreateOrReplace(ctx, RESOURCE_GROUP_EXPECTED, "order1", Order{
		Location: to.Ptr("eastus"),
		ID:       to.Ptr("/subscriptions/${SUBSCRIPTION_ID_EXPECTED}/resourceGroups/${RESOURCE_GROUP_EXPECTED}/providers/Azure.ResourceManager.OperationTemplates/orders/order1"),
		Name:     to.Ptr("order1"),
		Type:     to.Ptr("Azure.ResourceManager.Resources/orders"),
		Properties: &OrderProperties{
			ProvisioningState: to.Ptr("Succeeded"),
			ProductID:         to.Ptr("product1"),
			Amount:            to.Ptr(int32(1)),
		},
		SystemData: &SystemData{
			CreatedBy:          to.Ptr("AzureSDK"),
			CreatedByType:      &createdByType,
			CreatedAt:          &createdAt,
			LastModifiedBy:     to.Ptr("AzureSDK"),
			LastModifiedAt:     &lastModifiedAt,
			LastModifiedByType: &lastModifiedByType,
		},
	}, nil)
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestNewLroClient_BeginDelete(t *testing.T) {
	poller, err := clientFactory.NewLroClient().BeginDelete(context.Background(), RESOURCE_GROUP_EXPECTED, "order1", nil)
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestNewLroClient_BeginExport(t *testing.T) {
	body := ExportRequest{
		Format: to.Ptr("csv"),
	}
	poller, err := clientFactory.NewLroClient().BeginExport(context.Background(), RESOURCE_GROUP_EXPECTED, "order1", body, nil)
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}
