// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package enumconflictgroup_test

import (
	"context"
	"testing"

	"enumconflictgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestEnumConflictSecondOperationsClient_Second(t *testing.T) {
	enumClient, err := enumconflictgroup.NewEnumConflictClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)

	body := enumconflictgroup.SecondModel{
		Status:      to.Ptr(enumconflictgroup.SecondStatusRunning)?,
		Description: to.Ptr("test description"),
	}

	resp, err := enumClient.NewEnumConflictSecondOperationsClient().Second(context.Background(), body, nil)
	require.NoError(t, err)
	require.Equal(t, body.Status, resp.SecondModel.Status)
	require.Equal(t, body.Description, resp.SecondModel.Description)
}
