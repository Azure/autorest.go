//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nodiscgroup_test

import (
	"context"
	"nodiscgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestNotDiscriminatedClientGetValid(t *testing.T) {
	client, err := nodiscgroup.NewNotDiscriminatedClient(nil)
	require.NoError(t, err)
	resp, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nodiscgroup.Siamese{
		Age:   to.Ptr[int32](32),
		Name:  to.Ptr("abc"),
		Smart: to.Ptr(true),
	}, resp.Siamese)
}

func TestNotDiscriminatedClientPostValid(t *testing.T) {
	client, err := nodiscgroup.NewNotDiscriminatedClient(nil)
	require.NoError(t, err)
	resp, err := client.PostValid(context.Background(), nodiscgroup.Siamese{
		Age:   to.Ptr[int32](32),
		Name:  to.Ptr("abc"),
		Smart: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestNotDiscriminatedClientPutValid(t *testing.T) {
	client, err := nodiscgroup.NewNotDiscriminatedClient(nil)
	require.NoError(t, err)
	myCat := nodiscgroup.Siamese{
		Age:   to.Ptr[int32](9),
		Name:  to.Ptr("Luna"),
		Smart: to.Ptr(false),
	}
	resp, err := client.PutValid(context.Background(), myCat, nil)
	require.NoError(t, err)
	require.EqualValues(t, resp.Siamese, myCat)
}
