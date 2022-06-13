// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newStringClient() *StringClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewStringClient(pl)
}

func TestStringGetMBCS(t *testing.T) {
	client := newStringClient()
	result, err := client.GetMBCS(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr("啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€")); r != "" {
		t.Fatal(r)
	}
}

func TestStringPutMBCS(t *testing.T) {
	client := newStringClient()
	result, err := client.PutMBCS(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestStringGetBase64Encoded(t *testing.T) {
	client := newStringClient()
	result, err := client.GetBase64Encoded(context.Background(), nil)
	require.NoError(t, err)
	val := []byte("a string that gets encoded with base64")
	if r := cmp.Diff(result.Value, val); r != "" {
		t.Fatal(r)
	}
}

func TestStringGetBase64URLEncoded(t *testing.T) {
	client := newStringClient()
	result, err := client.GetBase64URLEncoded(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, []byte("a string that gets encoded with base64url")); r != "" {
		t.Fatal(r)
	}
}

func TestStringGetEmpty(t *testing.T) {
	client := newStringClient()
	result, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr("")); r != "" {
		t.Fatal(r)
	}
}

func TestStringGetNotProvided(t *testing.T) {
	client := newStringClient()
	result, err := client.GetNotProvided(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestStringGetNull(t *testing.T) {
	client := newStringClient()
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestStringGetNullBase64URLEncoded(t *testing.T) {
	client := newStringClient()
	result, err := client.GetNullBase64URLEncoded(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, ([]byte)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestStringGetWhitespace(t *testing.T) {
	client := newStringClient()
	result, err := client.GetWhitespace(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr("    Now is the time for all good men to come to the aid of their country    ")); r != "" {
		t.Fatal(r)
	}
}

func TestStringPutBase64URLEncoded(t *testing.T) {
	client := newStringClient()
	result, err := client.PutBase64URLEncoded(context.Background(), []byte("a string that gets encoded with base64url"), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestStringPutEmpty(t *testing.T) {
	client := newStringClient()
	result, err := client.PutEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestStringPutNull(t *testing.T) {
	t.Skip("missing x-nullable")
	client := newStringClient()
	result, err := client.PutNull(context.Background(), "", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestStringPutWhitespace(t *testing.T) {
	client := newStringClient()
	result, err := client.PutWhitespace(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
