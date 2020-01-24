// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package custombaseurlgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

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

// GetEmptyResponse ...
type GetEmptyResponse struct {
	StatusCode int
}
