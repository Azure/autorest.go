// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azalias_test

import (
	"azalias"
	"azalias/fake"
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeCreate(t *testing.T) {
	headerContent := []bool{true, false, false, true}
	queryContent := "foo"
	headerEnum := azalias.BooleanEnumEnabled
	queryEnum := azalias.BooleanEnumDisabled
	server := fake.Server{
		Create: func(ctx context.Context, headerBools []bool, stringQuery string, boolHeaderEnum azalias.BooleanEnum, options *azalias.ClientCreateOptions) (resp azfake.Responder[azalias.ClientCreateResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, headerContent, headerBools)
			require.EqualValues(t, queryContent, stringQuery)
			require.EqualValues(t, headerEnum, boolHeaderEnum)
			require.NotNil(t, options)
			require.NotNil(t, options.BoolHeaderEnum1)
			require.EqualValues(t, queryEnum, *options.BoolHeaderEnum1)
			resp.SetResponse(http.StatusCreated, azalias.ClientCreateResponse{}, nil)
			return
		},
	}
	client, err := azalias.NewClient("https://contoso.com", &azcore.ClientOptions{
		Transport: fake.NewServerTransport(&server),
	})
	require.NoError(t, err)
	_, err = client.Create(context.Background(), headerContent, queryContent, headerEnum, &azalias.ClientCreateOptions{
		BoolHeaderEnum1: to.Ptr(queryEnum),
	})
	require.NoError(t, err)
}

func TestFakeGetScript(t *testing.T) {
	headerContent := []int32{0, 2, 4}
	queryContent := []int64{3, 6, 9}
	explodedContent := []int64{9, 8, 7}
	explodedStrings := []string{"foo", "bar"}
	headerValue := int32(12345)
	timeContent, err := time.Parse(time.TimeOnly, "15:04:05.12345")
	headerStrings := []string{"bing", "bing"}
	require.NoError(t, err)
	server := fake.Server{
		GetScript: func(ctx context.Context, headerCounts []int32, queryCounts []int64, explodedStringStuff []string, numericHeader int32, headerTime time.Time, props azalias.GeoJSONObjectNamedCollection, someGroup azalias.SomeGroup, explodedGroup azalias.ExplodedGroup, options *azalias.ClientGetScriptOptions) (resp azfake.Responder[azalias.ClientGetScriptResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, headerContent, headerCounts)
			require.EqualValues(t, queryContent, queryCounts)
			require.EqualValues(t, explodedStringStuff, explodedStrings)
			require.EqualValues(t, headerValue, numericHeader)
			require.EqualValues(t, explodedContent, explodedGroup.ExplodedStuff)
			require.EqualValues(t, timeContent, headerTime)
			require.EqualValues(t, headerStrings, someGroup.HeaderStrings)
			resp.SetResponse(http.StatusOK, azalias.ClientGetScriptResponse{}, nil)
			return
		},
	}
	client, err := azalias.NewClient("https://contoso.com", &azcore.ClientOptions{
		Transport: fake.NewServerTransport(&server),
	})
	require.NoError(t, err)
	_, err = client.GetScript(context.Background(), headerContent, queryContent, explodedStrings, headerValue, timeContent, azalias.GeoJSONObjectNamedCollection{}, azalias.SomeGroup{
		HeaderStrings: headerStrings,
	}, azalias.ExplodedGroup{
		ExplodedStuff: explodedContent,
	}, nil)
	require.NoError(t, err)
}
