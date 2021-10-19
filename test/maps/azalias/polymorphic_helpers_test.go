package azalias

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
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
			CollectionName: stringPtr("all"),
			Objects: map[string]GeoJSONObjectClassification{
				"feature": &GeoJSONFeature{
					GeoJSONObject: GeoJSONObject{
						Type: GeoJSONObjectTypeGeoJSONFeature.ToPtr(),
					},
					GeoJSONFeatureData: GeoJSONFeatureData{
						ID:          stringPtr("id/feature"),
						FeatureType: stringPtr("some type"),
					},
				},
				"object": &GeoJSONObject{},
			},
		},
	}, {
		name: "round trip",
		input: GeoJSONObjectNamedCollection{
			CollectionName: stringPtr("all"),
			Objects: map[string]GeoJSONObjectClassification{
				"feature": &GeoJSONFeature{
					GeoJSONObject: GeoJSONObject{
						Type: GeoJSONObjectTypeGeoJSONFeature.ToPtr(),
					},
					GeoJSONFeatureData: GeoJSONFeatureData{
						ID:          stringPtr("id/feature"),
						FeatureType: stringPtr("some type"),
					},
				},
				"object": &GeoJSONObject{},
			},
		},
		expected: GeoJSONObjectNamedCollection{
			CollectionName: stringPtr("all"),
			Objects: map[string]GeoJSONObjectClassification{
				"feature": &GeoJSONFeature{
					GeoJSONObject: GeoJSONObject{
						Type: GeoJSONObjectTypeGeoJSONFeature.ToPtr(),
					},
					GeoJSONFeatureData: GeoJSONFeatureData{
						ID:          stringPtr("id/feature"),
						FeatureType: stringPtr("some type"),
					},
				},
				"object": &GeoJSONObject{},
			},
		},
	}, {
		name: "empty map",
		input: GeoJSONObjectNamedCollection{
			CollectionName: stringPtr("all"),
			Objects:        map[string]GeoJSONObjectClassification{},
		},
		expected: GeoJSONObjectNamedCollection{
			CollectionName: stringPtr("all"),
			Objects:        map[string]GeoJSONObjectClassification{},
		},
	}} {
		t.Run(testcase.name, func(t *testing.T) {
			b, err := json.Marshal(testcase.input)
			require.NoError(t, err)
			var output GeoJSONObjectNamedCollection
			err = json.Unmarshal(b, &output)
			require.NoError(t, err)
			assert.Equal(t, testcase.expected, output)
		})
	}
}

func stringPtr(s string) *string {
	return &s
}
