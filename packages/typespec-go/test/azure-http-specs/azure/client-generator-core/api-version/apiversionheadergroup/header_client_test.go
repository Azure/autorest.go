// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apiversionheadergroup_test

import (
	"context"
	"testing"

	"apiversionheadergroup"

	"github.com/stretchr/testify/require"
)

func TestHeaderClient_HeaderAPIVersion(t *testing.T) {
	client, err := apiversionheadergroup.NewHeaderClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	resp, err := client.HeaderAPIVersion(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
