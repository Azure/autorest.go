// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azalias

// ListItem - Detailed information for the alias.
type ListItem struct {
	// READ-ONLY; The id for the alias.
	AliasID *string

	// READ-ONLY; The created timestamp for the alias.
	CreatedTimestamp *string

	// READ-ONLY; The id for the creator data item that this alias references (could be null if the alias has not been assigned).
	CreatorDataItemID *string

	// READ-ONLY; The timestamp of the last time the alias was assigned.
	LastUpdatedTimestamp *string
}
