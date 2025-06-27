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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

type fakeCredential struct{}

func (mc fakeCredential) GetToken(context.Context, policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{Token: "https://security.microsoft.com/.default", ExpiresOn: time.Now().Add(time.Hour)}, nil
}

func TestUnionClientValidKey(t *testing.T) {
	cred := azcore.NewKeyCredential("valid-key")
	authPolicy := runtime.NewKeyCredentialPolicy(cred, "x-ms-api-key", &runtime.KeyCredentialPolicyOptions{
		InsecureAllowCredentialWithHTTP: true,
	})
	client, err := unionauthgroup.NewUnionClient(&azcore.ClientOptions{
		PerCallPolicies: []policy.Policy{authPolicy},
	})
	require.NoError(t, err)
	_, err = client.ValidKey(context.Background(), nil)
	require.NoError(t, err)
}

func TestUnionClientValidToken(t *testing.T) {
	authPolicy := runtime.NewBearerTokenPolicy(fakeCredential{}, []string{}, &policy.BearerTokenOptions{
		InsecureAllowCredentialWithHTTP: true,
	})
	client, err := unionauthgroup.NewUnionClient(&azcore.ClientOptions{
		PerCallPolicies: []policy.Policy{authPolicy},
	})
	require.NoError(t, err)
	_, err = client.ValidToken(context.Background(), nil)
	require.NoError(t, err)
}
