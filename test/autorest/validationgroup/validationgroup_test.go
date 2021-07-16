// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package validationgroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newAutoRestValidationTestClient() *AutoRestValidationTestClient {
	return NewAutoRestValidationTestClient(NewDefaultConnection(nil), "")
}

func TestValidationGetWithConstantInPath(t *testing.T) {
	client := newAutoRestValidationTestClient()
	result, err := client.GetWithConstantInPath(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetWithConstantInPath: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestValidationPostWithConstantInBody(t *testing.T) {
	client := newAutoRestValidationTestClient()
	result, err := client.PostWithConstantInBody(context.Background(), &AutoRestValidationTestPostWithConstantInBodyOptions{Body: &Product{
		Child: &ChildProduct{
			ConstProperty: to.StringPtr("constant")},
		ConstString: to.StringPtr("constant"),
		ConstInt:    to.Int32Ptr(0),
		ConstChild: &ConstantProduct{
			ConstProperty:  to.StringPtr("constant"),
			ConstProperty2: to.StringPtr("constant2")}}})
	if err != nil {
		t.Fatalf("PostWithConstantInBody: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestValidationValidationOfBody(t *testing.T) {
	t.Skip("need to confirm if this test will remain in the testserver and what values it's expecting")
	client := newAutoRestValidationTestClient()
	result, err := client.ValidationOfBody(context.Background(), "123", 150, &AutoRestValidationTestValidationOfBodyOptions{
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
	if err != nil {
		t.Fatalf("ValidationOfBody: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
