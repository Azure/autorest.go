// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package resources

// NestedClientCreateOrReplaceResponse contains the response from method NestedClient.BeginCreateOrReplace.
type NestedClientCreateOrReplaceResponse struct {
	// Nested child of Top Level Tracked Resource.
	NestedProxyResource
}

// NestedClientDeleteResponse contains the response from method NestedClient.BeginDelete.
type NestedClientDeleteResponse struct {
	// placeholder for future response values
}

// NestedClientGetResponse contains the response from method NestedClient.Get.
type NestedClientGetResponse struct {
	// Nested child of Top Level Tracked Resource.
	NestedProxyResource
}

// NestedClientListByTopLevelTrackedResourceResponse contains the response from method NestedClient.NewListByTopLevelTrackedResourcePager.
type NestedClientListByTopLevelTrackedResourceResponse struct {
	// The response of a NestedProxyResource list operation.
	NestedProxyResourceListResult
}

// NestedClientUpdateResponse contains the response from method NestedClient.BeginUpdate.
type NestedClientUpdateResponse struct {
	// Nested child of Top Level Tracked Resource.
	NestedProxyResource
}

// SingletonClientCreateOrUpdateResponse contains the response from method SingletonClient.BeginCreateOrUpdate.
type SingletonClientCreateOrUpdateResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	SingletonTrackedResource
}

// SingletonClientGetByResourceGroupResponse contains the response from method SingletonClient.GetByResourceGroup.
type SingletonClientGetByResourceGroupResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	SingletonTrackedResource
}

// SingletonClientListByResourceGroupResponse contains the response from method SingletonClient.NewListByResourceGroupPager.
type SingletonClientListByResourceGroupResponse struct {
	// The response of a SingletonTrackedResource list operation.
	SingletonTrackedResourceListResult
}

// SingletonClientUpdateResponse contains the response from method SingletonClient.Update.
type SingletonClientUpdateResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	SingletonTrackedResource
}

// TopLevelClientActionSyncResponse contains the response from method TopLevelClient.ActionSync.
type TopLevelClientActionSyncResponse struct {
	// placeholder for future response values
}

// TopLevelClientCreateOrReplaceResponse contains the response from method TopLevelClient.BeginCreateOrReplace.
type TopLevelClientCreateOrReplaceResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	TopLevelTrackedResource
}

// TopLevelClientDeleteResponse contains the response from method TopLevelClient.BeginDelete.
type TopLevelClientDeleteResponse struct {
	// placeholder for future response values
}

// TopLevelClientGetResponse contains the response from method TopLevelClient.Get.
type TopLevelClientGetResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	TopLevelTrackedResource
}

// TopLevelClientListByResourceGroupResponse contains the response from method TopLevelClient.NewListByResourceGroupPager.
type TopLevelClientListByResourceGroupResponse struct {
	// The response of a TopLevelTrackedResource list operation.
	TopLevelTrackedResourceListResult
}

// TopLevelClientListBySubscriptionResponse contains the response from method TopLevelClient.NewListBySubscriptionPager.
type TopLevelClientListBySubscriptionResponse struct {
	// The response of a TopLevelTrackedResource list operation.
	TopLevelTrackedResourceListResult
}

// TopLevelClientUpdateResponse contains the response from method TopLevelClient.BeginUpdate.
type TopLevelClientUpdateResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	TopLevelTrackedResource
}
