// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"errors"
	"generatortests"
	"net/http"
	"net/http/cookiejar"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newLROSClient() *LROsClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = time.Second
	options.Transport = httpClientWithCookieJar()
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &options)
	return NewLROsClient(pl)
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
	_, err = op.BeginDelete202Retry200(context.Background(), &LROsClientBeginDelete202Retry200Options{
		ResumeToken: rt,
	})
	require.Error(t, err)
}

func TestLROBeginDelete202NoRetry204(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDelete202NoRetry204(context.Background(), &LROsClientBeginDelete202NoRetry204Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROBeginDelete202Retry200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDelete202Retry200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDelete202Retry200(context.Background(), &LROsClientBeginDelete202Retry200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROBeginDelete204Succeeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDelete204Succeeded(context.Background(), nil)
	require.NoError(t, err)
	_, err = poller.ResumeToken()
	require.Error(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteAsyncNoHeaderInRetry(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteAsyncNoHeaderInRetry(context.Background(), &LROsClientBeginDeleteAsyncNoHeaderInRetryOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteAsyncNoRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteAsyncNoRetrySucceeded(context.Background(), &LROsClientBeginDeleteAsyncNoRetrySucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteAsyncRetryFailed(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteAsyncRetryFailed(context.Background(), &LROsClientBeginDeleteAsyncRetryFailedOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	res, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
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
	poller, err := op.BeginDeleteAsyncRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteAsyncRetrySucceeded(context.Background(), &LROsClientBeginDeleteAsyncRetrySucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROBeginDeleteAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteAsyncRetrycanceled(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteAsyncRetrycanceled(context.Background(), &LROsClientBeginDeleteAsyncRetrycanceledOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	res, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
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
	poller, err := op.BeginDeleteNoHeaderInRetry(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteNoHeaderInRetry(context.Background(), &LROsClientBeginDeleteNoHeaderInRetryOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), &LROsClientBeginDeleteProvisioning202Accepted200SucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROBeginDeleteProvisioning202DeletingFailed200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteProvisioning202DeletingFailed200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteProvisioning202DeletingFailed200(context.Background(), &LROsClientBeginDeleteProvisioning202DeletingFailed200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROBeginDeleteProvisioning202Deletingcanceled200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginDeleteProvisioning202Deletingcanceled200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginDeleteProvisioning202Deletingcanceled200(context.Background(), &LROsClientBeginDeleteProvisioning202Deletingcanceled200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.Error(t, err)
}

func TestLROBeginPatch201RetryWithAsyncHeader(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPatch201RetryWithAsyncHeader(context.Background(), Product{}, nil)
	require.NoError(t, err)
	res, err := resp.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("/lro/patch/201/retry/onlyAsyncHeader"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, res.Product)
}

func TestLROBeginPatch202RetryWithAsyncAndLocationHeader(t *testing.T) {
	t.Skip("https://github.com/Azure/autorest.testserver/pull/369")
	op := newLROSClient()
	resp, err := op.BeginPatch202RetryWithAsyncAndLocationHeader(context.Background(), Product{}, nil)
	require.NoError(t, err)
	res, err := resp.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("/lro/patch/202/retry/asyncAndLocationHeader"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, res.Product)
}

func TestLROBeginPost200WithPayload(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPost200WithPayload(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPost200WithPayload(context.Background(), &LROsClientBeginPost200WithPayloadOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, SKU{
		ID:   to.Ptr("1"),
		Name: to.Ptr("product"),
	}, pollResp.SKU)
}

func TestLROBeginPost202List(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPost202List(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPost202List(context.Background(), &LROsClientBeginPost202ListOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, []*Product{{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
	}}, pollResp.ProductArray)
}

func TestLROBeginPost202NoRetry204(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPost202NoRetry204(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPost202NoRetry204(context.Background(), &LROsClientBeginPost202NoRetry204Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROBeginPost202Retry200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPost202Retry200(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPost202Retry200(context.Background(), &LROsClientBeginPost202Retry200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
}

func TestLROBeginPostAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostAsyncNoRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostAsyncNoRetrySucceeded(context.Background(), &LROsClientBeginPostAsyncNoRetrySucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPostAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostAsyncRetryFailed(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostAsyncRetryFailed(context.Background(), &LROsClientBeginPostAsyncRetryFailedOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	res, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
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
	poller, err := op.BeginPostAsyncRetrySucceeded(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostAsyncRetrySucceeded(context.Background(), &LROsClientBeginPostAsyncRetrySucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPostAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostAsyncRetrycanceled(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostAsyncRetrycanceled(context.Background(), &LROsClientBeginPostAsyncRetrycanceledOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	res, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
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
	poller, err := op.BeginPostDoubleHeadersFinalAzureHeaderGet(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostDoubleHeadersFinalAzureHeaderGet(context.Background(), &LROsClientBeginPostDoubleHeadersFinalAzureHeaderGetOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID: to.Ptr("100"),
	}, pollResp.Product)
}

func TestLROBeginPostDoubleHeadersFinalAzureHeaderGetDefault(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostDoubleHeadersFinalAzureHeaderGetDefault(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostDoubleHeadersFinalAzureHeaderGetDefault(context.Background(), &LROsClientBeginPostDoubleHeadersFinalAzureHeaderGetDefaultOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
	}, pollResp.Product)
}

func TestLROBeginPostDoubleHeadersFinalLocationGet(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPostDoubleHeadersFinalLocationGet(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPostDoubleHeadersFinalLocationGet(context.Background(), &LROsClientBeginPostDoubleHeadersFinalLocationGetOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
	}, pollResp.Product)
}

func TestLROBeginPut200Acceptedcanceled200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut200Acceptedcanceled200(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPut200Acceptedcanceled200(context.Background(), Product{}, &LROsClientBeginPut200Acceptedcanceled200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
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
	poller, err := op.BeginPut200Succeeded(context.Background(), Product{}, nil)
	require.NoError(t, err)
	_, err = poller.ResumeToken()
	require.Error(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPut200SucceededNoState(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut200SucceededNoState(context.Background(), Product{}, nil)
	require.NoError(t, err)
	_, err = poller.ResumeToken()
	require.Error(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
	}, pollResp.Product)
}

// TODO check if this test should actually be returning a 200 or a 204
func TestLROBeginPut200UpdatingSucceeded204(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut200UpdatingSucceeded204(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPut200UpdatingSucceeded204(context.Background(), Product{}, &LROsClientBeginPut200UpdatingSucceeded204Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPut201CreatingFailed200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut201CreatingFailed200(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPut201CreatingFailed200(context.Background(), Product{}, &LROsClientBeginPut201CreatingFailed200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	_, err = poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
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
	poller, err := op.BeginPut201CreatingSucceeded200(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPut201CreatingSucceeded200(context.Background(), Product{}, &LROsClientBeginPut201CreatingSucceeded200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPut201Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut201Succeeded(context.Background(), Product{}, nil)
	require.NoError(t, err)
	res, err := resp.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, res.Product)
}

func TestLROBeginPut202Retry200(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPut202Retry200(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPut202Retry200(context.Background(), Product{}, &LROsClientBeginPut202Retry200Options{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
	}, pollResp.Product)
}

func TestLROBeginPutAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncNoHeaderInRetry(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncNoHeaderInRetry(context.Background(), Product{}, &LROsClientBeginPutAsyncNoHeaderInRetryOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPutAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncNoRetrySucceeded(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncNoRetrySucceeded(context.Background(), Product{}, &LROsClientBeginPutAsyncNoRetrySucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPutAsyncNoRetrycanceled(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncNoRetrycanceled(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncNoRetrycanceled(context.Background(), Product{}, &LROsClientBeginPutAsyncNoRetrycanceledOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	res, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
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
	poller, err := op.BeginPutAsyncNonResource(context.Background(), SKU{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncNonResource(context.Background(), SKU{}, &LROsClientBeginPutAsyncNonResourceOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, SKU{
		ID:   to.Ptr("100"),
		Name: to.Ptr("sku"),
	}, pollResp.SKU)
}

func TestLROBeginPutAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncRetryFailed(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncRetryFailed(context.Background(), Product{}, &LROsClientBeginPutAsyncRetryFailedOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	res, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
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
	poller, err := op.BeginPutAsyncRetrySucceeded(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncRetrySucceeded(context.Background(), Product{}, &LROsClientBeginPutAsyncRetrySucceededOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPutAsyncSubResource(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutAsyncSubResource(context.Background(), SubProduct{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutAsyncSubResource(context.Background(), SubProduct{}, &LROsClientBeginPutAsyncSubResourceOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, SubProduct{
		ID: to.Ptr("100"),
		Properties: &SubProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.SubProduct)
}

func TestLROBeginPutNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutNoHeaderInRetry(context.Background(), Product{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutNoHeaderInRetry(context.Background(), Product{}, &LROsClientBeginPutNoHeaderInRetryOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, Product{
		ID:   to.Ptr("100"),
		Name: to.Ptr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.Product)
}

func TestLROBeginPutNonResource(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutNonResource(context.Background(), SKU{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutNonResource(context.Background(), SKU{}, &LROsClientBeginPutNonResourceOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, SKU{
		ID:   to.Ptr("100"),
		Name: to.Ptr("sku"),
	}, pollResp.SKU)
}

func TestLROBeginPutSubResource(t *testing.T) {
	op := newLROSClient()
	poller, err := op.BeginPutSubResource(context.Background(), SubProduct{}, nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = op.BeginPutSubResource(context.Background(), SubProduct{}, &LROsClientBeginPutSubResourceOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pollResp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.Equal(t, SubProduct{
		ID: to.Ptr("100"),
		Properties: &SubProductProperties{
			ProvisioningState: to.Ptr("Succeeded"),
		},
	}, pollResp.SubProduct)
}
