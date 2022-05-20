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
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newHeaderClient() *HeaderClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewHeaderClient(pl)
}

// CustomNamedRequestID - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
func TestCustomNamedRequestID(t *testing.T) {
	client := newHeaderClient()
	result, err := client.CustomNamedRequestID(context.Background(), "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0", nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.FooRequestID, to.Ptr("123")); r != "" {
		t.Fatal(r)
	}
}

// CustomNamedRequestIDHead - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
func TestCustomNamedRequestIDHead(t *testing.T) {
	client := newHeaderClient()
	result, err := client.CustomNamedRequestIDHead(context.Background(), "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0", nil)
	require.NoError(t, err)
	if !result.Success {
		t.Fatal("expected success")
	}
	if r := cmp.Diff(result.FooRequestID, to.Ptr("123")); r != "" {
		t.Fatal(r)
	}
}

// CustomNamedRequestIDParamGrouping - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request, via a parameter group
func TestCustomNamedRequestIDParamGrouping(t *testing.T) {
	client := newHeaderClient()
	result, err := client.CustomNamedRequestIDParamGrouping(context.Background(), HeaderClientCustomNamedRequestIDParamGroupingParameters{
		FooClientRequestID: "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0",
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.FooRequestID, to.Ptr("123")); r != "" {
		t.Fatal(r)
	}
}
