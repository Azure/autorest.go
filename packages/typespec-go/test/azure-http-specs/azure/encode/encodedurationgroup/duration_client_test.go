// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package encodedurationgroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDurationClient_DurationConstant(t *testing.T) {
	client, err := NewDurationClient(nil)
	require.NoError(t, err)
	body := DurationModel{}
	result, err := client.DurationConstant(context.Background(), body, nil)
	require.NoError(t, err)
	require.Equal(t, DurationClientDurationConstantResponse{}, result)
}

func TestDurationClient_durationConstantCreateRequest(t *testing.T) {
	client := &DurationClient{}
	body := DurationModel{}
	req, err := client.durationConstantCreateRequest(context.Background(), body, nil)
	require.NoError(t, err)
	require.Equal(t, http.MethodPut, req.Raw().Method)
	require.Contains(t, req.Raw().URL.Path, "/azure/encode/duration/duration-constant")
	require.Equal(t, "application/json", req.Raw().Header.Get("Content-Type"))
}
