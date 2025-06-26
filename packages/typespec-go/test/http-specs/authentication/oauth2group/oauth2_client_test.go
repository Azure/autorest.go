// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package oauth2group

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

type fakeCredential struct{}

func (mc fakeCredential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{Token: "https://security.microsoft.com/.default", ExpiresOn: time.Now().Add(time.Hour)}, nil
}

func TestOAuth2ClientValid(t *testing.T) {
	authPolicy := runtime.NewBearerTokenPolicy(fakeCredential{}, []string{}, &policy.BearerTokenOptions{
		InsecureAllowCredentialWithHTTP: true,
	})
	client, err := NewOAuth2Client(&azcore.ClientOptions{
		PerCallPolicies: []policy.Policy{authPolicy},
	})
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOAuth2ClientInvalid(t *testing.T) {
	authPolicy := runtime.NewBearerTokenPolicy(fakeCredential{}, []string{}, &policy.BearerTokenOptions{
		InsecureAllowCredentialWithHTTP: true,
	})
	client, err := NewOAuth2Client(&azcore.ClientOptions{
		PerCallPolicies: []policy.Policy{authPolicy},
	})
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
	require.ErrorContains(t, err, "403")
	require.ErrorContains(t, err, "invalid")
}
