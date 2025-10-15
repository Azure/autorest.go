// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package previewversiongroup_test

import (
	"context"
	"testing"

	"previewversiongroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestPreviewVersionClient_GetWidget(t *testing.T) {
	const fakeID = "widget-123"
	client, err := previewversiongroup.NewPreviewVersionClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	ctx := context.Background()
	resp, err := client.GetWidget(ctx, fakeID, nil)

	require.NoError(t, err)
	require.Equal(t, "widget-123", *resp.ID)
	require.Equal(t, "Sample Widget", *resp.Name)
	require.Equal(t, "blue", *resp.Color)
}

func TestPreviewVersionClient_UpdateWidgetColor(t *testing.T) {
	const fakeID = "widget-123"
	const newColor = "red"
	client, err := previewversiongroup.NewPreviewVersionClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)

	colorUpdate := previewversiongroup.UpdateWidgetColorRequest{
		Color: to.Ptr(newColor),
	}

	resp, err := client.UpdateWidgetColor(context.Background(), fakeID, colorUpdate, nil)
	require.NoError(t, err)
	require.Equal(t, fakeID, *resp.ID)
	require.Equal(t, "Sample Widget", *resp.Name)
	require.Equal(t, newColor, *resp.Color)
}
