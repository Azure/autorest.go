// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newLRORetrysClient() *LRORetrysClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = time.Second
	options.Transport = httpClientWithCookieJar()
	return NewLRORetrysClient(&options)
}

func TestLRORetrysBeginDelete202Retry200(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginDelete202Retry200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LRORetrysClientDelete202Retry200Poller{}
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestLRORetrysBeginDeleteAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginDeleteAsyncRelativeRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LRORetrysClientDeleteAsyncRelativeRetrySucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLRORetrysBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LRORetrysClientDeleteProvisioning202Accepted200SucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	res, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(res.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
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
	poller = &LRORetrysClientPost202Retry200Poller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLRORetrysBeginPostAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginPostAsyncRelativeRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LRORetrysClientPostAsyncRelativeRetrySucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLRORetrysBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginPut201CreatingSucceeded200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LRORetrysClientPut201CreatingSucceeded200Poller{}
	require.NoError(t, poller.Resume(rt, op))
	res, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(res.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLRORetrysBeginPutAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	poller, err := op.BeginPutAsyncRelativeRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LRORetrysClientPutAsyncRelativeRetrySucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	res, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(res.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}
