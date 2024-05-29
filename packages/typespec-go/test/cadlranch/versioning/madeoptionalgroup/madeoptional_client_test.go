// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package madeoptionalgroup_test

import (
	"context"
	"madeoptionalgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func Test_Test(t *testing.T) {
	client, err := madeoptionalgroup.NewMadeOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.Test(context.Background(), madeoptionalgroup.TestModel{
		Prop: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Equal(t, madeoptionalgroup.MadeOptionalClientTestResponse{
		TestModel: madeoptionalgroup.TestModel{
			Prop: to.Ptr("foo"),
		},
	}, resp)
}
