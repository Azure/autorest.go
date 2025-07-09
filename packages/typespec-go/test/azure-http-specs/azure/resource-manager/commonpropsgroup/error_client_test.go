// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package commonpropsgroup_test

import (
	"commonpropsgroup"
	"encoding/json"
	"net/http"
	"regexp"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
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
	re := regexp.MustCompile(`(?s)\{.*\}`)
	jsonStr := re.FindString(err.Error())
	var errorResp ErrorResp
	err = json.Unmarshal([]byte(jsonStr), &errorResp)
	require.NoError(t, err)
	require.Contains(t, errorResp.Message, "Username should not contain only numbers")
	require.Equal(t, "general", errorResp.InnerError.ExceptionType)
	require.Equal(t, "BadRequest", errorResp.Code)
}

func TestErrorClient_GetForPredefinedError(t *testing.T) {
	resp, err := clientFactory.NewErrorClient().GetForPredefinedError(ctx, resourceGroupExpected, "confidential", nil)
	require.Error(t, err)
	require.NotNil(t, resp)
	var respErr *azcore.ResponseError
	require.ErrorAs(t, err, &respErr)
	require.EqualValues(t, http.StatusNotFound, respErr.StatusCode)
	require.Zero(t, resp)
}
