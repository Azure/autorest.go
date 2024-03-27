//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytesgroup_test

import (
	"bytesgroup"
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeaderClientBase64(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewHeaderClient().Base64(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientBase64URL(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewHeaderClient().Base64URL(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientBase64URLArray(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewHeaderClient().Base64URLArray(context.Background(), [][]byte{[]byte("test"), []byte("test")}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientDefault(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewHeaderClient().Default(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestPropertyClientBase64(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPropertyClient().Base64(context.Background(), bytesgroup.Base64BytesProperty{
		Value: []byte("test"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestPropertyClientBase64URL(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPropertyClient().Base64URL(context.Background(), bytesgroup.Base64URLBytesProperty{
		Value: []byte("test"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestPropertyClientBase64URLArray(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPropertyClient().Base64URLArray(context.Background(), bytesgroup.Base64URLArrayBytesProperty{
		Value: [][]byte{[]byte("test"), []byte("test")},
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, [][]byte{[]byte("test"), []byte("test")}, resp.Value)
}

func TestPropertyClientDefault(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPropertyClient().Default(context.Background(), bytesgroup.DefaultBytesProperty{
		Value: []byte("test"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestQueryClientBase64(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewQueryClient().Base64(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientBase64URL(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewQueryClient().Base64URL(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientBase64URLArray(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewQueryClient().Base64URLArray(context.Background(), [][]byte{[]byte("test"), []byte("test")}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientDefault(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewQueryClient().Default(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRequestBodyClientBase64(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewRequestBodyClient().Base64(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRequestBodyClientBase64URL(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewRequestBodyClient().Base64URL(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRequestBodyClientCustomContentType(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	pngFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.png", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer pngFile.Close()
	resp, err := client.NewRequestBodyClient().CustomContentType(context.Background(), pngFile, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRequestBodyClientDefault(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewRequestBodyClient().Default(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRequestBodyClientOctetStream(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	pngFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.png", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer pngFile.Close()
	resp, err := client.NewRequestBodyClient().OctetStream(context.Background(), pngFile, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResponseBodyClientBase64(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewResponseBodyClient().Base64(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestResponseBodyClientBase64URL(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewResponseBodyClient().Base64URL(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestResponseBodyClientCustomContent(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewResponseBodyClient().CustomContentType(context.Background(), nil)
	require.NoError(t, err)
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())
	pngFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.png", os.O_RDONLY, 0)
	require.NoError(t, err)
	pngBody, err := io.ReadAll(pngFile)
	require.NoError(t, err)
	require.NoError(t, pngFile.Close())
	require.EqualValues(t, pngBody, respBody)
}

func TestResponseBodyClientDefault(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewResponseBodyClient().Default(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestResponseBodyClientOctetStream(t *testing.T) {
	client, err := bytesgroup.NewBytesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewResponseBodyClient().OctetStream(context.Background(), nil)
	require.NoError(t, err)
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())
	pngFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.png", os.O_RDONLY, 0)
	require.NoError(t, err)
	pngBody, err := io.ReadAll(pngFile)
	require.NoError(t, err)
	require.NoError(t, pngFile.Close())
	require.EqualValues(t, pngBody, respBody)
}
