// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package templatesgroup

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestNewCheckNameAvailabilityClient_CheckGlobal(t *testing.T) {
	body := CheckNameAvailabilityRequest{
		Name: to.Ptr("checkName"),
		Type: to.Ptr("Microsoft.Web/site"),
	}
	_, err := clientFactory.NewCheckNameAvailabilityClient().CheckGlobal(ctx, body, nil)
	require.NoError(t, err)
}

func TestNewCheckNameAvailabilityClient_CheckLocal(t *testing.T) {
	body := CheckNameAvailabilityRequest{
		Name: to.Ptr("checkName"),
		Type: to.Ptr("Microsoft.Web/site"),
	}
	_, err := clientFactory.NewCheckNameAvailabilityClient().CheckLocal(ctx, locationExpected, body, nil)
	require.NoError(t, err)
}
