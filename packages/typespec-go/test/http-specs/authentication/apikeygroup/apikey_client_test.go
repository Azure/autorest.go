// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apikeygroup_test

import (
	"apikeygroup"
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/require"
)

func TestApiKeyClient_Invalid(t *testing.T) {
	client, err := apikeygroup.NewApiKeyClient(nil)
	require.NoError(t, err)
	contextWithAPIKey := context.Background()
	headers := http.Header{}
	headers.Set("x-ms-api-key", "invalid-key")
	ctx := policy.WithHTTPHeader(contextWithAPIKey, headers)
	resp, err := client.Invalid(ctx, &apikeygroup.APIKeyClientInvalidOptions{})
	require.Error(t, err)
	require.Zero(t, resp)
}

func TestApiKeyClient_OutputToInputOutput(t *testing.T) {
	client, err := apikeygroup.NewApiKeyClient(nil)
	require.NoError(t, err)
	contextWithAPIKey := context.Background()
	headers := http.Header{}
	headers.Set("x-ms-api-key", "valid-key")
	ctx := policy.WithHTTPHeader(contextWithAPIKey, headers)
	resp, err := client.Valid(ctx, &apikeygroup.APIKeyClientValidOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}
