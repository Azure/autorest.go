// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package additionalpropsgroup_test

import (
	"context"
	"generatortests"
	"generatortests/additionalpropsgroup"
	"generatortests/additionalpropsgroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestFakeCreateAPInProperties(t *testing.T) {
	server := fake.PetsServer{
		CreateAPInProperties: func(ctx context.Context, createParameters additionalpropsgroup.PetAPInProperties, options *additionalpropsgroup.PetsClientCreateAPInPropertiesOptions) (resp azfake.Responder[additionalpropsgroup.PetsClientCreateAPInPropertiesResponse], errResp azfake.ErrorResponder) {
			require.Nil(t, options)
			require.Nil(t, createParameters.Status) // read-only so not sent
			createParameters.Status = to.Ptr(false)
			resp.SetResponse(http.StatusOK, additionalpropsgroup.PetsClientCreateAPInPropertiesResponse{createParameters}, nil)
			return
		},
	}
	client, err := additionalpropsgroup.NewPetsClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewPetsServerTransport(&server),
	})
	require.NoError(t, err)
	result, err := client.CreateAPInProperties(context.Background(), additionalpropsgroup.PetAPInProperties{ID: to.Ptr[int32](4),
		Name: to.Ptr("Bunny"),
		AdditionalProperties: map[string]*float32{
			"height":   to.Ptr[float32](5.61),
			"weight":   to.Ptr[float32](599),
			"footsize": to.Ptr[float32](11.5),
		}}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.PetAPInProperties, additionalpropsgroup.PetAPInProperties{
		ID:     to.Ptr[int32](4),
		Name:   to.Ptr("Bunny"),
		Status: to.Ptr(false),
		AdditionalProperties: map[string]*float32{
			"height":   to.Ptr[float32](5.61),
			"weight":   to.Ptr[float32](599),
			"footsize": to.Ptr[float32](11.5),
		},
	}); r != "" {
		t.Fatal(r)
	}
}
