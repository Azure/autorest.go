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
	_, err = client.ErrorResponseStatusCode404(context.Background(), &statuscoderangegroup.StatusCodeRangeClientErrorResponseStatusCode404Options{})
	require.Error(t, err)
}

func TestStatusCodeRangeGroupClient_ErrorResponseStatusCodeInRange(t *testing.T) {
	client, err := statuscoderangegroup.NewStatusCodeRangeGrouppClient(nil)
	require.NoError(t, err)
	_, err = client.ErrorResponseStatusCodeInRange(context.Background(), &statuscoderangegroup.StatusCodeRangeClientErrorResponseStatusCodeInRangeOptions{})
	require.Error(t, err)

}
