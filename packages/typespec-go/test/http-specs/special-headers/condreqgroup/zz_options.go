// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package condreqgroup

import "time"

// ConditionalRequestClientHeadIfModifiedSinceOptions contains the optional parameters for the ConditionalRequestClient.HeadIfModifiedSince
// method.
type ConditionalRequestClientHeadIfModifiedSinceOptions struct {
	// A timestamp indicating the last modified time of the resource known to the
	// client. The operation will be performed only if the resource on the service has
	// been modified since the specified time.
	IfModifiedSince *time.Time
}

// ConditionalRequestClientPostIfMatchOptions contains the optional parameters for the ConditionalRequestClient.PostIfMatch
// method.
type ConditionalRequestClientPostIfMatchOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string
}

// ConditionalRequestClientPostIfNoneMatchOptions contains the optional parameters for the ConditionalRequestClient.PostIfNoneMatch
// method.
type ConditionalRequestClientPostIfNoneMatchOptions struct {
	// The request should only proceed if no entity matches this string.
	IfNoneMatch *string
}

// ConditionalRequestClientPostIfUnmodifiedSinceOptions contains the optional parameters for the ConditionalRequestClient.PostIfUnmodifiedSince
// method.
type ConditionalRequestClientPostIfUnmodifiedSinceOptions struct {
	// A timestamp indicating the last modified time of the resource known to the
	// client. The operation will be performed only if the resource on the service has
	// not been modified since the specified time.
	IfUnmodifiedSince *time.Time
}