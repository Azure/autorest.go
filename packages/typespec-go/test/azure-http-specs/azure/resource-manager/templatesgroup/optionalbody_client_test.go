// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package templatesgroup_test

import (
	"context"
	"fmt"
	"templatesgroup"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var validWidget = templatesgroup.Widget{
	ID:       to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Azure.ResourceManager.OperationTemplates/widgets/widget1", subscriptionIdExpected, resourceGroupExpected)),
	Name:     to.Ptr("widget1"),
	Type:     to.Ptr("Azure.ResourceManager.OperationTemplates/widgets"),
	Location: to.Ptr("eastus"),
	Properties: &templatesgroup.WidgetProperties{
		Name:              to.Ptr("widget1"),
		Description:       to.Ptr("A test widget"),
		ProvisioningState: to.Ptr("Succeeded"),
	},
	SystemData: &templatesgroup.SystemData{
		CreatedBy:     to.Ptr("AzureSDK"),
		CreatedByType: to.Ptr(templatesgroup.CreatedByTypeUser),
		CreatedAt: func() *time.Time {
			t, _ := time.Parse(time.RFC3339Nano, "2024-10-04T00:56:07.442Z")
			return &t
		}(),
		LastModifiedBy: to.Ptr("AzureSDK"),
		LastModifiedAt: func() *time.Time {
			t, _ := time.Parse(time.RFC3339Nano, "2024-10-04T00:56:07.442Z")
			return &t
		}(),
		LastModifiedByType: to.Ptr(templatesgroup.CreatedByTypeUser),
	},
}

func TestOptionalBodyClient_Get(t *testing.T) {
	client := clientFactory.NewOptionalBodyClient()
	require.NotNil(t, client)
	resp, err := client.Get(context.Background(), resourceGroupExpected, widgetName, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.EqualValues(t, validWidget, resp.Widget)
}

func TestOptionalBodyClient_Patch(t *testing.T) {
	client := clientFactory.NewOptionalBodyClient()
	require.NotNil(t, client)
	resp, err := client.Patch(context.Background(), resourceGroupExpected, widgetName, templatesgroup.Widget{}, &templatesgroup.OptionalBodyClientPatchOptions{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.EqualValues(t, validWidget, resp.Widget)

	widget := templatesgroup.Widget{
		Name: to.Ptr("widget1"),
		Properties: &templatesgroup.WidgetProperties{
			Name:        to.Ptr("updated-widget"),
			Description: to.Ptr("Updated description"),
		},
	}
	resp, err = client.Patch(context.Background(), resourceGroupExpected, widgetName, widget, &templatesgroup.OptionalBodyClientPatchOptions{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.EqualValues(t, widget.Name, resp.Properties.Name)
	require.EqualValues(t, widget.Properties.Description, resp.Properties.Description)
	require.EqualValues(t, widget.Name, resp.Name)
}

func TestOptionalBodyClient_Post(t *testing.T) {
	client := clientFactory.NewOptionalBodyClient()
	require.NotNil(t, client)
	resp, err := client.Post(context.Background(), resourceGroupExpected, widgetName, &templatesgroup.OptionalBodyClientPostOptions{})
	require.NoError(t, err)
	require.NotNil(t, resp.Result)
	require.EqualValues(t, "Action completed successfully", *resp.Result)

	resp, err = client.Post(context.Background(), resourceGroupExpected, widgetName, &templatesgroup.OptionalBodyClientPostOptions{
		Body: &templatesgroup.ActionRequest{
			ActionType: to.Ptr("perform"),
			Parameters: to.Ptr("test-parameters"),
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp.Result)
	require.EqualValues(t, "Action completed successfully with parameters", *resp.Result)
}

func TestOptionalBodyClient_ProviderPost(t *testing.T) {
	client := clientFactory.NewOptionalBodyClient()
	require.NotNil(t, client)
	resp, err := client.ProviderPost(context.Background(), &templatesgroup.OptionalBodyClientProviderPostOptions{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.EqualValues(t, "Changed to default allowance", *resp.Status)
	require.EqualValues(t, 50, *resp.TotalAllowed)

	resp, err = client.ProviderPost(context.Background(), &templatesgroup.OptionalBodyClientProviderPostOptions{
		Body: &templatesgroup.ChangeAllowanceRequest{
			Reason:       to.Ptr("Increased demand"),
			TotalAllowed: to.Ptr(int32(100)),
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.EqualValues(t, "Changed to requested allowance", *resp.Status)
	require.EqualValues(t, 100, *resp.TotalAllowed)
}
