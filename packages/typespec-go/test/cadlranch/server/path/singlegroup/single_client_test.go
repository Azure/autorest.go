// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package singlegroup_test

import (
	"context"
	"singlegroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleClient_MyOp(t *testing.T) {
	client, err := singlegroup.NewSingleClient(nil)
	require.NoError(t, err)
	resp, err := client.MyOp(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
