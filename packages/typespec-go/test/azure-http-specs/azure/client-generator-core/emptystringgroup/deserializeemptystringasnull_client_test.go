// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package emptystringgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeserializeEmptyStringAsNullClient_Get(t *testing.T) {
	client, err := NewDeserializeEmptyStringAsNullClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.Nil(t, resp.SampleURL)
}
