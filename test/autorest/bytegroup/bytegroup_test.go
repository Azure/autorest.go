// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegrouptest

import (
	"bytes"
	"context"
	"generatortests/autorest/generated/bytegroup"
	"testing"
)

func TestGetEmpty(t *testing.T) {
	client, err := bytegroup.NewByteClient("http://localhost:3000", nil)
	if err != nil {
		t.Fatalf("failed to create byte client: %v", err)
	}
	array, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	if !bytes.Equal(*array.Value, nil) {
		t.Fatalf("expected empty array, got %+v", *array.Value)
	}
}
