// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newLROSClient() *LROsClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = time.Second
	options.Transport = httpClientWithCookieJar()
	return NewLROsClient(&options)
}

func httpClientWithCookieJar() policy.Transporter {
	j, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = j
	return http.DefaultClient
}

func TestLROResumeWrongPoller(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp2 := LROsClientDelete202Retry200Poller{}
	require.Error(t, resp2.Resume(rt, op))
}

func TestLROBeginDelete202NoRetry204(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientDelete202NoRetry204Poller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDelete202Retry200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDelete202Retry200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientDelete202Retry200Poller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDelete204Succeeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDelete204Succeeded(context.Background(), nil)
	require.NoError(t, err)
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteAsyncNoHeaderInRetry(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientDeleteAsyncNoHeaderInRetryPoller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteAsyncNoRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientDeleteAsyncNoRetrySucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteAsyncRetryFailed(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientDeleteAsyncRetryFailedPoller{}
	require.NoError(t, poller.Resume(rt, op))
	res, err := poller.PollUntilDone(context.Background(), time.Second)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatal("expected azcore.ResponseError")
	} else if sc := respErr.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginDeleteAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteAsyncRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientDeleteAsyncRetrySucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteAsyncRetrycanceled(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientDeleteAsyncRetrycanceledPoller{}
	require.NoError(t, poller.Resume(rt, op))
	res, err := poller.PollUntilDone(context.Background(), time.Second)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatal("expected azcore.ResponseError")
	} else if sc := respErr.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginDeleteNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteNoHeaderInRetry(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientDeleteNoHeaderInRetryPoller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientDeleteProvisioning202Accepted200SucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteProvisioning202DeletingFailed200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteProvisioning202DeletingFailed200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROBeginDeleteProvisioning202Deletingcanceled200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteProvisioning202Deletingcanceled200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	result, err := poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestLROBeginPost200WithPayload(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPost200WithPayload(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPost200WithPayloadPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.SKU, SKU{
		ID:   to.StringPtr("1"),
		Name: to.StringPtr("product"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPost202List(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPost202List(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPost202ListPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.ProductArray, []*Product{
		{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPost202NoRetry204(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPost202NoRetry204(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPost202NoRetry204Poller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginPost202Retry200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPost202Retry200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPost202Retry200Poller{}
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginPostAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostAsyncNoRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPostAsyncNoRetrySucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostAsyncRetryFailed(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPostAsyncRetryFailedPoller{}
	require.NoError(t, poller.Resume(rt, op))
	res, err := poller.PollUntilDone(context.Background(), time.Second)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatal("expected azcore.ResponseError")
	} else if sc := respErr.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPostAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostAsyncRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPostAsyncRetrySucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostAsyncRetrycanceled(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPostAsyncRetrycanceledPoller{}
	require.NoError(t, poller.Resume(rt, op))
	res, err := poller.PollUntilDone(context.Background(), time.Second)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatal("expected azcore.ResponseError")
	} else if sc := respErr.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPostDoubleHeadersFinalAzureHeaderGet(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostDoubleHeadersFinalAzureHeaderGet(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPostDoubleHeadersFinalAzureHeaderGetPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID: to.StringPtr("100"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostDoubleHeadersFinalAzureHeaderGetDefault(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostDoubleHeadersFinalAzureHeaderGetDefault(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPostDoubleHeadersFinalAzureHeaderGetDefaultPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostDoubleHeadersFinalLocationGet(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostDoubleHeadersFinalLocationGet(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPostDoubleHeadersFinalLocationGetPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut200Acceptedcanceled200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut200Acceptedcanceled200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatal("expected azcore.ResponseError")
	} else if sc := respErr.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPut200Succeeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut200Succeeded(context.Background(), nil)
	require.NoError(t, err)
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("Expected an error but did not receive one")
	}
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut200SucceededNoState(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut200SucceededNoState(context.Background(), nil)
	require.NoError(t, err)
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("Expected an error but did not receive one")
	}
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}); r != "" {
		t.Fatal(r)
	}
}

// TODO check if this test should actually be returning a 200 or a 204
func TestLROBeginPut200UpdatingSucceeded204(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut200UpdatingSucceeded204(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPut200UpdatingSucceeded204Poller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut201CreatingFailed200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut201CreatingFailed200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	require.NoError(t, poller.Resume(rt, op))
	_, err = poller.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatal("expected azcore.ResponseError")
	} else if sc := respErr.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut201CreatingSucceeded200(context.Background(), &LROsClientBeginPut201CreatingSucceeded200Options{Product: &Product{}})
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPut201CreatingSucceeded200Poller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut202Retry200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut202Retry200(context.Background(), &LROsClientBeginPut202Retry200Options{Product: &Product{}})
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPut202Retry200Poller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncNoHeaderInRetry(context.Background(), &LROsClientBeginPutAsyncNoHeaderInRetryOptions{Product: &Product{}})
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPutAsyncNoHeaderInRetryPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncNoRetrySucceeded(context.Background(), &LROsClientBeginPutAsyncNoRetrySucceededOptions{Product: &Product{}})
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPutAsyncNoRetrySucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncNoRetrycanceled(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncNoRetrycanceled(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPutAsyncNoRetrycanceledPoller{}
	require.NoError(t, poller.Resume(rt, op))
	res, err := poller.PollUntilDone(context.Background(), time.Second)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatal("expected azcore.ResponseError")
	} else if sc := respErr.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPutAsyncNonResource(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncNonResource(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPutAsyncNonResourcePoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.SKU, SKU{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("sku"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncRetryFailed(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPutAsyncRetryFailedPoller{}
	require.NoError(t, poller.Resume(rt, op))
	res, err := poller.PollUntilDone(context.Background(), time.Second)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatal("expected azcore.ResponseError")
	} else if sc := respErr.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPutAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPutAsyncRetrySucceededPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncSubResource(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncSubResource(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPutAsyncSubResourcePoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.SubProduct, SubProduct{
		ID: to.StringPtr("100"),
		Properties: &SubProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutNoHeaderInRetry(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	poller, err := op.BeginPutNoHeaderInRetry(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPutNoHeaderInRetryPoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.Product, Product{
		ID: to.StringPtr("100"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutNonResource(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	poller, err := op.BeginPutNonResource(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPutNonResourcePoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.SKU, SKU{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("sku"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutSubResource(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	poller, err := op.BeginPutSubResource(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller = &LROsClientPutSubResourcePoller{}
	require.NoError(t, poller.Resume(rt, op))
	pollResp, err := poller.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	if r := cmp.Diff(pollResp.SubProduct, SubProduct{
		ID: to.StringPtr("100"),
		Properties: &SubProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}
