// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package durationgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newDurationClient() *DurationClient {
	return NewDurationClient(nil)
}

func TestGetInvalid(t *testing.T) {
	t.Skip("this does not apply to us meanwhile we do not parse durations")
	client := newDurationClient()
	_, err := client.GetInvalid(context.Background(), nil)
	require.Error(t, err)
}

func TestGetNull(t *testing.T) {
	client := newDurationClient()
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	var s *string
	if r := cmp.Diff(result.Value, s); r != "" {
		t.Fatal(r)
	}
}

func TestGetPositiveDuration(t *testing.T) {
	client := newDurationClient()
	result, err := client.GetPositiveDuration(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.StringPtr("P3Y6M4DT12H30M5S")); r != "" {
		t.Fatal(r)
	}
}

func TestPutPositiveDuration(t *testing.T) {
	client := newDurationClient()
	result, err := client.PutPositiveDuration(context.Background(), "P123DT22H14M12.011S", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
