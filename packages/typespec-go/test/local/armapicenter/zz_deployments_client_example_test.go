// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armapicenter_test

import (
	"armapicenter"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2024-03-15-preview/Deployments_CreateOrUpdate.json
func ExampleDeploymentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeploymentsClient().CreateOrUpdate(ctx, "contoso-resources", "contoso", "default", "echo-api", "production", armapicenter.Deployment{
		Properties: &armapicenter.DeploymentProperties{
			Title:         to.Ptr("Production deployment"),
			Description:   to.Ptr("Public cloud production deployment."),
			EnvironmentID: to.Ptr("/workspaces/default/environments/production"),
			DefinitionID:  to.Ptr("/workspaces/default/apis/echo-api/versions/2023-01-01/definitions/openapi"),
			State:         to.Ptr(armapicenter.DeploymentStateActive),
			Server: &armapicenter.DeploymentServer{
				RuntimeURI: []*string{
					to.Ptr("https://api.contoso.com"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapicenter.DeploymentsClientCreateOrUpdateResponse{
	// 	Deployment: &armapicenter.Deployment{
	// 		Type: to.Ptr("Microsoft.ApiCenter/services/apis/deployments"),
	// 		ID: to.Ptr("/subscriptions/a200340d-6b82-494d-9dbf-687ba6e33f9e/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/deployments/production"),
	// 		Name: to.Ptr("production"),
	// 		SystemData: &armapicenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128871Z"); return t}()),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.1288716Z"); return t}()),
	// 		},
	// 		Properties: &armapicenter.DeploymentProperties{
	// 			Title: to.Ptr("Production deployment"),
	// 			Description: to.Ptr("Public cloud production deployment."),
	// 			EnvironmentID: to.Ptr("/workspaces/default/environments/production"),
	// 			DefinitionID: to.Ptr("/workspaces/default/apis/echo-api/versions/2023-01-01/definitions/openapi"),
	// 			State: to.Ptr(armapicenter.DeploymentStateActive),
	// 			Server: &armapicenter.DeploymentServer{
	// 				RuntimeURI: []*string{
	// 					to.Ptr("https://api.contoso.com"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-03-15-preview/Deployments_Delete.json
func ExampleDeploymentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeploymentsClient().Delete(ctx, "contoso-resources", "contoso", "default", "echo-api", "production", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapicenter.DeploymentsClientDeleteResponse{
	// }
}

// Generated from example definition: 2024-03-15-preview/Deployments_Get.json
func ExampleDeploymentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeploymentsClient().Get(ctx, "contoso-resources", "contoso", "default", "echo-api", "production", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapicenter.DeploymentsClientGetResponse{
	// 	Deployment: &armapicenter.Deployment{
	// 		Type: to.Ptr("Microsoft.ApiCenter/services/apis/deployments"),
	// 		ID: to.Ptr("/subscriptions/a200340d-6b82-494d-9dbf-687ba6e33f9e/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/deployments/production"),
	// 		Name: to.Ptr("public"),
	// 		SystemData: &armapicenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128871Z"); return t}()),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.1288716Z"); return t}()),
	// 		},
	// 		Properties: &armapicenter.DeploymentProperties{
	// 			Title: to.Ptr("Production deployment"),
	// 			Description: to.Ptr("Public cloud production deployment."),
	// 			EnvironmentID: to.Ptr("/workspaces/default/environments/production"),
	// 			DefinitionID: to.Ptr("/workspaces/default/apis/echo-api/versions/2023-01-01/definitions/openapi"),
	// 			State: to.Ptr(armapicenter.DeploymentStateActive),
	// 			Server: &armapicenter.DeploymentServer{
	// 				RuntimeURI: []*string{
	// 					to.Ptr("https://api.contoso.com"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-03-15-preview/Deployments_Head.json
func ExampleDeploymentsClient_Head() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeploymentsClient().Head(ctx, "contoso-resources", "contoso", "default", "echo-api", "production", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapicenter.DeploymentsClientHeadResponse{
	// }
}

// Generated from example definition: 2024-03-15-preview/Deployments_List.json
func ExampleDeploymentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeploymentsClient().NewListPager("contoso-resources", "contoso", "default", "echo-api", nil)
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
		// page = armapicenter.DeploymentsClientListResponse{
		// 	DeploymentListResult: armapicenter.DeploymentListResult{
		// 		Value: []*armapicenter.Deployment{
		// 			{
		// 				Type: to.Ptr("Microsoft.ApiCenter/services/apis/deployments"),
		// 				ID: to.Ptr("/subscriptions/a200340d-6b82-494d-9dbf-687ba6e33f9e/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/deployments/production"),
		// 				Name: to.Ptr("public"),
		// 				SystemData: &armapicenter.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128871Z"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.1288716Z"); return t}()),
		// 				},
		// 				Properties: &armapicenter.DeploymentProperties{
		// 					Title: to.Ptr("Development"),
		// 					Description: to.Ptr("Public cloud production deployment."),
		// 					EnvironmentID: to.Ptr("/workspaces/default/environments/production"),
		// 					DefinitionID: to.Ptr("/workspaces/default/apis/echo-api/versions/2023-01-01/definitions/openapi"),
		// 					State: to.Ptr(armapicenter.DeploymentStateActive),
		// 					Server: &armapicenter.DeploymentServer{
		// 						RuntimeURI: []*string{
		// 							to.Ptr("https://api.contoso.com"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
