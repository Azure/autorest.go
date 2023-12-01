//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package extensibleenumgroup_test

import (
	"context"
	"extensibleenumgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtensibleClientGetKnownValue(t *testing.T) {
	client, err := extensibleenumgroup.NewExtensibleClient(nil)
	require.NoError(t, err)
	resp, err := client.GetKnownValue(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.Equal(t, extensibleenumgroup.DaysOfWeekExtensibleEnumMonday, *resp.Value)
}

func TestExtensibleClientGetUnknownValue(t *testing.T) {
	client, err := extensibleenumgroup.NewExtensibleClient(nil)
	require.NoError(t, err)
	resp, err := client.GetUnknownValue(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.Equal(t, extensibleenumgroup.DaysOfWeekExtensibleEnum("Weekend"), *resp.Value)
}

func TestExtensibleClientPutKnownValue(t *testing.T) {
	client, err := extensibleenumgroup.NewExtensibleClient(nil)
	require.NoError(t, err)
	resp, err := client.PutKnownValue(context.Background(), extensibleenumgroup.DaysOfWeekExtensibleEnumMonday, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestExtensibleClientPutUnknownValue(t *testing.T) {
	client, err := extensibleenumgroup.NewExtensibleClient(nil)
	require.NoError(t, err)
	resp, err := client.PutUnknownValue(context.Background(), extensibleenumgroup.DaysOfWeekExtensibleEnum("Weekend"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
