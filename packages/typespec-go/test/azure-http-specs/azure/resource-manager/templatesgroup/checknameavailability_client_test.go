// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// https://github.com/Azure/autorest.go/issues/1822
package templatesgroup_test

// import (
// 	"templatesgroup"
// 	"testing"

// 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
// 	"github.com/stretchr/testify/require"
// )

// func TestNewCheckNameAvailabilityClient_CheckGlobal(t *testing.T) {
// 	body := templatesgroup.CheckNameAvailabilityRequest{
// 		Name: to.Ptr("checkName"),
// 		Type: to.Ptr("Microsoft.Web/site"),
// 	}
// 	resp, err := clientFactory.NewCheckNameAvailabilityClient().CheckGlobal(ctx, body, nil)
// 	require.NoError(t, err)
// 	require.NotNil(t, resp)
// 	require.Equal(t, false, *resp.NameAvailable)
// 	require.Equal(t, templatesgroup.CheckNameAvailabilityReason("AlreadyExists"), *resp.Reason)
// 	require.Equal(t, "Hostname 'checkName' already exists. Please select a different name.", *resp.Message)
// }

// func TestNewCheckNameAvailabilityClient_CheckLocal(t *testing.T) {
// 	body := templatesgroup.CheckNameAvailabilityRequest{
// 		Name: to.Ptr("checkName"),
// 		Type: to.Ptr("Microsoft.Web/site"),
// 	}
// 	resp, err := clientFactory.NewCheckNameAvailabilityClient().CheckLocal(ctx, locationExpected, body, nil)
// 	require.NoError(t, err)
// 	require.NotNil(t, resp)
// 	require.Equal(t, false, *resp.NameAvailable)
// 	require.Equal(t, templatesgroup.CheckNameAvailabilityReason("AlreadyExists"), *resp.Reason)
// 	require.Equal(t, "Hostname 'checkName' already exists. Please select a different name.", *resp.Message)
// }

// func TestNewOperationsClient_NewListPager(t *testing.T) {
// 	options := &templatesgroup.OperationsClientListOptions{}
// 	pager := clientFactory.NewOperationsClient().NewListPager(options)
// 	require.NotNil(t, pager)
// 	for pager.More() {
// 		_, err := pager.NextPage(ctx)
// 		require.NoError(t, err)
// 	}
// }
