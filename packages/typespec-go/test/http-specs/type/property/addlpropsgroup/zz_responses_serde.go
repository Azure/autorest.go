// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package addlpropsgroup

// UnmarshalJSON implements the json.Unmarshaller interface for type AdditionalPropertiesExtendsUnknownDiscriminatedClientGetResponse.
func (a *AdditionalPropertiesExtendsUnknownDiscriminatedClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalExtendsUnknownAdditionalPropertiesDiscriminatedClassification(data)
	if err != nil {
		return err
	}
	a.ExtendsUnknownAdditionalPropertiesDiscriminatedClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AdditionalPropertiesIsUnknownDiscriminatedClientGetResponse.
func (a *AdditionalPropertiesIsUnknownDiscriminatedClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalIsUnknownAdditionalPropertiesDiscriminatedClassification(data)
	if err != nil {
		return err
	}
	a.IsUnknownAdditionalPropertiesDiscriminatedClassification = res
	return nil
}