// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlmultigroup

import (
	"context"
	"generatortests"
	"net/url"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func newQueriesClient(t *testing.T) *QueriesClient {
	client, err := NewQueriesClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func TestArrayStringMultiEmpty(t *testing.T) {
	client := newQueriesClient(t)
	result, err := client.ArrayStringMultiEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestURLMultiArrayStringMultiEmpty(t *testing.T) {
	client := newQueriesClient(t)
	result, err := client.ArrayStringMultiEmpty(context.Background(), &QueriesClientArrayStringMultiEmptyOptions{
		ArrayQuery: []string{},
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestURLMultiArrayStringMultiNull(t *testing.T) {
	client := newQueriesClient(t)
	result, err := client.ArrayStringMultiNull(context.Background(), &QueriesClientArrayStringMultiNullOptions{
		ArrayQuery: nil,
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestURLMultiArrayStringMultiValid(t *testing.T) {
	t.Skip("Cannot set nil for string value in string slice")
	client := newQueriesClient(t)
	result, err := client.ArrayStringMultiValid(context.Background(), &QueriesClientArrayStringMultiValidOptions{
		ArrayQuery: []string{
			"ArrayQuery1",
			url.QueryEscape("begin!*'();:@ &=+$,/?#[]end"),
			"",
			""},
	})
	require.NoError(t, err)
	require.Zero(t, result)
}
