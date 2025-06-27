// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apikeygroup_test

import (
	"apikeygroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func TestAPIKeyClientValid(t *testing.T) {
	cred := azcore.NewKeyCredential("valid-key")
	authPolicy := runtime.NewKeyCredentialPolicy(cred, "x-ms-api-key", &runtime.KeyCredentialPolicyOptions{
		InsecureAllowCredentialWithHTTP: true,
	})
	client, err := apikeygroup.NewAPIKeyClient(&azcore.ClientOptions{
		PerCallPolicies: []policy.Policy{authPolicy},
	})
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestAPIKeyClientInvalid(t *testing.T) {
	cred := azcore.NewKeyCredential("invalid-key")
	authPolicy := runtime.NewKeyCredentialPolicy(cred, "x-ms-api-key", &runtime.KeyCredentialPolicyOptions{
		InsecureAllowCredentialWithHTTP: true,
	})
	client, err := apikeygroup.NewAPIKeyClient(&azcore.ClientOptions{
		PerCallPolicies: []policy.Policy{authPolicy},
	})
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), nil)
	require.Zero(t, resp)
	require.ErrorContains(t, err, "403")
	require.ErrorContains(t, err, "invalid-api-key")
}
