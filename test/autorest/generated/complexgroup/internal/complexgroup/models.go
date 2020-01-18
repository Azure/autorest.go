// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

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

// ColorValues ...
func ColorValues() []ColorType {
	return []ColorType{ColorYellow, ColorMagenta, ColorCyan, ColorBlack}
}

// Error ...
type Error struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func newError(resp *azcore.Response) error {
	err := &Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (e *Error) Error() string {
	return e.Message
}

// Basic ..
type Basic struct {
	ID    int       `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Color ColorType `json:"color,omitempty"`
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
