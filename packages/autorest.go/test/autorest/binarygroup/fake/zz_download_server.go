// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/binarygroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DownloadServer is a fake server for instances of the binarygroup.DownloadClient type.
type DownloadServer struct {
	// ErrorStream is the fake for method DownloadClient.ErrorStream
	// HTTP status codes to indicate success: http.StatusOK
	ErrorStream func(ctx context.Context, options *binarygroup.DownloadClientErrorStreamOptions) (resp azfake.Responder[binarygroup.DownloadClientErrorStreamResponse], errResp azfake.ErrorResponder)
}

// NewDownloadServerTransport creates a new instance of DownloadServerTransport with the provided implementation.
// The returned DownloadServerTransport instance is connected to an instance of binarygroup.DownloadClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDownloadServerTransport(srv *DownloadServer) *DownloadServerTransport {
	return &DownloadServerTransport{srv: srv}
}

// DownloadServerTransport connects instances of binarygroup.DownloadClient to instances of DownloadServer.
// Don't use this type directly, use NewDownloadServerTransport instead.
type DownloadServerTransport struct {
	srv *DownloadServer
}

// Do implements the policy.Transporter interface for DownloadServerTransport.
func (d *DownloadServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DownloadServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if downloadServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = downloadServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DownloadClient.ErrorStream":
				res.resp, res.err = d.dispatchErrorStream(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (d *DownloadServerTransport) dispatchErrorStream(req *http.Request) (*http.Response, error) {
	if d.srv.ErrorStream == nil {
		return nil, &nonRetriableError{errors.New("fake for method ErrorStream not implemented")}
	}
	respr, errRespr := d.srv.ErrorStream(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: req.Header.Get("Content-Type"),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DownloadServerTransport
var downloadServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
