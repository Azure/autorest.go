// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apiversionpathgroup_test

import (
	"context"
	"testing"

	"apiversionpathgroup"

	"github.com/stretchr/testify/require"
)

func TestPathClient_PathAPIVersion(t *testing.T) {
	client, err := apiversionpathgroup.NewPathClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	resp, err := client.PathAPIVersion(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
