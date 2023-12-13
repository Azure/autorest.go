//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytesgroup_test

import (
	"bytesgroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeaderClientBase64(t *testing.T) {
	client, err := bytesgroup.NewHeaderClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientBase64URL(t *testing.T) {
	client, err := bytesgroup.NewHeaderClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64URL(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientBase64URLArray(t *testing.T) {
	client, err := bytesgroup.NewHeaderClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64URLArray(context.Background(), [][]byte{[]byte("test"), []byte("test")}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientDefault(t *testing.T) {
	client, err := bytesgroup.NewHeaderClient(nil)
	require.NoError(t, err)
	resp, err := client.Default(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestPropertyClientBase64(t *testing.T) {
	client, err := bytesgroup.NewPropertyClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64(context.Background(), bytesgroup.Base64BytesProperty{
		Value: []byte("test"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestPropertyClientBase64URL(t *testing.T) {
	client, err := bytesgroup.NewPropertyClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64URL(context.Background(), bytesgroup.Base64URLBytesProperty{
		Value: []byte("test"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestPropertyClientBase64URLArray(t *testing.T) {
	client, err := bytesgroup.NewPropertyClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64URLArray(context.Background(), bytesgroup.Base64URLArrayBytesProperty{
		Value: [][]byte{[]byte("test"), []byte("test")},
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, [][]byte{[]byte("test"), []byte("test")}, resp.Value)
}

func TestPropertyClientDefault(t *testing.T) {
	client, err := bytesgroup.NewPropertyClient(nil)
	require.NoError(t, err)
	resp, err := client.Default(context.Background(), bytesgroup.DefaultBytesProperty{
		Value: []byte("test"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestQueryClientBase64(t *testing.T) {
	client, err := bytesgroup.NewQueryClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientBase64URL(t *testing.T) {
	client, err := bytesgroup.NewQueryClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64URL(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientBase64URLArray(t *testing.T) {
	client, err := bytesgroup.NewQueryClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64URLArray(context.Background(), [][]byte{[]byte("test"), []byte("test")}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientDefault(t *testing.T) {
	client, err := bytesgroup.NewQueryClient(nil)
	require.NoError(t, err)
	resp, err := client.Default(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRequestBodyClientBase64(t *testing.T) {
	client, err := bytesgroup.NewRequestBodyClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRequestBodyClientBase64URL(t *testing.T) {
	client, err := bytesgroup.NewRequestBodyClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64URL(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRequestBodyClientCustomContentType(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/15")
}

func TestRequestBodyClientDefault(t *testing.T) {
	client, err := bytesgroup.NewRequestBodyClient(nil)
	require.NoError(t, err)
	resp, err := client.Default(context.Background(), []byte("test"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRequestBodyClientOctetStream(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/15")
}

func TestResponseBodyClientBase64(t *testing.T) {
	client, err := bytesgroup.NewResponseBodyClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestResponseBodyClientBase64URL(t *testing.T) {
	client, err := bytesgroup.NewResponseBodyClient(nil)
	require.NoError(t, err)
	resp, err := client.Base64URL(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestResponseBodyClientCustomContent(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/15")
}

func TestResponseBodyClientDefault(t *testing.T) {
	client, err := bytesgroup.NewResponseBodyClient(nil)
	require.NoError(t, err)
	resp, err := client.Default(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []byte("test"), resp.Value)
}

func TestResponseBodyClientOctetStream(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/15")
}
