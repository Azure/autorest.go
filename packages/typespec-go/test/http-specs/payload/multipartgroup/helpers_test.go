// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipartgroup_test

import (
	"multipartgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	jpgPath = "../../../../node_modules/@typespec/http-specs/assets/image.jpg"
	pngPath = "../../../../node_modules/@typespec/http-specs/assets/image.png"
)

func newClient(t *testing.T) *multipartgroup.MultiPartClient {
	t.Helper()
	client, err := multipartgroup.NewMultiPartClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	return client
}
