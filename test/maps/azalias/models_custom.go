//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azalias

import (
	"encoding/json"
	"fmt"
)

// GeoJSONObjectNamedCollection - A named collection of GeoJSON object
type GeoJSONObjectNamedCollection struct {
	// Name of the collection
	CollectionName *string `json:"collectionName,omitempty"`

	// Dictionary of
	Objects map[string]GeoJSONObjectClassification `json:"objects,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type GeoJSONObjectNamedCollection.
func (g GeoJSONObjectNamedCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "collectionName", g.CollectionName)
	populate(objectMap, "objects", g.Objects)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GeoJSONObjectNamedCollection.
func (g *GeoJSONObjectNamedCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "collectionName":
			err = unpopulate(val, "CollectionName", &g.CollectionName)
			delete(rawMsg, key)
		case "objects":
			g.Objects, err = unmarshalGeoJSONObjectClassificationMap(val)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}
