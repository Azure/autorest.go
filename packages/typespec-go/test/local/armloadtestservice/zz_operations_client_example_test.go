// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armloadtestservice_test

import (
	"armloadtestservice"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2023-12-01-preview/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armloadtestservice.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page = armloadtestservice.OperationsClientListResponse{
		// 	OperationListResult: armloadtestservice.OperationListResult{
		// 		Value: []*armloadtestservice.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.LoadTestService/loadTests/Write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armloadtestservice.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.LoadTestService"),
		// 					Resource: to.Ptr("loadTests"),
		// 					Operation: to.Ptr("Creates or updates the LoadTests"),
		// 					Description: to.Ptr("Set LoadTests"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.LoadTestService/loadTests/Delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armloadtestservice.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.LoadTestService"),
		// 					Resource: to.Ptr("loadTests"),
		// 					Operation: to.Ptr("Deletes the LoadTests"),
		// 					Description: to.Ptr("Delete LoadTests"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.LoadTestService/loadTests/Read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armloadtestservice.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.LoadTestService"),
		// 					Resource: to.Ptr("loadTests"),
		// 					Operation: to.Ptr("Reads the LoadTests"),
		// 					Description: to.Ptr("Read LoadTests"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}