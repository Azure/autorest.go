// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgrouptest

import (
	"context"
	"generatortests/autorest/generated/urlgroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func TestGetAllWithValues(t *testing.T) {
	client := urlgroup.NewDefaultClient(nil)
	grp := urlgroup.NewPathItemsClient(client, "globalStringPath", to.StringPtr("globalStringQuery"))
	result, err := grp.GetAllWithValues(context.Background(), "pathItemStringPath", "localStringPath", &urlgroup.PathItemsGetAllWithValuesOptions{
		LocalStringQuery:    to.StringPtr("localStringQuery"),
		PathItemStringQuery: to.StringPtr("pathItemStringQuery"),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestGetGlobalAndLocalQueryNull(t *testing.T) {
	client := urlgroup.NewDefaultClient(nil)
	grp := urlgroup.NewPathItemsClient(client, "globalStringPath", nil)
	result, err := grp.GetGlobalAndLocalQueryNull(context.Background(), "pathItemStringPath", "localStringPath", &urlgroup.PathItemsGetGlobalAndLocalQueryNullOptions{
		PathItemStringQuery: to.StringPtr("pathItemStringQuery"),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestGetGlobalQueryNull(t *testing.T) {
	client := urlgroup.NewDefaultClient(nil)
	grp := urlgroup.NewPathItemsClient(client, "globalStringPath", nil)
	result, err := grp.GetGlobalQueryNull(context.Background(), "pathItemStringPath", "localStringPath", &urlgroup.PathItemsGetGlobalQueryNullOptions{
		LocalStringQuery:    to.StringPtr("localStringQuery"),
		PathItemStringQuery: to.StringPtr("pathItemStringQuery"),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestGetLocalPathItemQueryNull(t *testing.T) {
	client := urlgroup.NewDefaultClient(nil)
	grp := urlgroup.NewPathItemsClient(client, "globalStringPath", to.StringPtr("globalStringQuery"))
	result, err := grp.GetLocalPathItemQueryNull(context.Background(), "pathItemStringPath", "localStringPath", nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
