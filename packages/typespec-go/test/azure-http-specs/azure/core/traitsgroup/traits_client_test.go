// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package traitsgroup_test

import (
	"context"
	"testing"
	"time"
	"traitsgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestTraitsClient_RepeatableAction(t *testing.T) {
	client, err := traitsgroup.NewTraitsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	resp, err := client.RepeatableAction(context.Background(), 1, traitsgroup.UserActionParam{
		UserActionValue: to.Ptr("test"),
	}, &traitsgroup.TraitsClientRepeatableActionOptions{
		RepeatabilityFirstSent: to.Ptr(time.Date(2023, time.November, 27, 11, 58, 0, 0, time.UTC)),
		RepeatabilityRequestID: to.Ptr("86aede1f-96fa-4e7f-b1e1-bf8a947cb804"),
	})
	require.NoError(t, err)
	require.NotNil(t, resp.UserActionResult)
	require.Equal(t, "test", *resp.UserActionResult)
}

func TestTraitsClient_SmokeTest(t *testing.T) {
	client, err := traitsgroup.NewTraitsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	require.NoError(t, err)
	const id = int32(1)
	resp, err := client.SmokeTest(context.Background(), id, "123", &traitsgroup.TraitsClientSmokeTestOptions{
		IfMatch:           to.Ptr("\"valid\""),
		IfNoneMatch:       to.Ptr("\"invalid\""),
		IfModifiedSince:   to.Ptr(time.Date(2021, time.August, 26, 14, 38, 0, 0, time.UTC)),
		IfUnmodifiedSince: to.Ptr(time.Date(2022, time.August, 26, 14, 38, 0, 0, time.UTC)),
		ClientRequestID:   to.Ptr("86aede1f-96fa-4e7f-b1e1-bf8a947cb804"),
	})
	require.NoError(t, err)
	require.NotNil(t, resp.ID)
	require.NotNil(t, resp.Name)
	require.Equal(t, id, *resp.ID)
	require.Equal(t, "Madge", *resp.Name)
}
