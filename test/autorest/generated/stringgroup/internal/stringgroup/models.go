// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup

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

// GetMBCSResponse ...
type GetMBCSResponse struct {
	StatusCode int
	Value      string `json:"value,omitempty"`
}

// PutMBCSResponse ...
type PutMBCSResponse struct {
	StatusCode int
}
