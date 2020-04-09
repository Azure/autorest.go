// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package numbergrouptest

import (
	"context"
	"generatortests/autorest/generated/validationgroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func getValidationOperations(t *testing.T) validationgroup.Operations {
	client, err := validationgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create validation client: %v", err)
	}
	return client.Operations("")
}

func TestValidationGetWithConstantInPath(t *testing.T) {
	client := getValidationOperations(t)
	result, err := client.GetWithConstantInPath(context.Background())
	if err != nil {
		t.Fatalf("GetWithConstantInPath: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestValidationPostWithConstantInBody(t *testing.T) {
	client := getValidationOperations(t)
	result, err := client.PostWithConstantInBody(context.Background(), &validationgroup.OperationsPostWithConstantInBodyOptions{Body: &validationgroup.Product{
		Child: &validationgroup.ChildProduct{
			ConstProperty: to.StringPtr("constant")},
		ConstString: to.StringPtr("constant"),
		ConstInt:    to.Int32Ptr(0),
		ConstChild: &validationgroup.ConstantProduct{
			ConstProperty:  to.StringPtr("constant"),
			ConstProperty2: to.StringPtr("constant2")}}})
	if err != nil {
		t.Fatalf("PostWithConstantInBody: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestValidationValidationOfBody(t *testing.T) {
	t.Skip("need to confirm if this test will remain in the testserver")
	client := getValidationOperations(t)
	result, err := client.ValidationOfBody(context.Background(), "123", 150, &validationgroup.OperationsValidationOfBodyOptions{
		Body: &validationgroup.Product{
			DisplayNames: &[]string{
				"displayname1",
				"displayname2",
				"displayname3",
				"displayname4",
				"displayname5",
				"displayname6",
				"displayname7"}},
	})
	if err != nil {
		t.Fatalf("ValidationOfBody: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}
