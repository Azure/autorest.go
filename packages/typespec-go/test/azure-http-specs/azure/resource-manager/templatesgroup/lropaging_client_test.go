// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package templatesgroup_test

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func TestNewLroPagingClient_BeginPostPagingLro(t *testing.T) {
	// Start the long-running operation
	poller, err := clientFactory.NewLroPagingClient().BeginPostPagingLro(ctx, resourceGroupExpected, "default", nil)
	require.NoError(t, err)

	// Poll until the LRO completes and get the pager
	pager, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.NotNil(t, pager)

	// Iterate through the first page
	hasFirstPage := pager.More()
	require.True(t, hasFirstPage, "Expected at least one page of results")

	page, err := pager.NextPage(context.Background())
	require.NoError(t, err)
	require.NotNil(t, page.Value)
	require.Len(t, page.Value, 1, "Expected first page to contain 1 product")

	// Validate first product
	product1 := page.Value[0]
	require.NotNil(t, product1.Name)
	require.Equal(t, "product1", *product1.Name)
	require.NotNil(t, product1.Properties)
	require.NotNil(t, product1.Properties.ProductID)
	require.Equal(t, "product1", *product1.Properties.ProductID)
	require.NotNil(t, product1.Properties.ProvisioningState)
	require.Equal(t, "Succeeded", *product1.Properties.ProvisioningState)

	// Iterate through the second page
	hasSecondPage := pager.More()
	require.True(t, hasSecondPage, "Expected a second page of results")

	page2, err := pager.NextPage(context.Background())
	require.NoError(t, err)
	require.NotNil(t, page2.Value)
	require.Len(t, page2.Value, 1, "Expected second page to contain 1 product")

	// Validate second product
	product2 := page2.Value[0]
	require.NotNil(t, product2.Name)
	require.Equal(t, "product2", *product2.Name)
	require.NotNil(t, product2.Properties)
	require.NotNil(t, product2.Properties.ProductID)
	require.Equal(t, "product2", *product2.Properties.ProductID)
	require.NotNil(t, product2.Properties.ProvisioningState)
	require.Equal(t, "Succeeded", *product2.Properties.ProvisioningState)

	// Verify no more pages
	hasMorePages := pager.More()
	require.False(t, hasMorePages, "Expected no more pages after the second page")
}
