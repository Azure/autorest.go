// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdevopsinfrastructure

// ImageVersionsClientListByImageResponse contains the response from method ImageVersionsClient.NewListByImagePager.
type ImageVersionsClientListByImageResponse struct {
	// The response of a ImageVersion list operation.
	ImageVersionListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	PagedOperation
}

// PoolsClientCreateOrUpdateResponse contains the response from method PoolsClient.BeginCreateOrUpdate.
type PoolsClientCreateOrUpdateResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	Pool
}

// PoolsClientDeleteResponse contains the response from method PoolsClient.BeginDelete.
type PoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// PoolsClientGetResponse contains the response from method PoolsClient.Get.
type PoolsClientGetResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	Pool
}

// PoolsClientListByResourceGroupResponse contains the response from method PoolsClient.NewListByResourceGroupPager.
type PoolsClientListByResourceGroupResponse struct {
	// The response of a Pool list operation.
	PoolListResult
}

// PoolsClientListBySubscriptionResponse contains the response from method PoolsClient.NewListBySubscriptionPager.
type PoolsClientListBySubscriptionResponse struct {
	// The response of a Pool list operation.
	PoolListResult
}

// PoolsClientUpdateResponse contains the response from method PoolsClient.BeginUpdate.
type PoolsClientUpdateResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	Pool
}

// ResourceDetailsClientListByPoolResponse contains the response from method ResourceDetailsClient.NewListByPoolPager.
type ResourceDetailsClientListByPoolResponse struct {
	// The response of a ResourceDetailsObject list operation.
	ResourceDetailsObjectListResult
}

// SKUClientListByLocationResponse contains the response from method SKUClient.NewListByLocationPager.
type SKUClientListByLocationResponse struct {
	// The response of a ResourceSku list operation.
	ResourceSKUListResult
}

// SubscriptionUsagesClientListByLocationResponse contains the response from method SubscriptionUsagesClient.NewListByLocationPager.
type SubscriptionUsagesClientListByLocationResponse struct {
	// The response of a Quota list operation.
	QuotaListResult
}