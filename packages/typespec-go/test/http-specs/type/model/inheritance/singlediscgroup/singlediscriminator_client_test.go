//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package singlediscgroup_test

import (
	"context"
	"singlediscgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestSingleDiscriminatorClientGetLegacyModel(t *testing.T) {
	client, err := singlediscgroup.NewSingleDiscriminatorClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.GetLegacyModel(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.DinosaurClassification)
	require.Equal(t, &singlediscgroup.TRex{
		Kind: to.Ptr("t-rex"),
		Size: to.Ptr[int32](20),
	}, resp.DinosaurClassification)
}

func TestSingleDiscriminatorClientGetMissingDiscriminator(t *testing.T) {
	client, err := singlediscgroup.NewSingleDiscriminatorClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.GetMissingDiscriminator(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.BirdClassification)
	require.Equal(t, &singlediscgroup.Bird{
		Wingspan: to.Ptr[int32](1),
	}, resp.BirdClassification)
}

func TestSingleDiscriminatorClientGetModel(t *testing.T) {
	client, err := singlediscgroup.NewSingleDiscriminatorClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.GetModel(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.BirdClassification)
	require.Equal(t, &singlediscgroup.Sparrow{
		Kind:     to.Ptr("sparrow"),
		Wingspan: to.Ptr[int32](1),
	}, resp.BirdClassification)
}

func TestSingleDiscriminatorClientGetRecursiveModel(t *testing.T) {
	client, err := singlediscgroup.NewSingleDiscriminatorClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.GetRecursiveModel(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.BirdClassification)
	require.Equal(t, &singlediscgroup.Eagle{
		Kind:     to.Ptr("eagle"),
		Wingspan: to.Ptr[int32](5),
		Partner: &singlediscgroup.Goose{
			Kind:     to.Ptr("goose"),
			Wingspan: to.Ptr[int32](2),
		},
		Friends: []singlediscgroup.BirdClassification{
			&singlediscgroup.SeaGull{
				Kind:     to.Ptr("seagull"),
				Wingspan: to.Ptr[int32](2),
			},
		},
		Hate: map[string]singlediscgroup.BirdClassification{
			"key3": &singlediscgroup.Sparrow{
				Kind:     to.Ptr("sparrow"),
				Wingspan: to.Ptr[int32](1),
			},
		},
	}, resp.BirdClassification)
}

func TestSingleDiscriminatorClientGetWrongDiscriminator(t *testing.T) {
	client, err := singlediscgroup.NewSingleDiscriminatorClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.GetWrongDiscriminator(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.BirdClassification)
	require.Equal(t, &singlediscgroup.Bird{
		Kind:     to.Ptr("wrongKind"),
		Wingspan: to.Ptr[int32](1),
	}, resp.BirdClassification)
}

func TestSingleDiscriminatorClientPutModel(t *testing.T) {
	client, err := singlediscgroup.NewSingleDiscriminatorClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.PutModel(context.Background(), &singlediscgroup.Sparrow{
		Kind:     to.Ptr("sparrow"),
		Wingspan: to.Ptr[int32](1),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestSingleDiscriminatorClientPutRecursiveModel(t *testing.T) {
	client, err := singlediscgroup.NewSingleDiscriminatorClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.PutRecursiveModel(context.Background(), &singlediscgroup.Eagle{
		Kind:     to.Ptr("eagle"),
		Wingspan: to.Ptr[int32](5),
		Partner: &singlediscgroup.Goose{
			Kind:     to.Ptr("goose"),
			Wingspan: to.Ptr[int32](2),
		},
		Friends: []singlediscgroup.BirdClassification{
			&singlediscgroup.SeaGull{
				Kind:     to.Ptr("seagull"),
				Wingspan: to.Ptr[int32](2),
			},
		},
		Hate: map[string]singlediscgroup.BirdClassification{
			"key3": &singlediscgroup.Sparrow{
				Kind:     to.Ptr("sparrow"),
				Wingspan: to.Ptr[int32](1),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
