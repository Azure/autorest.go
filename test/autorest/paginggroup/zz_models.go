//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package paginggroup

// CustomParameterGroup contains a group of parameters for the PagingClient.GetMultiplePagesFragmentWithGroupingNextLink method.
type CustomParameterGroup struct {
	// Sets the api version to use.
	APIVersion string
	// Sets the tenant to use.
	Tenant string
}

type ODataProductResult struct {
	ODataNextLink *string    `json:"odata.nextLink,omitempty"`
	Values        []*Product `json:"values,omitempty"`
}

// PagingClientBeginGetMultiplePagesLROOptions contains the optional parameters for the PagingClient.BeginGetMultiplePagesLRO
// method.
type PagingClientBeginGetMultiplePagesLROOptions struct {
	ClientRequestID *string
	// Sets the maximum number of items to return in the response.
	Maxresults *int32
	// Resumes the LRO from the provided token.
	ResumeToken string
	// Sets the maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.
	Timeout *int32
}

// PagingClientDuplicateParamsOptions contains the optional parameters for the PagingClient.DuplicateParams method.
type PagingClientDuplicateParamsOptions struct {
	// OData filter options. Pass in 'foo'
	Filter *string
}

// PagingClientFirstResponseEmptyOptions contains the optional parameters for the PagingClient.FirstResponseEmpty method.
type PagingClientFirstResponseEmptyOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetMultiplePagesFailureOptions contains the optional parameters for the PagingClient.GetMultiplePagesFailure
// method.
type PagingClientGetMultiplePagesFailureOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetMultiplePagesFailureURIOptions contains the optional parameters for the PagingClient.GetMultiplePagesFailureURI
// method.
type PagingClientGetMultiplePagesFailureURIOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetMultiplePagesFragmentNextLinkOptions contains the optional parameters for the PagingClient.GetMultiplePagesFragmentNextLink
// method.
type PagingClientGetMultiplePagesFragmentNextLinkOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetMultiplePagesFragmentWithGroupingNextLinkOptions contains the optional parameters for the PagingClient.GetMultiplePagesFragmentWithGroupingNextLink
// method.
type PagingClientGetMultiplePagesFragmentWithGroupingNextLinkOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetMultiplePagesOptions contains the optional parameters for the PagingClient.GetMultiplePages method.
type PagingClientGetMultiplePagesOptions struct {
	ClientRequestID *string
	// Sets the maximum number of items to return in the response.
	Maxresults *int32
	// Sets the maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.
	Timeout *int32
}

// PagingClientGetMultiplePagesRetryFirstOptions contains the optional parameters for the PagingClient.GetMultiplePagesRetryFirst
// method.
type PagingClientGetMultiplePagesRetryFirstOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetMultiplePagesRetrySecondOptions contains the optional parameters for the PagingClient.GetMultiplePagesRetrySecond
// method.
type PagingClientGetMultiplePagesRetrySecondOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetMultiplePagesWithOffsetOptions contains the optional parameters for the PagingClient.GetMultiplePagesWithOffset
// method.
type PagingClientGetMultiplePagesWithOffsetOptions struct {
	ClientRequestID *string
	// Sets the maximum number of items to return in the response.
	Maxresults *int32
	// Offset of return value
	Offset int32
	// Sets the maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.
	Timeout *int32
}

// PagingClientGetNoItemNamePagesOptions contains the optional parameters for the PagingClient.GetNoItemNamePages method.
type PagingClientGetNoItemNamePagesOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetNullNextLinkNamePagesOptions contains the optional parameters for the PagingClient.GetNullNextLinkNamePages
// method.
type PagingClientGetNullNextLinkNamePagesOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetODataMultiplePagesOptions contains the optional parameters for the PagingClient.GetODataMultiplePages method.
type PagingClientGetODataMultiplePagesOptions struct {
	ClientRequestID *string
	// Sets the maximum number of items to return in the response.
	Maxresults *int32
	// Sets the maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.
	Timeout *int32
}

// PagingClientGetPagingModelWithItemNameWithXMSClientNameOptions contains the optional parameters for the PagingClient.GetPagingModelWithItemNameWithXMSClientName
// method.
type PagingClientGetPagingModelWithItemNameWithXMSClientNameOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetSinglePagesFailureOptions contains the optional parameters for the PagingClient.GetSinglePagesFailure method.
type PagingClientGetSinglePagesFailureOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetSinglePagesOptions contains the optional parameters for the PagingClient.GetSinglePages method.
type PagingClientGetSinglePagesOptions struct {
	// placeholder for future optional parameters
}

// PagingClientGetWithQueryParamsOptions contains the optional parameters for the PagingClient.GetWithQueryParams method.
type PagingClientGetWithQueryParamsOptions struct {
	// placeholder for future optional parameters
}

type Product struct {
	Properties *ProductProperties `json:"properties,omitempty"`
}

type ProductProperties struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ProductResult struct {
	NextLink *string    `json:"nextLink,omitempty"`
	Values   []*Product `json:"values,omitempty"`
}

type ProductResultValue struct {
	NextLink *string    `json:"nextLink,omitempty"`
	Value    []*Product `json:"value,omitempty"`
}

type ProductResultValueWithXMSClientName struct {
	Indexes  []*Product `json:"values,omitempty"`
	NextLink *string    `json:"nextLink,omitempty"`
}