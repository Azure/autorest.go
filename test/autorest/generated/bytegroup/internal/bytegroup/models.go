// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegroup

// GetEmptyResponse ...
type GetEmptyResponse struct {
	StatusCode int
	Value      []byte `json:"value,omitempty"`
}

// GetInvalidResponse ...
type GetInvalidResponse struct {
	StatusCode int
	Value      []byte `json:"value,omitempty"`
}

// GetNonASCIIResponse ...
type GetNonASCIIResponse struct {
	StatusCode int
	Value      []byte `json:"value,omitempty"`
}

// GetNullResponse ...
type GetNullResponse struct {
	StatusCode int
	Value      []byte `json:"value,omitempty"`
}

// PutNonASCIIResponse ...
type PutNonASCIIResponse struct {
	StatusCode int
}
