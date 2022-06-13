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

func newImplicitClient() *ImplicitClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewImplicitClient("", "", nil, pl)
}

func TestImplicitGetOptionalGlobalQuery(t *testing.T) {
	client := newImplicitClient()
	result, err := client.GetOptionalGlobalQuery(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitGetRequiredGlobalPath(t *testing.T) {
	t.Skip("Cannot set nil for string parameter so test invalid for Go")
	client := newImplicitClient()
	result, err := client.GetRequiredGlobalPath(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitGetRequiredGlobalQuery(t *testing.T) {
	t.Skip("Cannot set nil for string parameter so test invalid for Go")
	client := newImplicitClient()
	result, err := client.GetRequiredGlobalQuery(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitGetRequiredPath(t *testing.T) {
	t.Skip("Cannot set nil for string parameter so test invalid for Go")
	client := newImplicitClient()
	result, err := client.GetRequiredPath(context.Background(), "", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitPutOptionalBody(t *testing.T) {
	client := newImplicitClient()
	result, err := client.PutOptionalBody(context.Background(), "", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitPutOptionalHeader(t *testing.T) {
	client := newImplicitClient()
	result, err := client.PutOptionalHeader(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestImplicitPutOptionalQuery(t *testing.T) {
	client := newImplicitClient()
	result, err := client.PutOptionalQuery(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
