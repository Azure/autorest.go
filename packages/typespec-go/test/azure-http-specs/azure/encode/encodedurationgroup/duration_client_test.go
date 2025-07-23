// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package encodedurationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDurationClient_DurationConstant(t *testing.T) {
	client, err := NewDurationClient("http://localhost:3000", nil)
	require.NoError(t, err)
	input := "1.02:59:59.5000000"
	body := DurationModel{
		Input: &input,
	}
	result, err := client.DurationConstant(context.Background(), body, nil)
	require.NoError(t, err)
	require.Equal(t, DurationClientDurationConstantResponse{}, result)
}
