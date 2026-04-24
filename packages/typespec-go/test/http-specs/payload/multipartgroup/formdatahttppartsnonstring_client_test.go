// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipartgroup_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormDataHTTPPartsNonStringClient_Float(t *testing.T) {
	client := newClient(t)
	resp, err := client.NewMultiPartFormDataClient().NewMultiPartFormDataHTTPPartsClient().NewMultiPartFormDataHTTPPartsNonStringClient().Float(context.Background(), 0.5, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
