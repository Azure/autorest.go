// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"projectednamegroup"
	"strings"
	"sync"
)

// ProjectedNameServer is a fake server for instances of the projectednamegroup.ProjectedNameClient type.
type ProjectedNameServer struct {
	// ModelServer contains the fakes for client ModelClient
	ModelServer ModelServer

	// PropertyServer contains the fakes for client PropertyClient
	PropertyServer PropertyServer

	// ClientName is the fake for method ProjectedNameClient.ClientName
	// HTTP status codes to indicate success: http.StatusNoContent
	ClientName func(ctx context.Context, options *projectednamegroup.ProjectedNameClientClientNameOptions) (resp azfake.Responder[projectednamegroup.ProjectedNameClientClientNameResponse], errResp azfake.ErrorResponder)

	// Parameter is the fake for method ProjectedNameClient.Parameter
	// HTTP status codes to indicate success: http.StatusNoContent
	Parameter func(ctx context.Context, clientName string, options *projectednamegroup.ProjectedNameClientParameterOptions) (resp azfake.Responder[projectednamegroup.ProjectedNameClientParameterResponse], errResp azfake.ErrorResponder)
}

// NewProjectedNameServerTransport creates a new instance of ProjectedNameServerTransport with the provided implementation.
// The returned ProjectedNameServerTransport instance is connected to an instance of projectednamegroup.ProjectedNameClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProjectedNameServerTransport(srv *ProjectedNameServer) *ProjectedNameServerTransport {
	return &ProjectedNameServerTransport{srv: srv}
}

// ProjectedNameServerTransport connects instances of projectednamegroup.ProjectedNameClient to instances of ProjectedNameServer.
// Don't use this type directly, use NewProjectedNameServerTransport instead.
type ProjectedNameServerTransport struct {
	srv              *ProjectedNameServer
	trMu             sync.Mutex
	trModelServer    *ModelServerTransport
	trPropertyServer *PropertyServerTransport
}

// Do implements the policy.Transporter interface for ProjectedNameServerTransport.
func (p *ProjectedNameServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	if client := method[:strings.Index(method, ".")]; client != "ProjectedNameClient" {
		return p.dispatchToClientFake(req, client)
	}
	return p.dispatchToMethodFake(req, method)
}

func (p *ProjectedNameServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "ModelClient":
		initServer(&p.trMu, &p.trModelServer, func() *ModelServerTransport {
			return NewModelServerTransport(&p.srv.ModelServer)
		})
		resp, err = p.trModelServer.Do(req)
	case "PropertyClient":
		initServer(&p.trMu, &p.trPropertyServer, func() *PropertyServerTransport {
			return NewPropertyServerTransport(&p.srv.PropertyServer)
		})
		resp, err = p.trPropertyServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}

func (p *ProjectedNameServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ProjectedNameClient.ClientName":
		resp, err = p.dispatchClientName(req)
	case "ProjectedNameClient.Parameter":
		resp, err = p.dispatchParameter(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (p *ProjectedNameServerTransport) dispatchClientName(req *http.Request) (*http.Response, error) {
	if p.srv.ClientName == nil {
		return nil, &nonRetriableError{errors.New("fake for method ClientName not implemented")}
	}
	respr, errRespr := p.srv.ClientName(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProjectedNameServerTransport) dispatchParameter(req *http.Request) (*http.Response, error) {
	if p.srv.Parameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method Parameter not implemented")}
	}
	qp := req.URL.Query()
	clientNameParam, err := url.QueryUnescape(qp.Get("default-name"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Parameter(req.Context(), clientNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
