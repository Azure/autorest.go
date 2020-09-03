// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package additionalpropsgrouptest

import (
	"context"
	"encoding/base64"
	"generatortests/autorest/generated/additionalpropsgroup"
	"generatortests/helpers"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newPetsClient() additionalpropsgroup.PetsOperations {
	return additionalpropsgroup.NewPetsClient(additionalpropsgroup.NewDefaultClient(nil))
}

// CreateApInProperties - Create a Pet which contains more properties than what is defined.
func TestCreateApInProperties(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateApInProperties(context.Background(), additionalpropsgroup.PetApInProperties{
		ID:   to.Int32Ptr(4),
		Name: to.StringPtr("Bunny"),
		AdditionalProperties: &map[string]float32{
			"height":   5.61,
			"weight":   599,
			"footsize": 11.5,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.PetApInProperties, &additionalpropsgroup.PetApInProperties{
		ID:     to.Int32Ptr(4),
		Name:   to.StringPtr("Bunny"),
		Status: to.BoolPtr(true),
		AdditionalProperties: &map[string]float32{
			"height":   5.61,
			"weight":   599,
			"footsize": 11.5,
		},
	})
}

// CreateApInPropertiesWithApstring - Create a Pet which contains more properties than what is defined.
func TestCreateApInPropertiesWithApstring(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateApInPropertiesWithApstring(context.Background(), additionalpropsgroup.PetApInPropertiesWithApstring{
		ID:            to.Int32Ptr(5),
		Name:          to.StringPtr("Funny"),
		OdataLocation: to.StringPtr("westus"),
		AdditionalProperties: &map[string]string{
			"color": "red",
			"city":  "Seattle",
			"food":  "tikka masala",
		},
		AdditionalProperties1: &map[string]float32{
			"height":   5.61,
			"weight":   599,
			"footsize": 11.5,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.PetApInPropertiesWithApstring, &additionalpropsgroup.PetApInPropertiesWithApstring{
		ID:            to.Int32Ptr(5),
		Name:          to.StringPtr("Funny"),
		OdataLocation: to.StringPtr("westus"),
		Status:        to.BoolPtr(true),
		AdditionalProperties: &map[string]string{
			"color": "red",
			"city":  "Seattle",
			"food":  "tikka masala",
		},
		AdditionalProperties1: &map[string]float32{
			"height":   5.61,
			"weight":   599,
			"footsize": 11.5,
		},
	})
}

// CreateApObject - Create a Pet which contains more properties than what is defined.
func TestCreateApObject(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateApObject(context.Background(), additionalpropsgroup.PetApObject{
		ID:   to.Int32Ptr(2),
		Name: to.StringPtr("Hira"),
		AdditionalProperties: &map[string]interface{}{
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
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.PetApObject, &additionalpropsgroup.PetApObject{
		ID:     to.Int32Ptr(2),
		Name:   to.StringPtr("Hira"),
		Status: to.BoolPtr(true),
		AdditionalProperties: &map[string]interface{}{
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
	})
}

// CreateApString - Create a Pet which contains more properties than what is defined.
func TestCreateApString(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateApString(context.Background(), additionalpropsgroup.PetApString{
		ID:   to.Int32Ptr(3),
		Name: to.StringPtr("Tommy"),
		AdditionalProperties: &map[string]string{
			"color":  "red",
			"weight": "10 kg",
			"city":   "Bombay",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.PetApString, &additionalpropsgroup.PetApString{
		ID:     to.Int32Ptr(3),
		Name:   to.StringPtr("Tommy"),
		Status: to.BoolPtr(true),
		AdditionalProperties: &map[string]string{
			"color":  "red",
			"weight": "10 kg",
			"city":   "Bombay",
		},
	})
}

// CreateApTrue - Create a Pet which contains more properties than what is defined.
func TestCreateApTrue(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateApTrue(context.Background(), additionalpropsgroup.PetApTrue{
		ID:   to.Int32Ptr(1),
		Name: to.StringPtr("Puppy"),
		AdditionalProperties: &map[string]interface{}{
			"birthdate": "2017-12-13T02:29:51Z",
			"complexProperty": map[string]interface{}{
				"color": "Red",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.PetApTrue, &additionalpropsgroup.PetApTrue{
		ID:     to.Int32Ptr(1),
		Name:   to.StringPtr("Puppy"),
		Status: to.BoolPtr(true),
		AdditionalProperties: &map[string]interface{}{
			"birthdate": "2017-12-13T02:29:51Z",
			"complexProperty": map[string]interface{}{
				"color": "Red",
			},
		},
	})
}

// CreateCatApTrue - Create a CatAPTrue which contains more properties than what is defined.
func TestCreateCatApTrue(t *testing.T) {
	client := newPetsClient()
	result, err := client.CreateCatApTrue(context.Background(), additionalpropsgroup.CatApTrue{
		PetApTrue: additionalpropsgroup.PetApTrue{
			ID:   to.Int32Ptr(1),
			Name: to.StringPtr("Lisa"),
			AdditionalProperties: &map[string]interface{}{
				"birthdate": "2017-12-13T02:29:51Z",
				"complexProperty": map[string]interface{}{
					"color": "Red",
				},
			},
		},
		Friendly: to.BoolPtr(true),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.CatApTrue, &additionalpropsgroup.CatApTrue{
		PetApTrue: additionalpropsgroup.PetApTrue{
			ID:     to.Int32Ptr(1),
			Name:   to.StringPtr("Lisa"),
			Status: to.BoolPtr(true),
			AdditionalProperties: &map[string]interface{}{
				"birthdate": "2017-12-13T02:29:51Z",
				"complexProperty": map[string]interface{}{
					"color": "Red",
				},
			},
		},
		Friendly: to.BoolPtr(true),
	})
}
