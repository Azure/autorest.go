// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// MultiPartServer is a fake server for instances of the multipartgroup.MultiPartClient type.
type MultiPartServer struct {
	// MultiPartFormDataServer contains the fakes for client MultiPartFormDataClient
	MultiPartFormDataServer MultiPartFormDataServer
}

// NewMultiPartServerTransport creates a new instance of MultiPartServerTransport with the provided implementation.
// The returned MultiPartServerTransport instance is connected to an instance of multipartgroup.MultiPartClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMultiPartServerTransport(srv *MultiPartServer) *MultiPartServerTransport {
	return &MultiPartServerTransport{srv: srv}
}

// MultiPartServerTransport connects instances of multipartgroup.MultiPartClient to instances of MultiPartServer.
// Don't use this type directly, use NewMultiPartServerTransport instead.
type MultiPartServerTransport struct {
	srv                       *MultiPartServer
	trMu                      sync.Mutex
	trMultiPartFormDataServer *MultiPartFormDataServerTransport
}

// Do implements the policy.Transporter interface for MultiPartServerTransport.
func (m *MultiPartServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (m *MultiPartServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "MultiPartFormDataClient":
		initServer(&m.trMu, &m.trMultiPartFormDataServer, func() *MultiPartFormDataServerTransport {
			return NewMultiPartFormDataServerTransport(&m.srv.MultiPartFormDataServer)
		})
		resp, err = m.trMultiPartFormDataServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}
