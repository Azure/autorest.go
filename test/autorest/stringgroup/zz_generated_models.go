// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"fmt"
	"net/http"
)

// ByteArrayResponse is the response envelope for operations that return a []byte type.
type ByteArrayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Value       *[]byte
}

// ColorsResponse is the response envelope for operations that return a Colors type.
type ColorsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Value       *Colors
}

// EnumGetNotExpandableOptions contains the optional parameters for the Enum.GetNotExpandable method.
type EnumGetNotExpandableOptions struct {
	// placeholder for future optional parameters
}

// EnumGetReferencedConstantOptions contains the optional parameters for the Enum.GetReferencedConstant method.
type EnumGetReferencedConstantOptions struct {
	// placeholder for future optional parameters
}

// EnumGetReferencedOptions contains the optional parameters for the Enum.GetReferenced method.
type EnumGetReferencedOptions struct {
	// placeholder for future optional parameters
}

// EnumPutNotExpandableOptions contains the optional parameters for the Enum.PutNotExpandable method.
type EnumPutNotExpandableOptions struct {
	// placeholder for future optional parameters
}

// EnumPutReferencedConstantOptions contains the optional parameters for the Enum.PutReferencedConstant method.
type EnumPutReferencedConstantOptions struct {
	// placeholder for future optional parameters
}

// EnumPutReferencedOptions contains the optional parameters for the Enum.PutReferenced method.
type EnumPutReferencedOptions struct {
	// placeholder for future optional parameters
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
func (e Error) Error() string {
	msg := ""
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if e.Status != nil {
		msg += fmt.Sprintf("Status: %v\n", *e.Status)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

type RefColorConstant struct {
	// Referenced Color Constant Description.
	ColorConstant *string `json:"ColorConstant,omitempty"`

	// Sample string.
	Field1 *string `json:"field1,omitempty"`
}

// RefColorConstantResponse is the response envelope for operations that return a RefColorConstant type.
type RefColorConstantResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse      *http.Response
	RefColorConstant *RefColorConstant
}

// StringGetBase64EncodedOptions contains the optional parameters for the String.GetBase64Encoded method.
type StringGetBase64EncodedOptions struct {
	// placeholder for future optional parameters
}

// StringGetBase64URLEncodedOptions contains the optional parameters for the String.GetBase64URLEncoded method.
type StringGetBase64URLEncodedOptions struct {
	// placeholder for future optional parameters
}

// StringGetEmptyOptions contains the optional parameters for the String.GetEmpty method.
type StringGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// StringGetMBCSOptions contains the optional parameters for the String.GetMBCS method.
type StringGetMBCSOptions struct {
	// placeholder for future optional parameters
}

// StringGetNotProvidedOptions contains the optional parameters for the String.GetNotProvided method.
type StringGetNotProvidedOptions struct {
	// placeholder for future optional parameters
}

// StringGetNullBase64URLEncodedOptions contains the optional parameters for the String.GetNullBase64URLEncoded method.
type StringGetNullBase64URLEncodedOptions struct {
	// placeholder for future optional parameters
}

// StringGetNullOptions contains the optional parameters for the String.GetNull method.
type StringGetNullOptions struct {
	// placeholder for future optional parameters
}

// StringGetWhitespaceOptions contains the optional parameters for the String.GetWhitespace method.
type StringGetWhitespaceOptions struct {
	// placeholder for future optional parameters
}

// StringPutBase64URLEncodedOptions contains the optional parameters for the String.PutBase64URLEncoded method.
type StringPutBase64URLEncodedOptions struct {
	// placeholder for future optional parameters
}

// StringPutEmptyOptions contains the optional parameters for the String.PutEmpty method.
type StringPutEmptyOptions struct {
	// placeholder for future optional parameters
}

// StringPutMBCSOptions contains the optional parameters for the String.PutMBCS method.
type StringPutMBCSOptions struct {
	// placeholder for future optional parameters
}

// StringPutNullOptions contains the optional parameters for the String.PutNull method.
type StringPutNullOptions struct {
	// string body
	StringBody *string
}

// StringPutWhitespaceOptions contains the optional parameters for the String.PutWhitespace method.
type StringPutWhitespaceOptions struct {
	// placeholder for future optional parameters
}

// StringResponse is the response envelope for operations that return a string type.
type StringResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Value       *string
}