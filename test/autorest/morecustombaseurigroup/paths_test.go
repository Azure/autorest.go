// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package morecustombaseurigrouptest

import (
	"context"
	"generatortests/autorest/generated/morecustombaseurigroup"
	"net/http"
	"reflect"
	"testing"
)

func getMoreCustomBaseURIClient(t *testing.T) morecustombaseurigroup.PathsOperations {
	client, err := morecustombaseurigroup.NewClient("http://localhost:3000", nil)
	if err != nil {
		t.Fatalf("failed to create more custom base URL client: %v", err)
	}
	return client.PathsOperations()
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
}

func TestGetEmpty(t *testing.T) {
	client := getMoreCustomBaseURIClient(t)
	result, err := client.GetEmpty(context.Background(), "", "", "key1", "v1", "test12")
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	expected := &morecustombaseurigroup.PathsGetEmptyResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
