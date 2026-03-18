// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup_test

import (
	"context"
	"net/http"
	"testing"
	"xmlgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func TestXMLXMLErrorValueClient_Get(t *testing.T) {
	client, err := xmlgroup.NewXMLClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewXMLXMLErrorValueClient().Get(context.Background(), nil)
	var respErr *azcore.ResponseError
	require.Error(t, err)
	require.ErrorAs(t, err, &respErr)
	require.EqualValues(t, http.StatusBadRequest, respErr.StatusCode)
	require.EqualValues(t, "400", respErr.ErrorCode)
	require.Zero(t, resp)
}
