// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package statuscoderangegroup_test

import (
	"context"
	"statuscoderangegroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatusCodeRangeGroupClient_ErrorResponseStatusCode404(t *testing.T) {
	client, err := statuscoderangegroup.NewStatusCodeRangeGrouppClient(nil)
	require.NoError(t, err)
	resp, err := client.ErrorResponseStatusCode404(context.Background(), &statuscoderangegroup.StatusCodeRangeClientErrorResponseStatusCode404Options{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "404 Not Found")
	require.Equal(t, statuscoderangegroup.StatusCodeRangeClientErrorResponseStatusCode404Response{}, resp)
}

func TestStatusCodeRangeGroupClient_ErrorResponseStatusCodeInRange(t *testing.T) {
	client, err := statuscoderangegroup.NewStatusCodeRangeGrouppClient(nil)
	require.NoError(t, err)
	resp, err := client.ErrorResponseStatusCodeInRange(context.Background(), &statuscoderangegroup.StatusCodeRangeClientErrorResponseStatusCodeInRangeOptions{})
	require.Error(t, err)
	require.Equal(t, statuscoderangegroup.StatusCodeRangeClientErrorResponseStatusCodeInRangeResponse{}, resp)
}
