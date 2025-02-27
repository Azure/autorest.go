// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paramgroupinggroup

// FirstParameterGroup contains a group of parameters for the ParameterGroupingClient.PostMultiParamGroups method.
type FirstParameterGroup struct {
	HeaderOne *string

	// Query parameter with default
	QueryOne *int32
}

// Grouper contains a group of parameters for the ParameterGroupingClient.GroupWithConstant method.
type Grouper struct {
	// A grouped parameter that is a constant.. Specifying any value will set the value to foo.
	GroupedConstant *string

	// Optional parameter part of a parameter grouping.
	GroupedParameter *string
}

// ParameterGroupingClientGroupWithConstantOptions contains the optional parameters for the ParameterGroupingClient.GroupWithConstant
// method.
type ParameterGroupingClientGroupWithConstantOptions struct {
	// placeholder for future optional parameters
}

// ParameterGroupingClientPostMultiParamGroupsOptions contains the optional parameters for the ParameterGroupingClient.PostMultiParamGroups
// method.
type ParameterGroupingClientPostMultiParamGroupsOptions struct {
	// placeholder for future optional parameters
}

// ParameterGroupingClientPostMultiParamGroupsSecondParamGroup contains a group of parameters for the ParameterGroupingClient.PostMultiParamGroups
// method.
type ParameterGroupingClientPostMultiParamGroupsSecondParamGroup struct {
	HeaderTwo *string

	// Query parameter with default
	QueryTwo *int32
}

// ParameterGroupingClientPostOptionalOptions contains the optional parameters for the ParameterGroupingClient.PostOptional
// method.
type ParameterGroupingClientPostOptionalOptions struct {
	// placeholder for future optional parameters
}

// ParameterGroupingClientPostOptionalParameters contains a group of parameters for the ParameterGroupingClient.PostOptional
// method.
type ParameterGroupingClientPostOptionalParameters struct {
	CustomHeader *string

	// Query parameter with default
	Query *int32
}

// ParameterGroupingClientPostRequiredOptions contains the optional parameters for the ParameterGroupingClient.PostRequired
// method.
type ParameterGroupingClientPostRequiredOptions struct {
	// placeholder for future optional parameters
}

// ParameterGroupingClientPostRequiredParameters contains a group of parameters for the ParameterGroupingClient.PostRequired
// method.
type ParameterGroupingClientPostRequiredParameters struct {
	Body         int32
	CustomHeader *string

	// Path parameter
	Path string

	// Query parameter with default
	Query *int32
}

// ParameterGroupingClientPostReservedWordsOptions contains the optional parameters for the ParameterGroupingClient.PostReservedWords
// method.
type ParameterGroupingClientPostReservedWordsOptions struct {
	// placeholder for future optional parameters
}

// ParameterGroupingClientPostReservedWordsParameters contains a group of parameters for the ParameterGroupingClient.PostReservedWords
// method.
type ParameterGroupingClientPostReservedWordsParameters struct {
	// 'accept' is a reserved word. Pass in 'yes' to pass.
	Accept *string

	// 'from' is a reserved word. Pass in 'bob' to pass.
	From *string
}

// ParameterGroupingClientPostSharedParameterGroupObjectOptions contains the optional parameters for the ParameterGroupingClient.PostSharedParameterGroupObject
// method.
type ParameterGroupingClientPostSharedParameterGroupObjectOptions struct {
	// placeholder for future optional parameters
}
