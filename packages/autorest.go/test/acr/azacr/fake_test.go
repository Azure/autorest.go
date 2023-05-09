//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azacr_test

import (
	"azacr"
	"azacr/fake"
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var authSvr = fake.AuthenticationServer{
	ExchangeAADAccessTokenForAcrRefreshToken: func(ctx context.Context, grantType azacr.PostContentSchemaGrantType, service string, options *azacr.AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenOptions) (resp azfake.Responder[azacr.AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenResponse], err azfake.ErrorResponder) {
		testingT := ctx.Value(ctxKeyTestingT{}).(*testing.T)
		require.Equal(testingT, azacr.PostContentSchemaGrantTypeAccessToken, grantType)
		require.Equal(testingT, "fake_service", service)
		require.NotNil(testingT, options)
		require.Equal(testingT, "access_token", *options.AccessToken)
		require.Equal(testingT, "refresh_token", *options.RefreshToken)
		require.Equal(testingT, "tenant", *options.Tenant)
		resp.SetResponse(http.StatusOK, azacr.AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenResponse{
			RefreshToken: azacr.RefreshToken{
				RefreshToken: to.Ptr("boom"),
			},
		}, nil)
		return
	},
}

type ctxKeyTestingT struct{}

func TestFakeExchangeAADAccessTokenForAcrRefreshToken(t *testing.T) {
	client, err := azacr.NewAuthenticationClient("https://contoso.com/fake/thing", &azacr.AuthenticationClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewAuthenticationServerTransport(&authSvr),
		},
	})
	require.NoError(t, err)

	ctx := context.WithValue(context.Background(), ctxKeyTestingT{}, t)
	resp, err := client.ExchangeAADAccessTokenForAcrRefreshToken(ctx, azacr.PostContentSchemaGrantTypeAccessToken, "fake_service", &azacr.AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenOptions{
		AccessToken:  to.Ptr("access_token"),
		RefreshToken: to.Ptr("refresh_token"),
		Tenant:       to.Ptr("tenant"),
	})

	require.NoError(t, err)
	require.Equal(t, "boom", *resp.RefreshToken.RefreshToken)
}
