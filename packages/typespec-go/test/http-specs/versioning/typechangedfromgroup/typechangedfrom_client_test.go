// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package typechangedfromgroup_test

import (
	"context"
	"testing"
	"typechangedfromgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func Test_Test(t *testing.T) {
	client, err := typechangedfromgroup.NewTypeChangedFromClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Test(context.Background(), typechangedfromgroup.TestModel{
		Prop:        to.Ptr("foo"),
		ChangedProp: to.Ptr("bar"),
	}, "baz", nil)
	require.NoError(t, err)
	require.Equal(t, typechangedfromgroup.TypeChangedFromClientTestResponse{
		TestModel: typechangedfromgroup.TestModel{
			Prop:        to.Ptr("foo"),
			ChangedProp: to.Ptr("bar"),
		},
	}, resp)
}
