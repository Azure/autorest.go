// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package overridegroup_test

import (
	"context"
	"overridegroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOverrideGroupParametersClient_Group(t *testing.T) {
	client, err := overridegroup.NewOverrideClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOverrideGroupParametersClient().Group(context.Background(), overridegroup.GroupParametersOptions{
		Param1: "param1",
		Param2: "param2",
	}, &overridegroup.OverrideGroupParametersClientGroupOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}
