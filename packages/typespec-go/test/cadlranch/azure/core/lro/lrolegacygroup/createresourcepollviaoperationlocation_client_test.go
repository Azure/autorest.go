// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrolegacygroup_test

import (
	"context"
	"lrolegacygroup"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestCreateResourcePollViaOperationLocationClient_BeginCreateJob(t *testing.T) {
	client, err := lrolegacygroup.NewLegacyClient(nil)
	require.NoError(t, err)
	poller, err := client.NewLegacyCreateResourcePollViaOperationLocationClient().BeginCreateJob(context.Background(), lrolegacygroup.JobData{
		Comment: to.Ptr("async job"),
	}, nil)
	require.NoError(t, err)
	resp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: time.Second,
	})
	require.NoError(t, err)
	require.Equal(t, lrolegacygroup.JobResult{
		Comment: to.Ptr("async job"),
		JobID:   to.Ptr("job1"),
		Status:  to.Ptr(lrolegacygroup.JobStatusSucceeded),
		Results: []*string{
			to.Ptr("job1 result"),
		},
	}, resp.JobResult)
}

func TestCreateResourcePollViaOperationLocationClient_GetJob(t *testing.T) {
	client, err := lrolegacygroup.NewLegacyClient(nil)
	require.NoError(t, err)
	resp, err := client.NewLegacyCreateResourcePollViaOperationLocationClient().GetJob(context.Background(), "job1", nil)
	require.NoError(t, err)
	require.Equal(t, lrolegacygroup.JobResult{
		Comment: to.Ptr("async job"),
		JobID:   to.Ptr("job1"),
		Status:  to.Ptr(lrolegacygroup.JobStatusSucceeded),
		Results: []*string{
			to.Ptr("job1 result"),
		},
	}, resp.JobResult)
}
