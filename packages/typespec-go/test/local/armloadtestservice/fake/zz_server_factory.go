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

// ServerFactory is a fake server for instances of the armloadtestservice.ClientFactory type.
type ServerFactory struct {
	// LoadTestMappingsServer contains the fakes for client LoadTestMappingsClient
	LoadTestMappingsServer LoadTestMappingsServer

	// LoadTestProfileMappingsServer contains the fakes for client LoadTestProfileMappingsClient
	LoadTestProfileMappingsServer LoadTestProfileMappingsServer

	// LoadTestsServer contains the fakes for client LoadTestsClient
	LoadTestsServer LoadTestsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// QuotasServer contains the fakes for client QuotasClient
	QuotasServer QuotasServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armloadtestservice.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armloadtestservice.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                             *ServerFactory
	trMu                            sync.Mutex
	trLoadTestMappingsServer        *LoadTestMappingsServerTransport
	trLoadTestProfileMappingsServer *LoadTestProfileMappingsServerTransport
	trLoadTestsServer               *LoadTestsServerTransport
	trOperationsServer              *OperationsServerTransport
	trQuotasServer                  *QuotasServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "LoadTestMappingsClient":
		initServer(s, &s.trLoadTestMappingsServer, func() *LoadTestMappingsServerTransport {
			return NewLoadTestMappingsServerTransport(&s.srv.LoadTestMappingsServer)
		})
		resp, err = s.trLoadTestMappingsServer.Do(req)
	case "LoadTestProfileMappingsClient":
		initServer(s, &s.trLoadTestProfileMappingsServer, func() *LoadTestProfileMappingsServerTransport {
			return NewLoadTestProfileMappingsServerTransport(&s.srv.LoadTestProfileMappingsServer)
		})
		resp, err = s.trLoadTestProfileMappingsServer.Do(req)
	case "LoadTestsClient":
		initServer(s, &s.trLoadTestsServer, func() *LoadTestsServerTransport { return NewLoadTestsServerTransport(&s.srv.LoadTestsServer) })
		resp, err = s.trLoadTestsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "QuotasClient":
		initServer(s, &s.trQuotasServer, func() *QuotasServerTransport { return NewQuotasServerTransport(&s.srv.QuotasServer) })
		resp, err = s.trQuotasServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}