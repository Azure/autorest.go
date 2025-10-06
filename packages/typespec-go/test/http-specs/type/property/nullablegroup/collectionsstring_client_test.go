// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nullablegroup_test

import (
	"context"
	"nullablegroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestCollectionsStringClientGetNonNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewNullableCollectionsStringClient().GetNonNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.CollectionsStringProperty{
		NullableProperty: []*string{to.Ptr("hello"), to.Ptr("world")},
		RequiredProperty: to.Ptr("foo"),
	}, resp.CollectionsStringProperty)
}

func TestCollectionsStringClientGetNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewNullableCollectionsStringClient().GetNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.CollectionsStringProperty{
		RequiredProperty: to.Ptr("foo"),
	}, resp.CollectionsStringProperty)
}

func TestCollectionsStringClientPatchNonNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewNullableCollectionsStringClient().PatchNonNull(context.Background(), nullablegroup.CollectionsStringProperty{
		NullableProperty: []*string{to.Ptr("hello"), to.Ptr("world")},
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
}

func TestCollectionsStringClientPatchNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewNullableCollectionsStringClient().PatchNull(context.Background(), nullablegroup.CollectionsStringProperty{
		NullableProperty: azcore.NullValue[[]*string](),
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
}
