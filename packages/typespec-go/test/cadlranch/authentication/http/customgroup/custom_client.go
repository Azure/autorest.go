// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package customgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewCustomClientWithKeyCredential(cred *azcore.KeyCredential, options *azcore.ClientOptions) (*CustomClient, error) {
	internal, err := azcore.NewClient("customgroup", "v0.1.0", runtime.PipelineOptions{
		PerCall: []policy.Policy{
			runtime.NewKeyCredentialPolicy(cred, "Authorization", &runtime.KeyCredentialPolicyOptions{
				InsecureAllowCredentialWithHTTP: true,
				Prefix:                          "SharedAccessKey ",
			}),
		},
	}, options)
	if err != nil {
		return nil, err
	}
	return &CustomClient{
		internal: internal,
	}, nil
}
