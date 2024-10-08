// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package twoopgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// TwoOperationGroupClient contains the methods for the TwoOperationGroup group.
// Don't use this type directly, use a constructor function instead.
type TwoOperationGroupClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// NewTwoOperationGroupGroup1Client creates a new instance of [TwoOperationGroupGroup1Client].
func (client *TwoOperationGroupClient) NewTwoOperationGroupGroup1Client() *TwoOperationGroupGroup1Client {
	return &TwoOperationGroupGroup1Client{
		internal: client.internal,
		endpoint: client.endpoint,
		client:   client.client,
	}
}

// NewTwoOperationGroupGroup2Client creates a new instance of [TwoOperationGroupGroup2Client].
func (client *TwoOperationGroupClient) NewTwoOperationGroupGroup2Client() *TwoOperationGroupGroup2Client {
	return &TwoOperationGroupGroup2Client{
		internal: client.internal,
		endpoint: client.endpoint,
		client:   client.client,
	}
}
