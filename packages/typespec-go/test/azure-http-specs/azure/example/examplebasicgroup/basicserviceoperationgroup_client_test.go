// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package examplebasicgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicServiceOperationGroupClient_Basic_Success(t *testing.T) {
	client, err := NewBasicServiceOperationGroupClient(nil)
	require.NoError(t, err)
	reqBody := ActionRequest{
		StringProperty: toString("text"),
		ModelProperty: &Model{
			Int32Property:   func(i int32) *int32 { return &i }(1),
			Float32Property: func(f float32) *float32 { return &f }(1.5),
			EnumProperty:    func(e Enum) *Enum { return &e }(EnumEnumValue1),
		},
		ArrayProperty: []*string{toString("item")},
		RecordProperty: map[string]*string{
			"record": toString("value"),
		},
	}
	_, err = client.Basic(context.Background(), "query", "header", reqBody, nil)
	require.NoError(t, err)
}


func toString(v string) *string {
	return &v
}
