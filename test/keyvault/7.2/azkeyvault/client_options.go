//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azkeyvault

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ClientOptions contains the optional settings for Client.
type ClientOptions struct {
	azcore.ClientOptions
}

// HSMSecurityDomainClientOptions contains the optional settings for Client.
type HSMSecurityDomainClientOptions struct {
	azcore.ClientOptions
}

// RoleAssignmentsClientOptions contains the optional settings for Client.
type RoleAssignmentsClientOptions struct {
	azcore.ClientOptions
}

// RoleDefinitionsClientOptions contains the optional settings for Client.
type RoleDefinitionsClientOptions struct {
	azcore.ClientOptions
}
