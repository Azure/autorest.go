// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package unionauthgroup_test

import (
	"context"
	"testing"
	"unionauthgroup"

	"github.com/stretchr/testify/require"
)

func TestUnionAuthGroupClient_ValidKey(t *testing.T) {
	client, err := unionauthgroup.NewunionauthgroupClient(nil)
	require.NoError(t, err)
	resp, err := client.ValidKey(context.Background(), &unionauthgroup.UnionClientValidKeyOptions{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "Expected valid-key but got undefined")
	require.Zero(t, resp)
}

func TestUnionAuthGroupClient_ValidToken(t *testing.T) {
	client, err := unionauthgroup.NewunionauthgroupClient(nil)
	require.NoError(t, err)
	resp, err := client.ValidToken(context.Background(), &unionauthgroup.UnionClientValidTokenOptions{})
	require.Error(t, err)
	require.Zero(t, resp)
}
