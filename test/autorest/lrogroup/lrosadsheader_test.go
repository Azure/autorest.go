// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func newLrosaDsClient() *LROSADsClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = time.Second
	options.Transport = httpClientWithCookieJar()
	return NewLROSADsClient(&options)
}

func TestLROSADSBeginDelete202NonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginDelete202NonRetry400(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginDelete202RetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDelete202RetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginDelete204Succeeded(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginDelete204Succeeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if rt != "" {
		t.Fatal("expected an empty resume token")
	}
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROSADSBeginDeleteAsyncRelativeRetry400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginDeleteAsyncRelativeRetry400(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginDeleteAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDeleteAsyncRelativeRetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginDeleteAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginDeleteAsyncRelativeRetryInvalidJSONPolling(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginDeleteAsyncRelativeRetryNoStatus(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginDeleteAsyncRelativeRetryNoStatus(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginDeleteNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDeleteNonRetry400(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPost202NoLocation(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPost202NoLocation(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPost202NonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPost202NonRetry400(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPost202RetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPost202RetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPostAsyncRelativeRetry400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPostAsyncRelativeRetry400(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPostAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPostAsyncRelativeRetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPostAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPostAsyncRelativeRetryInvalidJSONPolling(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPostAsyncRelativeRetryNoPayload(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPostAsyncRelativeRetryNoPayload(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPostNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPostNonRetry400(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPut200InvalidJSON(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPut200InvalidJSON(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPutAsyncRelativeRetry400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutAsyncRelativeRetry400(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPutAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPutAsyncRelativeRetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPutAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutAsyncRelativeRetryInvalidJSONPolling(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPutAsyncRelativeRetryNoStatus(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutAsyncRelativeRetryNoStatus(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPutAsyncRelativeRetryNoStatusPayload(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutAsyncRelativeRetryNoStatusPayload(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPutError201NoProvisioningStatePayload(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutError201NoProvisioningStatePayload(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPutNonRetry201Creating400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutNonRetry201Creating400(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPutNonRetry201Creating400InvalidJSON(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutNonRetry201Creating400InvalidJSON(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROSADSBeginPutNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPutNonRetry400(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}
