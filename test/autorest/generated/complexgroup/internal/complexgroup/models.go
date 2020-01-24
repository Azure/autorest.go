// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ColorType is an enumerated type for complex group color strings
type ColorType string

const (
	// ColorYellow ...
	ColorYellow ColorType = "YELLOW"
	// ColorMagenta ...
	ColorMagenta ColorType = "Magenta"
	// ColorCyan ...
	ColorCyan ColorType = "cyan"
	// ColorBlack ...
	ColorBlack ColorType = "blacK"
)

// PossibleColorValues ...
func PossibleColorValues() []ColorType {
	return []ColorType{ColorYellow, ColorMagenta, ColorCyan, ColorBlack}
}

// Error ...
type Error struct {
	Status  *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

func newError(resp *azcore.Response) error {
	err := &Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (e *Error) Error() string {
	return *e.Message
}

// Basic ..
type Basic struct {
	ID    *int       `json:"id,omitempty"`
	Name  *string    `json:"name,omitempty"`
	Color *ColorType `json:"color,omitempty"`
}

// IntWrapper ...
type IntWrapper struct {
	Field1 *int32 `json:"field1,omitempty"`
	Field2 *int32 `json:"field2,omitempty"`
}

// LongWrapper ...
type LongWrapper struct {
	Field1 *int64 `json:"field1,omitempty"`
	Field2 *int64 `json:"field2,omitempty"`
}

// FloatWrapper ...
type FloatWrapper struct {
	Field1 *float32 `json:"field1,omitempty"`
	Field2 *float32 `json:"field2,omitempty"`
}

// DoubleWrapper ..
type DoubleWrapper struct {
	Field1                                                                          *float64 `json:"field1,omitempty"`
	Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose *float64 `json:"field_56_zeros_after_the_dot_and_negative_zero_before_dot_and_this_is_a_long_field_name_on_purpose,omitempty"`
}

// BooleanWrapper ...
type BooleanWrapper struct {
	FieldTrue  *bool `json:"field_true,omitempty"`
	FieldFalse *bool `json:"field_false,omitempty"`
}

// StringWrapper ...
type StringWrapper struct {
	Field *string `json:"field,omitempty"`
	Empty *string `json:"empty,omitempty"`
	Null  *string `json:"null,omitempty"`
}

// DateWrapper ...
type DateWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Leap  *time.Time `json:"leap,omitempty"`
}

// GetValidResponse ...
type GetValidResponse struct {
	StatusCode int
	Basic      *Basic
}

// PutValidResponse ...
type PutValidResponse struct {
	StatusCode int
}

// GetInvalidResponse ...
type GetInvalidResponse struct {
	StatusCode int
	Basic      *Basic
}

// GetEmptyResponse ...
type GetEmptyResponse struct {
	StatusCode int
	Basic      *Basic
}

// GetNullResponse ...
type GetNullResponse struct {
	StatusCode int
	Basic      *Basic
}

// GetNotProvidedResponse ...
type GetNotProvidedResponse struct {
	StatusCode int
	Basic      *Basic
}

// GetIntResponse ...
type GetIntResponse struct {
	StatusCode int
	IntWrapper *IntWrapper
}

// PutIntResponse ...
type PutIntResponse struct {
	StatusCode int
	IntWrapper *IntWrapper
}

// GetLongResponse ...
type GetLongResponse struct {
	StatusCode  int
	LongWrapper *LongWrapper
}

// PutLongResponse ...
type PutLongResponse struct {
	StatusCode  int
	LongWrapper *LongWrapper
}

// GetFloatResponse ...
type GetFloatResponse struct {
	StatusCode   int
	FloatWrapper *FloatWrapper
}

// PutFloatResponse ...
type PutFloatResponse struct {
	StatusCode   int
	FloatWrapper *FloatWrapper
}

// GetDoubleResponse ...
type GetDoubleResponse struct {
	StatusCode    int
	DoubleWrapper *DoubleWrapper
}

// PutDoubleResponse ...
type PutDoubleResponse struct {
	StatusCode    int
	DoubleWrapper *DoubleWrapper
}

// GetBoolResponse ...
type GetBoolResponse struct {
	StatusCode     int
	BooleanWrapper *BooleanWrapper
}

// PutBoolResponse ...
type PutBoolResponse struct {
	StatusCode     int
	BooleanWrapper *BooleanWrapper
}

// GetStringResponse ...
type GetStringResponse struct {
	StatusCode    int
	StringWrapper *StringWrapper
}

// PutStringResponse ...
type PutStringResponse struct {
	StatusCode    int
	StringWrapper *StringWrapper
}

// GetDateResponse ...
type GetDateResponse struct {
	StatusCode  int
	DateWrapper *DateWrapper
}

// PutDateResponse ...
type PutDateResponse struct {
	StatusCode  int
	DateWrapper *DateWrapper
}
