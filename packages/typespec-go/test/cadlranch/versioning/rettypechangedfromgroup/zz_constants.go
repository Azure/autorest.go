// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package rettypechangedfromgroup

// Versions - The version of the API.
type Versions string

const (
	// VersionsV1 - The version v1.
	VersionsV1 Versions = "v1"
	// VersionsV2 - The version v2.
	VersionsV2 Versions = "v2"
)

// PossibleVersionsValues returns the possible values for the Versions const type.
func PossibleVersionsValues() []Versions {
	return []Versions{
		VersionsV1,
		VersionsV2,
	}
}
