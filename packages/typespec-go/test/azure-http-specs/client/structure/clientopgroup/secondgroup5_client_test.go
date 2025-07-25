// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientopgroup_test

import (
	"clientopgroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecondGroup5Client_Six(t *testing.T) {
	client, err := clientopgroup.NewSecondGroup5Client("http://localhost:3000", clientopgroup.ClientTypeClientOperationGroup, nil)
	require.Nil(t, err)
	require.NotNil(t, client)
	resp, err := client.Six(context.Background(), &clientopgroup.SecondGroup5ClientSixOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}
