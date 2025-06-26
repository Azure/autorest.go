// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package customgroup_test

import (
	"context"
	"customgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func TestCustomClientValid(t *testing.T) {
	cred := azcore.NewKeyCredential("SharedAccessKey valid-key")
	authPolicy := runtime.NewKeyCredentialPolicy(cred, "Authorization", &runtime.KeyCredentialPolicyOptions{
		InsecureAllowCredentialWithHTTP: true,
	})
	client, err := customgroup.NewCustomClient(&azcore.ClientOptions{
		PerCallPolicies: []policy.Policy{authPolicy},
	})
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestCustomClientInvalid(t *testing.T) {
	cred := azcore.NewKeyCredential("SharedAccessKey invalid-key")
	authPolicy := runtime.NewKeyCredentialPolicy(cred, "Authorization", &runtime.KeyCredentialPolicyOptions{
		InsecureAllowCredentialWithHTTP: true,
	})
	client, err := customgroup.NewCustomClient(&azcore.ClientOptions{
		PerCallPolicies: []policy.Policy{authPolicy},
	})
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
	require.ErrorContains(t, err, "403")
	require.ErrorContains(t, err, "invalid-api-key")
}
