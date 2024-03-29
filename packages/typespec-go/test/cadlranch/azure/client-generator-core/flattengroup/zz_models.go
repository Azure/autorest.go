// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package flattengroup

// This is the child model to be flattened. And it has flattened property as well.
type ChildFlattenModel struct {
	// REQUIRED
	Properties *ChildModel

	// REQUIRED
	Summary *string
}

// This is the child model to be flattened.
type ChildModel struct {
	// REQUIRED
	Age *int32

	// REQUIRED
	Description *string
}

// This is the model with one level of flattening.
type FlattenModel struct {
	// REQUIRED
	Name *string

	// REQUIRED
	Properties *ChildModel
}

// This is the model with two levels of flattening.
type NestedFlattenModel struct {
	// REQUIRED
	Name *string

	// REQUIRED
	Properties *ChildFlattenModel
}
