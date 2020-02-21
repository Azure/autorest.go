// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgrouptest

import (
	"context"
	"generatortests/autorest/generated/urlgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getPathItemsClient(t *testing.T) urlgroup.PathItemsOperations {
	client, err := urlgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create enum client: %v", err)
	}
	return client.PathItemsOperations("globalStringPath", "globalStringQuery")
}

func TestGetAllWithValues(t *testing.T) {
	client := getPathItemsClient(t)
	result, err := client.GetAllWithValues(context.Background(), "pathItemStringPath", "pathItemStringQuery", "localStringPath", "localStringQuery")
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestGetGlobalAndLocalQueryNull(t *testing.T) {
	t.Skip("need optional string params")
	client := getPathItemsClient(t)
	result, err := client.GetGlobalAndLocalQueryNull(context.Background(), "pathItemStringPath", "pathItemStringQuery", "localStringPath", "")
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestGetGlobalQueryNull(t *testing.T) {
	t.Skip("need optional string params")
	client := getPathItemsClient(t)
	result, err := client.GetGlobalQueryNull(context.Background(), "pathItemStringPath", "pathItemStringQuery", "localStringPath", "localStringQuery")
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestGetLocalPathItemQueryNull(t *testing.T) {
	t.Skip("need optional string params")
	client := getPathItemsClient(t)
	result, err := client.GetLocalPathItemQueryNull(context.Background(), "pathItemStringPath", "", "localStringPath", "")
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}
