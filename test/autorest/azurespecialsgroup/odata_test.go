// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func newODataClient() *ODataClient {
	return NewODataClient(nil)
}

// GetWithFilter - Specify filter parameter with value '$filter=id gt 5 and name eq 'foo'&$orderby=id&$top=10'
func TestGetWithFilter(t *testing.T) {
	client := newODataClient()
	result, err := client.GetWithFilter(context.Background(), &ODataGetWithFilterOptions{
		Filter:  to.StringPtr("id gt 5 and name eq 'foo'"),
		Orderby: to.StringPtr("id"),
		Top:     to.Int32Ptr(10),
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
