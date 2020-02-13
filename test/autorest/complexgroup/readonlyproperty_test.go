// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"net/http"
	"testing"
)

func getReadonlypropertyOperations(t *testing.T) complexgroup.ReadonlypropertyOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.ReadonlypropertyOperations()
}

func TestReadonlypropertyGetValid(t *testing.T) {
	client := getReadonlypropertyOperations(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	id, size := "1234", int32(2)
	expected := &complexgroup.ReadonlypropertyGetValidResponse{
		StatusCode:  http.StatusOK,
		ReadonlyObj: &complexgroup.ReadonlyObj{ID: &id, Size: &size},
	}
	deepEqualOrFatal(t, result, expected)
}

// func TestReadonlypropertyPutValid(t *testing.T) {
// 	client := getReadonlypropertyOperations(t)
// 	id, size := "1234", int32(2)
// 	result, err := client.PutValid(context.Background(), complexgroup.ReadonlyObj{ID: &id, Size: &size})
// 	if err != nil {
// 		t.Fatalf("PutValid: %v", err)
// 	}
// 	expected := &complexgroup.ReadonlypropertyPutValidResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }
