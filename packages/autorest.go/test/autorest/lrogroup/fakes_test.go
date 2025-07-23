// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup_test

import (
	"context"
	"errors"
	"generatortests"
	"generatortests/lrogroup"
	"generatortests/lrogroup/fake"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeBeginDelete202NoRetry204(t *testing.T) {
	resourceID := "the_id"
	server := fake.LROsServer{
		BeginDelete202NoRetry204: func(ctx context.Context, options *lrogroup.LROsClientBeginDelete202NoRetry204Options) (resp azfake.PollerResponder[lrogroup.LROsClientDelete202NoRetry204Response], errResp azfake.ErrorResponder) {
			resp.AddNonTerminalResponse(http.StatusAccepted, nil)
			resp.SetTerminalResponse(http.StatusOK, lrogroup.LROsClientDelete202NoRetry204Response{
				lrogroup.Product{
					ID: &resourceID,
				},
			}, nil)
			return
		},
	}
	client, err := lrogroup.NewLROsClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewLROsServerTransport(&server),
	})
	require.NoError(t, err)
	poller, err := client.BeginDelete202NoRetry204(context.Background(), nil)
	require.NoError(t, err)
	token, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = client.BeginDelete202NoRetry204(context.Background(), &lrogroup.LROsClientBeginDelete202NoRetry204Options{
		ResumeToken: token,
	})
	require.NoError(t, err)
	resp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: time.Second,
	})
	require.NoError(t, err)
	require.NotNil(t, resp.ID)
	require.EqualValues(t, resourceID, *resp.ID)
}

func TestFakeBeginDelete204Succeeded(t *testing.T) {
	server := fake.LROsServer{
		BeginDelete204Succeeded: func(ctx context.Context, options *lrogroup.LROsClientBeginDelete204SucceededOptions) (resp azfake.PollerResponder[lrogroup.LROsClientDelete204SucceededResponse], errResp azfake.ErrorResponder) {
			resp.SetTerminalResponse(http.StatusNoContent, lrogroup.LROsClientDelete204SucceededResponse{}, nil)
			return
		},
	}
	client, err := lrogroup.NewLROsClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewLROsServerTransport(&server),
	})
	require.NoError(t, err)
	poller, err := client.BeginDelete204Succeeded(context.Background(), nil)
	require.NoError(t, err)
	token, err := poller.ResumeToken()
	require.Error(t, err)
	require.Zero(t, token)
	resp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: time.Second,
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFakeBeginDeleteAsyncRetryFailed(t *testing.T) {
	const errorCode = "YouCantDoThat"
	server := fake.LROsServer{
		BeginDeleteAsyncRetryFailed: func(ctx context.Context, options *lrogroup.LROsClientBeginDeleteAsyncRetryFailedOptions) (resp azfake.PollerResponder[lrogroup.LROsClientDeleteAsyncRetryFailedResponse], errResp azfake.ErrorResponder) {
			resp.AddNonTerminalResponse(http.StatusAccepted, nil)
			resp.AddNonTerminalResponse(http.StatusAccepted, nil)
			resp.SetTerminalError(http.StatusForbidden, errorCode)
			return
		},
	}
	client, err := lrogroup.NewLROsClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewLROsServerTransport(&server),
	})
	require.NoError(t, err)
	poller, err := client.BeginDeleteAsyncRetryFailed(context.Background(), nil)
	require.NoError(t, err)
	token, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = client.BeginDeleteAsyncRetryFailed(context.Background(), &lrogroup.LROsClientBeginDeleteAsyncRetryFailedOptions{
		ResumeToken: token,
	})
	require.NoError(t, err)
	resp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: time.Second,
	})
	require.Error(t, err)
	var respErr *azcore.ResponseError
	require.ErrorAs(t, err, &respErr)
	require.EqualValues(t, errorCode, respErr.ErrorCode)
	require.EqualValues(t, http.StatusForbidden, respErr.StatusCode)
	require.Zero(t, resp)
}

func TestFakeBeginPut200Succeeded(t *testing.T) {
	tagsIn := map[string]*string{
		"key1": to.Ptr("value1"),
		"key2": to.Ptr("value2"),
	}
	tagsOut := map[string]*string{
		"key3": to.Ptr("value3"),
		"key4": to.Ptr("value4"),
	}
	createdID := "the-id"
	server := fake.LROsServer{
		BeginPut200Succeeded: func(ctx context.Context, product lrogroup.Product, options *lrogroup.LROsClientBeginPut200SucceededOptions) (resp azfake.PollerResponder[lrogroup.LROsClientPut200SucceededResponse], errResp azfake.ErrorResponder) {
			require.NotNil(t, product.ID)
			require.EqualValues(t, "input", *product.ID)
			require.NotNil(t, product.Location)
			require.EqualValues(t, "here", *product.Location)
			require.EqualValues(t, tagsIn, product.Tags)
			resp.AddNonTerminalResponse(http.StatusOK, nil)
			resp.AddPollingError(errors.New("transient_error"))
			resp.AddNonTerminalResponse(http.StatusOK, nil)
			resp.SetTerminalResponse(http.StatusOK, lrogroup.LROsClientPut200SucceededResponse{
				Product: lrogroup.Product{
					ID:   &createdID,
					Tags: tagsOut,
				},
			}, nil)
			return
		},
	}
	client, err := lrogroup.NewLROsClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewLROsServerTransport(&server),
	})
	require.NoError(t, err)
	poller, err := client.BeginPut200Succeeded(context.Background(), lrogroup.Product{
		ID:       to.Ptr("input"),
		Location: to.Ptr("here"),
		Tags:     tagsIn,
	}, nil)
	require.NoError(t, err)
	token, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = client.BeginPut200Succeeded(context.Background(), lrogroup.Product{}, &lrogroup.LROsClientBeginPut200SucceededOptions{
		ResumeToken: token,
	})
	require.NoError(t, err)

	iterations := 0
	for {
		if poller.Done() {
			resp, err := poller.Result(context.Background())
			require.NoError(t, err)
			require.NotNil(t, resp.ID)
			require.EqualValues(t, createdID, *resp.ID)
			require.EqualValues(t, tagsOut, resp.Tags)
			break
		}

		resp, err := poller.Poll(context.Background())
		iterations++

		if err != nil {
			require.Nil(t, resp)
			require.EqualValues(t, "transient_error", err.Error())
			continue
		}
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.EqualValues(t, http.StatusOK, resp.StatusCode)
	}

	require.EqualValues(t, 3, iterations)
}
