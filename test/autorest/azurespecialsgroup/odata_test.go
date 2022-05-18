// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newODataClient() *ODataClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewODataClient(pl)
}

// GetWithFilter - Specify filter parameter with value '$filter=id gt 5 and name eq 'foo'&$orderby=id&$top=10'
func TestGetWithFilter(t *testing.T) {
	client := newODataClient()
	result, err := client.GetWithFilter(context.Background(), &ODataClientGetWithFilterOptions{
		Filter:  to.Ptr("id gt 5 and name eq 'foo'"),
		Orderby: to.Ptr("id"),
		Top:     to.Ptr[int32](10),
	})
	require.NoError(t, err)
	require.Zero(t, result)
}
