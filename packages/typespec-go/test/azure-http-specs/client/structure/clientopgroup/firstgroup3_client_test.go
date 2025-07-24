// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientopgroup_test

import (
	"clientopgroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstGroup3Client_Three(t *testing.T) {
	client, err := clientopgroup.NewFirstGroup3Client("http://localhost:3000", clientopgroup.ClientTypeClientOperationGroup, nil)
	require.Nil(t, err)
	require.NotNil(t, client)
	resp, err := client.Three(context.Background(), &clientopgroup.FirstGroup3ClientThreeOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}

func TestFirstGroup3Client_Two(t *testing.T) {
	client, err := clientopgroup.NewFirstGroup3Client("http://localhost:3000", clientopgroup.ClientTypeClientOperationGroup, nil)
	require.Nil(t, err)
	require.NotNil(t, client)
	resp, err := client.Two(context.Background(), &clientopgroup.FirstGroup3ClientTwoOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}
