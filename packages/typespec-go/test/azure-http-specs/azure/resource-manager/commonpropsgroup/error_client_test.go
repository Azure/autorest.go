// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package commonpropsgroup_test

import (
	"commonpropsgroup"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestErrorClient_CreateForUserDefinedError(t *testing.T) {
	resp, err := clientFactory.NewErrorClient().CreateForUserDefinedError(ctx, resourceGroupExpected, "confidential", commonpropsgroup.ConfidentialResource{
		Location: to.Ptr("eastus"),
		Properties: &commonpropsgroup.ConfidentialResourceProperties{
			Username: to.Ptr("00"),
		},
	}, nil)
	require.Error(t, err)
	require.NotNil(t, resp)
	var respErr *azcore.ResponseError
	require.ErrorAs(t, err, &respErr)
	require.EqualValues(t, http.StatusBadRequest, respErr.StatusCode)
	require.Zero(t, resp)
	bodyBytes := make([]byte, respErr.RawResponse.ContentLength)
	_, readErr := respErr.RawResponse.Body.Read(bodyBytes)
	require.NoError(t, readErr)
	bodyString := string(bodyBytes)
	require.Contains(t, bodyString, "Username should not contain only numbers.")
}

func TestErrorClient_GetForPredefinedError(t *testing.T) {
	resp, err := clientFactory.NewErrorClient().GetForPredefinedError(ctx, resourceGroupExpected, "confidential", nil)
	require.Error(t, err)
	require.NotNil(t, resp)
	var respErr *azcore.ResponseError
	require.ErrorAs(t, err, &respErr)
	require.EqualValues(t, http.StatusNotFound, respErr.StatusCode)
	require.Zero(t, resp)
	bodyBytes := make([]byte, respErr.RawResponse.ContentLength)
	_, readErr := respErr.RawResponse.Body.Read(bodyBytes)
	require.NoError(t, readErr)
	bodyString := string(bodyBytes)
	require.Contains(t, bodyString, "The Resource 'Azure.ResourceManager.CommonProperties/confidentialResources/confidential' under resource group 'test-rg' was not found.")
}
