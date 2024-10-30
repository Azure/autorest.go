// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package fake

import "net/http"

type interceptorFunc func(*http.Request) (*http.Response, error, bool)

func (i interceptorFunc) Do(req *http.Request) (*http.Response, error, bool) {
	return i(req)
}

func SetContainerRegistryServerInterceptor(interceptor func(*http.Request) (*http.Response, error, bool)) {
	if interceptor == nil {
		containerRegistryServerTransportInterceptor = nil
	} else {
		containerRegistryServerTransportInterceptor = interceptorFunc(interceptor)
	}
}
