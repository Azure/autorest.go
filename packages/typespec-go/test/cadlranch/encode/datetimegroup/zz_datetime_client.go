// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package datetimegroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// DatetimeClient - Test for encode decorator on datetime.
// Don't use this type directly, use a constructor function instead.
type DatetimeClient struct {
	internal *azcore.Client
}

// NewHeaderClient creates a new instance of [HeaderClient].
func (client *DatetimeClient) NewHeaderClient() *HeaderClient {
	return &HeaderClient{
		internal: client.internal,
	}
}

// NewPropertyClient creates a new instance of [PropertyClient].
func (client *DatetimeClient) NewPropertyClient() *PropertyClient {
	return &PropertyClient{
		internal: client.internal,
	}
}

// NewQueryClient creates a new instance of [QueryClient].
func (client *DatetimeClient) NewQueryClient() *QueryClient {
	return &QueryClient{
		internal: client.internal,
	}
}

// NewResponseHeaderClient creates a new instance of [ResponseHeaderClient].
func (client *DatetimeClient) NewResponseHeaderClient() *ResponseHeaderClient {
	return &ResponseHeaderClient{
		internal: client.internal,
	}
}