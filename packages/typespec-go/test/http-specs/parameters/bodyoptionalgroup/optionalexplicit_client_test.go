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

func TestOptionalExplicitClient_Omit(t *testing.T) {
	client, err := bodyoptionalgroup.NewBodyOptionalityClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewBodyOptionalityOptionalExplicitClient().Omit(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalExplicitClient_Set(t *testing.T) {
	client, err := bodyoptionalgroup.NewBodyOptionalityClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewBodyOptionalityOptionalExplicitClient().Set(context.Background(), &bodyoptionalgroup.BodyOptionalityOptionalExplicitClientSetOptions{
		Body: &bodyoptionalgroup.BodyModel{
			Name: to.Ptr("foo"),
		},
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}
