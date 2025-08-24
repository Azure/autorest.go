// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package overridegroup_test

import (
	"context"
	"overridegroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOverrideReorderParametersClient_Reorder(t *testing.T) {
	client, err := overridegroup.NewOverrideClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOverrideReorderParametersClient().Reorder(context.Background(), "param2", "param1", &overridegroup.OverrideReorderParametersClientReorderOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}
