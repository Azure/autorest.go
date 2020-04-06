// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"generatortests/autorest/generated/complexgroup"
	"testing"
)

func getPolymorphismOperations(t *testing.T) complexgroup.PolymorphismOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.PolymorphismOperations()
}

/*func TestPolymorphismGetValid(t *testing.T) {
	client := getPolymorphismOperations(t)

	fishtype, length, iswild, location, species := "smart_salmon", float32(1), true, "alaska", "king"
	// TODO the fields related to the FishModel siblings on SmartSalmon are commented due to unmarshalling that's pending
	// sLen, sBirth, sAge, sSpecies := float64(20), time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC), int32(6), "predator"
	// ssLen, ssBirth, ssAge, ssSpecies := float64(10), time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC), int32(105), "dangerous"
	// gsLen, gsBirth, gsAge, gsSpecies, gsJaw := float64(30), time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC), int32(1), "scary", int32(5)
	// gsColor := complexgroup.GoblinSharkColorPink

	var ss = complexgroup.SmartSalmon{
		Fishtype: &fishtype,
		Length:   &length,
		Iswild:   &iswild,
		Location: &location,
		Species:  &species,
		// AdditionalProperties: map[string]interface{}{
		// 	"additionalProperty1": float64(1),
		// 	"additionalProperty2": false,
		// 	"additionalProperty3": "hello",
		// 	"additionalProperty4": map[string]interface{}{
		// 		"a": float64(1),
		// 		"b": float64(2),
		// 	},
		// 	"additionalProperty5": []interface{}{float64(1), float64(3)},
		// },
		// Siblings: &[]complexgroup.FishModel{
		// 	complexgroup.Shark{
		// 		Length:   &sLen,
		// 		Birthday: &sBirth,
		// 		Age:      &sAge,
		// 		Species:  &sSpecies,
		// 	},
		// 	complexgroup.Sawshark{
		// 		Length:   &ssLen,
		// 		Birthday: &ssBirth,
		// 		Age:      &ssAge,
		// 		Species:  &ssSpecies,
		// 		Picture:  &[]byte{255, 255, 255, 255, 254},
		// 	},
		// 	complexgroup.Goblinshark{
		// 		Length:   &gsLen,
		// 		Birthday: &gsBirth,
		// 		Age:      &gsAge,
		// 		Species:  &gsSpecies,
		// 		Jawsize:  &gsJaw,
		// 		Color:    &gsColor,
		// 	},
		// },
	}

	result, err := client.GetComplicated(context.Background())
	if err != nil {
		t.Fatalf("GetComplicated: %v", err)
	}
	expected := &complexgroup.PolymorphismGetComplicatedResponse{
		StatusCode: http.StatusOK,
		Salmon:     ss,
	}
	deepEqualOrFatal(t, result, expected)
}*/
