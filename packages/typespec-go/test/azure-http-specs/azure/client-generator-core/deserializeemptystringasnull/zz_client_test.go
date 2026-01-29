// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package deserializeemptystringasnull_test

import (
	"context"
	"deserializeemptystringasnull"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	client, err := deserializeemptystringasnull.NewClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, "", *resp.SampleURL)
}
