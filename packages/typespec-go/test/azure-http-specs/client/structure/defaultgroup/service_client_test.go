//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package defaultgroup_test

import (
	"context"
	"defaultgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceClient_One(t *testing.T) {
	client, err := defaultgroup.NewServiceClientWithNoCredential("http://localhost:3000", defaultgroup.ClientTypeDefault, nil)
	require.NoError(t, err)
	resp, err := client.One(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestServiceClient_Two(t *testing.T) {
	client, err := defaultgroup.NewServiceClientWithNoCredential("http://localhost:3000", defaultgroup.ClientTypeDefault, nil)
	require.NoError(t, err)
	resp, err := client.Two(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
