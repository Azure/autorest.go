// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package oauth2group

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewOAuth2Client(cred azcore.TokenCredential, options *azcore.ClientOptions) (*OAuth2Client, error) {
	internal, err := azcore.NewClient("oauth2group", "v0.1.0", runtime.PipelineOptions{
		PerCall: []policy.Policy{
			runtime.NewBearerTokenPolicy(cred, []string{
				"https://security.microsoft.com/.default",
			}, &policy.BearerTokenOptions{
				InsecureAllowCredentialWithHTTP: true,
			}),
		},
	}, options)
	if err != nil {
		return nil, err
	}
	return &OAuth2Client{
		internal: internal,
	}, nil
}
