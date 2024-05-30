// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apikeygroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewAPIKeyClientWithKeyCredential(cred *azcore.KeyCredential, options *azcore.ClientOptions) (*APIKeyClient, error) {
	internal, err := azcore.NewClient("apikeygroup", "v0.1.0", runtime.PipelineOptions{
		PerCall: []policy.Policy{
			runtime.NewKeyCredentialPolicy(cred, "x-ms-api-key", &runtime.KeyCredentialPolicyOptions{
				InsecureAllowCredentialWithHTTP: true,
			}),
		},
	}, options)
	if err != nil {
		return nil, err
	}
	return &APIKeyClient{
		internal: internal,
	}, nil
}
