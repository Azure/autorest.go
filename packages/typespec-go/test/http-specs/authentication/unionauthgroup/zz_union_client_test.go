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
	require.NoError(t, err)
	require.Zero(t, resp)
}