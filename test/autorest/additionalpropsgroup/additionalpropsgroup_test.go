// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package additionalpropsgroup

import (
	"context"
	"encoding/base64"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newPetsClient() *PetsClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewPetsClient(pl)
}

// CreateAPInProperties - Create a Pet which contains more properties than what is defined.
func TestCreateAPInProperties(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateAPInProperties(context.Background(), PetAPInProperties{
		ID:   to.Ptr[int32](4),
		Name: to.Ptr("Bunny"),
		AdditionalProperties: map[string]*float32{
			"height":   to.Ptr[float32](5.61),
			"weight":   to.Ptr[float32](599),
			"footsize": to.Ptr[float32](11.5),
		},
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.PetAPInProperties, PetAPInProperties{
		ID:     to.Ptr[int32](4),
		Name:   to.Ptr("Bunny"),
		Status: to.Ptr(true),
		AdditionalProperties: map[string]*float32{
			"height":   to.Ptr[float32](5.61),
			"weight":   to.Ptr[float32](599),
			"footsize": to.Ptr[float32](11.5),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// CreateAPInPropertiesWithAPString - Create a Pet which contains more properties than what is defined.
func TestCreateAPInPropertiesWithAPString(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateAPInPropertiesWithAPString(context.Background(), PetAPInPropertiesWithAPString{
		ID:            to.Ptr[int32](5),
		Name:          to.Ptr("Funny"),
		ODataLocation: to.Ptr("westus"),
		AdditionalProperties: map[string]*string{
			"color": to.Ptr("red"),
			"city":  to.Ptr("Seattle"),
			"food":  to.Ptr("tikka masala"),
		},
		AdditionalProperties1: map[string]*float32{
			"height":   to.Ptr[float32](5.61),
			"weight":   to.Ptr[float32](599),
			"footsize": to.Ptr[float32](11.5),
		},
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.PetAPInPropertiesWithAPString, PetAPInPropertiesWithAPString{
		ID:            to.Ptr[int32](5),
		Name:          to.Ptr("Funny"),
		ODataLocation: to.Ptr("westus"),
		Status:        to.Ptr(true),
		AdditionalProperties: map[string]*string{
			"color": to.Ptr("red"),
			"city":  to.Ptr("Seattle"),
			"food":  to.Ptr("tikka masala"),
		},
		AdditionalProperties1: map[string]*float32{
			"height":   to.Ptr[float32](5.61),
			"weight":   to.Ptr[float32](599),
			"footsize": to.Ptr[float32](11.5),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// CreateAPObject - Create a Pet which contains more properties than what is defined.
func TestCreateAPObject(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateAPObject(context.Background(), PetAPObject{
		ID:   to.Ptr[int32](2),
		Name: to.Ptr("Hira"),
		AdditionalProperties: map[string]interface{}{
			"siblings": []interface{}{
				map[string]interface{}{
					"id":        float64(1),
					"name":      "Puppy",
					"birthdate": "2017-12-13T02:29:51Z",
					"complexProperty": map[string]interface{}{
						"color": "Red",
					},
				},
			},
			"picture": base64.StdEncoding.EncodeToString([]byte{255, 255, 255, 255, 254}),
		},
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.PetAPObject, PetAPObject{
		ID:     to.Ptr[int32](2),
		Name:   to.Ptr("Hira"),
		Status: to.Ptr(true),
		AdditionalProperties: map[string]interface{}{
			"siblings": []interface{}{
				map[string]interface{}{
					"id":        float64(1),
					"name":      "Puppy",
					"birthdate": "2017-12-13T02:29:51Z",
					"complexProperty": map[string]interface{}{
						"color": "Red",
					},
				},
			},
			"picture": base64.StdEncoding.EncodeToString([]byte{255, 255, 255, 255, 254}),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// CreateAPString - Create a Pet which contains more properties than what is defined.
func TestCreateAPString(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateAPString(context.Background(), PetAPString{
		ID:   to.Ptr[int32](3),
		Name: to.Ptr("Tommy"),
		AdditionalProperties: map[string]*string{
			"color":  to.Ptr("red"),
			"weight": to.Ptr("10 kg"),
			"city":   to.Ptr("Bombay"),
		},
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.PetAPString, PetAPString{
		ID:     to.Ptr[int32](3),
		Name:   to.Ptr("Tommy"),
		Status: to.Ptr(true),
		AdditionalProperties: map[string]*string{
			"color":  to.Ptr("red"),
			"weight": to.Ptr("10 kg"),
			"city":   to.Ptr("Bombay"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// CreateAPTrue - Create a Pet which contains more properties than what is defined.
func TestCreateAPTrue(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateAPTrue(context.Background(), PetAPTrue{
		ID:   to.Ptr[int32](1),
		Name: to.Ptr("Puppy"),
		AdditionalProperties: map[string]interface{}{
			"birthdate": "2017-12-13T02:29:51Z",
			"complexProperty": map[string]interface{}{
				"color": "Red",
			},
		},
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.PetAPTrue, PetAPTrue{
		ID:     to.Ptr[int32](1),
		Name:   to.Ptr("Puppy"),
		Status: to.Ptr(true),
		AdditionalProperties: map[string]interface{}{
			"birthdate": "2017-12-13T02:29:51Z",
			"complexProperty": map[string]interface{}{
				"color": "Red",
			},
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// CreateCatAPTrue - Create a CatAPTrue which contains more properties than what is defined.
func TestCreateCatAPTrue(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateCatAPTrue(context.Background(), CatAPTrue{
		ID:   to.Ptr[int32](1),
		Name: to.Ptr("Lisa"),
		AdditionalProperties: map[string]interface{}{
			"birthdate": "2017-12-13T02:29:51Z",
			"complexProperty": map[string]interface{}{
				"color": "Red",
			},
		},
		Friendly: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.CatAPTrue, CatAPTrue{
		ID:     to.Ptr[int32](1),
		Name:   to.Ptr("Lisa"),
		Status: to.Ptr(true),
		AdditionalProperties: map[string]interface{}{
			"birthdate": "2017-12-13T02:29:51Z",
			"complexProperty": map[string]interface{}{
				"color": "Red",
			},
		},
		Friendly: to.Ptr(true),
	}); r != "" {
		t.Fatal(r)
	}
}
