// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package oauth2group_test

import (
	"context"
	"oauth2group"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOauth2GroupClient_Invalid(t *testing.T) {
	client, err := oauth2group.NewOauth2groupClient(nil)
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), &oauth2group.OAuth2ClientInvalidOptions{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid-grant")
	require.Zero(t, resp)
}

func TestOauth2GroupClient_Valid(t *testing.T) {
	client, err := oauth2group.NewOauth2groupClient(nil)
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), &oauth2group.OAuth2ClientValidOptions{})
	require.Error(t, err)
	require.Zero(t, resp)
}
