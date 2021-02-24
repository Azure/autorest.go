// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func TestGetAllWithValues(t *testing.T) {
	client := NewDefaultConnection(nil)
	grp := NewPathItemsClient(client, "globalStringPath", to.StringPtr("globalStringQuery"))
	result, err := grp.GetAllWithValues(context.Background(), "pathItemStringPath", "localStringPath", &PathItemsGetAllWithValuesOptions{
		LocalStringQuery:    to.StringPtr("localStringQuery"),
		PathItemStringQuery: to.StringPtr("pathItemStringQuery"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestGetGlobalAndLocalQueryNull(t *testing.T) {
	client := NewDefaultConnection(nil)
	grp := NewPathItemsClient(client, "globalStringPath", nil)
	result, err := grp.GetGlobalAndLocalQueryNull(context.Background(), "pathItemStringPath", "localStringPath", &PathItemsGetGlobalAndLocalQueryNullOptions{
		PathItemStringQuery: to.StringPtr("pathItemStringQuery"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestGetGlobalQueryNull(t *testing.T) {
	client := NewDefaultConnection(nil)
	grp := NewPathItemsClient(client, "globalStringPath", nil)
	result, err := grp.GetGlobalQueryNull(context.Background(), "pathItemStringPath", "localStringPath", &PathItemsGetGlobalQueryNullOptions{
		LocalStringQuery:    to.StringPtr("localStringQuery"),
		PathItemStringQuery: to.StringPtr("pathItemStringQuery"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestGetLocalPathItemQueryNull(t *testing.T) {
	client := NewDefaultConnection(nil)
	grp := NewPathItemsClient(client, "globalStringPath", to.StringPtr("globalStringQuery"))
	result, err := grp.GetLocalPathItemQueryNull(context.Background(), "pathItemStringPath", "localStringPath", nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
