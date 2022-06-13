// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"generatortests"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newLRORetrysClient() *LRORetrysClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = time.Second
	options.Transport = httpClientWithCookieJar()
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &options)
	return NewLRORetrysClient(pl)
}

func TestLRORetrysBeginDelete202Retry200(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginDelete202Retry200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDelete202Retry200(context.Background(), &LRORetrysClientBeginDelete202Retry200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	result, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestLRORetrysBeginDeleteAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginDeleteAsyncRelativeRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteAsyncRelativeRetrySucceeded(context.Background(), &LRORetrysClientBeginDeleteAsyncRelativeRetrySucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLRORetrysBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), &LRORetrysClientBeginDeleteProvisioning202Accepted200SucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	res, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	if r := cmp.Diff(res.Product, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLRORetrysBeginPost202Retry200(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginPost202Retry200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPost202Retry200(context.Background(), &LRORetrysClientBeginPost202Retry200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLRORetrysBeginPostAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginPostAsyncRelativeRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostAsyncRelativeRetrySucceeded(context.Background(), &LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLRORetrysBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginPut201CreatingSucceeded200(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPut201CreatingSucceeded200(context.Background(), Product{}, &LRORetrysClientBeginPut201CreatingSucceeded200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	res, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	if r := cmp.Diff(res.Product, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLRORetrysBeginPutAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginPutAsyncRelativeRetrySucceeded(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncRelativeRetrySucceeded(context.Background(), Product{}, &LRORetrysClientBeginPutAsyncRelativeRetrySucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	res, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	if r := cmp.Diff(res.Product, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}
