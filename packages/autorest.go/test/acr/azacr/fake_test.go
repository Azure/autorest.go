//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azacr_test

import (
	"azacr"
	"azacr/fake"
	"context"
	"io"
	"net"
	"net/http"
	"strings"
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
	client, err := azacr.NewAuthenticationClient("https://contoso.com/fake/thing", &azcore.ClientOptions{
		Transport: fake.NewAuthenticationServerTransport(&authSvr),
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

func TestFakeUpdateTagAttributes(t *testing.T) {
	const (
		theName      = "theName"
		theReference = ";encoded$#reference/"
	)
	props := azacr.TagWriteableProperties{
		CanList:  to.Ptr(false),
		CanWrite: to.Ptr(true),
	}
	server := fake.ContainerRegistryServer{
		UpdateTagAttributes: func(ctx context.Context, name, reference string, value azacr.TagWriteableProperties, options *azacr.ContainerRegistryClientUpdateTagAttributesOptions) (resp azfake.Responder[azacr.ContainerRegistryClientUpdateTagAttributesResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, theName, name)
			require.EqualValues(t, theReference, reference)
			require.EqualValues(t, props, value)
			resp.SetResponse(http.StatusOK, azacr.ContainerRegistryClientUpdateTagAttributesResponse{}, nil)
			return
		},
	}
	client, err := azacr.NewContainerRegistryClient("https://contoso.com/fake/thing", &azcore.ClientOptions{
		Transport: fake.NewContainerRegistryServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.UpdateTagAttributes(context.Background(), theName, theReference, props, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFakeAuthenticationServerInterceptor(t *testing.T) {
	server := fake.ContainerRegistryServer{
		CheckDockerV2Support: func(ctx context.Context, options *azacr.ContainerRegistryClientCheckDockerV2SupportOptions) (resp azfake.Responder[azacr.ContainerRegistryClientCheckDockerV2SupportResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, azacr.ContainerRegistryClientCheckDockerV2SupportResponse{}, nil)
			return
		},
	}
	client, err := azacr.NewContainerRegistryClient("https://contoso.com/fake/thing", &azcore.ClientOptions{
		Transport: fake.NewContainerRegistryServerTransport(&server),
	})
	require.NoError(t, err)

	// intercept to return ResponseError
	fake.SetContainerRegistryServerInterceptor(func(req *http.Request) (*http.Response, error, bool) {
		resp := &http.Response{
			Request:    req,
			Status:     "Fake ResponseError",
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(strings.NewReader(`{"error":{"code":"Intercepted"}}`)),
			Header:     http.Header{},
		}

		return resp, nil, true
	})
	_, err = client.CheckDockerV2Support(context.Background(), nil)
	var respErr *azcore.ResponseError
	require.ErrorAs(t, err, &respErr)
	require.EqualValues(t, "Intercepted", respErr.ErrorCode)

	// intercept to return error
	fake.SetContainerRegistryServerInterceptor(func(*http.Request) (*http.Response, error, bool) {
		return nil, net.ErrClosed, true
	})
	_, err = client.CheckDockerV2Support(context.Background(), nil)
	require.ErrorIs(t, err, net.ErrClosed)

	// no intercept
	fake.SetContainerRegistryServerInterceptor(func(*http.Request) (*http.Response, error, bool) {
		return nil, nil, false
	})
	_, err = client.CheckDockerV2Support(context.Background(), nil)
	require.NoError(t, err)

	// nil intercept
	fake.SetContainerRegistryServerInterceptor(nil)
	_, err = client.CheckDockerV2Support(context.Background(), nil)
	require.NoError(t, err)
}
