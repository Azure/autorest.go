// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package xmlgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// XMLClient - Sends and receives bodies in XML format.
// Don't use this type directly, use a constructor function instead.
type XMLClient struct {
	internal *azcore.Client
}

// NewXMLModelWithArrayOfModelValueClient creates a new instance of [XMLModelWithArrayOfModelValueClient].
func (client *XMLClient) NewXMLModelWithArrayOfModelValueClient() *XMLModelWithArrayOfModelValueClient {
	return &XMLModelWithArrayOfModelValueClient{
		internal: client.internal,
	}
}

// NewXMLModelWithAttributesValueClient creates a new instance of [XMLModelWithAttributesValueClient].
func (client *XMLClient) NewXMLModelWithAttributesValueClient() *XMLModelWithAttributesValueClient {
	return &XMLModelWithAttributesValueClient{
		internal: client.internal,
	}
}

// NewXMLModelWithDictionaryValueClient creates a new instance of [XMLModelWithDictionaryValueClient].
func (client *XMLClient) NewXMLModelWithDictionaryValueClient() *XMLModelWithDictionaryValueClient {
	return &XMLModelWithDictionaryValueClient{
		internal: client.internal,
	}
}

// NewXMLModelWithEmptyArrayValueClient creates a new instance of [XMLModelWithEmptyArrayValueClient].
func (client *XMLClient) NewXMLModelWithEmptyArrayValueClient() *XMLModelWithEmptyArrayValueClient {
	return &XMLModelWithEmptyArrayValueClient{
		internal: client.internal,
	}
}

// NewXMLModelWithEncodedNamesValueClient creates a new instance of [XMLModelWithEncodedNamesValueClient].
func (client *XMLClient) NewXMLModelWithEncodedNamesValueClient() *XMLModelWithEncodedNamesValueClient {
	return &XMLModelWithEncodedNamesValueClient{
		internal: client.internal,
	}
}

// NewXMLModelWithOptionalFieldValueClient creates a new instance of [XMLModelWithOptionalFieldValueClient].
func (client *XMLClient) NewXMLModelWithOptionalFieldValueClient() *XMLModelWithOptionalFieldValueClient {
	return &XMLModelWithOptionalFieldValueClient{
		internal: client.internal,
	}
}

// NewXMLModelWithRenamedArraysValueClient creates a new instance of [XMLModelWithRenamedArraysValueClient].
func (client *XMLClient) NewXMLModelWithRenamedArraysValueClient() *XMLModelWithRenamedArraysValueClient {
	return &XMLModelWithRenamedArraysValueClient{
		internal: client.internal,
	}
}

// NewXMLModelWithRenamedFieldsValueClient creates a new instance of [XMLModelWithRenamedFieldsValueClient].
func (client *XMLClient) NewXMLModelWithRenamedFieldsValueClient() *XMLModelWithRenamedFieldsValueClient {
	return &XMLModelWithRenamedFieldsValueClient{
		internal: client.internal,
	}
}

// NewXMLModelWithSimpleArraysValueClient creates a new instance of [XMLModelWithSimpleArraysValueClient].
func (client *XMLClient) NewXMLModelWithSimpleArraysValueClient() *XMLModelWithSimpleArraysValueClient {
	return &XMLModelWithSimpleArraysValueClient{
		internal: client.internal,
	}
}

// NewXMLModelWithTextValueClient creates a new instance of [XMLModelWithTextValueClient].
func (client *XMLClient) NewXMLModelWithTextValueClient() *XMLModelWithTextValueClient {
	return &XMLModelWithTextValueClient{
		internal: client.internal,
	}
}

// NewXMLModelWithUnwrappedArrayValueClient creates a new instance of [XMLModelWithUnwrappedArrayValueClient].
func (client *XMLClient) NewXMLModelWithUnwrappedArrayValueClient() *XMLModelWithUnwrappedArrayValueClient {
	return &XMLModelWithUnwrappedArrayValueClient{
		internal: client.internal,
	}
}

// NewXMLSimpleModelValueClient creates a new instance of [XMLSimpleModelValueClient].
func (client *XMLClient) NewXMLSimpleModelValueClient() *XMLSimpleModelValueClient {
	return &XMLSimpleModelValueClient{
		internal: client.internal,
	}
}