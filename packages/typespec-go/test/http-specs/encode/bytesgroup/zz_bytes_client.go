// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bytesgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// BytesClient - Test for encode decorator on bytes.
// Don't use this type directly, use a constructor function instead.
type BytesClient struct {
	internal *azcore.Client
}

// NewBytesHeaderClient creates a new instance of [BytesHeaderClient].
func (client *BytesClient) NewBytesHeaderClient() *BytesHeaderClient {
	return &BytesHeaderClient{
		internal: client.internal,
	}
}

// NewBytesPropertyClient creates a new instance of [BytesPropertyClient].
func (client *BytesClient) NewBytesPropertyClient() *BytesPropertyClient {
	return &BytesPropertyClient{
		internal: client.internal,
	}
}

// NewBytesQueryClient creates a new instance of [BytesQueryClient].
func (client *BytesClient) NewBytesQueryClient() *BytesQueryClient {
	return &BytesQueryClient{
		internal: client.internal,
	}
}

// NewBytesRequestBodyClient creates a new instance of [BytesRequestBodyClient].
func (client *BytesClient) NewBytesRequestBodyClient() *BytesRequestBodyClient {
	return &BytesRequestBodyClient{
		internal: client.internal,
	}
}

// NewBytesResponseBodyClient creates a new instance of [BytesResponseBodyClient].
func (client *BytesClient) NewBytesResponseBodyClient() *BytesResponseBodyClient {
	return &BytesResponseBodyClient{
		internal: client.internal,
	}
}