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
	"lrorpcgroup"
	"net/http"
)

// RpcServer is a fake server for instances of the lrorpcgroup.RpcClient type.
type RpcServer struct {
	// BeginLongRunningRPC is the fake for method RpcClient.BeginLongRunningRPC
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginLongRunningRPC func(ctx context.Context, body lrorpcgroup.GenerationOptions, options *lrorpcgroup.RpcClientLongRunningRPCOptions) (resp azfake.PollerResponder[lrorpcgroup.RpcClientLongRunningRPCResponse], errResp azfake.ErrorResponder)
}

// NewRpcServerTransport creates a new instance of RpcServerTransport with the provided implementation.
// The returned RpcServerTransport instance is connected to an instance of lrorpcgroup.RpcClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRpcServerTransport(srv *RpcServer) *RpcServerTransport {
	return &RpcServerTransport{
		srv:                 srv,
		beginLongRunningRPC: newTracker[azfake.PollerResponder[lrorpcgroup.RpcClientLongRunningRPCResponse]](),
	}
}

// RpcServerTransport connects instances of lrorpcgroup.RpcClient to instances of RpcServer.
// Don't use this type directly, use NewRpcServerTransport instead.
type RpcServerTransport struct {
	srv                 *RpcServer
	beginLongRunningRPC *tracker[azfake.PollerResponder[lrorpcgroup.RpcClientLongRunningRPCResponse]]
}

// Do implements the policy.Transporter interface for RpcServerTransport.
func (r *RpcServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RpcServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "RpcClient.BeginLongRunningRPC":
		resp, err = r.dispatchBeginLongRunningRPC(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (r *RpcServerTransport) dispatchBeginLongRunningRPC(req *http.Request) (*http.Response, error) {
	if r.srv.BeginLongRunningRPC == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginLongRunningRPC not implemented")}
	}
	beginLongRunningRPC := r.beginLongRunningRPC.get(req)
	if beginLongRunningRPC == nil {
		body, err := server.UnmarshalRequestAsJSON[lrorpcgroup.GenerationOptions](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginLongRunningRPC(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginLongRunningRPC = &respr
		r.beginLongRunningRPC.add(req, beginLongRunningRPC)
	}

	resp, err := server.PollerResponderNext(beginLongRunningRPC, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		r.beginLongRunningRPC.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginLongRunningRPC) {
		r.beginLongRunningRPC.remove(req)
	}

	return resp, nil
}
