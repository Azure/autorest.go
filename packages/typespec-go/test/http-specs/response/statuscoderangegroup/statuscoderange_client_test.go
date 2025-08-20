// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package statuscoderangegroup_test

import (
	"context"
	"statuscoderangegroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatusCodeRangeClientErrorResponseStatusCode404(t *testing.T) {
	client, err := statuscoderangegroup.NewStatusCodeRangeClient("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.ErrorResponseStatusCode404(context.Background(), nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "404")
	require.Contains(t, err.Error(), `"code": "not-found"`)
	require.Contains(t, err.Error(), `"resourceId": "resource1"`)
}

func TestStatusCodeRangeClientErrorResponseStatusCodeInRange(t *testing.T) {
	client, err := statuscoderangegroup.NewStatusCodeRangeClient("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.ErrorResponseStatusCodeInRange(context.Background(), nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "494")
	require.Contains(t, err.Error(), `"code": "request-header-too-large"`)
	require.Contains(t, err.Error(), `"message": "Request header too large"`)
}
