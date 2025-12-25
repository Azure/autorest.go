// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package methodsubscriptionidgroup_test

import (
	"context"
	"methodsubscriptionidgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestOperationsClient_NewListPager(t *testing.T) {
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	require.NotNil(t, pager)
	validOperation := methodsubscriptionidgroup.Operation{
		Name:         to.Ptr("Azure.ResourceManager.MethodSubscriptionId/services/read"),
		IsDataAction: to.Ptr(false),
		Display: &methodsubscriptionidgroup.OperationDisplay{
			Provider:    to.Ptr("Azure.ResourceManager.MethodSubscriptionId"),
			Resource:    to.Ptr("services"),
			Operation:   to.Ptr("Lists services"),
			Description: to.Ptr("Lists registered services"),
		},
	}
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotNil(t, page)
		require.NotEmpty(t, page.Value)
		for _, op := range page.Value {
			require.Equal(t, validOperation, *op)
		}
	}
}
