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
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelPropertiesClient().SameAsModel(context.Background(), specialwordsgroup.SameAsModel{
		SameAsModel: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelPropertiesClient_DictMethods(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelPropertiesClient().DictMethods(context.Background(), specialwordsgroup.DictMethods{
		Clear:      to.Ptr("ok"),
		Copy:       to.Ptr("ok"),
		Get:        to.Ptr("ok"),
		Items:      to.Ptr("ok"),
		Keys:       to.Ptr("ok"),
		Pop:        to.Ptr("ok"),
		Popitem:    to.Ptr("ok"),
		Setdefault: to.Ptr("ok"),
		Update:     to.Ptr("ok"),
		Values:     to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
