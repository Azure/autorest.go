// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armapicenter

// APIKind - The kind of the API
type APIKind string

const (
	// APIKindGraphql - A Graph query language Api
	APIKindGraphql APIKind = "graphql"
	// APIKindGrpc - A gRPC Api
	APIKindGrpc APIKind = "grpc"
	// APIKindRest - A Representational State Transfer Api
	APIKindRest APIKind = "rest"
	// APIKindSoap - A SOAP Api
	APIKindSoap APIKind = "soap"
	// APIKindWebhook - Web Hook
	APIKindWebhook APIKind = "webhook"
	// APIKindWebsocket - Web Socket
	APIKindWebsocket APIKind = "websocket"
)

// PossibleAPIKindValues returns the possible values for the APIKind const type.
func PossibleAPIKindValues() []APIKind {
	return []APIKind{
		APIKindGraphql,
		APIKindGrpc,
		APIKindRest,
		APIKindSoap,
		APIKindWebhook,
		APIKindWebsocket,
	}
}

// APISpecExportResultFormat - Result format for exported Api spec
type APISpecExportResultFormat string

const (
	// APISpecExportResultFormatInline - The inlined content of a specification document.
	APISpecExportResultFormatInline APISpecExportResultFormat = "inline"
	// APISpecExportResultFormatLink - The link to the result of the export operation. The URL is valid for 5 minutes.
	APISpecExportResultFormatLink APISpecExportResultFormat = "link"
)

// PossibleAPISpecExportResultFormatValues returns the possible values for the APISpecExportResultFormat const type.
func PossibleAPISpecExportResultFormatValues() []APISpecExportResultFormat {
	return []APISpecExportResultFormat{
		APISpecExportResultFormatInline,
		APISpecExportResultFormatLink,
	}
}

// APISpecImportSourceFormat - Source format for imported Api spec
type APISpecImportSourceFormat string

const (
	// APISpecImportSourceFormatInline - The inlined content of a specification document.
	APISpecImportSourceFormatInline APISpecImportSourceFormat = "inline"
	// APISpecImportSourceFormatLink - The link to a specification document hosted on a publicly accessible internet
	// address.
	APISpecImportSourceFormatLink APISpecImportSourceFormat = "link"
)

// PossibleAPISpecImportSourceFormatValues returns the possible values for the APISpecImportSourceFormat const type.
func PossibleAPISpecImportSourceFormatValues() []APISpecImportSourceFormat {
	return []APISpecImportSourceFormat{
		APISpecImportSourceFormatInline,
		APISpecImportSourceFormatLink,
	}
}

// ActionType - Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	// ActionTypeInternal - Actions are for internal-only APIs.
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DeploymentState - State of the Deployment
type DeploymentState string

const (
	// DeploymentStateActive - Active State
	DeploymentStateActive DeploymentState = "active"
	// DeploymentStateInactive - Inactive State
	DeploymentStateInactive DeploymentState = "inactive"
)

// PossibleDeploymentStateValues returns the possible values for the DeploymentState const type.
func PossibleDeploymentStateValues() []DeploymentState {
	return []DeploymentState{
		DeploymentStateActive,
		DeploymentStateInactive,
	}
}

// EnvironmentKind - The kind of environment
type EnvironmentKind string

const (
	// EnvironmentKindDevelopment - Development environment
	EnvironmentKindDevelopment EnvironmentKind = "development"
	// EnvironmentKindProduction - Production environment
	EnvironmentKindProduction EnvironmentKind = "production"
	// EnvironmentKindStaging - Staging environment
	EnvironmentKindStaging EnvironmentKind = "staging"
	// EnvironmentKindTesting - Testing environment
	EnvironmentKindTesting EnvironmentKind = "testing"
)

// PossibleEnvironmentKindValues returns the possible values for the EnvironmentKind const type.
func PossibleEnvironmentKindValues() []EnvironmentKind {
	return []EnvironmentKind{
		EnvironmentKindDevelopment,
		EnvironmentKindProduction,
		EnvironmentKindStaging,
		EnvironmentKindTesting,
	}
}

// EnvironmentServerType - The type of environment server
type EnvironmentServerType string

const (
	// EnvironmentServerTypeAWSAPIGateway - AWS Api Gateway server
	EnvironmentServerTypeAWSAPIGateway EnvironmentServerType = "AWS API Gateway"
	// EnvironmentServerTypeApigeeAPIManagement - Apigee server
	EnvironmentServerTypeApigeeAPIManagement EnvironmentServerType = "Apigee API Management"
	// EnvironmentServerTypeAzureAPIManagement - Api Management Server
	EnvironmentServerTypeAzureAPIManagement EnvironmentServerType = "Azure API Management"
	// EnvironmentServerTypeAzureComputeService - Compute server
	EnvironmentServerTypeAzureComputeService EnvironmentServerType = "Azure compute service"
	// EnvironmentServerTypeKongAPIGateway - Kong API Gateway server
	EnvironmentServerTypeKongAPIGateway EnvironmentServerType = "Kong API Gateway"
	// EnvironmentServerTypeKubernetes - Kubernetes server
	EnvironmentServerTypeKubernetes EnvironmentServerType = "Kubernetes"
	// EnvironmentServerTypeMuleSoftAPIManagement - Mulesoft Api Management server
	EnvironmentServerTypeMuleSoftAPIManagement EnvironmentServerType = "MuleSoft API Management"
)

