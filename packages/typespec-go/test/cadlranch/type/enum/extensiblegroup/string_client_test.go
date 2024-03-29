//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package extensiblegroup_test

import (
	"context"
	"extensiblegroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtensibleClientGetKnownValue(t *testing.T) {
	client, err := extensiblegroup.NewExtensibleClient(nil)
	require.NoError(t, err)
	resp, err := client.NewExtensibleStringClient().GetKnownValue(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.Equal(t, extensiblegroup.DaysOfWeekExtensibleEnumMonday, *resp.Value)
}

func TestExtensibleClientGetUnknownValue(t *testing.T) {
	client, err := extensiblegroup.NewExtensibleClient(nil)
	require.NoError(t, err)
	resp, err := client.NewExtensibleStringClient().GetUnknownValue(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.Equal(t, extensiblegroup.DaysOfWeekExtensibleEnum("Weekend"), *resp.Value)
}

func TestExtensibleClientPutKnownValue(t *testing.T) {
	client, err := extensiblegroup.NewExtensibleClient(nil)
	require.NoError(t, err)
	resp, err := client.NewExtensibleStringClient().PutKnownValue(context.Background(), extensiblegroup.DaysOfWeekExtensibleEnumMonday, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestExtensibleClientPutUnknownValue(t *testing.T) {
	client, err := extensiblegroup.NewExtensibleClient(nil)
	require.NoError(t, err)
	resp, err := client.NewExtensibleStringClient().PutUnknownValue(context.Background(), extensiblegroup.DaysOfWeekExtensibleEnum("Weekend"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
