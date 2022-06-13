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
	"github.com/stretchr/testify/require"
)

func newLrosaDsClient() *LROSADsClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = time.Second
	options.Transport = httpClientWithCookieJar()
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &options)
	return NewLROSADsClient(pl)
}

func TestLROSADSBeginDelete202NonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginDelete202NonRetry400(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDelete202NonRetry400(context.Background(), &LROSADsClientBeginDelete202NonRetry400Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginDelete202RetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDelete202RetryInvalidHeader(context.Background(), nil)
	require.Error(t, err)
}

func TestLROSADSBeginDelete204Succeeded(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginDelete204Succeeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.Error(t, err)
	require.Empty(t, rt)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROSADSBeginDeleteAsyncRelativeRetry400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginDeleteAsyncRelativeRetry400(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteAsyncRelativeRetry400(context.Background(), &LROSADsClientBeginDeleteAsyncRelativeRetry400Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginDeleteAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDeleteAsyncRelativeRetryInvalidHeader(context.Background(), nil)
	require.Error(t, err)
}

func TestLROSADSBeginDeleteAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginDeleteAsyncRelativeRetryInvalidJSONPolling(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteAsyncRelativeRetryInvalidJSONPolling(context.Background(), &LROSADsClientBeginDeleteAsyncRelativeRetryInvalidJSONPollingOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginDeleteAsyncRelativeRetryNoStatus(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginDeleteAsyncRelativeRetryNoStatus(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteAsyncRelativeRetryNoStatus(context.Background(), &LROSADsClientBeginDeleteAsyncRelativeRetryNoStatusOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginDeleteNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDeleteNonRetry400(context.Background(), nil)
	require.Error(t, err)
}

func TestLROSADSBeginPost202NoLocation(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPost202NoLocation(context.Background(), nil)
	require.Error(t, err)
}

func TestLROSADSBeginPost202NonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPost202NonRetry400(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPost202NonRetry400(context.Background(), &LROSADsClientBeginPost202NonRetry400Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPost202RetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPost202RetryInvalidHeader(context.Background(), nil)
	require.Error(t, err)
}

func TestLROSADSBeginPostAsyncRelativeRetry400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPostAsyncRelativeRetry400(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostAsyncRelativeRetry400(context.Background(), &LROSADsClientBeginPostAsyncRelativeRetry400Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPostAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPostAsyncRelativeRetryInvalidHeader(context.Background(), nil)
	require.Error(t, err)
}

func TestLROSADSBeginPostAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPostAsyncRelativeRetryInvalidJSONPolling(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostAsyncRelativeRetryInvalidJSONPolling(context.Background(), &LROSADsClientBeginPostAsyncRelativeRetryInvalidJSONPollingOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPostAsyncRelativeRetryNoPayload(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPostAsyncRelativeRetryNoPayload(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostAsyncRelativeRetryNoPayload(context.Background(), &LROSADsClientBeginPostAsyncRelativeRetryNoPayloadOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPostNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPostNonRetry400(context.Background(), nil)
	require.Error(t, err)
}

func TestLROSADSBeginPut200InvalidJSON(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPut200InvalidJSON(context.Background(), Product{}, nil)
	require.Error(t, err)
}

func TestLROSADSBeginPutAsyncRelativeRetry400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutAsyncRelativeRetry400(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncRelativeRetry400(context.Background(), Product{}, &LROSADsClientBeginPutAsyncRelativeRetry400Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPutAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPutAsyncRelativeRetryInvalidHeader(context.Background(), Product{}, nil)
	require.Error(t, err)
}

func TestLROSADSBeginPutAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutAsyncRelativeRetryInvalidJSONPolling(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncRelativeRetryInvalidJSONPolling(context.Background(), Product{}, &LROSADsClientBeginPutAsyncRelativeRetryInvalidJSONPollingOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPutAsyncRelativeRetryNoStatus(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutAsyncRelativeRetryNoStatus(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncRelativeRetryNoStatus(context.Background(), Product{}, &LROSADsClientBeginPutAsyncRelativeRetryNoStatusOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPutAsyncRelativeRetryNoStatusPayload(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutAsyncRelativeRetryNoStatusPayload(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncRelativeRetryNoStatusPayload(context.Background(), Product{}, &LROSADsClientBeginPutAsyncRelativeRetryNoStatusPayloadOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPutError201NoProvisioningStatePayload(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutError201NoProvisioningStatePayload(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutError201NoProvisioningStatePayload(context.Background(), Product{}, &LROSADsClientBeginPutError201NoProvisioningStatePayloadOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPutNonRetry201Creating400(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutNonRetry201Creating400(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutNonRetry201Creating400(context.Background(), Product{}, &LROSADsClientBeginPutNonRetry201Creating400Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPutNonRetry201Creating400InvalidJSON(t *testing.T) {
	op := newLrosaDsClient()
	poller, err := op.BeginPutNonRetry201Creating400InvalidJSON(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutNonRetry201Creating400InvalidJSON(context.Background(), Product{}, &LROSADsClientBeginPutNonRetry201Creating400InvalidJSONOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROSADSBeginPutNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPutNonRetry400(context.Background(), Product{}, nil)
	require.Error(t, err)
}
