// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package overridegroup_test

import (
	"context"
	"overridegroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestOverrideRemoveOptionalParametersClient_Group(t *testing.T) {
	client, err := overridegroup.NewOverrideClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOverrideRemoveOptionalParameterClient().RemoveOptional(context.Background(), "param1", &overridegroup.OverrideRemoveOptionalParameterClientRemoveOptionalOptions{to.Ptr("param2")})
	require.NoError(t, err)
	require.Zero(t, resp)
}
