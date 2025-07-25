// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package datetimegroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// DatetimeClient - Test for encode decorator on datetime.
// Don't use this type directly, use a constructor function instead.
type DatetimeClient struct {
	internal *azcore.Client
	endpoint string
}

// NewDatetimeHeaderClient creates a new instance of [DatetimeHeaderClient].
func (client *DatetimeClient) NewDatetimeHeaderClient() *DatetimeHeaderClient {
	return &DatetimeHeaderClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewDatetimePropertyClient creates a new instance of [DatetimePropertyClient].
func (client *DatetimeClient) NewDatetimePropertyClient() *DatetimePropertyClient {
	return &DatetimePropertyClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewDatetimeQueryClient creates a new instance of [DatetimeQueryClient].
func (client *DatetimeClient) NewDatetimeQueryClient() *DatetimeQueryClient {
	return &DatetimeQueryClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewDatetimeResponseHeaderClient creates a new instance of [DatetimeResponseHeaderClient].
func (client *DatetimeClient) NewDatetimeResponseHeaderClient() *DatetimeResponseHeaderClient {
	return &DatetimeResponseHeaderClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}
