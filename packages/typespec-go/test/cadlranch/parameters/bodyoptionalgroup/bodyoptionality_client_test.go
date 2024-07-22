//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bodyoptionalgroup_test

import (
	"bodyoptionalgroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestBodyOptionalityClient_RequiredExplicit(t *testing.T) {
	client, err := bodyoptionalgroup.NewBodyOptionalityClient(nil)
	require.NoError(t, err)
	resp, err := client.RequiredExplicit(context.Background(), bodyoptionalgroup.BodyModel{
		Name: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestBodyOptionalityClient_RequiredImplicit(t *testing.T) {
	client, err := bodyoptionalgroup.NewBodyOptionalityClient(nil)
	require.NoError(t, err)
	resp, err := client.RequiredImplicit(context.Background(), "foo", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
