// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armlargeinstance_test

import (
	"armlargeinstance"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2024-08-01-preview/AzureLargeInstance_Create.json
func ExampleClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("f0f4887f-d13c-4943-a8ba-d7da28d2a3fd", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Create(ctx, "myResourceGroup", "myALInstance", armlargeinstance.AzureLargeInstance{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"testkey": to.Ptr("testvalue"),
		},
		Properties: &armlargeinstance.Properties{
			AzureLargeInstanceID:    to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
			PowerState:              to.Ptr(armlargeinstance.PowerStateEnumStarted),
			ProximityPlacementGroup: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Compute/proximityPlacementGroups/myplacementgroup"),
			HwRevision:              to.Ptr("Rev 3"),
			HardwareProfile: &armlargeinstance.HardwareProfile{
				HardwareType:           to.Ptr(armlargeinstance.HardwareTypeNamesEnumCiscoUCS),
				AzureLargeInstanceSize: to.Ptr(armlargeinstance.SizeNamesEnumS72),
			},
			NetworkProfile: &armlargeinstance.NetworkProfile{
				NetworkInterfaces: []*armlargeinstance.IPAddress{
					{
						IPAddress: to.Ptr("100.100.100.100"),
					},
				},
				CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuit"),
			},
			StorageProfile: &armlargeinstance.StorageProfile{
				NfsIPAddress: to.Ptr("200.200.200.200"),
			},
			OSProfile: &armlargeinstance.OsProfile{
				ComputerName: to.Ptr("myComputerName"),
				OSType:       to.Ptr("SUSE"),
				Version:      to.Ptr("12 SP1"),
				SSHPublicKey: to.Ptr("{ssh-rsa public key}"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armlargeinstance.ClientCreateResponse{
	// 	AzureLargeInstance: &armlargeinstance.AzureLargeInstance{
	// 		ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myALInstance"),
	// 		Name: to.Ptr("myALInstance"),
	// 		Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeInstances"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"testkey": to.Ptr("testvalue"),
	// 		},
	// 		SystemData: &armlargeinstance.SystemData{
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-20T23:10:22.6828621Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.0000000Z"); return t}()),
	// 		},
	// 		Properties: &armlargeinstance.Properties{
	// 			AzureLargeInstanceID: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
	// 			PowerState: to.Ptr(armlargeinstance.PowerStateEnumStarted),
	// 			ProximityPlacementGroup: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Compute/proximityPlacementGroups/myplacementgroup"),
	// 			HwRevision: to.Ptr("Rev 3"),
	// 			HardwareProfile: &armlargeinstance.HardwareProfile{
	// 				HardwareType: to.Ptr(armlargeinstance.HardwareTypeNamesEnumCiscoUCS),
	// 				AzureLargeInstanceSize: to.Ptr(armlargeinstance.SizeNamesEnumS72),
	// 			},
	// 			NetworkProfile: &armlargeinstance.NetworkProfile{
	// 				NetworkInterfaces: []*armlargeinstance.IPAddress{
	// 					{
	// 						IPAddress: to.Ptr("100.100.100.100"),
	// 					},
	// 				},
	// 				CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuit"),
	// 			},
	// 			StorageProfile: &armlargeinstance.StorageProfile{
	// 				NfsIPAddress: to.Ptr("200.200.200.200"),
	// 			},
	// 			OSProfile: &armlargeinstance.OsProfile{
	// 				ComputerName: to.Ptr("myComputerName"),
	// 				OSType: to.Ptr("SUSE"),
	// 				Version: to.Ptr("12 SP1"),
	// 				SSHPublicKey: to.Ptr("{ssh-rsa public key}"),
	// 			},
	// 			ProvisioningState: to.Ptr(armlargeinstance.ProvisioningStatesEnumSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-08-01-preview/AzureLargeInstance_Delete.json
func ExampleClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("f0f4887f-d13c-4943-a8ba-d7da28d2a3fd", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Delete(ctx, "myResourceGroup", "myAzureLargeInstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armlargeinstance.ClientDeleteResponse{
	// }
}

// Generated from example definition: 2024-08-01-preview/AzureLargeInstance_Get.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("f0f4887f-d13c-4943-a8ba-d7da28d2a3fd", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "myResourceGroup", "myAzureLargeInstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armlargeinstance.ClientGetResponse{
	// 	AzureLargeInstance: &armlargeinstance.AzureLargeInstance{
	// 		ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myAzureLargeInstance"),
	// 		Location: to.Ptr("westus2"),
	// 		Name: to.Ptr("myAzureLargeInstance"),
	// 		Tags: map[string]*string{
	// 			"key": to.Ptr("value"),
	// 		},
	// 		Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeInstances"),
	// 		SystemData: &armlargeinstance.SystemData{
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-20T23:10:22.6828621Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.0000000Z"); return t}()),
	// 		},
	// 		Properties: &armlargeinstance.Properties{
	// 			AzureLargeInstanceID: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
	// 			PowerState: to.Ptr(armlargeinstance.PowerStateEnumRestarting),
	// 			HwRevision: to.Ptr("Rev 4.2"),
	// 			HardwareProfile: &armlargeinstance.HardwareProfile{
	// 				HardwareType: to.Ptr(armlargeinstance.HardwareTypeNamesEnumCiscoUCS),
	// 				AzureLargeInstanceSize: to.Ptr(armlargeinstance.SizeNamesEnumS72),
	// 			},
	// 			NetworkProfile: &armlargeinstance.NetworkProfile{
	// 				NetworkInterfaces: []*armlargeinstance.IPAddress{
	// 					{
	// 						IPAddress: to.Ptr("123.123.123.123"),
	// 					},
	// 				},
	// 				CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuits/myCircuitId"),
	// 			},
	// 			StorageProfile: &armlargeinstance.StorageProfile{
	// 				NfsIPAddress: to.Ptr("123.123.119.123"),
	// 			},
	// 			OSProfile: &armlargeinstance.OsProfile{
	// 				ComputerName: to.Ptr("myComputerName"),
	// 				OSType: to.Ptr("SLES 12 SP2"),
	// 				Version: to.Ptr("12 SP2"),
	// 			},
	// 			ProvisioningState: to.Ptr(armlargeinstance.ProvisioningStatesEnumSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-08-01-preview/AzureLargeInstance_ListByResourceGroup.json
func ExampleClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("f0f4887f-d13c-4943-a8ba-d7da28d2a3fd", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page = armlargeinstance.ClientListByResourceGroupResponse{
		// 	ListResult: armlargeinstance.ListResult{
		// 		Value: []*armlargeinstance.AzureLargeInstance{
		// 			{
		// 				ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myAzureLargeMetalInstance1"),
		// 				Name: to.Ptr("myAzureLargeMetalInstance1"),
		// 				Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeInstances"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"key": to.Ptr("value"),
		// 				},
		// 				SystemData: &armlargeinstance.SystemData{
		// 					CreatedBy: to.Ptr("user@microsoft.com"),
		// 					CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-20T23:10:22.6828621Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.0000000Z"); return t}()),
		// 				},
		// 				Properties: &armlargeinstance.Properties{
		// 					AzureLargeInstanceID: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
		// 					PowerState: to.Ptr(armlargeinstance.PowerStateEnumStarted),
		// 					ProximityPlacementGroup: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Compute/proximityPlacementGroups/myplacementgroup"),
		// 					HwRevision: to.Ptr("Rev 3"),
		// 					HardwareProfile: &armlargeinstance.HardwareProfile{
		// 						HardwareType: to.Ptr(armlargeinstance.HardwareTypeNamesEnumCiscoUCS),
		// 						AzureLargeInstanceSize: to.Ptr(armlargeinstance.SizeNamesEnumS72),
		// 					},
		// 					NetworkProfile: &armlargeinstance.NetworkProfile{
		// 						NetworkInterfaces: []*armlargeinstance.IPAddress{
		// 							{
		// 								IPAddress: to.Ptr("100.100.100.100"),
		// 							},
		// 						},
		// 						CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuit"),
		// 					},
		// 					StorageProfile: &armlargeinstance.StorageProfile{
		// 						NfsIPAddress: to.Ptr("200.200.200.200"),
		// 					},
		// 					OSProfile: &armlargeinstance.OsProfile{
		// 						ComputerName: to.Ptr("myComputerName1"),
		// 						OSType: to.Ptr("SUSE"),
		// 						Version: to.Ptr("12 SP1"),
		// 						SSHPublicKey: to.Ptr("{ssh-rsa public key}"),
		// 					},
		// 					ProvisioningState: to.Ptr(armlargeinstance.ProvisioningStatesEnumSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myABMInstance2"),
		// 				Name: to.Ptr("myABMInstance2"),
		// 				Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeInstances"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"key": to.Ptr("value"),
		// 				},
		// 				SystemData: &armlargeinstance.SystemData{
		// 					CreatedBy: to.Ptr("user@microsoft.com"),
		// 					CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.0000000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-13T08:01:22.1234567Z"); return t}()),
		// 				},
		// 				Properties: &armlargeinstance.Properties{
		// 					AzureLargeInstanceID: to.Ptr("589bce49-9fe6-4dc8-82df-cf6ae25e0cb9"),
		// 					PowerState: to.Ptr(armlargeinstance.PowerStateEnumStarted),
		// 					ProximityPlacementGroup: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Compute/proximityPlacementGroups/myplacementgroup"),
		// 					HwRevision: to.Ptr("Rev 3"),
		// 					HardwareProfile: &armlargeinstance.HardwareProfile{
		// 						HardwareType: to.Ptr(armlargeinstance.HardwareTypeNamesEnumHPE),
		// 						AzureLargeInstanceSize: to.Ptr(armlargeinstance.SizeNamesEnumS384),
		// 					},
		// 					NetworkProfile: &armlargeinstance.NetworkProfile{
		// 						NetworkInterfaces: []*armlargeinstance.IPAddress{
		// 							{
		// 								IPAddress: to.Ptr("100.100.100.101"),
		// 							},
		// 						},
		// 						CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuit"),
		// 					},
		// 					StorageProfile: &armlargeinstance.StorageProfile{
		// 						NfsIPAddress: to.Ptr("200.200.200.201"),
		// 					},
		// 					OSProfile: &armlargeinstance.OsProfile{
		// 						ComputerName: to.Ptr("myComputerName2"),
		// 						OSType: to.Ptr("SUSE"),
		// 						Version: to.Ptr("12 SP1"),
		// 						SSHPublicKey: to.Ptr("{ssh-rsa public key}"),
		// 					},
		// 					ProvisioningState: to.Ptr(armlargeinstance.ProvisioningStatesEnumSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2024-08-01-preview/AzureLargeInstance_ListBySubscription.json
func ExampleClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("f0f4887f-d13c-4943-a8ba-d7da28d2a3fd", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListBySubscriptionPager(nil)
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
		// page = armlargeinstance.ClientListBySubscriptionResponse{
		// 	ListResult: armlargeinstance.ListResult{
		// 		Value: []*armlargeinstance.AzureLargeInstance{
		// 			{
		// 				ID: to.Ptr("/subscriptions/57d3422f-467a-448e-b798-ebf490849542/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myAzureLargeInstance1"),
		// 				Location: to.Ptr("westus2"),
		// 				Name: to.Ptr("myAzureLargeInstance1"),
		// 				Tags: map[string]*string{
		// 					"key": to.Ptr("value"),
		// 				},
		// 				SystemData: &armlargeinstance.SystemData{
		// 					CreatedBy: to.Ptr("user@microsoft.com"),
		// 					CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.0000000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-13T08:01:22.1234567Z"); return t}()),
		// 				},
		// 				Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeInstances"),
		// 				Properties: &armlargeinstance.Properties{
		// 					AzureLargeInstanceID: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
		// 					PowerState: to.Ptr(armlargeinstance.PowerStateEnumRestarting),
		// 					HwRevision: to.Ptr("Rev 4.2"),
		// 					HardwareProfile: &armlargeinstance.HardwareProfile{
		// 						HardwareType: to.Ptr(armlargeinstance.HardwareTypeNamesEnumCiscoUCS),
		// 						AzureLargeInstanceSize: to.Ptr(armlargeinstance.SizeNamesEnumS72),
		// 					},
		// 					NetworkProfile: &armlargeinstance.NetworkProfile{
		// 						NetworkInterfaces: []*armlargeinstance.IPAddress{
		// 							{
		// 								IPAddress: to.Ptr("123.123.123.123"),
		// 							},
		// 						},
		// 						CircuitID: to.Ptr("/subscriptions/57d3422f-467a-448e-b798-ebf490849542/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuits/myCircuitId"),
		// 					},
		// 					StorageProfile: &armlargeinstance.StorageProfile{
		// 						NfsIPAddress: to.Ptr("123.123.119.123"),
		// 					},
		// 					OSProfile: &armlargeinstance.OsProfile{
		// 						ComputerName: to.Ptr("myComputerName"),
		// 						OSType: to.Ptr("SLES 12 SP2"),
		// 						Version: to.Ptr("12 SP2"),
		// 					},
		// 					ProvisioningState: to.Ptr(armlargeinstance.ProvisioningStatesEnumSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myAzureLargeInstance2"),
		// 				Location: to.Ptr("westus2"),
		// 				Name: to.Ptr("myAzureLargeInstance2"),
		// 				Tags: map[string]*string{
		// 					"key": to.Ptr("value"),
		// 				},
		// 				SystemData: &armlargeinstance.SystemData{
		// 					CreatedBy: to.Ptr("user@microsoft.com"),
		// 					CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.0000000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-13T08:01:22.1234567Z"); return t}()),
		// 				},
		// 				Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeInstances"),
		// 				Properties: &armlargeinstance.Properties{
		// 					AzureLargeInstanceID: to.Ptr("589bce49-9fe6-4dc8-82df-cf6ae25e0cb9"),
		// 					PowerState: to.Ptr(armlargeinstance.PowerStateEnumRestarting),
		// 					HwRevision: to.Ptr("Rev 4.2"),
		// 					HardwareProfile: &armlargeinstance.HardwareProfile{
		// 						HardwareType: to.Ptr(armlargeinstance.HardwareTypeNamesEnumCiscoUCS),
		// 						AzureLargeInstanceSize: to.Ptr(armlargeinstance.SizeNamesEnumS72),
		// 					},
		// 					NetworkProfile: &armlargeinstance.NetworkProfile{
		// 						NetworkInterfaces: []*armlargeinstance.IPAddress{
		// 							{
		// 								IPAddress: to.Ptr("123.123.123.123"),
		// 							},
		// 						},
		// 						CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuits/myCircuitId"),
		// 					},
		// 					StorageProfile: &armlargeinstance.StorageProfile{
		// 						NfsIPAddress: to.Ptr("123.123.119.123"),
		// 					},
		// 					OSProfile: &armlargeinstance.OsProfile{
		// 						ComputerName: to.Ptr("myComputerName2"),
		// 						OSType: to.Ptr("SLES 12 SP2"),
		// 						Version: to.Ptr("12 SP2"),
		// 					},
		// 					ProvisioningState: to.Ptr(armlargeinstance.ProvisioningStatesEnumSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2024-08-01-preview/AzureLargeInstance_Restart.json
func ExampleClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginRestart(ctx, "myResourceGroup", "myALInstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armlargeinstance.ClientRestartResponse{
	// }
}

// Generated from example definition: 2024-08-01-preview/AzureLargeInstance_Shutdown.json
func ExampleClient_BeginShutdown() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginShutdown(ctx, "myResourceGroup", "myALInstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armlargeinstance.ClientShutdownResponse{
	// 	OperationStatusResult: &armlargeinstance.OperationStatusResult{
	// 		Name: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 		Status: to.Ptr("InProgress"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-04T21:17:24.9052926Z"); return t}()),
	// 		Error: &armlargeinstance.ErrorDetail{
	// 			Code: to.Ptr(""),
	// 			Message: to.Ptr(""),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-08-01-preview/AzureLargeInstance_Start.json
func ExampleClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginStart(ctx, "myResourceGroup", "myALInstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armlargeinstance.ClientStartResponse{
	// 	OperationStatusResult: &armlargeinstance.OperationStatusResult{
	// 		Name: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 		Status: to.Ptr("InProgress"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-04T21:17:24.9052926Z"); return t}()),
	// 		Error: &armlargeinstance.ErrorDetail{
	// 			Code: to.Ptr(""),
	// 			Message: to.Ptr(""),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-08-01-preview/AzureLargeInstance_PatchTags.json
func ExampleClient_Update_azureLargeInstanceUpdateTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("f0f4887f-d13c-4943-a8ba-d7da28d2a3fd", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Update(ctx, "myResourceGroup", "myALInstance", armlargeinstance.TagsUpdate{
		Tags: map[string]*string{
			"testkey": to.Ptr("testvalue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armlargeinstance.ClientUpdateResponse{
	// 	AzureLargeInstance: &armlargeinstance.AzureLargeInstance{
	// 		ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myALInstance"),
	// 		Name: to.Ptr("myALInstance"),
	// 		Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeInstances"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"testkey": to.Ptr("testvalue"),
	// 		},
	// 		SystemData: &armlargeinstance.SystemData{
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-20T23:10:22.6828621Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.0000000Z"); return t}()),
	// 		},
	// 		Properties: &armlargeinstance.Properties{
	// 			AzureLargeInstanceID: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
	// 			PowerState: to.Ptr(armlargeinstance.PowerStateEnumStarted),
	// 			ProximityPlacementGroup: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Compute/proximityPlacementGroups/myplacementgroup"),
	// 			HwRevision: to.Ptr("Rev 3"),
	// 			HardwareProfile: &armlargeinstance.HardwareProfile{
	// 				HardwareType: to.Ptr(armlargeinstance.HardwareTypeNamesEnumCiscoUCS),
	// 				AzureLargeInstanceSize: to.Ptr(armlargeinstance.SizeNamesEnumS72),
	// 			},
	// 			NetworkProfile: &armlargeinstance.NetworkProfile{
	// 				NetworkInterfaces: []*armlargeinstance.IPAddress{
	// 					{
	// 						IPAddress: to.Ptr("100.100.100.100"),
	// 					},
	// 				},
	// 				CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuit"),
	// 			},
	// 			StorageProfile: &armlargeinstance.StorageProfile{
	// 				NfsIPAddress: to.Ptr("200.200.200.200"),
	// 			},
	// 			OSProfile: &armlargeinstance.OsProfile{
	// 				ComputerName: to.Ptr("myComputerName"),
	// 				OSType: to.Ptr("SUSE"),
	// 				Version: to.Ptr("12 SP1"),
	// 				SSHPublicKey: to.Ptr("{ssh-rsa public key}"),
	// 			},
	// 			ProvisioningState: to.Ptr(armlargeinstance.ProvisioningStatesEnumSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-08-01-preview/AzureLargeInstance_PatchTags_Delete.json
func ExampleClient_Update_azureLargeInstanceDeleteTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("f0f4887f-d13c-4943-a8ba-d7da28d2a3fd", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Update(ctx, "myResourceGroup", "myALInstance", armlargeinstance.TagsUpdate{
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armlargeinstance.ClientUpdateResponse{
	// 	AzureLargeInstance: &armlargeinstance.AzureLargeInstance{
	// 		ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myALInstance"),
	// 		Name: to.Ptr("myALInstance"),
	// 		Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeInstances"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		SystemData: &armlargeinstance.SystemData{
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-20T23:10:22.6828621Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.0000000Z"); return t}()),
	// 		},
	// 		Properties: &armlargeinstance.Properties{
	// 			AzureLargeInstanceID: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
	// 			PowerState: to.Ptr(armlargeinstance.PowerStateEnumStarted),
	// 			ProximityPlacementGroup: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Compute/proximityPlacementGroups/myplacementgroup"),
	// 			HwRevision: to.Ptr("Rev 3"),
	// 			HardwareProfile: &armlargeinstance.HardwareProfile{
	// 				HardwareType: to.Ptr(armlargeinstance.HardwareTypeNamesEnumCiscoUCS),
	// 				AzureLargeInstanceSize: to.Ptr(armlargeinstance.SizeNamesEnumS72),
	// 			},
	// 			NetworkProfile: &armlargeinstance.NetworkProfile{
	// 				NetworkInterfaces: []*armlargeinstance.IPAddress{
	// 					{
	// 						IPAddress: to.Ptr("100.100.100.100"),
	// 					},
	// 				},
	// 				CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuit"),
	// 			},
	// 			StorageProfile: &armlargeinstance.StorageProfile{
	// 				NfsIPAddress: to.Ptr("200.200.200.200"),
	// 			},
	// 			OSProfile: &armlargeinstance.OsProfile{
	// 				ComputerName: to.Ptr("myComputerName"),
	// 				OSType: to.Ptr("SUSE"),
	// 				Version: to.Ptr("12 SP1"),
	// 				SSHPublicKey: to.Ptr("{ssh-rsa public key}"),
	// 			},
	// 			ProvisioningState: to.Ptr(armlargeinstance.ProvisioningStatesEnumSucceeded),
	// 		},
	// 	},
	// }
}
