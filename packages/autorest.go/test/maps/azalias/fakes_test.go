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
	"github.com/stretchr/testify/require"
)

func TestFakeGetScript(t *testing.T) {
	headerContent := []int32{0, 2, 4}
	queryContent := []int64{3, 6, 9}
	explodedContent := []int64{9, 8, 7}
	timeContent, err := time.Parse(time.TimeOnly, "15:04:05.12345")
	require.NoError(t, err)
	server := fake.Server{
		GetScript: func(ctx context.Context, headerCounts []int32, queryCounts []int64, headerTime time.Time, props azalias.GeoJSONObjectNamedCollection, explodedGroup azalias.ExplodedGroup, options *azalias.ClientGetScriptOptions) (resp azfake.Responder[azalias.ClientGetScriptResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, headerContent, headerCounts)
			require.EqualValues(t, queryContent, queryCounts)
			require.EqualValues(t, explodedContent, explodedGroup.ExplodedStuff)
			require.EqualValues(t, timeContent, headerTime)
			resp.SetResponse(http.StatusOK, azalias.ClientGetScriptResponse{}, nil)
			return
		},
	}
	client, err := azalias.NewClient("https://contoso.com", &azcore.ClientOptions{
		Transport: fake.NewServerTransport(&server),
	})
	require.NoError(t, err)
	_, err = client.GetScript(context.Background(), headerContent, queryContent, timeContent, azalias.GeoJSONObjectNamedCollection{}, azalias.ExplodedGroup{
		ExplodedStuff: explodedContent,
	}, nil)
	require.NoError(t, err)
}
