// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package basicgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// PagedUser - Paged collection of User items
type PagedUser struct {
	// REQUIRED; The User items on this page
	Value []*User

	// The link to the next page of items
	NextLink *string
}

// User - Details about a user.
type User struct {
	// REQUIRED; The user's name.
	Name *string

	// The user's order list
	Orders []*UserOrder

	// READ-ONLY; The entity tag for this resource.
	Etag *azcore.ETag

	// READ-ONLY; The user's id.
	ID *int32
}

type UserList struct {
	// REQUIRED
	Users []*User
}

// UserOrder for testing list with expand.
type UserOrder struct {
	// REQUIRED; The user's order detail
	Detail *string

	// REQUIRED; The user's id.
	UserID *int32

	// READ-ONLY; The user's id.
	ID *int32
}
