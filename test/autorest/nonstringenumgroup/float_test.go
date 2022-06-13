// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonstringenumgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newFloatClient() *FloatClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewFloatClient(pl)
}

// Get - Get a float enum
func TestFloatGet(t *testing.T) {
	client := newFloatClient()
	result, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr(FloatEnumFourHundredTwentyNine1)); r != "" {
		t.Fatal(r)
	}
}

// Put - Put a float enum
func TestFloatPut(t *testing.T) {
	client := newFloatClient()
	result, err := client.Put(context.Background(), FloatEnumTwoHundred4, nil)
	require.NoError(t, err)
	if *result.Value != "Nice job posting a float enum" {
		t.Fatalf("unexpected value %s", *result.Value)
	}
}
