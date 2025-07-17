// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newDictionaryClient(t *testing.T) *DictionaryClient {
	client, err := NewDictionaryClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewDictionaryClient(endpoint string, options *azcore.ClientOptions) (*DictionaryClient, error) {
	client, err := azcore.NewClient("complexgroup.DictionaryClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &DictionaryClient{internal: client, endpoint: endpoint}, nil
}

func TestDictionaryGetEmpty(t *testing.T) {
	client := newDictionaryClient(t)
	result, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.DictionaryWrapper, DictionaryWrapper{DefaultProgram: map[string]*string{}}); r != "" {
		t.Fatal(r)
	}
}

func TestDictionaryGetNotProvided(t *testing.T) {
	client := newDictionaryClient(t)
	result, err := client.GetNotProvided(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.DictionaryWrapper, DictionaryWrapper{}); r != "" {
		t.Fatal(r)
	}
}

func TestDictionaryGetNull(t *testing.T) {
	client := newDictionaryClient(t)
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.DictionaryWrapper, DictionaryWrapper{}); r != "" {
		t.Fatal(r)
	}
}

func TestDictionaryGetValid(t *testing.T) {
	client := newDictionaryClient(t)
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	s1, s2, s3, s4 := "notepad", "mspaint", "excel", ""
	val := DictionaryWrapper{DefaultProgram: map[string]*string{"txt": &s1, "bmp": &s2, "xls": &s3, "exe": &s4, "": nil}}
	if r := cmp.Diff(result.DictionaryWrapper, val); r != "" {
		t.Fatal(r)
	}
}

func TestDictionaryPutEmpty(t *testing.T) {
	client := newDictionaryClient(t)
	result, err := client.PutEmpty(context.Background(), DictionaryWrapper{DefaultProgram: map[string]*string{}}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestDictionaryPutValid(t *testing.T) {
	client := newDictionaryClient(t)
	s1, s2, s3, s4 := "notepad", "mspaint", "excel", ""
	result, err := client.PutValid(context.Background(), DictionaryWrapper{DefaultProgram: map[string]*string{"txt": &s1, "bmp": &s2, "xls": &s3, "exe": &s4, "": nil}}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
