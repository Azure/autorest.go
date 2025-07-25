// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package valuetypesgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// ValueTypesClient - Illustrates various property types for models
// Don't use this type directly, use a constructor function instead.
type ValueTypesClient struct {
	internal *azcore.Client
	endpoint string
}

// NewValueTypesBooleanClient creates a new instance of [ValueTypesBooleanClient].
func (client *ValueTypesClient) NewValueTypesBooleanClient() *ValueTypesBooleanClient {
	return &ValueTypesBooleanClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesBooleanLiteralClient creates a new instance of [ValueTypesBooleanLiteralClient].
func (client *ValueTypesClient) NewValueTypesBooleanLiteralClient() *ValueTypesBooleanLiteralClient {
	return &ValueTypesBooleanLiteralClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesBytesClient creates a new instance of [ValueTypesBytesClient].
func (client *ValueTypesClient) NewValueTypesBytesClient() *ValueTypesBytesClient {
	return &ValueTypesBytesClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesCollectionsIntClient creates a new instance of [ValueTypesCollectionsIntClient].
func (client *ValueTypesClient) NewValueTypesCollectionsIntClient() *ValueTypesCollectionsIntClient {
	return &ValueTypesCollectionsIntClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesCollectionsModelClient creates a new instance of [ValueTypesCollectionsModelClient].
func (client *ValueTypesClient) NewValueTypesCollectionsModelClient() *ValueTypesCollectionsModelClient {
	return &ValueTypesCollectionsModelClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesCollectionsStringClient creates a new instance of [ValueTypesCollectionsStringClient].
func (client *ValueTypesClient) NewValueTypesCollectionsStringClient() *ValueTypesCollectionsStringClient {
	return &ValueTypesCollectionsStringClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesDatetimeClient creates a new instance of [ValueTypesDatetimeClient].
func (client *ValueTypesClient) NewValueTypesDatetimeClient() *ValueTypesDatetimeClient {
	return &ValueTypesDatetimeClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesDecimal128Client creates a new instance of [ValueTypesDecimal128Client].
func (client *ValueTypesClient) NewValueTypesDecimal128Client() *ValueTypesDecimal128Client {
	return &ValueTypesDecimal128Client{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesDecimalClient creates a new instance of [ValueTypesDecimalClient].
func (client *ValueTypesClient) NewValueTypesDecimalClient() *ValueTypesDecimalClient {
	return &ValueTypesDecimalClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesDictionaryStringClient creates a new instance of [ValueTypesDictionaryStringClient].
func (client *ValueTypesClient) NewValueTypesDictionaryStringClient() *ValueTypesDictionaryStringClient {
	return &ValueTypesDictionaryStringClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesDurationClient creates a new instance of [ValueTypesDurationClient].
func (client *ValueTypesClient) NewValueTypesDurationClient() *ValueTypesDurationClient {
	return &ValueTypesDurationClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesEnumClient creates a new instance of [ValueTypesEnumClient].
func (client *ValueTypesClient) NewValueTypesEnumClient() *ValueTypesEnumClient {
	return &ValueTypesEnumClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesExtensibleEnumClient creates a new instance of [ValueTypesExtensibleEnumClient].
func (client *ValueTypesClient) NewValueTypesExtensibleEnumClient() *ValueTypesExtensibleEnumClient {
	return &ValueTypesExtensibleEnumClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesFloatClient creates a new instance of [ValueTypesFloatClient].
func (client *ValueTypesClient) NewValueTypesFloatClient() *ValueTypesFloatClient {
	return &ValueTypesFloatClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesFloatLiteralClient creates a new instance of [ValueTypesFloatLiteralClient].
func (client *ValueTypesClient) NewValueTypesFloatLiteralClient() *ValueTypesFloatLiteralClient {
	return &ValueTypesFloatLiteralClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesIntClient creates a new instance of [ValueTypesIntClient].
func (client *ValueTypesClient) NewValueTypesIntClient() *ValueTypesIntClient {
	return &ValueTypesIntClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesIntLiteralClient creates a new instance of [ValueTypesIntLiteralClient].
func (client *ValueTypesClient) NewValueTypesIntLiteralClient() *ValueTypesIntLiteralClient {
	return &ValueTypesIntLiteralClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesModelClient creates a new instance of [ValueTypesModelClient].
func (client *ValueTypesClient) NewValueTypesModelClient() *ValueTypesModelClient {
	return &ValueTypesModelClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesNeverClient creates a new instance of [ValueTypesNeverClient].
func (client *ValueTypesClient) NewValueTypesNeverClient() *ValueTypesNeverClient {
	return &ValueTypesNeverClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesStringClient creates a new instance of [ValueTypesStringClient].
func (client *ValueTypesClient) NewValueTypesStringClient() *ValueTypesStringClient {
	return &ValueTypesStringClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesStringLiteralClient creates a new instance of [ValueTypesStringLiteralClient].
func (client *ValueTypesClient) NewValueTypesStringLiteralClient() *ValueTypesStringLiteralClient {
	return &ValueTypesStringLiteralClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesUnionEnumValueClient creates a new instance of [ValueTypesUnionEnumValueClient].
func (client *ValueTypesClient) NewValueTypesUnionEnumValueClient() *ValueTypesUnionEnumValueClient {
	return &ValueTypesUnionEnumValueClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesUnionFloatLiteralClient creates a new instance of [ValueTypesUnionFloatLiteralClient].
func (client *ValueTypesClient) NewValueTypesUnionFloatLiteralClient() *ValueTypesUnionFloatLiteralClient {
	return &ValueTypesUnionFloatLiteralClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesUnionIntLiteralClient creates a new instance of [ValueTypesUnionIntLiteralClient].
func (client *ValueTypesClient) NewValueTypesUnionIntLiteralClient() *ValueTypesUnionIntLiteralClient {
	return &ValueTypesUnionIntLiteralClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesUnionStringLiteralClient creates a new instance of [ValueTypesUnionStringLiteralClient].
func (client *ValueTypesClient) NewValueTypesUnionStringLiteralClient() *ValueTypesUnionStringLiteralClient {
	return &ValueTypesUnionStringLiteralClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesUnknownArrayClient creates a new instance of [ValueTypesUnknownArrayClient].
func (client *ValueTypesClient) NewValueTypesUnknownArrayClient() *ValueTypesUnknownArrayClient {
	return &ValueTypesUnknownArrayClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesUnknownDictClient creates a new instance of [ValueTypesUnknownDictClient].
func (client *ValueTypesClient) NewValueTypesUnknownDictClient() *ValueTypesUnknownDictClient {
	return &ValueTypesUnknownDictClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesUnknownIntClient creates a new instance of [ValueTypesUnknownIntClient].
func (client *ValueTypesClient) NewValueTypesUnknownIntClient() *ValueTypesUnknownIntClient {
	return &ValueTypesUnknownIntClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// NewValueTypesUnknownStringClient creates a new instance of [ValueTypesUnknownStringClient].
func (client *ValueTypesClient) NewValueTypesUnknownStringClient() *ValueTypesUnknownStringClient {
	return &ValueTypesUnknownStringClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}
