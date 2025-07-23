// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package examplebasicgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestBasicServiceOperationGroupClient_Basic_Success(t *testing.T) {
	client, err := NewBasicServiceOperationGroupClient("http://localhost:3000", nil)
	require.NoError(t, err)
	reqBody := ActionRequest{
		StringProperty: to.Ptr("text"),
		ModelProperty: &Model{
			Int32Property:   to.Ptr(int32(1)),
			Float32Property: to.Ptr(float32(1.5)),
			EnumProperty:    to.Ptr(EnumEnumValue1),
		},
		ArrayProperty: []*string{to.Ptr("item")},
		RecordProperty: map[string]*string{
			"record": to.Ptr("value"),
		},
	}
	_, err = client.Basic(context.Background(), "query", "header", reqBody, nil)
	require.NoError(t, err)
}
