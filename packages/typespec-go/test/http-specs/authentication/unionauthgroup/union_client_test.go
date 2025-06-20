// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package unionauthgroup_test

import (
	"context"
	"net/http"
	"testing"
	"unionauthgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/require"
)

func TestUnionAuthGroupClient_ValidKey(t *testing.T) {
	client, err := unionauthgroup.NewunionauthgroupClient(nil)
	require.NoError(t, err)
	ctxInit := context.Background()
	headers := http.Header{}
	headers.Set("x-ms-api-key", "valid-key")
	ctx := policy.WithHTTPHeader(ctxInit, headers)
	resp, err := client.ValidKey(ctx, &unionauthgroup.UnionClientValidKeyOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestUnionAuthGroupClient_ValidToken(t *testing.T) {
	client, err := unionauthgroup.NewunionauthgroupClient(nil)
	require.NoError(t, err)
	ctxInit := context.Background()
	headers := http.Header{}
	headers.Set("authorization", "Bearer https://security.microsoft.com/.default")
	ctx := policy.WithHTTPHeader(ctxInit, headers)
	resp, err := client.ValidToken(ctx, &unionauthgroup.UnionClientValidTokenOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}
