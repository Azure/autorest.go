// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package validationgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newAutoRestValidationTestClient() *AutoRestValidationTestClient {
	return NewAutoRestValidationTestClient("", nil)
}

func TestValidationGetWithConstantInPath(t *testing.T) {
	client := newAutoRestValidationTestClient()
	result, err := client.GetWithConstantInPath(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestValidationPostWithConstantInBody(t *testing.T) {
	client := newAutoRestValidationTestClient()
	product := Product{
		Child: &ChildProduct{
			ConstProperty: to.StringPtr("constant")},
		ConstString: to.StringPtr("constant"),
		ConstInt:    to.Int32Ptr(0),
		ConstChild: &ConstantProduct{
			ConstProperty:  to.StringPtr("constant"),
			ConstProperty2: to.StringPtr("constant2")}}
	result, err := client.PostWithConstantInBody(context.Background(), &AutoRestValidationTestClientPostWithConstantInBodyOptions{Body: &product})
	require.NoError(t, err)
	if r := cmp.Diff(product, result.Product); r != "" {
		t.Fatal(r)
	}
}

func TestValidationValidationOfBody(t *testing.T) {
	t.Skip("need to confirm if this test will remain in the testserver and what values it's expecting")
	client := newAutoRestValidationTestClient()
	result, err := client.ValidationOfBody(context.Background(), "123", 150, &AutoRestValidationTestClientValidationOfBodyOptions{
		Body: &Product{
			DisplayNames: []*string{
				to.StringPtr("displayname1"),
				to.StringPtr("displayname2"),
				to.StringPtr("displayname3"),
				to.StringPtr("displayname4"),
				to.StringPtr("displayname5"),
				to.StringPtr("displayname6"),
				to.StringPtr("displayname7")}},
	})
	require.NoError(t, err)
	require.Zero(t, result)
}
