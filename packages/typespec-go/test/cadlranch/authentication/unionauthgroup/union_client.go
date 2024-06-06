// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package unionauthgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type UnionClientOptions struct {
	azcore.ClientOptions
}

func NewUnionClient(cred azcore.TokenCredential, options *UnionClientOptions) (*UnionClient, error) {
	return newUnionClient(runtime.NewBearerTokenPolicy(cred, []string{
		"https://security.microsoft.com/.default",
	}, &policy.BearerTokenOptions{
		InsecureAllowCredentialWithHTTP: true,
	}), options)
}

func NewUnionClientWithKeyCredential(cred *azcore.KeyCredential, options *UnionClientOptions) (*UnionClient, error) {
	return newUnionClient(runtime.NewKeyCredentialPolicy(cred, "x-ms-api-key", &runtime.KeyCredentialPolicyOptions{
		InsecureAllowCredentialWithHTTP: true,
	}), options)
}

func newUnionClient(credPolicy policy.Policy, options *UnionClientOptions) (*UnionClient, error) {
	if options == nil {
		options = &UnionClientOptions{}
	}
	internal, err := azcore.NewClient("unionauthgroup", "v0.1.0", runtime.PipelineOptions{
		PerCall: []policy.Policy{credPolicy},
	}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	return &UnionClient{
		internal: internal,
	}, nil
}
