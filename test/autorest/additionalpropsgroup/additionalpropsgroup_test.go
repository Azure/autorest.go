// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package additionalpropsgroup

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newPetsClient() *PetsClient {
	return NewPetsClient(nil)
}

// CreateAPInProperties - Create a Pet which contains more properties than what is defined.
func TestCreateAPInProperties(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateAPInProperties(context.Background(), PetAPInProperties{
		ID:   to.Int32Ptr(4),
		Name: to.StringPtr("Bunny"),
		AdditionalProperties: map[string]*float32{
			"height":   to.Float32Ptr(5.61),
			"weight":   to.Float32Ptr(599),
			"footsize": to.Float32Ptr(11.5),
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.PetAPInProperties, PetAPInProperties{
		ID:     to.Int32Ptr(4),
		Name:   to.StringPtr("Bunny"),
		Status: to.BoolPtr(true),
		AdditionalProperties: map[string]*float32{
			"height":   to.Float32Ptr(5.61),
			"weight":   to.Float32Ptr(599),
			"footsize": to.Float32Ptr(11.5),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// CreateAPInPropertiesWithAPString - Create a Pet which contains more properties than what is defined.
func TestCreateAPInPropertiesWithAPString(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateAPInPropertiesWithAPString(context.Background(), PetAPInPropertiesWithAPString{
		ID:            to.Int32Ptr(5),
		Name:          to.StringPtr("Funny"),
		ODataLocation: to.StringPtr("westus"),
		AdditionalProperties: map[string]*string{
			"color": to.StringPtr("red"),
			"city":  to.StringPtr("Seattle"),
			"food":  to.StringPtr("tikka masala"),
		},
		AdditionalProperties1: map[string]*float32{
			"height":   to.Float32Ptr(5.61),
			"weight":   to.Float32Ptr(599),
			"footsize": to.Float32Ptr(11.5),
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.PetAPInPropertiesWithAPString, PetAPInPropertiesWithAPString{
		ID:            to.Int32Ptr(5),
		Name:          to.StringPtr("Funny"),
		ODataLocation: to.StringPtr("westus"),
		Status:        to.BoolPtr(true),
		AdditionalProperties: map[string]*string{
			"color": to.StringPtr("red"),
			"city":  to.StringPtr("Seattle"),
			"food":  to.StringPtr("tikka masala"),
		},
		AdditionalProperties1: map[string]*float32{
			"height":   to.Float32Ptr(5.61),
			"weight":   to.Float32Ptr(599),
			"footsize": to.Float32Ptr(11.5),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// CreateAPObject - Create a Pet which contains more properties than what is defined.
func TestCreateAPObject(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateAPObject(context.Background(), PetAPObject{
		ID:   to.Int32Ptr(2),
		Name: to.StringPtr("Hira"),
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
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.PetAPObject, PetAPObject{
		ID:     to.Int32Ptr(2),
		Name:   to.StringPtr("Hira"),
		Status: to.BoolPtr(true),
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
		ID:   to.Int32Ptr(3),
		Name: to.StringPtr("Tommy"),
		AdditionalProperties: map[string]*string{
			"color":  to.StringPtr("red"),
			"weight": to.StringPtr("10 kg"),
			"city":   to.StringPtr("Bombay"),
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.PetAPString, PetAPString{
		ID:     to.Int32Ptr(3),
		Name:   to.StringPtr("Tommy"),
		Status: to.BoolPtr(true),
		AdditionalProperties: map[string]*string{
			"color":  to.StringPtr("red"),
			"weight": to.StringPtr("10 kg"),
			"city":   to.StringPtr("Bombay"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// CreateAPTrue - Create a Pet which contains more properties than what is defined.
func TestCreateAPTrue(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateAPTrue(context.Background(), PetAPTrue{
		ID:   to.Int32Ptr(1),
		Name: to.StringPtr("Puppy"),
		AdditionalProperties: map[string]interface{}{
			"birthdate": "2017-12-13T02:29:51Z",
			"complexProperty": map[string]interface{}{
				"color": "Red",
			},
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.PetAPTrue, PetAPTrue{
		ID:     to.Int32Ptr(1),
		Name:   to.StringPtr("Puppy"),
		Status: to.BoolPtr(true),
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
		PetAPTrue: PetAPTrue{
			ID:   to.Int32Ptr(1),
			Name: to.StringPtr("Lisa"),
			AdditionalProperties: map[string]interface{}{
				"birthdate": "2017-12-13T02:29:51Z",
				"complexProperty": map[string]interface{}{
					"color": "Red",
				},
			},
		},
		Friendly: to.BoolPtr(true),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.CatAPTrue, CatAPTrue{
		PetAPTrue: PetAPTrue{
			ID:     to.Int32Ptr(1),
			Name:   to.StringPtr("Lisa"),
			Status: to.BoolPtr(true),
			AdditionalProperties: map[string]interface{}{
				"birthdate": "2017-12-13T02:29:51Z",
				"complexProperty": map[string]interface{}{
					"color": "Red",
				},
			},
		},
		Friendly: to.BoolPtr(true),
	}); r != "" {
		t.Fatal(r)
	}
}
