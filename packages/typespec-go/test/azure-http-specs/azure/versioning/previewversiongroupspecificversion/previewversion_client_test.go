// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package previewversiongroupspecificversion_test

import (
	"context"
	"previewversiongroupspecificversion"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestPreviewVersionClient_ListWidgets(t *testing.T) {
	client, err := previewversiongroupspecificversion.NewPreviewVersionClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	opts := &previewversiongroupspecificversion.PreviewVersionClientListWidgetsOptions{
		Name: to.Ptr("test"),
	}
	resp, err := client.ListWidgets(context.Background(), opts)
	require.NoError(t, err)
	require.Equal(t, "widget-1", *resp.Widgets[0].ID)
	require.Equal(t, "test", *resp.Widgets[0].Name)
}
