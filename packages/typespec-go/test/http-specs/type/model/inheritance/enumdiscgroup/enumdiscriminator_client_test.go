//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package enumdiscgroup_test

import (
	"context"
	"enumdiscgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestEnumDiscriminatorClientGetExtensibleModel(t *testing.T) {
	client, err := enumdiscgroup.NewEnumDiscriminatorClient(nil)
	require.NoError(t, err)
	resp, err := client.GetExtensibleModel(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.DogClassification)
	require.Equal(t, &enumdiscgroup.Golden{
		Kind:   to.Ptr(enumdiscgroup.DogKindGolden),
		Weight: to.Ptr[int32](10),
	}, resp.DogClassification)
}

func TestEnumDiscriminatorClientGetExtensibleModelMissingDiscriminator(t *testing.T) {
	client, err := enumdiscgroup.NewEnumDiscriminatorClient(nil)
	require.NoError(t, err)
	resp, err := client.GetExtensibleModelMissingDiscriminator(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.DogClassification)
	require.Equal(t, &enumdiscgroup.Dog{
		Weight: to.Ptr[int32](10),
	}, resp.DogClassification)
}

func TestEnumDiscriminatorClientGetExtensibleModelWrongDiscriminator(t *testing.T) {
	client, err := enumdiscgroup.NewEnumDiscriminatorClient(nil)
	require.NoError(t, err)
	resp, err := client.GetExtensibleModelWrongDiscriminator(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.DogClassification)
	require.Equal(t, &enumdiscgroup.Dog{
		Kind:   to.Ptr(enumdiscgroup.DogKind("wrongKind")),
		Weight: to.Ptr[int32](8),
	}, resp.DogClassification)
}

func TestEnumDiscriminatorClientGetFixedModel(t *testing.T) {
	client, err := enumdiscgroup.NewEnumDiscriminatorClient(nil)
	require.NoError(t, err)
	resp, err := client.GetFixedModel(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.SnakeClassification)
	require.Equal(t, &enumdiscgroup.Cobra{
		Kind:   to.Ptr(enumdiscgroup.SnakeKindCobra),
		Length: to.Ptr[int32](10),
	}, resp.SnakeClassification)
}

func TestEnumDiscriminatorClientGetFixedModelMissingDiscriminator(t *testing.T) {
	client, err := enumdiscgroup.NewEnumDiscriminatorClient(nil)
	require.NoError(t, err)
	resp, err := client.GetFixedModelMissingDiscriminator(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.SnakeClassification)
	require.Equal(t, &enumdiscgroup.Snake{
		Length: to.Ptr[int32](10),
	}, resp.SnakeClassification)
}

func TestEnumDiscriminatorClientGetFixedModelWrongDiscriminator(t *testing.T) {
	client, err := enumdiscgroup.NewEnumDiscriminatorClient(nil)
	require.NoError(t, err)
	resp, err := client.GetFixedModelWrongDiscriminator(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.SnakeClassification)
	require.Equal(t, &enumdiscgroup.Snake{
		Kind:   to.Ptr(enumdiscgroup.SnakeKind("wrongKind")),
		Length: to.Ptr[int32](8),
	}, resp.SnakeClassification)
}

func TestEnumDiscriminatorClientPutExtensibleModel(t *testing.T) {
	client, err := enumdiscgroup.NewEnumDiscriminatorClient(nil)
	require.NoError(t, err)
	resp, err := client.PutExtensibleModel(context.Background(), &enumdiscgroup.Golden{
		Kind:   to.Ptr(enumdiscgroup.DogKindGolden),
		Weight: to.Ptr[int32](10),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestEnumDiscriminatorClientPutFixedModel(t *testing.T) {
	client, err := enumdiscgroup.NewEnumDiscriminatorClient(nil)
	require.NoError(t, err)
	resp, err := client.PutFixedModel(context.Background(), &enumdiscgroup.Cobra{
		Kind:   to.Ptr(enumdiscgroup.SnakeKindCobra),
		Length: to.Ptr[int32](10),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
