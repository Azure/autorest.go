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
	headerBoolsContent := []bool{true, false, false, true}
	stringQueryContent := "foo"
	boolHeaderEnumContent := azalias.BooleanEnumEnabled
	unixTimeQueryContent := time.Unix(1460505600, 0).UTC()
	headerEnumContent := azalias.SomeEnumOne
	queryEnumContent := azalias.SomeEnumThree
	BoolHeaderEnum1Content := azalias.BooleanEnumDisabled
	server := fake.Server{
		Create: func(ctx context.Context, headerBools []bool, stringQuery string, boolHeaderEnum azalias.BooleanEnum, unixTimeQuery time.Time, headerEnum azalias.SomeEnum, queryEnum azalias.SomeEnum, options *azalias.CreateOptions) (resp azfake.Responder[azalias.CreateResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, headerBoolsContent, headerBools)
			require.EqualValues(t, stringQueryContent, stringQuery)
			require.EqualValues(t, boolHeaderEnumContent, boolHeaderEnum)
			require.EqualValues(t, unixTimeQueryContent, unixTimeQuery)
			require.EqualValues(t, headerEnumContent, headerEnum)
			require.EqualValues(t, queryEnumContent, queryEnum)
			require.NotNil(t, options)
			require.NotNil(t, options.BoolHeaderEnum1)
			require.EqualValues(t, BoolHeaderEnum1Content, *options.BoolHeaderEnum1)
			resp.SetResponse(http.StatusCreated, azalias.CreateResponse{}, nil)
			return
		},
	}
	client, err := azalias.NewClient("https://contoso.com", &azcore.ClientOptions{
		Transport: fake.NewServerTransport(&server),
	})
	require.NoError(t, err)
	_, err = client.Create(context.Background(), headerBoolsContent, stringQueryContent, boolHeaderEnumContent, unixTimeQueryContent, headerEnumContent, queryEnumContent, &azalias.CreateOptions{
		BoolHeaderEnum1: to.Ptr(BoolHeaderEnum1Content),
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
		GetScript: func(ctx context.Context, headerCounts []int32, queryCounts []int64, explodedStringStuff []string, numericHeader int32, headerTime time.Time, props azalias.GeoJSONObjectNamedCollection, someGroup azalias.SomeGroup, explodedGroup azalias.ExplodedGroup, options *azalias.GetScriptOptions) (resp azfake.Responder[azalias.GetScriptResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, headerContent, headerCounts)
			require.EqualValues(t, queryContent, queryCounts)
			require.EqualValues(t, explodedStringStuff, explodedStrings)
			require.EqualValues(t, headerValue, numericHeader)
			require.EqualValues(t, explodedContent, explodedGroup.ExplodedStuff)
			require.EqualValues(t, timeContent, headerTime)
			require.EqualValues(t, headerStrings, someGroup.HeaderStrings)
			resp.SetResponse(http.StatusOK, azalias.GetScriptResponse{}, nil)
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
