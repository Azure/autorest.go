// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package validationgrouptest

import (
	"context"
	"generatortests/autorest/generated/validationgroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func TestValidationGetWithConstantInPath(t *testing.T) {
	client := validationgroup.NewDefaultClient(nil).AutoRestValidationTestOperations("")
	result, err := client.GetWithConstantInPath(context.Background())
	if err != nil {
		t.Fatalf("GetWithConstantInPath: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestValidationPostWithConstantInBody(t *testing.T) {
	client := validationgroup.NewDefaultClient(nil).AutoRestValidationTestOperations("")
	result, err := client.PostWithConstantInBody(context.Background(), &validationgroup.AutoRestValidationTestPostWithConstantInBodyOptions{Body: &validationgroup.Product{
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
	t.Skip("need to confirm if this test will remain in the testserver and what values it's expecting")
	client := validationgroup.NewDefaultClient(nil).AutoRestValidationTestOperations("")
	result, err := client.ValidationOfBody(context.Background(), "123", 150, &validationgroup.AutoRestValidationTestValidationOfBodyOptions{
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
