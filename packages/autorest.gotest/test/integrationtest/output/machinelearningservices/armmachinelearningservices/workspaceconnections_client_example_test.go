//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceConnection/list.json
func ExampleWorkspaceConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceConnectionsClient().NewListPager("resourceGroup-1", "workspace-1", &armmachinelearningservices.WorkspaceConnectionsClientListOptions{Target: to.Ptr("www.facebook.com"),
		Category: to.Ptr("ACR"),
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PaginatedWorkspaceConnectionsList = armmachinelearningservices.PaginatedWorkspaceConnectionsList{
		// 	Value: []*armmachinelearningservices.WorkspaceConnection{
		// 		{
		// 			Name: to.Ptr("connection-1"),
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/linkedWorkspaces/connection-1"),
		// 			Properties: &armmachinelearningservices.WorkspaceConnectionProps{
		// 				AuthType: to.Ptr("PAT"),
		// 				Category: to.Ptr("ACR"),
		// 				Target: to.Ptr("www.facebook.com"),
		// 				Value: to.Ptr("secrets"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("connection-2"),
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/linkedWorkspaces/connection-2"),
		// 			Properties: &armmachinelearningservices.WorkspaceConnectionProps{
		// 				AuthType: to.Ptr("PAT"),
		// 				Category: to.Ptr("ACR"),
		// 				Target: to.Ptr("www.facebook.com"),
		// 				Value: to.Ptr("secrets"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceConnection/create.json
func ExampleWorkspaceConnectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceConnectionsClient().Create(ctx, "resourceGroup-1", "workspace-1", "connection-1", armmachinelearningservices.WorkspaceConnection{
		Properties: &armmachinelearningservices.WorkspaceConnectionProps{
			AuthType: to.Ptr("PAT"),
			Category: to.Ptr("ACR"),
			Target:   to.Ptr("www.facebook.com"),
			Value:    to.Ptr("secrets"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkspaceConnection = armmachinelearningservices.WorkspaceConnection{
	// 	Name: to.Ptr("connection-1"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/connections/connection-1"),
	// 	Properties: &armmachinelearningservices.WorkspaceConnectionProps{
	// 		AuthType: to.Ptr("PAT"),
	// 		Category: to.Ptr("ACR"),
	// 		Target: to.Ptr("www.facebook.com"),
	// 		Value: to.Ptr("secrets"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceConnection/get.json
func ExampleWorkspaceConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceConnectionsClient().Get(ctx, "resourceGroup-1", "workspace-1", "connection-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkspaceConnection = armmachinelearningservices.WorkspaceConnection{
	// 	Name: to.Ptr("connection-1"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/connections/connection-1"),
	// 	Properties: &armmachinelearningservices.WorkspaceConnectionProps{
	// 		AuthType: to.Ptr("PAT"),
	// 		Category: to.Ptr("ACR"),
	// 		Target: to.Ptr("www.facebook.com"),
	// 		Value: to.Ptr("secrets"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceConnection/delete.json
func ExampleWorkspaceConnectionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspaceConnectionsClient().Delete(ctx, "resourceGroup-1", "workspace-1", "connection-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
