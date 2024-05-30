// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package unionauthgroup_test

import (
	"context"
	"testing"
	"time"
	"unionauthgroup"

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

func TestUnionClient_ValidKey(t *testing.T) {
	client, err := unionauthgroup.NewUnionClientWithKeyCredential(azcore.NewKeyCredential("valid-key"), nil)
	require.NoError(t, err)
	resp, err := client.ValidKey(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestUnionClient_ValidToken(t *testing.T) {
	client, err := unionauthgroup.NewUnionClient(&FakeToken{
		TokenValue: "https://security.microsoft.com/.default",
	}, nil)
	require.NoError(t, err)
	resp, err := client.ValidToken(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
