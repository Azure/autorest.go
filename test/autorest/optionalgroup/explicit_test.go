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

func newExplicitClient() *ExplicitClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewExplicitClient(pl)
}

func TestExplicitPostOptionalArrayHeader(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalArrayHeader(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestExplicitPostOptionalArrayParameter(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalArrayParameter(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestExplicitPostOptionalArrayProperty(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalArrayProperty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestExplicitPostOptionalClassParameter(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalClassParameter(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestExplicitPostOptionalClassProperty(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalClassProperty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestExplicitPostOptionalIntegerHeader(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalIntegerHeader(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestExplicitPostOptionalIntegerParameter(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalIntegerParameter(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestExplicitPostOptionalIntegerProperty(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalIntegerProperty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestExplicitPostOptionalStringHeader(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalStringHeader(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestExplicitPostOptionalStringParameter(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalStringParameter(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestExplicitPostOptionalStringProperty(t *testing.T) {
	client := newExplicitClient()
	result, err := client.PostOptionalStringProperty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// TODO the goal of this test is to throw an exception but nils are acceptable for  []strings in go
func TestExplicitPostRequiredArrayHeader(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredArrayHeader(context.Background(), nil, nil)
	require.Error(t, err)
	require.Zero(t, result)
}

// TODO the goal of this test is to throw an exception but nils are acceptable for  []strings in go
func TestExplicitPostRequiredArrayParameter(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredArrayParameter(context.Background(), nil, nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestExplicitPostRequiredArrayProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredArrayProperty(context.Background(), ArrayWrapper{Value: nil}, nil)
	require.Error(t, err)
	require.Zero(t, result)
}

// TODO check this test
func TestExplicitPostRequiredClassParameter(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredClassParameter(context.Background(), Product{}, nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestExplicitPostRequiredClassProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredClassProperty(context.Background(), ClassWrapper{}, nil)
	require.Error(t, err)
	require.Zero(t, result)
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredIntegerHeader(t *testing.T) {
	t.Skip("cannot set nil for int32 in Go")
	client := newExplicitClient()
	result, err := client.PostRequiredIntegerHeader(context.Background(), 0, nil)
	require.Error(t, err)
	require.Zero(t, result)
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredIntegerParameter(t *testing.T) {
	t.Skip("cannot set nil for int32 in Go")
	client := newExplicitClient()
	result, err := client.PostRequiredIntegerParameter(context.Background(), 0, nil)
	require.Error(t, err)
	require.Zero(t, result)
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredIntegerProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredIntegerProperty(context.Background(), IntWrapper{}, nil)
	require.Error(t, err)
	require.Zero(t, result)
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredStringHeader(t *testing.T) {
	t.Skip("cannot set nil for string in Go")
	client := newExplicitClient()
	result, err := client.PostRequiredStringHeader(context.Background(), "", nil)
	require.Error(t, err)
	require.Zero(t, result)
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredStringParameter(t *testing.T) {
	t.Skip("cannot set nil for string in Go")
	client := newExplicitClient()
	result, err := client.PostRequiredStringParameter(context.Background(), "", nil)
	require.Error(t, err)
	require.Zero(t, result)
}

// TODO check this test is does pass if we query the endpoint but that is not the expected behavior
func TestExplicitPostRequiredStringProperty(t *testing.T) {
	t.Skip("are not validating parameters in track2")
	client := newExplicitClient()
	result, err := client.PostRequiredStringProperty(context.Background(), StringWrapper{}, nil)
	require.Error(t, err)
	require.Zero(t, result)
}
