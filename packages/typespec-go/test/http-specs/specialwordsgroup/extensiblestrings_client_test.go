// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package specialwordsgroup_test

import (
	"context"
	"specialwordsgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtensibleStringsClient_PutExtensibleStringValue(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsExtensibleStringsClient().PutExtensibleStringValue(context.Background(), specialwordsgroup.ExtensibleStringClass, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, specialwordsgroup.ExtensibleStringClass, *resp.Value)
}
