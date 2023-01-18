// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newPathItemsClient(t *testing.T, globalStringPath string, globalStringQuery *string) *PathItemsClient {
	client, err := NewPathItemsClient(globalStringPath, globalStringQuery, nil)
	require.NoError(t, err)
	return client
}

func NewPathItemsClient(globalStringPath string, globalStringQuery *string, options *azcore.ClientOptions) (*PathItemsClient, error) {
	client, err := azcore.NewClient("urlgroup.PathItemsClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &PathItemsClient{
		internal:          client,
		globalStringPath:  globalStringPath,
		globalStringQuery: globalStringQuery,
	}, nil
}

func TestGetAllWithValues(t *testing.T) {
	grp := newPathItemsClient(t, "globalStringPath", to.Ptr("globalStringQuery"))
	result, err := grp.GetAllWithValues(context.Background(), "pathItemStringPath", "localStringPath", &PathItemsClientGetAllWithValuesOptions{
		LocalStringQuery:    to.Ptr("localStringQuery"),
		PathItemStringQuery: to.Ptr("pathItemStringQuery"),
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestGetGlobalAndLocalQueryNull(t *testing.T) {
	grp := newPathItemsClient(t, "globalStringPath", nil)
	result, err := grp.GetGlobalAndLocalQueryNull(context.Background(), "pathItemStringPath", "localStringPath", &PathItemsClientGetGlobalAndLocalQueryNullOptions{
		PathItemStringQuery: to.Ptr("pathItemStringQuery"),
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestGetGlobalQueryNull(t *testing.T) {
	grp := newPathItemsClient(t, "globalStringPath", nil)
	result, err := grp.GetGlobalQueryNull(context.Background(), "pathItemStringPath", "localStringPath", &PathItemsClientGetGlobalQueryNullOptions{
		LocalStringQuery:    to.Ptr("localStringQuery"),
		PathItemStringQuery: to.Ptr("pathItemStringQuery"),
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestGetLocalPathItemQueryNull(t *testing.T) {
	grp := newPathItemsClient(t, "globalStringPath", to.Ptr("globalStringQuery"))
	result, err := grp.GetLocalPathItemQueryNull(context.Background(), "pathItemStringPath", "localStringPath", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
