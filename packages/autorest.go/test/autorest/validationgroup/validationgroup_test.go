// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package validationgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newAutoRestValidationTestClient(t *testing.T) *AutoRestValidationTestClient {
	client, err := NewAutoRestValidationTestClient(generatortests.Host, "", &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewAutoRestValidationTestClient(endpoint string, subscriptionID string, options *azcore.ClientOptions) (*AutoRestValidationTestClient, error) {
	client, err := azcore.NewClient("validationgroup.AutoRestValidationTestClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &AutoRestValidationTestClient{
		internal:       client,
		endpoint:       endpoint,
		subscriptionID: subscriptionID,
	}, nil
}

func TestValidationGetWithConstantInPath(t *testing.T) {
	client := newAutoRestValidationTestClient(t)
	result, err := client.GetWithConstantInPath(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestValidationPostWithConstantInBody(t *testing.T) {
	client := newAutoRestValidationTestClient(t)
	product := Product{
		Child:      &ChildProduct{},
		ConstChild: &ConstantProduct{},
	}
	result, err := client.PostWithConstantInBody(context.Background(), &AutoRestValidationTestClientPostWithConstantInBodyOptions{Body: &product})
	require.NoError(t, err)
	require.Equal(t, Product{
		Child: &ChildProduct{
			ConstProperty: to.Ptr("constant"),
		},
		ConstChild: &ConstantProduct{
			ConstProperty:  to.Ptr("constant"),
			ConstProperty2: to.Ptr("constant2"),
		},
		ConstInt:          to.Ptr[int32](0),
		ConstString:       to.Ptr("constant"),
		ConstStringAsEnum: to.Ptr("constant_string_as_enum"),
	}, result.Product)
}

func TestValidationValidationOfBody(t *testing.T) {
	t.Skip("need to confirm if this test will remain in the testserver and what values it's expecting")
	client := newAutoRestValidationTestClient(t)
	result, err := client.ValidationOfBody(context.Background(), "123", 150, Product{
		DisplayNames: []*string{
			to.Ptr("displayname1"),
			to.Ptr("displayname2"),
			to.Ptr("displayname3"),
			to.Ptr("displayname4"),
			to.Ptr("displayname5"),
			to.Ptr("displayname6"),
			to.Ptr("displayname7")}}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
