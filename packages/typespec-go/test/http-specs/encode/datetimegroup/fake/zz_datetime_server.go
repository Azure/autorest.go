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

// DatetimeServer is a fake server for instances of the datetimegroup.DatetimeClient type.
type DatetimeServer struct {
	// DatetimeHeaderServer contains the fakes for client DatetimeHeaderClient
	DatetimeHeaderServer DatetimeHeaderServer

	// DatetimePropertyServer contains the fakes for client DatetimePropertyClient
	DatetimePropertyServer DatetimePropertyServer

	// DatetimeQueryServer contains the fakes for client DatetimeQueryClient
	DatetimeQueryServer DatetimeQueryServer

	// DatetimeResponseHeaderServer contains the fakes for client DatetimeResponseHeaderClient
	DatetimeResponseHeaderServer DatetimeResponseHeaderServer
}

// NewDatetimeServerTransport creates a new instance of DatetimeServerTransport with the provided implementation.
// The returned DatetimeServerTransport instance is connected to an instance of datetimegroup.DatetimeClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatetimeServerTransport(srv *DatetimeServer) *DatetimeServerTransport {
	return &DatetimeServerTransport{srv: srv}
}

// DatetimeServerTransport connects instances of datetimegroup.DatetimeClient to instances of DatetimeServer.
// Don't use this type directly, use NewDatetimeServerTransport instead.
type DatetimeServerTransport struct {
	srv                            *DatetimeServer
	trMu                           sync.Mutex
	trDatetimeHeaderServer         *DatetimeHeaderServerTransport
	trDatetimePropertyServer       *DatetimePropertyServerTransport
	trDatetimeQueryServer          *DatetimeQueryServerTransport
	trDatetimeResponseHeaderServer *DatetimeResponseHeaderServerTransport
}

// Do implements the policy.Transporter interface for DatetimeServerTransport.
func (d *DatetimeServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (d *DatetimeServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "DatetimeHeaderClient":
		initServer(&d.trMu, &d.trDatetimeHeaderServer, func() *DatetimeHeaderServerTransport {
			return NewDatetimeHeaderServerTransport(&d.srv.DatetimeHeaderServer)
		})
		resp, err = d.trDatetimeHeaderServer.Do(req)
	case "DatetimePropertyClient":
		initServer(&d.trMu, &d.trDatetimePropertyServer, func() *DatetimePropertyServerTransport {
			return NewDatetimePropertyServerTransport(&d.srv.DatetimePropertyServer)
		})
		resp, err = d.trDatetimePropertyServer.Do(req)
	case "DatetimeQueryClient":
		initServer(&d.trMu, &d.trDatetimeQueryServer, func() *DatetimeQueryServerTransport {
			return NewDatetimeQueryServerTransport(&d.srv.DatetimeQueryServer)
		})
		resp, err = d.trDatetimeQueryServer.Do(req)
	case "DatetimeResponseHeaderClient":
		initServer(&d.trMu, &d.trDatetimeResponseHeaderServer, func() *DatetimeResponseHeaderServerTransport {
			return NewDatetimeResponseHeaderServerTransport(&d.srv.DatetimeResponseHeaderServer)
		})
		resp, err = d.trDatetimeResponseHeaderServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}

// set this to conditionally intercept incoming requests to DatetimeServerTransport
var datetimeServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
