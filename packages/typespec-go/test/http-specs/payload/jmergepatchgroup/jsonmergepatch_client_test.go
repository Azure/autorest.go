// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package jmergepatchgroup_test

import (
	"context"
	"jmergepatchgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestJsonMergePatchClient_CreateResource(t *testing.T) {
	client, err := jmergepatchgroup.NewJSONMergePatchClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.CreateResource(context.Background(), jmergepatchgroup.Resource{
		Name:        to.Ptr("Madge"),
		Description: to.Ptr("desc"),
		Map: map[string]*jmergepatchgroup.InnerModel{
			"key": {
				Name:        to.Ptr("InnerMadge"),
				Description: to.Ptr("innerDesc"),
			},
		},
		Array: []*jmergepatchgroup.InnerModel{
			{
				Name:        to.Ptr("InnerMadge"),
				Description: to.Ptr("innerDesc"),
			},
		},
		IntValue:   to.Ptr[int32](1),
		FloatValue: to.Ptr[float32](1.25),
		InnerModel: &jmergepatchgroup.InnerModel{
			Name:        to.Ptr("InnerMadge"),
			Description: to.Ptr("innerDesc"),
		},
		IntArray: []*int32{
			to.Ptr[int32](1),
			to.Ptr[int32](2),
			to.Ptr[int32](3),
		},
	}, nil)
	require.NoError(t, err)
	require.Equal(t, jmergepatchgroup.Resource{
		Name:        to.Ptr("Madge"),
		Description: to.Ptr("desc"),
		Map: map[string]*jmergepatchgroup.InnerModel{
			"key": {
				Name:        to.Ptr("InnerMadge"),
				Description: to.Ptr("innerDesc"),
			},
		},
		Array: []*jmergepatchgroup.InnerModel{
			{
				Name:        to.Ptr("InnerMadge"),
				Description: to.Ptr("innerDesc"),
			},
		},
		IntValue:   to.Ptr[int32](1),
		FloatValue: to.Ptr[float32](1.25),
		InnerModel: &jmergepatchgroup.InnerModel{
			Name:        to.Ptr("InnerMadge"),
			Description: to.Ptr("innerDesc"),
		},
		IntArray: []*int32{
			to.Ptr[int32](1),
			to.Ptr[int32](2),
			to.Ptr[int32](3),
		},
	}, resp.Resource)
}

func TestJsonMergePatchClient_UpdateOptionalResource(t *testing.T) {
	client, err := jmergepatchgroup.NewJSONMergePatchClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.UpdateOptionalResource(context.Background(),
		jmergepatchgroup.ResourcePatch{
			Description: azcore.NullValue[*string](),
			Map: map[string]*jmergepatchgroup.InnerModel{
				"key": {
					Description: azcore.NullValue[*string](),
				},
				"key2": nil,
			},
			Array:      azcore.NullValue[[]*jmergepatchgroup.InnerModel](),
			IntValue:   azcore.NullValue[*int32](),
			FloatValue: azcore.NullValue[*float32](),
			InnerModel: azcore.NullValue[*jmergepatchgroup.InnerModel](),
			IntArray:   azcore.NullValue[[]*int32](),
		}, nil)
	require.NoError(t, err)
	require.Equal(t, jmergepatchgroup.Resource{
		Name: to.Ptr("Madge"),
		Map: map[string]*jmergepatchgroup.InnerModel{
			"key": {
				Name: to.Ptr("InnerMadge"),
			},
		},
	}, resp.Resource)
}

func TestJsonMergePatchClient_UpdateResource(t *testing.T) {
	client, err := jmergepatchgroup.NewJSONMergePatchClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.UpdateResource(context.Background(), jmergepatchgroup.ResourcePatch{
		Description: azcore.NullValue[*string](),
		Map: map[string]*jmergepatchgroup.InnerModel{
			"key": {
				Description: azcore.NullValue[*string](),
			},
			"key2": nil,
		},
		Array:      azcore.NullValue[[]*jmergepatchgroup.InnerModel](),
		IntValue:   azcore.NullValue[*int32](),
		FloatValue: azcore.NullValue[*float32](),
		InnerModel: azcore.NullValue[*jmergepatchgroup.InnerModel](),
		IntArray:   azcore.NullValue[[]*int32](),
	}, nil)
	require.NoError(t, err)
	require.Equal(t, jmergepatchgroup.Resource{
		Name: to.Ptr("Madge"),
		Map: map[string]*jmergepatchgroup.InnerModel{
			"key": {
				Name: to.Ptr("InnerMadge"),
			},
		},
	}, resp.Resource)
}
