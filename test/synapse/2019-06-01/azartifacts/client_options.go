//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azartifacts

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// BigDataPoolsClientOptions contains the optional settings for Client.
type BigDataPoolsClientOptions struct {
	azcore.ClientOptions
}

// DataFlowClientOptions contains the optional settings for Client.
type DataFlowClientOptions struct {
	azcore.ClientOptions
}

// DataFlowDebugSessionClientOptions contains the optional settings for Client.
type DataFlowDebugSessionClientOptions struct {
	azcore.ClientOptions
}

// DatasetClientOptions contains the optional settings for Client.
type DatasetClientOptions struct {
	azcore.ClientOptions
}

// IntegrationRuntimesClientOptions contains the optional settings for Client.
type IntegrationRuntimesClientOptions struct {
	azcore.ClientOptions
}

// LibraryClientOptions contains the optional settings for Client.
type LibraryClientOptions struct {
	azcore.ClientOptions
}

// LinkedServiceClientOptions contains the optional settings for Client.
type LinkedServiceClientOptions struct {
	azcore.ClientOptions
}

// NotebookClientOptions contains the optional settings for Client.
type NotebookClientOptions struct {
	azcore.ClientOptions
}

// PipelineClientOptions contains the optional settings for Client.
type PipelineClientOptions struct {
	azcore.ClientOptions
}

// PipelineRunClientOptions contains the optional settings for Client.
type PipelineRunClientOptions struct {
	azcore.ClientOptions
}

// SparkJobDefinitionClientOptions contains the optional settings for Client.
type SparkJobDefinitionClientOptions struct {
	azcore.ClientOptions
}

// SQLPoolsClientOptions contains the optional settings for Client.
type SQLPoolsClientOptions struct {
	azcore.ClientOptions
}

// SQLScriptClientOptions contains the optional settings for Client.
type SQLScriptClientOptions struct {
	azcore.ClientOptions
}

// TriggerClientOptions contains the optional settings for Client.
type TriggerClientOptions struct {
	azcore.ClientOptions
}

// TriggerRunClientOptions contains the optional settings for Client.
type TriggerRunClientOptions struct {
	azcore.ClientOptions
}

// WorkspaceClientOptions contains the optional settings for Client.
type WorkspaceClientOptions struct {
	azcore.ClientOptions
}

// WorkspaceGitRepoManagementClientOptions contains the optional settings for Client.
type WorkspaceGitRepoManagementClientOptions struct {
	azcore.ClientOptions
}
