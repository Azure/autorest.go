// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package commonpropsgroup_test

import (
	"commonpropsgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

type ErrorResp struct {
	Code       string     `json:"code"`
	Message    string     `json:"message"`
	InnerError InnerError `json:"innererror"`
}

type InnerError struct {
	ExceptionType string `json:"exceptiontype"`
}

func TestErrorClient_CreateForUserDefinedError(t *testing.T) {
	resp, err := clientFactory.NewErrorClient().CreateForUserDefinedError(ctx, resourceGroupExpected, "confidential", commonpropsgroup.ConfidentialResource{
		Location: to.Ptr("eastus"),
		Properties: &commonpropsgroup.ConfidentialResourceProperties{
			Username: to.Ptr("00"),
		},
	}, nil)
	require.Error(t, err)
	require.NotNil(t, resp)
	require.Contains(t, err.Error(), "Username should not contain only numbers.")
	require.Contains(t, err.Error(), "\"exceptiontype\": \"general\"")
	require.Contains(t, err.Error(), "BadRequest")

}

func TestErrorClient_GetForPredefinedError(t *testing.T) {
	resp, err := clientFactory.NewErrorClient().GetForPredefinedError(ctx, resourceGroupExpected, "confidential", nil)
	require.Error(t, err)
	require.NotNil(t, resp)
	require.Contains(t, err.Error(), "The Resource 'Azure.ResourceManager.CommonProperties/confidentialResources/confidential' under resource group 'test-rg' was not found.")
	require.Contains(t, err.Error(), "ResourceNotFound")
}
