// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdatabasewatcher

const host = "https://management.azure.com"

const (
	moduleName    = "armdatabasewatcher"
	moduleVersion = "v0.1.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
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

// KustoOfferingType - The type of Kusto offering.
type KustoOfferingType string

const (
	// KustoOfferingTypeAdx - The Azure Data Explorer cluster Kusto offering.
	KustoOfferingTypeAdx KustoOfferingType = "adx"
	// KustoOfferingTypeFabric - The Fabric Real-Time Analytics Kusto offering.
	KustoOfferingTypeFabric KustoOfferingType = "fabric"
	// KustoOfferingTypeFree - The free Azure Data Explorer cluster Kusto offering.
	KustoOfferingTypeFree KustoOfferingType = "free"
)

// PossibleKustoOfferingTypeValues returns the possible values for the KustoOfferingType const type.
func PossibleKustoOfferingTypeValues() []KustoOfferingType {
	return []KustoOfferingType{
		KustoOfferingTypeAdx,
		KustoOfferingTypeFabric,
		KustoOfferingTypeFree,
	}
}

// ManagedIdentityType - The kind of managed identity assigned to this resource.
type ManagedIdentityType string

const (
	// ManagedIdentityTypeNone - No managed identity.
	ManagedIdentityTypeNone ManagedIdentityType = "None"
	// ManagedIdentityTypeSystemAndUserAssigned - System and user assigned managed identity.
	ManagedIdentityTypeSystemAndUserAssigned ManagedIdentityType = "SystemAssigned, UserAssigned"
	// ManagedIdentityTypeSystemAssigned - System assigned managed identity.
	ManagedIdentityTypeSystemAssigned ManagedIdentityType = "SystemAssigned"
	// ManagedIdentityTypeUserAssigned - User assigned managed identity.
	ManagedIdentityTypeUserAssigned ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns the possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{
		ManagedIdentityTypeNone,
		ManagedIdentityTypeSystemAndUserAssigned,
		ManagedIdentityTypeSystemAssigned,
		ManagedIdentityTypeUserAssigned,
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

// ProvisioningState - The status of the last provisioning operation performed on the resource.
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

// ResourceProvisioningState - The provisioning state of a resource type.
type ResourceProvisioningState string

const (
	// ResourceProvisioningStateCanceled - Resource creation was canceled.
	ResourceProvisioningStateCanceled ResourceProvisioningState = "Canceled"
	// ResourceProvisioningStateFailed - Resource creation failed.
	ResourceProvisioningStateFailed ResourceProvisioningState = "Failed"
	// ResourceProvisioningStateSucceeded - Resource has been created.
	ResourceProvisioningStateSucceeded ResourceProvisioningState = "Succeeded"
)

// PossibleResourceProvisioningStateValues returns the possible values for the ResourceProvisioningState const type.
func PossibleResourceProvisioningStateValues() []ResourceProvisioningState {
	return []ResourceProvisioningState{
		ResourceProvisioningStateCanceled,
		ResourceProvisioningStateFailed,
		ResourceProvisioningStateSucceeded,
	}
}

// SharedPrivateLinkResourceStatus - Status of the shared private link resource. Can be Pending, Approved, Rejected or Disconnected.
type SharedPrivateLinkResourceStatus string

const (
	// SharedPrivateLinkResourceStatusApproved - The shared private link connection request was approved by the resource owner.
	SharedPrivateLinkResourceStatusApproved SharedPrivateLinkResourceStatus = "Approved"
	// SharedPrivateLinkResourceStatusDisconnected - The shared private link connection request was disconnected by the resource
	// owner.
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	// SharedPrivateLinkResourceStatusPending - The shared private link connection request was not yet authorized by the resource
	// owner.
	SharedPrivateLinkResourceStatusPending SharedPrivateLinkResourceStatus = "Pending"
	// SharedPrivateLinkResourceStatusRejected - The shared private link connection request was rejected by the resource owner.
	SharedPrivateLinkResourceStatusRejected SharedPrivateLinkResourceStatus = "Rejected"
)

// PossibleSharedPrivateLinkResourceStatusValues returns the possible values for the SharedPrivateLinkResourceStatus const type.
func PossibleSharedPrivateLinkResourceStatusValues() []SharedPrivateLinkResourceStatus {
	return []SharedPrivateLinkResourceStatus{
		SharedPrivateLinkResourceStatusApproved,
		SharedPrivateLinkResourceStatusDisconnected,
		SharedPrivateLinkResourceStatusPending,
		SharedPrivateLinkResourceStatusRejected,
	}
}

// TargetAuthenticationType - The type of authentication to use when connecting to a target.
type TargetAuthenticationType string

const (
	// TargetAuthenticationTypeAAD - The Azure Active Directory authentication.
	TargetAuthenticationTypeAAD TargetAuthenticationType = "Aad"
	// TargetAuthenticationTypeSQL - The SQL password authentication.
	TargetAuthenticationTypeSQL TargetAuthenticationType = "Sql"
)

// PossibleTargetAuthenticationTypeValues returns the possible values for the TargetAuthenticationType const type.
func PossibleTargetAuthenticationTypeValues() []TargetAuthenticationType {
	return []TargetAuthenticationType{
		TargetAuthenticationTypeAAD,
		TargetAuthenticationTypeSQL,
	}
}

// WatcherStatus - The monitoring collection status of a watcher.
type WatcherStatus string

const (
	// WatcherStatusDeleting - Denotes the watcher is in a deleting state.
	WatcherStatusDeleting WatcherStatus = "Deleting"
	// WatcherStatusRunning - Denotes the watcher is in a running state.
	WatcherStatusRunning WatcherStatus = "Running"
	// WatcherStatusStarting - Denotes the watcher is in a starting state.
	WatcherStatusStarting WatcherStatus = "Starting"
	// WatcherStatusStopped - Denotes the watcher is in a stopped state.
	WatcherStatusStopped WatcherStatus = "Stopped"
	// WatcherStatusStopping - Denotes the watcher is in a stopping state.
	WatcherStatusStopping WatcherStatus = "Stopping"
)

// PossibleWatcherStatusValues returns the possible values for the WatcherStatus const type.
func PossibleWatcherStatusValues() []WatcherStatus {
	return []WatcherStatus{
		WatcherStatusDeleting,
		WatcherStatusRunning,
		WatcherStatusStarting,
		WatcherStatusStopped,
		WatcherStatusStopping,
	}
}