// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package documentationgroup_test

import (
	"context"
	"documentationgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTextFormattingClient_BoldText(t *testing.T) {
	client, err := documentationgroup.NewDocumentationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDocumentationTextFormattingClient().BoldText(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestTextFormattingClient_CombinedFormatting(t *testing.T) {
	client, err := documentationgroup.NewDocumentationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDocumentationTextFormattingClient().CombinedFormatting(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestTextFormattingClient_ItalicText(t *testing.T) {
	client, err := documentationgroup.NewDocumentationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDocumentationTextFormattingClient().ItalicText(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
