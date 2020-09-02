// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgrouptest

import (
	"context"
	"generatortests/autorest/generated/azurespecialsgroup"
	"generatortests/helpers"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

// CustomNamedRequestID - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
func TestCustomNamedRequestID(t *testing.T) {
	client := azurespecialsgroup.NewDefaultClient(nil).HeaderOperations()
	result, err := client.CustomNamedRequestID(context.Background(), "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.FooRequestID, to.StringPtr("123"))
}

// CustomNamedRequestIDHead - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
func TestCustomNamedRequestIDHead(t *testing.T) {
	client := azurespecialsgroup.NewDefaultClient(nil).HeaderOperations()
	result, err := client.CustomNamedRequestIDHead(context.Background(), "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.FooRequestID, to.StringPtr("123"))
}

// CustomNamedRequestIDParamGrouping - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request, via a parameter group
func TestCustomNamedRequestIDParamGrouping(t *testing.T) {
	client := azurespecialsgroup.NewDefaultClient(nil).HeaderOperations()
	result, err := client.CustomNamedRequestIDParamGrouping(context.Background(), azurespecialsgroup.HeaderCustomNamedRequestIDParamGroupingParameters{
		FooClientRequestId: "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0",
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.FooRequestID, to.StringPtr("123"))
}
