// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup_test

import (
	"context"
	"testing"
	"xmlgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestXMLModelWithRenamedFieldsValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithRenamedFieldsValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlgroup.ModelWithRenamedFields{
		InputData: &xmlgroup.SimpleModel{
			Name: to.Ptr("foo"),
			Age:  to.Ptr[int32](123),
		},
		OutputData: &xmlgroup.SimpleModel{
			Name: to.Ptr("bar"),
			Age:  to.Ptr[int32](456),
		},
	}, resp.ModelWithRenamedFields)
}

func TestXMLModelWithRenamedFieldsValueClient_Put(t *testing.T) {
	client, err := xmlgroup.NewXMLClient(nil)
	require.NoError(t, err)
	resp, err := client.NewXMLModelWithRenamedFieldsValueClient().Put(context.Background(), xmlgroup.ModelWithRenamedFields{
		InputData: &xmlgroup.SimpleModel{
			Name: to.Ptr("foo"),
			Age:  to.Ptr[int32](123),
		},
		OutputData: &xmlgroup.SimpleModel{
			Name: to.Ptr("bar"),
			Age:  to.Ptr[int32](456),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
