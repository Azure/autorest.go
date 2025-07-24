// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientopgroup_test

import (
	"clientopgroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstGroup4Client_Four(t *testing.T) {
	client, err := clientopgroup.NewFirstGroup4Client("http://localhost:3000", clientopgroup.ClientTypeClientOperationGroup, nil)
	require.Nil(t, err)
	require.NotNil(t, client)
	resp, err := client.Four(context.Background(), &clientopgroup.FirstGroup4ClientFourOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}
