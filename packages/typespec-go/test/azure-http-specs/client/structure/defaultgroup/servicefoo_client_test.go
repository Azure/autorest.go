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

func TestServiceFooClient_Four(t *testing.T) {
	client, err := defaultgroup.NewServiceClientWithNoCredential("http://localhost:3000", defaultgroup.ClientTypeDefault, nil)
	require.NoError(t, err)
	resp, err := client.NewServiceFooClient().Four(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestServiceFooClient_Seven(t *testing.T) {
	client, err := defaultgroup.NewServiceClientWithNoCredential("http://localhost:3000", defaultgroup.ClientTypeDefault, nil)
	require.NoError(t, err)
	resp, err := client.NewServiceBazClient().NewServiceBazFooClient().Seven(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestServiceFooClient_Three(t *testing.T) {
	client, err := defaultgroup.NewServiceClientWithNoCredential("http://localhost:3000", defaultgroup.ClientTypeDefault, nil)
	require.NoError(t, err)
	resp, err := client.NewServiceFooClient().Three(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
