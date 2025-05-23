// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armloadtestservice_test

import (
	"armloadtestservice"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2023-12-01-preview/Quotas_CheckAvailability.json
func ExampleQuotasClient_CheckAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armloadtestservice.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQuotasClient("00000000-0000-0000-0000-000000000000").CheckAvailability(ctx, "westus", "testQuotaBucket", armloadtestservice.QuotaBucketRequest{
		Properties: &armloadtestservice.QuotaBucketRequestProperties{
			CurrentUsage: to.Ptr[int32](20),
			CurrentQuota: to.Ptr[int32](40),
			NewQuota:     to.Ptr[int32](50),
			Dimensions: &armloadtestservice.QuotaBucketRequestPropertiesDimensions{
				SubscriptionID: to.Ptr("testsubscriptionId"),
				Location:       to.Ptr("westus"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armloadtestservice.QuotasClientCheckAvailabilityResponse{
	// 	CheckQuotaAvailabilityResponse: &armloadtestservice.CheckQuotaAvailabilityResponse{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.LoadTestService/locations/westus/quotas/testQuotaBucket"),
	// 		Name: to.Ptr("testQuotaBucket"),
	// 		Type: to.Ptr("Microsoft.LoadTestService/locations/quotas"),
	// 		Properties: &armloadtestservice.CheckQuotaAvailabilityResponseProperties{
	// 			IsAvailable: to.Ptr(false),
	// 			AvailabilityStatus: to.Ptr("The requested quota is currently unavailable. Please request for different quota, or upgrade subscription offer type and try again later."),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2023-12-01-preview/Quotas_Get.json
func ExampleQuotasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armloadtestservice.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQuotasClient("00000000-0000-0000-0000-000000000000").Get(ctx, "westus", "testQuotaBucket", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armloadtestservice.QuotasClientGetResponse{
	// 	QuotaResource: &armloadtestservice.QuotaResource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.LoadTestService/locations/westus/quotas/testQuotaBucket"),
	// 		Name: to.Ptr("testQuotaBucket"),
	// 		Type: to.Ptr("Microsoft.LoadTestService/locations/quotas"),
	// 		Properties: &armloadtestservice.QuotaResourceProperties{
	// 			Limit: to.Ptr[int32](50),
	// 			Usage: to.Ptr[int32](20),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2023-12-01-preview/Quotas_List.json
func ExampleQuotasClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armloadtestservice.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQuotasClient("00000000-0000-0000-0000-000000000000").NewListPager("westus", nil)
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
		// page = armloadtestservice.QuotasClientListResponse{
		// 	QuotaResourceListResult: armloadtestservice.QuotaResourceListResult{
		// 		Value: []*armloadtestservice.QuotaResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.LoadTestService/locations/westus/quotas/testQuotaBucket"),
		// 				Name: to.Ptr("testQuotaBucket"),
		// 				Type: to.Ptr("Microsoft.LoadTestService/locations/quotas"),
		// 				Properties: &armloadtestservice.QuotaResourceProperties{
		// 					Limit: to.Ptr[int32](50),
		// 					Usage: to.Ptr[int32](20),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
