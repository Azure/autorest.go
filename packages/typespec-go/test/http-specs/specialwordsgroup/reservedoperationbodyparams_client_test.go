// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package specialwordsgroup_test

import (
	"context"
	"specialwordsgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReservedOperationBodyParamsClient_WithItems(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsReservedOperationBodyParamsClient().WithItems(context.Background(), []string{"item"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
