// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package methodsubscriptionidgroup_test

import (
	"context"
	"methodsubscriptionidgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
)

var (
	ctx           context.Context
	clientFactory *methodsubscriptionidgroup.ClientFactory

	subscriptionIdExpected = "00000000-0000-0000-0000-000000000000"
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	clientFactory, _ = methodsubscriptionidgroup.NewClientFactory(subscriptionIdExpected, &azfake.TokenCredential{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud.Configuration{
				Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
					cloud.ResourceManager: {
						Audience: "fake_audience",
						Endpoint: "http://localhost:3000",
					},
				},
			},
			InsecureAllowCredentialWithHTTP: true,
		},
	})

	m.Run()
}
