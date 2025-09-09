// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package oauth2group_test

import (
	"context"
	"net/http"
	"oauth2group"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/require"
)

type fakeCredential struct{}

func (mc fakeCredential) GetToken(context.Context, policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{Token: "https://security.microsoft.com/.default", ExpiresOn: time.Now().Add(time.Hour)}, nil
}

func TestInvalid(t *testing.T) {
	client, err := oauth2group.NewOAuth2Client("http://localhost:3000", &fakeCredential{}, &oauth2group.OAuth2ClientOptions{
		ClientOptions: azcore.ClientOptions{
			InsecureAllowCredentialWithHTTP: true,
		},
	})
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), nil)
	var respErr *azcore.ResponseError
	require.ErrorAs(t, err, &respErr)
	require.EqualValues(t, http.StatusForbidden, respErr.StatusCode)
	require.Zero(t, resp)
}

func TestValid(t *testing.T) {
	client, err := oauth2group.NewOAuth2Client("http://localhost:3000", &fakeCredential{}, &oauth2group.OAuth2ClientOptions{
		ClientOptions: azcore.ClientOptions{
			InsecureAllowCredentialWithHTTP: true,
		},
	})
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
