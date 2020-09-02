// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogrouptest

import (
	"context"
	"generatortests/autorest/generated/lrogroup"
	"generatortests/helpers"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newLRORetrysClient() lrogroup.LroRetrysOperations {
	options := lrogroup.DefaultClientOptions()
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	return lrogroup.NewLroRetrysClient(lrogroup.NewDefaultClient(&options))
}

func TestLRORetrysBeginDelete202Retry200(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginDelete202Retry200(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDelete202Retry200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 200)
}

func TestLRORetrysBeginDeleteAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginDeleteAsyncRelativeRetrySucceeded(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRelativeRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 200)
}

func TestLRORetrysBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteProvisioning202Accepted200Succeeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &lrogroup.Product{
		Resource: lrogroup.Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &lrogroup.ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLRORetrysBeginPost202Retry200(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginPost202Retry200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost202Retry200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 200)
}

func TestLRORetrysBeginPostAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginPostAsyncRelativeRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostAsyncRelativeRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 200)
}

func TestLRORetrysBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginPut201CreatingSucceeded200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut201CreatingSucceeded200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &lrogroup.Product{
		Resource: lrogroup.Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &lrogroup.ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLRORetrysBeginPutAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginPutAsyncRelativeRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncRelativeRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &lrogroup.Product{
		Resource: lrogroup.Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &lrogroup.ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}
