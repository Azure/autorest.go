// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// Error ...
type Error struct {
	Status  string
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

// Basic ..
type Basic struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

// GetValidResponse ...
type GetValidResponse struct {
	StatusCode int
	Value      Basic
}

// PutValidResponse ...
type PutValidResponse struct {
	StatusCode int
}
