// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package durationgroup

import (
	"context"
	"generatortests/autorest/generated/durationgroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func getDurationOperations(t *testing.T) durationgroup.DurationOperations {
	client, err := durationgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client.DurationOperations()
}

func TestGetInvalid(t *testing.T) {
	t.Skip("this does not apply to us meanwhile we do not parse durations")
	client := getDurationOperations(t)
	_, err := client.GetInvalid(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestGetNull(t *testing.T) {
	client := getDurationOperations(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var s *string
	helpers.DeepEqualOrFatal(t, result.Value, s)
}

func TestGetPositiveDuration(t *testing.T) {
	client := getDurationOperations(t)
	result, err := client.GetPositiveDuration(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("P3Y6M4DT12H30M5S"))
}

func TestPutPositiveDuration(t *testing.T) {
	client := getDurationOperations(t)
	result, err := client.PutPositiveDuration(context.Background(), "P123DT22H14M12.011S")
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
