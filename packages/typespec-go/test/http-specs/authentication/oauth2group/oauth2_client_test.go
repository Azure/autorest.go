// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package oauth2group_test

import (
	"context"
	"net/http"
	"oauth2group"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/require"
)

func TestOauth2GroupClient_Invalid(t *testing.T) {
	client, err := oauth2group.NewOauth2groupClient(nil)
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), &oauth2group.OAuth2ClientInvalidOptions{})
	require.Error(t, err)
	require.Zero(t, resp)
}

func TestOauth2GroupClient_Valid(t *testing.T) {
	client, err := oauth2group.NewOauth2groupClient(nil)
	require.NoError(t, err)
	ctxInit := context.Background()
	headers := http.Header{}
	headers.Set("authorization", "Bearer https://security.microsoft.com/.default")
	ctx := policy.WithHTTPHeader(ctxInit, headers)
	resp, err := client.Valid(ctx, &oauth2group.OAuth2ClientValidOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}
