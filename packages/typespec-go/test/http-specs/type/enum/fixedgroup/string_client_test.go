//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package fixedgroup_test

import (
	"context"
	"fixedgroup"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func TestFixedClientGetKnownValue(t *testing.T) {
	client, err := fixedgroup.NewFixedClient(nil)
	require.NoError(t, err)
	resp, err := client.NewFixedStringClient().GetKnownValue(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, fixedgroup.DaysOfWeekEnumMonday, *resp.Value)
}

func TestFixedClientPutKnownValue(t *testing.T) {
	client, err := fixedgroup.NewFixedClient(nil)
	require.NoError(t, err)
	resp, err := client.NewFixedStringClient().PutKnownValue(context.Background(), fixedgroup.DaysOfWeekEnumMonday, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFixedClientPutUnknownValue(t *testing.T) {
	client, err := fixedgroup.NewFixedClient(nil)
	require.NoError(t, err)
	resp, err := client.NewFixedStringClient().PutUnknownValue(context.Background(), fixedgroup.DaysOfWeekEnum("Weekend"), nil)
	var respErr *azcore.ResponseError
	require.ErrorAs(t, err, &respErr)
	require.EqualValues(t, http.StatusInternalServerError, respErr.StatusCode)
	require.Zero(t, resp)
}
