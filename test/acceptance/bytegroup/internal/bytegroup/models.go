// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ByteArray ...
type ByteArray struct {
	azcore.Response `json:"-"`
	Value           *[]byte `json:"value,omitempty"`
}
