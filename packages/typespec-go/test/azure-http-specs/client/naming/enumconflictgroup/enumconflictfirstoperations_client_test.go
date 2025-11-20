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

func TestEnumConflictFirstOperationsClient_First(t *testing.T) {
	enumClient, err := enumconflictgroup.NewEnumConflictClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)

	body := enumconflictgroup.FirstModel{
		Status: to.Ptr(enumconflictgroup.StatusActive),
		Name:   to.Ptr("test"),
	}

	resp, err := enumClient.NewEnumConflictFirstOperationsClient().First(context.Background(), body, nil)
	require.NoError(t, err)
	require.Equal(t, body.Status, resp.Status)
	require.Equal(t, body.Name, resp.Name)
}
