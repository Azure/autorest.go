//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package collectionfmtgroup_test

import (
	"collectionfmtgroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeaderClient_CSV(t *testing.T) {
	client, err := collectionfmtgroup.NewCollectionFormatClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewCollectionFormatHeaderClient().CSV(context.Background(), []string{"blue", "red", "green"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
