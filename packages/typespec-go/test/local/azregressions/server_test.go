// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azregressions_test

import (
	"azregressions"
	"azregressions/fake"
	"context"
	"testing"
	"time"

	"net/http"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newTestClient(t *testing.T) *azregressions.Client {
	t.Helper()
	client, err := azregressions.NewClientWithNoCredential("https://fake.endpoint", &azregressions.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Transport: fake.NewServerTransport(&fake.Server{
				GetQueue: func(_ context.Context, _ *azregressions.ClientGetQueueOptions) (resp azfake.Responder[azregressions.ClientGetQueueResponse], errResp azfake.ErrorResponder) {
					return
				},
			}),
		},
	})
	require.NoError(t, err)
	return client
}

// TestCancelledContextPanic reproduces issue https://github.com/Azure/azure-sdk-for-go/issues/25895.
// When the context is already cancelled before calling a fake server method, the outer select in
// dispatchToMethodFake returns immediately on <-req.Context().Done(), defer close(resultChan)
// fires, and then the goroutine tries to send on the now-closed channel, causing a panic.
func TestCancelledContextPanic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel before calling any method

	client := newTestClient(t)

	for range 500 {
		_, err := client.GetQueue(ctx, nil)
		require.ErrorIs(t, err, context.Canceled)
	}
}

// TestContextCancelDuringDispatchRace reproduces the race condition identified in
// https://github.com/Azure/azure-sdk-for-go/pull/26444#discussion_r3010484579.
// When the context is cancelled while the goroutine is still dispatching, there is a data race
// between the deferred close(resultChan) in the outer function and the goroutine's send on resultChan.
func TestContextCancelDuringDispatchRace(t *testing.T) {
	client := newTestClient(t)

	for range 500 {
		// use a near-zero timeout so the context expires while the goroutine is mid-dispatch,
		// triggering the race between close(resultChan) and the goroutine's send
		ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
		_, _ = client.GetQueue(ctx, nil)
		cancel()
	}
}

// TestGetIntegerFake verifies that the GetInteger fake server round-trips an integer
// value correctly through the text/plain response body.
func TestGetIntegerFake(t *testing.T) {
	const expected int64 = 42
	client, err := azregressions.NewClientWithNoCredential("https://fake.endpoint", &azregressions.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Transport: fake.NewServerTransport(&fake.Server{
				GetInteger: func(_ context.Context, _ *azregressions.ClientGetIntegerOptions) (resp azfake.Responder[azregressions.ClientGetIntegerResponse], errResp azfake.ErrorResponder) {
					resp.SetResponse(http.StatusOK, azregressions.ClientGetIntegerResponse{
						ContentType: to.Ptr("text/plain; charset=utf-8"),
						Value:       to.Ptr(expected),
					}, nil)
					return
				},
			}),
		},
	})
	require.NoError(t, err)

	result, err := client.GetInteger(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result.Value)
	require.Equal(t, expected, *result.Value)
	require.NotNil(t, result.ContentType)
	require.Equal(t, "text/plain; charset=utf-8", *result.ContentType)
}

// TestDoubleDecodeQueryParam verifies that query parameter values containing
// percent-encoded sequences (e.g. "foo%20bar") are not double-decoded by the fake server.
// The server's dispatchDoubleDecode calls url.QueryUnescape on a value that is already
// decoded by req.URL.Query(), so "foo%20bar" incorrectly becomes "foo bar".
func TestDoubleDecodeQueryParam(t *testing.T) {
	const pathValue = "foo%20bar"
	const queryValue = "baz%20qux"

	var receivedPath, receivedQuery string
	client, err := azregressions.NewClientWithNoCredential("https://fake.endpoint", &azregressions.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Transport: fake.NewServerTransport(&fake.Server{
				DoubleDecode: func(_ context.Context, pathParam string, query string, _ *azregressions.ClientDoubleDecodeOptions) (resp azfake.Responder[azregressions.ClientDoubleDecodeResponse], errResp azfake.ErrorResponder) {
					receivedPath = pathParam
					receivedQuery = query
					resp.SetResponse(http.StatusNoContent, azregressions.ClientDoubleDecodeResponse{}, nil)
					return
				},
			}),
		},
	})
	require.NoError(t, err)

	_, err = client.DoubleDecode(context.Background(), pathValue, queryValue, nil)
	require.NoError(t, err)
	require.Equal(t, pathValue, receivedPath, "path param was double-decoded")
	require.Equal(t, queryValue, receivedQuery, "query param was double-decoded")
}
