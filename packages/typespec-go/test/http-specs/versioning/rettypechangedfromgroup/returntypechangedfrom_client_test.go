// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package rettypechangedfromgroup_test

import (
	"context"
	"rettypechangedfromgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Test(t *testing.T) {
	client, err := rettypechangedfromgroup.NewReturnTypeChangedFromClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Test(context.Background(), "test", nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, "test", *resp.Value)
}
