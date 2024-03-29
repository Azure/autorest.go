// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"encoding/json"
	"singlediscgroup"
)

func unmarshalBirdClassification(rawMsg json.RawMessage) (singlediscgroup.BirdClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b singlediscgroup.BirdClassification
	switch m["kind"] {
	case "eagle":
		b = &singlediscgroup.Eagle{}
	case "goose":
		b = &singlediscgroup.Goose{}
	case "seagull":
		b = &singlediscgroup.SeaGull{}
	case "sparrow":
		b = &singlediscgroup.Sparrow{}
	default:
		b = &singlediscgroup.Bird{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
