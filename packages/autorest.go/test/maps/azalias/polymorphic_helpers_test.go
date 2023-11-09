// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azalias

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestGeoObjectNamedCollectionRoundTrip(t *testing.T) {
	for _, testcase := range []struct {
		name     string
		input    interface{}
		expected GeoJSONObjectNamedCollection
	}{{
		// Round tripping from a raw map confirms the JSON shape.
		name: "round trip from raw map",
		input: map[string]interface{}{
			"collectionName": "all",
			"objects": map[string]interface{}{
				"feature": map[string]interface{}{
					"type":        GeoJSONObjectTypeGeoJSONFeature,
					"id":          "id/feature",
					"featureType": "some type",
				},
				"object": map[string]interface{}{},
			},
		},
		expected: GeoJSONObjectNamedCollection{
			CollectionName: to.Ptr("all"),
			Objects: map[string]GeoJSONObjectClassification{
				"feature": &GeoJSONFeature{
					Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature),
					ID:          to.Ptr("id/feature"),
					FeatureType: to.Ptr("some type"),
				},
				"object": &GeoJSONObject{},
			},
		},
	}, {
		name: "round trip",
		input: GeoJSONObjectNamedCollection{
			CollectionName: to.Ptr("all"),
			Objects: map[string]GeoJSONObjectClassification{
				"feature": &GeoJSONFeature{
					Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature),
					ID:          to.Ptr("id/feature"),
					FeatureType: to.Ptr("some type"),
				},
				"object": &GeoJSONObject{},
			},
		},
		expected: GeoJSONObjectNamedCollection{
			CollectionName: to.Ptr("all"),
			Objects: map[string]GeoJSONObjectClassification{
				"feature": &GeoJSONFeature{
					Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature),
					ID:          to.Ptr("id/feature"),
					FeatureType: to.Ptr("some type"),
					Setting:     to.Ptr(DataSettingTwo),
				},
				"object": &GeoJSONObject{},
			},
		},
	}, {
		name: "empty map",
		input: GeoJSONObjectNamedCollection{
			CollectionName: to.Ptr("all"),
			Objects:        map[string]GeoJSONObjectClassification{},
		},
		expected: GeoJSONObjectNamedCollection{
			CollectionName: to.Ptr("all"),
			Objects:        map[string]GeoJSONObjectClassification{},
		},
	}} {
		t.Run(testcase.name, func(t *testing.T) {
			b, err := json.Marshal(testcase.input)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			var output GeoJSONObjectNamedCollection
			err = json.Unmarshal(b, &output)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !reflect.DeepEqual(testcase.expected, output) {
				t.Errorf("expected %#v, saw %#v", testcase.expected, output)
			}
		})
	}
}

func TestInterfaceRoundTrip(t *testing.T) {
	props1 := ScheduleCreateOrUpdateProperties{
		Aliases:     []string{"foo"},
		Description: to.Ptr("funky"),
		Interval:    false,
		StartTime:   to.Ptr(time.Now().UTC()),
	}
	b, err := json.Marshal(props1)
	if err != nil {
		t.Fatal(err)
	}
	var props2 ScheduleCreateOrUpdateProperties
	err = json.Unmarshal(b, &props2)
	if err != nil {
		t.Fatal(err)
	}
	if props2.Interval == nil {
		t.Fatal("props2.Interval is nil")
	}
	if *props1.Description != *props2.Description {
		t.Fatalf("expected %v, got %v", *props1.Description, *props2.Description)
	}
	i1, ok := props1.Interval.(bool)
	require.True(t, ok)
	i2, ok := props2.Interval.(bool)
	require.True(t, ok)
	if i1 != i2 {
		t.Fatalf("expected %v, got %v", props1.Interval, props2.Interval)
	}
	if *props1.StartTime != *props2.StartTime {
		t.Fatalf("expected %v, got %v", *props1.StartTime, *props2.StartTime)
	}
	if props1.Aliases[0] != props2.Aliases[0] {
		t.Fatalf("expected %v, got %v", props1.Aliases[0], props2.Aliases[0])
	}
}

