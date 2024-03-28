// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package contentneggroup_test

import (
	"contentneggroup"
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSameBodyClient_GetAvatarAsJPEG(t *testing.T) {
	client, err := contentneggroup.NewContentNegotiationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSameBodyClient().GetAvatarAsJPEG(context.Background(), nil)
	require.NoError(t, err)
	jpgResp, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	jpgFile, err := os.ReadFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.jpg")
	require.NoError(t, err)
	require.EqualValues(t, jpgFile, jpgResp)
}

func TestSameBodyClient_GetAvatarAsPNG(t *testing.T) {
	client, err := contentneggroup.NewContentNegotiationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSameBodyClient().GetAvatarAsPNG(context.Background(), nil)
	require.NoError(t, err)
	pngResp, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	pngFile, err := os.ReadFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.png")
	require.NoError(t, err)
	require.EqualValues(t, pngFile, pngResp)
}
