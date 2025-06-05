// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonresourcegroup

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestNewNonResourceOperationsClient_Create(t *testing.T) {
	body := NonResource{
		ID:   to.Ptr("id"),
		Name: to.Ptr("hello"),
		Type: to.Ptr("nonResource"),
	}

	resp, err := clientFactory.NewNonResourceOperationsClient().Create(ctx, locationExpected, "parameter", body, nil)
	require.NoError(t, err)
	require.Equal(t, resp.NonResource.ID, to.Ptr("id"))
	require.Equal(t, resp.NonResource.Name, to.Ptr("hello"))
	require.Equal(t, resp.NonResource.Type, to.Ptr("nonResource"))
}

func TestNewNonResourceOperationsClient_Get(t *testing.T) {
	resp, err := clientFactory.NewNonResourceOperationsClient().Get(ctx, locationExpected, "parameter", nil)
	require.NoError(t, err)
	require.Equal(t, resp.NonResource.ID, to.Ptr("id"))
	require.Equal(t, resp.NonResource.Name, to.Ptr("hello"))
	require.Equal(t, resp.NonResource.Type, to.Ptr("nonResource"))
}