// PossibleEnvironmentServerTypeValues returns the possible values for the EnvironmentServerType const type.
func PossibleEnvironmentServerTypeValues() []EnvironmentServerType {
	return []EnvironmentServerType{
		EnvironmentServerTypeAWSAPIGateway,
		EnvironmentServerTypeApigeeAPIManagement,
		EnvironmentServerTypeAzureAPIManagement,
		EnvironmentServerTypeAzureComputeService,
		EnvironmentServerTypeKongAPIGateway,
		EnvironmentServerTypeKubernetes,
		EnvironmentServerTypeMuleSoftAPIManagement,
	}
}

// LifecycleStage - The stage of the Api development lifecycle
type LifecycleStage string

const (
	// LifecycleStageDeprecated - deprecated stage
	LifecycleStageDeprecated LifecycleStage = "deprecated"
	// LifecycleStageDesign - design stage
	LifecycleStageDesign LifecycleStage = "design"
	// LifecycleStageDevelopment - development stage
	LifecycleStageDevelopment LifecycleStage = "development"
	// LifecycleStagePreview - In preview
	LifecycleStagePreview LifecycleStage = "preview"
	// LifecycleStageProduction - In production
	LifecycleStageProduction LifecycleStage = "production"
	// LifecycleStageRetired - Retired stage
	LifecycleStageRetired LifecycleStage = "retired"
	// LifecycleStageTesting - testing stage
	LifecycleStageTesting LifecycleStage = "testing"
)

// PossibleLifecycleStageValues returns the possible values for the LifecycleStage const type.
func PossibleLifecycleStageValues() []LifecycleStage {
	return []LifecycleStage{
		LifecycleStageDeprecated,
		LifecycleStageDesign,
		LifecycleStageDevelopment,
		LifecycleStagePreview,
		LifecycleStageProduction,
		LifecycleStageRetired,
		LifecycleStageTesting,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	// ManagedServiceIdentityTypeNone - No managed identity.
	ManagedServiceIdentityTypeNone ManagedServiceIdentityType = "None"
	// ManagedServiceIdentityTypeSystemAssigned - System assigned managed identity.
	ManagedServiceIdentityTypeSystemAssigned ManagedServiceIdentityType = "SystemAssigned"
	// ManagedServiceIdentityTypeSystemAssignedUserAssigned - System and user assigned managed identity.
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	// ManagedServiceIdentityTypeUserAssigned - User assigned managed identity.
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// MetadataAssignmentEntity - Assignment entity for Metadata
type MetadataAssignmentEntity string

const (
	// MetadataAssignmentEntityAPI - Assigned to API
	MetadataAssignmentEntityAPI MetadataAssignmentEntity = "api"
	// MetadataAssignmentEntityDeployment - Assigned to Deployment
	MetadataAssignmentEntityDeployment MetadataAssignmentEntity = "deployment"
	// MetadataAssignmentEntityEnvironment - Assigned to Environment
	MetadataAssignmentEntityEnvironment MetadataAssignmentEntity = "environment"
)

// PossibleMetadataAssignmentEntityValues returns the possible values for the MetadataAssignmentEntity const type.
func PossibleMetadataAssignmentEntityValues() []MetadataAssignmentEntity {
	return []MetadataAssignmentEntity{
		MetadataAssignmentEntityAPI,
		MetadataAssignmentEntityDeployment,
		MetadataAssignmentEntityEnvironment,
	}
}

// MetadataSchemaExportFormat - The format for schema export
type MetadataSchemaExportFormat string

const (
	// MetadataSchemaExportFormatInline - The inlined content of a schema document.
	MetadataSchemaExportFormatInline MetadataSchemaExportFormat = "inline"
	// MetadataSchemaExportFormatLink - The link to a schema document. The URL is valid for 5 minutes.
	MetadataSchemaExportFormatLink MetadataSchemaExportFormat = "link"
)

// PossibleMetadataSchemaExportFormatValues returns the possible values for the MetadataSchemaExportFormat const type.
func PossibleMetadataSchemaExportFormatValues() []MetadataSchemaExportFormat {
	return []MetadataSchemaExportFormat{
		MetadataSchemaExportFormatInline,
		MetadataSchemaExportFormatLink,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	// OriginSystem - Indicates the operation is initiated by a system.
	OriginSystem Origin = "system"
	// OriginUser - Indicates the operation is initiated by a user.
	OriginUser Origin = "user"
	// OriginUserSystem - Indicates the operation is initiated by a user or system.
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - The provisioning state of the resource
type ProvisioningState string

const (
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}
