//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package specialwordsgroup_test

import (
	"context"
	"specialwordsgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestModelPropertiesClient_SameAsModel(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelPropertiesClient().SameAsModel(context.Background(), specialwordsgroup.SameAsModel{
		SameAsModel: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
