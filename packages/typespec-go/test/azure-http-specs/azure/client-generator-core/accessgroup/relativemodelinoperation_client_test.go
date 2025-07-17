// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package accessgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestRelativeModelInOperationClient_discriminator(t *testing.T) {
	client, err := NewAccessClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewAccessRelativeModelInOperationClient().discriminator(context.Background(), "real", nil)
	require.NoError(t, err)
	require.Equal(t, &realModel{
		Kind: to.Ptr("real"),
		Name: to.Ptr("Madge"),
	}, resp.abstractModelClassification)
}

func TestRelativeModelInOperationClient_operation(t *testing.T) {
	client, err := NewAccessClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewAccessRelativeModelInOperationClient().operation(context.Background(), "Madge", nil)
	require.NoError(t, err)
	require.Equal(t, outerModel{
		Name: to.Ptr("Madge"),
		Inner: &innerModel{
			Name: to.Ptr("Madge"),
		},
	}, resp.outerModel)
}
