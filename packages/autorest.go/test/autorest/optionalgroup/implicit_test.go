// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package optionalgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newImplicitClient(t *testing.T) *ImplicitClient {
	client, err := NewImplicitClient("", "", nil, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewImplicitClient(equiredGlobalPath string, requiredGlobalQuery string, optionalGlobalQuery *int32, options *azcore.ClientOptions) (*ImplicitClient, error) {
	client, err := azcore.NewClient("optionalgroup.ImplicitClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ImplicitClient{internal: client}, nil
}

func TestImplicitGetOptionalGlobalQuery(t *testing.T) {
	client := newImplicitClient(t)
	result, err := client.GetOptionalGlobalQuery(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitGetRequiredGlobalPath(t *testing.T) {
	t.Skip("Cannot set nil for string parameter so test invalid for Go")
	client := newImplicitClient(t)
	result, err := client.GetRequiredGlobalPath(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitGetRequiredGlobalQuery(t *testing.T) {
	t.Skip("Cannot set nil for string parameter so test invalid for Go")
	client := newImplicitClient(t)
	result, err := client.GetRequiredGlobalQuery(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitGetRequiredPath(t *testing.T) {
	t.Skip("Cannot set nil for string parameter so test invalid for Go")
	client := newImplicitClient(t)
	result, err := client.GetRequiredPath(context.Background(), "", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitPutOptionalBody(t *testing.T) {
	client := newImplicitClient(t)
	result, err := client.PutOptionalBody(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitPutOptionalHeader(t *testing.T) {
	client := newImplicitClient(t)
	result, err := client.PutOptionalHeader(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitPutOptionalQuery(t *testing.T) {
	client := newImplicitClient(t)
	result, err := client.PutOptionalQuery(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
