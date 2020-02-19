// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headergrouptest

import (
	"context"
	"generatortests/autorest/generated/headergroup"
	"net/http"
	"reflect"
	"testing"
)

func getHeaderClient(t *testing.T) headergroup.HeaderOperations {
	client, err := headergroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create header client: %v", err)
	}
	return client.HeaderOperations()
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
}

// func TestHeaderCustomRequestID(t *testing.T) {
// 	client := getHeaderClient(t)
// 	result, err := client.CustomRequestID(context.Background())
// 	if err != nil {
// 		t.Fatalf("CustomRequestID: %v", err)
// 	}
// 	expected := &headergroup.HeaderCustomRequestIDResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

func TestHeaderParamBool(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ParamBool(context.Background(), "false", false)
	if err != nil {
		t.Fatalf("ParamBool: %v", err)
	}
	expected := &headergroup.HeaderParamBoolResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
