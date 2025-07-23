package traitsgroup

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestTraitsClient_RepeatableAction(t *testing.T) {
	client, err := NewTraitsClient(nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	id := int32(1)
	rfc7231 := "Mon, 27 Nov 2023 11:58:00 GMT"
	layout := time.RFC1123
	tt, err := time.Parse(layout, rfc7231)
	require.NoError(t, err)
	client.endpoint = "http://localhost:3000"
	resp, err := client.RepeatableAction(context.Background(), id, UserActionParam{
		UserActionValue: to.Ptr("test"),
	}, &TraitsClientRepeatableActionOptions{
		RepeatabilityFirstSent: to.Ptr(tt),
		RepeatabilityRequestID: to.Ptr("86aede1f-96fa-4e7f-b1e1-bf8a947cb804"),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "test", *resp.UserActionResponse.UserActionResult)
}

func TestTraitsClient_SmokeTest(t *testing.T) {
	client, err := NewTraitsClient(nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	require.NoError(t, err)
	id := int32(1)
	foo := "123"
	rfc7231_01 := "Thu, 26 Aug 2021 14:38:00 GMT"
	rfc7231_02 := "Fri, 26 Aug 2022 14:38:00 GMT"
	layout := time.RFC1123
	t1, err := time.Parse(layout, rfc7231_01)
	require.NoError(t, err)
	t2, err := time.Parse(layout, rfc7231_02)
	require.NoError(t, err)
	client.endpoint = "http://localhost:3000"
	resp, err := client.SmokeTest(context.Background(), id, foo, &TraitsClientSmokeTestOptions{
		IfMatch:           to.Ptr("\"valid\""),
		IfNoneMatch:       to.Ptr("\"invalid\""),
		IfModifiedSince:   to.Ptr(t1),
		IfUnmodifiedSince: to.Ptr(t2),
		ClientRequestID:   to.Ptr("86aede1f-96fa-4e7f-b1e1-bf8a947cb804"),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, id, *resp.User.ID)
	require.Equal(t, "Madge", *resp.User.Name)
}
