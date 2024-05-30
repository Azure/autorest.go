// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package oauth2group_test

import (
	"context"
	"oauth2group"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/require"
)

type FakeToken struct {
	TokenValue string
}

func (f *FakeToken) GetToken(context.Context, policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{
		ExpiresOn: time.Now().Add(time.Hour),
		Token:     f.TokenValue,
	}, nil
}

func TestOAuth2Client_Invalid(t *testing.T) {
	client, err := oauth2group.NewOAuth2Client(&FakeToken{
		TokenValue: "invalid-value",
	}, nil)
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

func TestOAuth2Client_Valid(t *testing.T) {
	client, err := oauth2group.NewOAuth2Client(&FakeToken{
		TokenValue: "https://security.microsoft.com/.default",
	}, nil)
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
