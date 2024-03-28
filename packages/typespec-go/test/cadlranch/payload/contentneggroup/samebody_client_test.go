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
	jpgRawFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.jpg", os.O_RDONLY, 0)
	require.NoError(t, err)
	jpgFile, err := io.ReadAll(jpgRawFile)
	require.NoError(t, err)
	require.NoError(t, jpgRawFile.Close())
	require.EqualValues(t, jpgFile, jpgResp)
}

func TestSameBodyClient_GetAvatarAsPNG(t *testing.T) {
	client, err := contentneggroup.NewContentNegotiationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSameBodyClient().GetAvatarAsPNG(context.Background(), nil)
	require.NoError(t, err)
	pngResp, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	pngRawFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.png", os.O_RDONLY, 0)
	require.NoError(t, err)
	pngFile, err := io.ReadAll(pngRawFile)
	require.NoError(t, err)
	require.NoError(t, pngRawFile.Close())
	require.EqualValues(t, pngFile, pngResp)
}
