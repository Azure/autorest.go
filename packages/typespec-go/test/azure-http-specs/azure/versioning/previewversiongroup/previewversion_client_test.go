// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package previewversiongroup_test

import (
	"context"
	"fmt"
	"net/http"
	"previewversiongroup"
	"previewversiongroup/fake"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestPreviewVersionClient_GetWidget(t *testing.T) {
	const fakeID = "widget-123"
	fakeWidget := previewversiongroup.Widget{
		ID:    to.Ptr(fakeID),
		Name:  to.Ptr("Sample Widget"),
		Color: to.Ptr("blue"),
	}

	server := fake.PreviewVersionServer{
		GetWidget: func(ctx context.Context, id string, options *previewversiongroup.PreviewVersionClientGetWidgetOptions) (resp azfake.Responder[previewversiongroup.PreviewVersionClientGetWidgetResponse], errResp azfake.ErrorResponder) {
			if id != fakeID {
				errResp.SetError(fmt.Errorf("unexpected widget id: %s", id))
				return
			}

			resp.SetResponse(http.StatusOK, previewversiongroup.PreviewVersionClientGetWidgetResponse{
				Widget: fakeWidget,
			}, nil)
			return
		},
	}

	client, err := previewversiongroup.NewPreviewVersionClientWithNoCredential("http://localhost:3000", &previewversiongroup.PreviewVersionClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPreviewVersionServerTransport(&server),
		},
	})
	require.NoError(t, err)

	resp, err := client.GetWidget(context.Background(), fakeID, nil)
	require.NoError(t, err)
	require.Equal(t, "widget-123", *resp.ID)
	require.Equal(t, "Sample Widget", *resp.Name)
	require.Equal(t, "blue", *resp.Color)
}

func TestPreviewVersionClient_ListWidgets(t *testing.T) {
	fakeWidgets := []*previewversiongroup.Widget{
		{
			ID:    to.Ptr("widget-1"),
			Name:  to.Ptr("test"),
			Color: to.Ptr("red"),
		},
		{
			ID:    to.Ptr("widget-2"),
			Name:  to.Ptr("Blue Widget"),
			Color: to.Ptr("blue"),
		},
	}

	server := fake.PreviewVersionServer{
		ListWidgets: func(ctx context.Context, options *previewversiongroup.PreviewVersionClientListWidgetsOptions) (resp azfake.Responder[previewversiongroup.PreviewVersionClientListWidgetsResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, previewversiongroup.PreviewVersionClientListWidgetsResponse{
				ListWidgetsResponse: previewversiongroup.ListWidgetsResponse{
					Widgets: fakeWidgets,
				},
			}, nil)
			return
		},
	}

	client, err := previewversiongroup.NewPreviewVersionClientWithNoCredential("http://localhost:3000", &previewversiongroup.PreviewVersionClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPreviewVersionServerTransport(&server),
		},
	})
	require.NoError(t, err)

	resp, err := client.ListWidgets(context.Background(), nil)
	require.NoError(t, err)
	require.Len(t, resp.Widgets, 2)
	require.Equal(t, "widget-1", *resp.Widgets[0].ID)
	require.Equal(t, "test", *resp.Widgets[0].Name)
	require.Equal(t, "red", *resp.Widgets[0].Color)
}

func TestPreviewVersionClient_UpdateWidgetColor(t *testing.T) {
	const fakeID = "widget-1"
	const newColor = "green"
	updatedWidget := previewversiongroup.Widget{
		ID:    to.Ptr(fakeID),
		Name:  to.Ptr("test"),
		Color: to.Ptr(newColor),
	}

	server := fake.PreviewVersionServer{
		UpdateWidgetColor: func(ctx context.Context, id string, colorUpdate previewversiongroup.UpdateWidgetColorRequest, options *previewversiongroup.PreviewVersionClientUpdateWidgetColorOptions) (resp azfake.Responder[previewversiongroup.PreviewVersionClientUpdateWidgetColorResponse], errResp azfake.ErrorResponder) {
			if id != fakeID {
				errResp.SetError(fmt.Errorf("unexpected widget id: %s", id))
				return
			}
			require.NotNil(t, colorUpdate.Color)
			require.Equal(t, newColor, *colorUpdate.Color)

			resp.SetResponse(http.StatusOK, previewversiongroup.PreviewVersionClientUpdateWidgetColorResponse{
				Widget: updatedWidget,
			}, nil)
			return
		},
	}

	client, err := previewversiongroup.NewPreviewVersionClientWithNoCredential("https://localhost", &previewversiongroup.PreviewVersionClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPreviewVersionServerTransport(&server),
		},
	})
	require.NoError(t, err)

	colorUpdate := previewversiongroup.UpdateWidgetColorRequest{
		Color: to.Ptr(newColor),
	}

	resp, err := client.UpdateWidgetColor(context.Background(), fakeID, colorUpdate, nil)
	require.NoError(t, err)
	require.Equal(t, fakeID, *resp.ID)
	require.Equal(t, "test", *resp.Name)
	require.Equal(t, newColor, *resp.Color)
}
