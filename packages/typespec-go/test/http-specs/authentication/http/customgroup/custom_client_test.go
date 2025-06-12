// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package customgroup_test

import (
	"context"
	"customgroup"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/require"
)

func TestCustomClient_Invalid_SharedAccessKey(t *testing.T) {
	client, err := customgroup.NewCustomClient(nil)
	require.NoError(t, err)
	ctxInit := context.Background()
	headers := http.Header{}
	headers.Set("authorization", "SharedAccessKey invalid-key")
	ctx := policy.WithHTTPHeader(ctxInit, headers)
	resp, err := client.Invalid(ctx, nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

func TestCustomClient_Valid_WithValidKey(t *testing.T) {
	client, err := customgroup.NewCustomClient(nil)
	require.NoError(t, err)
	ctxInit := context.Background()
	headers := http.Header{}
	headers.Set("authorization", "SharedAccessKey valid-key")
	ctx := policy.WithHTTPHeader(ctxInit, headers)
	resp, err := client.Valid(ctx, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
