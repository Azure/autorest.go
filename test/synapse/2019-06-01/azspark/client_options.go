//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azspark

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// BatchClientOptions contains the optional settings for Client.
type BatchClientOptions struct {
	azcore.ClientOptions
}

// SessionClientOptions contains the optional settings for Client.
type SessionClientOptions struct {
	azcore.ClientOptions
}
