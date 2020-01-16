// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// Error ...
type Error struct {
	Status  int32
	Message string
}

func newError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (e Error) Error() string {
	return e.Message
}

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
