//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package aztables

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ClientOptions contains the optional settings for Client.
type ClientOptions struct {
	azcore.ClientOptions
}

// ServiceClientOptions contains the optional settings for Client.
type ServiceClientOptions struct {
	azcore.ClientOptions
}
