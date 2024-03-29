// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package lrorpcgroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Error.
func (e Error) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "innererror", e.Innererror)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Error.
func (e *Error) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "innererror":
			err = unpopulate(val, "Innererror", &e.Innererror)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GenerationOptions.
func (g GenerationOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "prompt", g.Prompt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GenerationOptions.
func (g *GenerationOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "prompt":
			err = unpopulate(val, "Prompt", &g.Prompt)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GenerationResponse.
func (g GenerationResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", g.Error)
	populate(objectMap, "id", g.ID)
	populate(objectMap, "result", g.Result)
	populate(objectMap, "status", g.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GenerationResponse.
func (g *GenerationResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &g.Error)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &g.ID)
			delete(rawMsg, key)
		case "result":
			err = unpopulate(val, "Result", &g.Result)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &g.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GenerationResult.
func (g GenerationResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "data", g.Data)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GenerationResult.
func (g *GenerationResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "data":
			err = unpopulate(val, "Data", &g.Data)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type InnerError.
func (i InnerError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", i.Code)
	populate(objectMap, "innererror", i.Innererror)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InnerError.
func (i *InnerError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &i.Code)
			delete(rawMsg, key)
		case "innererror":
			err = unpopulate(val, "Innererror", &i.Innererror)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
