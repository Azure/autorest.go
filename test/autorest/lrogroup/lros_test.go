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
	resp, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp2 := LROsClientDelete202Retry200PollerResponse{}
	if err = resp2.Resume(context.Background(), op, rt); err == nil {
		t.Fatal("expected an error but did not find receive one")
	}
}

func TestLROBeginDelete202NoRetry204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientDelete202NoRetry204PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDelete202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202Retry200(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientDelete202Retry200PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDelete204Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete204Succeeded(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	_, err = poller.ResumeToken()
	require.Error(t, err)
	_, err = resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncNoHeaderInRetry(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientDeleteAsyncNoHeaderInRetryPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncNoRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientDeleteAsyncNoRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetryFailed(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientDeleteAsyncRetryFailedPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
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
	resp, err := op.BeginDeleteAsyncRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientDeleteAsyncRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetrycanceled(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientDeleteAsyncRetrycanceledPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
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
	resp, err := op.BeginDeleteNoHeaderInRetry(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientDeleteNoHeaderInRetryPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientDeleteProvisioning202Accepted200SucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginDeleteProvisioning202DeletingFailed200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202DeletingFailed200(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	if err = resp.Resume(context.Background(), op, rt); err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROBeginDeleteProvisioning202Deletingcanceled200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202Deletingcanceled200(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	if err = resp.Resume(context.Background(), op, rt); err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROBeginPatch201RetryWithAsyncHeader(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPatch201RetryWithAsyncHeader(context.Background(), nil)
	require.NoError(t, err)
	res, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("/lro/patch/201/retry/onlyAsyncHeader"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, res.Product)
}

func TestLROBeginPatch202RetryWithAsyncAndLocationHeader(t *testing.T) {
	t.Skip("https://github.com/Azure/autorest.testserver/pull/369")
	op := newLROSClient()
	resp, err := op.BeginPatch202RetryWithAsyncAndLocationHeader(context.Background(), nil)
	require.NoError(t, err)
	res, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("/lro/patch/202/retry/asyncAndLocationHeader"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, res.Product)
}

func TestLROBeginPost200WithPayload(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost200WithPayload(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPost200WithPayloadPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, SKU{
		ID:   to.StringPtr("1"),
		Name: to.StringPtr("product"),
	}, pollResp.SKU)
}

func TestLROBeginPost202List(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202List(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPost202ListPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, []*Product{{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}}, pollResp.ProductArray)
}

func TestLROBeginPost202NoRetry204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202NoRetry204(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPost202NoRetry204PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginPost202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202Retry200(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPost202Retry200PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
}

func TestLROBeginPostAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncNoRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPostAsyncNoRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPostAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetryFailed(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPostAsyncRetryFailedPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
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
	resp, err := op.BeginPostAsyncRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPostAsyncRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPostAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetrycanceled(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPostAsyncRetrycanceledPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
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
	resp, err := op.BeginPostDoubleHeadersFinalAzureHeaderGet(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPostDoubleHeadersFinalAzureHeaderGetPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID: to.StringPtr("100"),
	}, pollResp.Product)
}

func TestLROBeginPostDoubleHeadersFinalAzureHeaderGetDefault(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalAzureHeaderGetDefault(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPostDoubleHeadersFinalAzureHeaderGetDefaultPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}, pollResp.Product)
}

func TestLROBeginPostDoubleHeadersFinalLocationGet(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalLocationGet(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPostDoubleHeadersFinalLocationGetPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}, pollResp.Product)
}

func TestLROBeginPut200Acceptedcanceled200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200Acceptedcanceled200(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	if err = resp.Resume(context.Background(), op, rt); err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatal("expected azcore.ResponseError")
	} else if sc := respErr.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPut200Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200Succeeded(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	_, err = poller.ResumeToken()
	require.Error(t, err)
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPut200SucceededNoState(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200SucceededNoState(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	_, err = poller.ResumeToken()
	require.Error(t, err)
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}, pollResp.Product)
}

// TODO check if this test should actually be returning a 200 or a 204
func TestLROBeginPut200UpdatingSucceeded204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200UpdatingSucceeded204(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPut200UpdatingSucceeded204PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPut201CreatingFailed200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut201CreatingFailed200(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	if err = resp.Resume(context.Background(), op, rt); err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	var respErr *azcore.ResponseError
	if !errors.As(err, &respErr) {
		t.Fatal("expected azcore.ResponseError")
	} else if sc := respErr.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut201CreatingSucceeded200(context.Background(), &LROsClientBeginPut201CreatingSucceeded200Options{Product: &Product{}})
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPut201CreatingSucceeded200PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPut201Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut201Succeeded(context.Background(), nil)
	require.NoError(t, err)
	res, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, res.Product)
}

func TestLROBeginPut202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut202Retry200(context.Background(), &LROsClientBeginPut202Retry200Options{Product: &Product{}})
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPut202Retry200PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}, pollResp.Product)
}

func TestLROBeginPutAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoHeaderInRetry(context.Background(), &LROsClientBeginPutAsyncNoHeaderInRetryOptions{Product: &Product{}})
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPutAsyncNoHeaderInRetryPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPutAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoRetrySucceeded(context.Background(), &LROsClientBeginPutAsyncNoRetrySucceededOptions{Product: &Product{}})
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPutAsyncNoRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPutAsyncNoRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoRetrycanceled(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPutAsyncNoRetrycanceledPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
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
	resp, err := op.BeginPutAsyncNonResource(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPutAsyncNonResourcePollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, SKU{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("sku"),
	}, pollResp.SKU)
}

func TestLROBeginPutAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncRetryFailed(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPutAsyncRetryFailedPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), time.Second)
	require.Error(t, err)
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
	resp, err := op.BeginPutAsyncRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPutAsyncRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPutAsyncSubResource(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncSubResource(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPutAsyncSubResourcePollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, SubProduct{
		ID: to.StringPtr("100"),
		Properties: &SubProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.SubProduct)
}

func TestLROBeginPutNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutNoHeaderInRetry(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPutNoHeaderInRetryPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPutNonResource(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutNonResource(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPutNonResourcePollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, SKU{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("sku"),
	}, pollResp.SKU)
}

func TestLROBeginPutSubResource(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutSubResource(context.Background(), nil)
	require.NoError(t, err)
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	resp = LROsClientPutSubResourcePollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), time.Second)
	require.NoError(t, err)
	require.Equal(t, SubProduct{
		ID: to.StringPtr("100"),
		Properties: &SubProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}, pollResp.SubProduct)
}
