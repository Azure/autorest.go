// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipleservicesgroup_test

import (
	"context"
	"multipleservicesgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceAOperationsClient_OpA(t *testing.T) {
	client, err := multipleservicesgroup.NewServiceAClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewServiceAOperationsClient().OpA(context.Background(), nil)
	require.NoError(t, err)
}

func TestServiceASubNamespaceClient_SubOpA(t *testing.T) {
	client, err := multipleservicesgroup.NewServiceAClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewServiceASubNamespaceClient().SubOpA(context.Background(), nil)
	require.NoError(t, err)
}

func TestServiceBOperationsClient_OpB(t *testing.T) {
	client, err := multipleservicesgroup.NewServiceBClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewServiceBOperationsClient().OpB(context.Background(), nil)
	require.NoError(t, err)
}

func TestServiceBSubNamespaceClient_SubOpB(t *testing.T) {
	client, err := multipleservicesgroup.NewServiceBClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewServiceBSubNamespaceClient().SubOpB(context.Background(), nil)
	require.NoError(t, err)
}
