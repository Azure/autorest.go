// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcommunitymanagement_test

import (
	"armcommunitymanagement"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2023-11-01/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunitymanagement.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armcommunitymanagement.OperationsClientListResponse{
		// 	OperationListResult: armcommunitymanagement.OperationListResult{
		// 		Value: []*armcommunitymanagement.Operation{
		// 			{
		// 				Display: &armcommunitymanagement.OperationDisplay{
		// 					Description: to.Ptr("List CommunityTrainings Resource resources by subscription ID"),
		// 					Operation: to.Ptr("CommunityTrainings_ListBySubscription"),
		// 					Provider: to.Ptr("Microsoft.Community"),
		// 					Resource: to.Ptr("communityTrainings"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Name: to.Ptr("Microsoft.Community/communityTrainings/read"),
		// 			},
		// 			{
		// 				Display: &armcommunitymanagement.OperationDisplay{
		// 					Description: to.Ptr("List CommunityTrainings Resource resources by resource group"),
		// 					Operation: to.Ptr("CommunityTrainings_ListByResourceGroup"),
		// 					Provider: to.Ptr("Microsoft.Community"),
		// 					Resource: to.Ptr("communityTrainings"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Name: to.Ptr("Microsoft.Community/communityTrainings/read"),
		// 			},
		// 			{
		// 				Display: &armcommunitymanagement.OperationDisplay{
		// 					Description: to.Ptr("Get a CommunityTrainings Resource"),
		// 					Operation: to.Ptr("CommunityTrainings_Get"),
		// 					Provider: to.Ptr("Microsoft.Community"),
		// 					Resource: to.Ptr("communityTrainings"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Name: to.Ptr("Microsoft.Community/communityTrainings/read"),
		// 			},
		// 			{
		// 				Display: &armcommunitymanagement.OperationDisplay{
		// 					Description: to.Ptr("Create a CommunityTrainings Resource"),
		// 					Operation: to.Ptr("CommunityTrainings_Create"),
		// 					Provider: to.Ptr("Microsoft.Community"),
		// 					Resource: to.Ptr("communityTrainings"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Name: to.Ptr("Microsoft.Community/communityTrainings/write"),
		// 			},
		// 			{
		// 				Display: &armcommunitymanagement.OperationDisplay{
		// 					Description: to.Ptr("Delete a CommunityTraining Resource"),
		// 					Operation: to.Ptr("CommunityTrainings_Delete"),
		// 					Provider: to.Ptr("Microsoft.Community"),
		// 					Resource: to.Ptr("communityTrainings"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Name: to.Ptr("Microsoft.Community/communityTrainings/delete"),
		// 			},
		// 			{
		// 				Display: &armcommunitymanagement.OperationDisplay{
		// 					Description: to.Ptr("Update a CommunityTrainings Resource"),
		// 					Operation: to.Ptr("CommunityTrainings_Update"),
		// 					Provider: to.Ptr("Microsoft.Community"),
		// 					Resource: to.Ptr("communityTrainings"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Name: to.Ptr("Microsoft.Community/communityTrainings/write"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