func TestInterfaceNil(t *testing.T) {
	props1 := ScheduleCreateOrUpdateProperties{
		Description: to.Ptr("funky"),
		StartTime:   to.Ptr(time.Now().UTC()),
	}
	b, err := json.Marshal(props1)
	if err != nil {
		t.Fatal(err)
	}
	var props2 ScheduleCreateOrUpdateProperties
	err = json.Unmarshal(b, &props2)
	if err != nil {
		t.Fatal(err)
	}
	if *props1.Description != *props2.Description {
		t.Fatalf("expected %v, got %v", *props1.Description, *props2.Description)
	}
	if *props1.StartTime != *props2.StartTime {
		t.Fatalf("expected %v, got %v", *props1.StartTime, *props2.StartTime)
	}
	if props2.Interval != nil {
		t.Fatal("expected nil Interval")
	}
	if props2.Aliases != nil {
		t.Fatal("expected nil Aliases")
	}
}

func TestInterfaceJSONNull(t *testing.T) {
	props1 := TypeWithRawJSON{
		AnyObject: azcore.NullValue[*any](),
	}
	b, err := json.Marshal(props1)
	require.NoError(t, err)
	require.Contains(t, string(b), `"anyObject":null`)
	require.NotContains(t, string(b), "anything")
}

func TestGeoJSONRecursiveDisciminators(t *testing.T) {
	obj1 := GeoJSONRecursiveDisciminators{
		CombinedOne: []map[string]map[string]GeoJSONObjectClassification{
			{
				"entry": {
					"thing": &GeoJSONFeature{
						FeatureType: to.Ptr("slice-of-map-of-map-of-discriminators"),
						ID:          to.Ptr("entry-one"),
						Setting:     to.Ptr(DataSettingOne),
						Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature), // set by marshaller but set here to simplify the test
					},
				},
			},
		},
		CombinedThree: map[string][]map[string]GeoJSONObjectClassification{
			"entry": {
				{
					"thing": &GeoJSONFeature{
						FeatureType: to.Ptr("map-of-slice-of-map-of-discriminators"),
						ID:          to.Ptr("entry-one"),
						Setting:     to.Ptr(DataSettingOne),
						Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature), // set by marshaller but set here to simplify the test
					},
				},
			},
		},
		CombinedTwo: map[string]map[string][]GeoJSONObjectClassification{
			"entry": {
				"thing": {
					&GeoJSONFeature{
						FeatureType: to.Ptr("map-of-map-of-slice-of-discriminators"),
						ID:          to.Ptr("0-0"),
						Setting:     to.Ptr(DataSettingOne),
						Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature), // set by marshaller but set here to simplify the test
					},
					&GeoJSONFeature{
						FeatureType: to.Ptr("map-of-map-of-slice-of-discriminators"),
						ID:          to.Ptr("0-1"),
						Setting:     to.Ptr(DataSettingTwo),
						Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature), // set by marshaller but set here to simplify the test
					},
				},
			},
		},
		Items: [][]GeoJSONObjectClassification{
			{
				&GeoJSONFeature{
					FeatureType: to.Ptr("slice-of-slice-discriminators"),
					ID:          to.Ptr("0-0"),
					Setting:     to.Ptr(DataSettingOne),
					Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature), // set by marshaller but set here to simplify the test
				},
				&GeoJSONFeature{
					FeatureType: to.Ptr("slice-of-slice-discriminators"),
					ID:          to.Ptr("0-1"),
					Setting:     to.Ptr(DataSettingTwo),
					Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature), // set by marshaller but set here to simplify the test
				},
			},
			{
				&GeoJSONFeature{
					FeatureType: to.Ptr("slice-of-slice-discriminators"),
					ID:          to.Ptr("1-0"),
					Setting:     to.Ptr(DataSettingTwo),
					Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature), // set by marshaller but set here to simplify the test
				},
				&GeoJSONFeature{
					FeatureType: to.Ptr("slice-of-slice-discriminators"),
					ID:          to.Ptr("1-1"),
					Setting:     to.Ptr(DataSettingThree),
					Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature), // set by marshaller but set here to simplify the test
				},
			},
		},
		Objects: map[string]map[string]GeoJSONObjectClassification{
			"entry": {
				"thing": &GeoJSONFeature{
					FeatureType: to.Ptr("map-of-map-of-discriminators"),
					ID:          to.Ptr("entry-one"),
					Setting:     to.Ptr(DataSettingOne),
					Type:        to.Ptr(GeoJSONObjectTypeGeoJSONFeature), // set by marshaller but set here to simplify the test
				},
			},
		},
	}
	b, err := json.Marshal(obj1)
	require.NoError(t, err)
	var obj2 GeoJSONRecursiveDisciminators
	require.NoError(t, json.Unmarshal(b, &obj2))
	require.EqualValues(t, obj1, obj2)
}
