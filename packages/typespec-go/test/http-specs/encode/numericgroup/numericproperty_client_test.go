// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package numericgroup_test

import (
	"context"
	"numericgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestNumericPropertyClient_SafeintAsString(t *testing.T) {
	client, err := numericgroup.NewNumericClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewNumericPropertyClient().SafeintAsString(context.Background(), numericgroup.SafeintAsStringProperty{Value: to.Ptr(int64(10000000000))}, nil)
	require.NoError(t, err)
	require.Equal(t, numericgroup.SafeintAsStringProperty{
		Value: to.Ptr(int64(10000000000)),
	}, resp.SafeintAsStringProperty)
}

func TestNumericPropertyClient_Uint32AsStringOptional(t *testing.T) {
	client, err := numericgroup.NewNumericClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewNumericPropertyClient().Uint32AsStringOptional(context.Background(), numericgroup.Uint32AsStringProperty{Value: to.Ptr(uint32(1))}, nil)
	require.NoError(t, err)
	require.Equal(t, numericgroup.Uint32AsStringProperty{
		Value: to.Ptr(uint32(1)),
	}, resp.Uint32AsStringProperty)
}

func TestNumericPropertyClient_Uint8AsString(t *testing.T) {
	client, err := numericgroup.NewNumericClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewNumericPropertyClient().Uint8AsString(context.Background(), numericgroup.Uint8AsStringProperty{Value: to.Ptr(uint8(255))}, nil)
	require.NoError(t, err)
	require.Equal(t, numericgroup.Uint8AsStringProperty{
		Value: to.Ptr(uint8(255)),
	}, resp.Uint8AsStringProperty)
}
