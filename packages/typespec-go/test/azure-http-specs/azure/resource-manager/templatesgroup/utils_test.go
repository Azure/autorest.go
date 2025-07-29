// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package templatesgroup_test

import (
	"context"
	"fmt"
	"net/http"
	"templatesgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

var (
	ctx           context.Context
	clientFactory *templatesgroup.ClientFactory

	subscriptionIdExpected = "00000000-0000-0000-0000-000000000000"
	locationExpected       = "eastus"
	resourceGroupExpected  = "test-rg"
	widgetName             = "widget1"
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	clientFactory, _ = templatesgroup.NewClientFactory(subscriptionIdExpected, &azfake.TokenCredential{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			InsecureAllowCredentialWithHTTP: true,
			PerCallPolicies: []policy.Policy{
				&cadlranchPolicy{},
			},
		},
	})

	m.Run()
}

type cadlranchPolicy struct {
	useHttps  bool
	proxyPort int
}

func (p cadlranchPolicy) scheme() string {
	if p.useHttps {
		return "https"
	}
	return "http"
}

func (p cadlranchPolicy) host() string {
	if p.proxyPort != 0 {
		return fmt.Sprintf("localhost:%d", p.proxyPort)
	}
	return "localhost:3000"
}

func (p *cadlranchPolicy) Do(req *policy.Request) (*http.Response, error) {
	oriSchema := req.Raw().URL.Scheme
	oriHost := req.Raw().URL.Host

	// don't modify the original request
	cp := *req

	cpURL := *cp.Raw().URL
	cp.Raw().URL = &cpURL
	cp.Raw().Header = req.Raw().Header.Clone()

	cp.Raw().URL.Scheme = p.scheme()
	cp.Raw().URL.Host = p.host()
	cp.Raw().Host = p.host()
	req = &cp

	resp, err := req.Next()
	if resp != nil {
		resp.Request.URL.Scheme = oriSchema
		resp.Request.URL.Host = oriHost
	}

	return resp, err
}
