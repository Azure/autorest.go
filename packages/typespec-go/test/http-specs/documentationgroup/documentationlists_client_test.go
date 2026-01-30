// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package documentationgroup_test

import (
	"context"
	"documentationgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListsClient_BulletPointsModel(t *testing.T) {
	client, err := documentationgroup.NewDocumentationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	enumValue := documentationgroup.BulletPointsEnum("Simple")
	resp, err := client.NewDocumentationListsClient().BulletPointsModel(context.Background(), documentationgroup.BulletPointsModel{Prop: &enumValue}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestListsClient_Numbered(t *testing.T) {
	client, err := documentationgroup.NewDocumentationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDocumentationListsClient().Numbered(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestListsClient_BulletPointsOp(t *testing.T) {
	client, err := documentationgroup.NewDocumentationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDocumentationListsClient().BulletPointsOp(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
