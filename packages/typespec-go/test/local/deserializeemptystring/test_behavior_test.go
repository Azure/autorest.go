package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Mock the generated ResponseModel for testing
type ResponseModel struct {
	SampleURL      *string `json:"sampleUrl"`
	NormalProperty *string `json:"normalProperty"`
}

// Mock the generated UnmarshalJSON method with @deserializeEmptyStringAsNull behavior
func (r *ResponseModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "normalProperty":
			err = unpopulate(val, "NormalProperty", &r.NormalProperty)
		case "sampleUrl":
			// This is the special handling for @deserializeEmptyStringAsNull
			if val != nil && string(val) != "null" && string(val) != "\"\"" {
				err = unpopulate(val, "SampleURL", &r.SampleURL)
			}
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
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

func TestDeserializeEmptyStringAsNull(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *string
	}{
		{
			name:     "Empty string should be treated as null",
			input:    `{"sampleUrl": "", "normalProperty": "value"}`,
			expected: nil,
		},
		{
			name:     "Non-empty string should be preserved",
			input:    `{"sampleUrl": "https://example.com", "normalProperty": "value"}`,
			expected: func() *string { s := "https://example.com"; return &s }(),
		},
		{
			name:     "Null should remain null",
			input:    `{"sampleUrl": null, "normalProperty": "value"}`,
			expected: nil,
		},
		{
			name:     "Missing property should remain nil",
			input:    `{"normalProperty": "value"}`,
			expected: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var model ResponseModel
			err := json.Unmarshal([]byte(tc.input), &model)
			if err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			if tc.expected == nil {
				if model.SampleURL != nil {
					t.Errorf("Expected SampleURL to be nil, but got %v", *model.SampleURL)
				}
			} else {
				if model.SampleURL == nil {
					t.Errorf("Expected SampleURL to be %v, but got nil", *tc.expected)
				} else if *model.SampleURL != *tc.expected {
					t.Errorf("Expected SampleURL to be %v, but got %v", *tc.expected, *model.SampleURL)
				}
			}
		})
	}
}

func TestNormalPropertyStillWorks(t *testing.T) {
	input := `{"sampleUrl": "", "normalProperty": ""}`
	var model ResponseModel
	err := json.Unmarshal([]byte(input), &model)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// SampleURL should be nil due to @deserializeEmptyStringAsNull
	if model.SampleURL != nil {
		t.Errorf("Expected SampleURL to be nil due to empty string, but got %v", *model.SampleURL)
	}

	// NormalProperty should be an empty string (normal behavior)
	if model.NormalProperty == nil {
		t.Errorf("Expected NormalProperty to be empty string, but got nil")
	} else if *model.NormalProperty != "" {
		t.Errorf("Expected NormalProperty to be empty string, but got %v", *model.NormalProperty)
	}
}