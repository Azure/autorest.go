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

type fakeCredential struct{}

func (mc fakeCredential) GetToken(context.Context, policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{Token: "https://security.microsoft.com/.default", ExpiresOn: time.Now().Add(time.Hour)}, nil
}

func TestValidKey(t *testing.T) {
	t.Skip("API key credential is not supported")
}

func TestValidToken(t *testing.T) {
	client, err := unionauthgroup.NewUnionClient("http://localhost:3000", &fakeCredential{}, &unionauthgroup.UnionClientOptions{
		ClientOptions: azcore.ClientOptions{
			InsecureAllowCredentialWithHTTP: true,
		},
	})
	require.NoError(t, err)
	resp, err := client.ValidToken(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
