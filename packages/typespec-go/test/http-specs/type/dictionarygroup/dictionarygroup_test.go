//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package dictionarygroup_test

import (
	"context"
	"dictionarygroup"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestBooleanValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryBooleanValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*bool{
		"k1": to.Ptr(true),
		"k2": to.Ptr(false),
	}, resp.Value)
}

func TestBooleanValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryBooleanValueClient().Put(context.Background(), map[string]*bool{
		"k1": to.Ptr(true),
		"k2": to.Ptr(false),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestDatetimeValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryDatetimeValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*time.Time{
		"k1": to.Ptr(time.Date(2022, time.August, 26, 18, 38, 0, 0, time.UTC)),
	}, resp.Value)
}

func TestDatetimeValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryDatetimeValueClient().Put(context.Background(), map[string]*time.Time{
		"k1": to.Ptr(time.Date(2022, time.August, 26, 18, 38, 0, 0, time.UTC)),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestDurationValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryDurationValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*string{
		"k1": to.Ptr("P123DT22H14M12.011S"),
	}, resp.Value)
}

func TestDurationValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryDurationValueClient().Put(context.Background(), map[string]*string{
		"k1": to.Ptr("P123DT22H14M12.011S"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFloat32ValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryFloat32ValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*float32{
		"k1": to.Ptr[float32](43.125),
	}, resp.Value)
}

func TestFloat32ValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryFloat32ValueClient().Put(context.Background(), map[string]*float32{
		"k1": to.Ptr[float32](43.125),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestInt32ValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryInt32ValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*int32{
		"k1": to.Ptr[int32](1),
		"k2": to.Ptr[int32](2),
	}, resp.Value)
}

func TestInt32ValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryInt32ValueClient().Put(context.Background(), map[string]*int32{
		"k1": to.Ptr[int32](1),
		"k2": to.Ptr[int32](2),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestInt64ValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryInt64ValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*int64{
		"k1": to.Ptr[int64](9007199254740991),
		"k2": to.Ptr[int64](-9007199254740991),
	}, resp.Value)
}

func TestInt64ValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryInt64ValueClient().Put(context.Background(), map[string]*int64{
		"k1": to.Ptr[int64](9007199254740991),
		"k2": to.Ptr[int64](-9007199254740991),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryModelValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*dictionarygroup.InnerModel{
		"k1": {
			Property: to.Ptr("hello"),
		},
		"k2": {
			Property: to.Ptr("world"),
		},
	}, resp.Value)
}

func TestModelValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryModelValueClient().Put(context.Background(), map[string]*dictionarygroup.InnerModel{
		"k1": {
			Property: to.Ptr("hello"),
		},
		"k2": {
			Property: to.Ptr("world"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestNullableFloatValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryNullableFloatValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*float32{
		"k1": to.Ptr[float32](1.25),
		"k2": to.Ptr[float32](0.5),
		"k3": nil,
	}, resp.Value)
}

func TestNullableFloatValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryNullableFloatValueClient().Put(context.Background(), map[string]*float32{
		"k1": to.Ptr[float32](1.25),
		"k2": to.Ptr[float32](0.5),
		"k3": nil,
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRecursiveModelValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryRecursiveModelValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*dictionarygroup.InnerModel{
		"k1": {
			Property: to.Ptr("hello"),
			Children: map[string]*dictionarygroup.InnerModel{},
		},
		"k2": {
			Property: to.Ptr("world"),
			Children: map[string]*dictionarygroup.InnerModel{
				"k2.1": {
					Property: to.Ptr("inner world"),
				},
			},
		},
	}, resp.Value)
}

func TestRecursiveModelValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryRecursiveModelValueClient().Put(context.Background(), map[string]*dictionarygroup.InnerModel{
		"k1": {
			Property: to.Ptr("hello"),
			Children: map[string]*dictionarygroup.InnerModel{},
		},
		"k2": {
			Property: to.Ptr("world"),
			Children: map[string]*dictionarygroup.InnerModel{
				"k2.1": {
					Property: to.Ptr("inner world"),
				},
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestStringValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryStringValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]*string{
		"k1": to.Ptr("hello"),
		"k2": to.Ptr(""),
	}, resp.Value)
}

func TestStringValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryStringValueClient().Put(context.Background(), map[string]*string{
		"k1": to.Ptr("hello"),
		"k2": to.Ptr(""),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestUnknownValueClientGet(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryUnknownValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, map[string]any{
		"k1": float64(1),
		"k2": "hello",
		"k3": nil,
	}, resp.Value)
}

func TestUnknownValueClientPut(t *testing.T) {
	client, err := dictionarygroup.NewDictionaryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDictionaryUnknownValueClient().Put(context.Background(), map[string]any{
		"k1": float64(1),
		"k2": "hello",
		"k3": nil,
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
