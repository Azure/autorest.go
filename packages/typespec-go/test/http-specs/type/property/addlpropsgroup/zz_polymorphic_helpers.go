// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package addlpropsgroup

import "encoding/json"

func unmarshalExtendsUnknownAdditionalPropertiesDiscriminatedClassification(rawMsg json.RawMessage) (ExtendsUnknownAdditionalPropertiesDiscriminatedClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ExtendsUnknownAdditionalPropertiesDiscriminatedClassification
	switch m["kind"] {
	case "derived":
		b = &ExtendsUnknownAdditionalPropertiesDiscriminatedDerived{}
	default:
		b = &ExtendsUnknownAdditionalPropertiesDiscriminated{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalIsUnknownAdditionalPropertiesDiscriminatedClassification(rawMsg json.RawMessage) (IsUnknownAdditionalPropertiesDiscriminatedClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b IsUnknownAdditionalPropertiesDiscriminatedClassification
	switch m["kind"] {
	case "derived":
		b = &IsUnknownAdditionalPropertiesDiscriminatedDerived{}
	default:
		b = &IsUnknownAdditionalPropertiesDiscriminated{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
