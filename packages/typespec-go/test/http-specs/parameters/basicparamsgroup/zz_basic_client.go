// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package basicparamsgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// BasicClient - Test for basic parameters cases.
// Don't use this type directly, use a constructor function instead.
type BasicClient struct {
	internal *azcore.Client
	endpoint string
}

// NewBasicExplicitBodyClient creates a new instance of [BasicExplicitBodyClient].
func (client *BasicClient) NewBasicExplicitBodyClient() *BasicExplicitBodyClient {
	return &BasicExplicitBodyClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewBasicImplicitBodyClient creates a new instance of [BasicImplicitBodyClient].
func (client *BasicClient) NewBasicImplicitBodyClient() *BasicImplicitBodyClient {
	return &BasicImplicitBodyClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}
