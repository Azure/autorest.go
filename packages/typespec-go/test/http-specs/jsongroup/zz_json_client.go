// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package jsongroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// JSONClient - Projection
// Don't use this type directly, use a constructor function instead.
type JSONClient struct {
	internal *azcore.Client
}

// NewJSONPropertyClient creates a new instance of [JSONPropertyClient].
func (client *JSONClient) NewJSONPropertyClient() *JSONPropertyClient {
	return &JSONPropertyClient{
		internal: client.internal,
	}
}