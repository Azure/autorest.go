// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegrouptest

import (
	"context"
	"generatortests/autorest/generated/bytegroup"
	"net/http"
	"reflect"
	"testing"
)

func TestGetEmpty(t *testing.T) {
	client, err := bytegroup.NewByteClient(nil)
	if err != nil {
		t.Fatalf("failed to create byte client: %v", err)
	}
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	expected := &bytegroup.GetEmptyResponse{
		StatusCode: http.StatusOK,
		Value:      &[]byte{},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected empty array, got %+v", result)
	}
}
