// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"armbillingbenefits"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"regexp"
)

// ReservationOrderAliasResponsesServer is a fake server for instances of the armbillingbenefits.ReservationOrderAliasResponsesClient type.
type ReservationOrderAliasResponsesServer struct {
	// BeginCreate is the fake for method ReservationOrderAliasResponsesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, reservationOrderAliasName string, body armbillingbenefits.ReservationOrderAliasRequest, options *armbillingbenefits.ReservationOrderAliasResponsesClientBeginCreateOptions) (resp azfake.PollerResponder[armbillingbenefits.ReservationOrderAliasResponsesClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ReservationOrderAliasResponsesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, reservationOrderAliasName string, options *armbillingbenefits.ReservationOrderAliasResponsesClientGetOptions) (resp azfake.Responder[armbillingbenefits.ReservationOrderAliasResponsesClientGetResponse], errResp azfake.ErrorResponder)
}

// NewReservationOrderAliasResponsesServerTransport creates a new instance of ReservationOrderAliasResponsesServerTransport with the provided implementation.
// The returned ReservationOrderAliasResponsesServerTransport instance is connected to an instance of armbillingbenefits.ReservationOrderAliasResponsesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReservationOrderAliasResponsesServerTransport(srv *ReservationOrderAliasResponsesServer) *ReservationOrderAliasResponsesServerTransport {
	return &ReservationOrderAliasResponsesServerTransport{
		srv:         srv,
		beginCreate: newTracker[azfake.PollerResponder[armbillingbenefits.ReservationOrderAliasResponsesClientCreateResponse]](),
	}
}

// ReservationOrderAliasResponsesServerTransport connects instances of armbillingbenefits.ReservationOrderAliasResponsesClient to instances of ReservationOrderAliasResponsesServer.
// Don't use this type directly, use NewReservationOrderAliasResponsesServerTransport instead.
type ReservationOrderAliasResponsesServerTransport struct {
	srv         *ReservationOrderAliasResponsesServer
	beginCreate *tracker[azfake.PollerResponder[armbillingbenefits.ReservationOrderAliasResponsesClientCreateResponse]]
}

// Do implements the policy.Transporter interface for ReservationOrderAliasResponsesServerTransport.
func (r *ReservationOrderAliasResponsesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ReservationOrderAliasResponsesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if reservationOrderAliasResponsesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = reservationOrderAliasResponsesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ReservationOrderAliasResponsesClient.BeginCreate":
				res.resp, res.err = r.dispatchBeginCreate(req)
			case "ReservationOrderAliasResponsesClient.Get":
				res.resp, res.err = r.dispatchGet(req)
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

func (r *ReservationOrderAliasResponsesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := r.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/providers/Microsoft\.BillingBenefits/reservationOrderAliases/(?P<reservationOrderAliasName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armbillingbenefits.ReservationOrderAliasRequest](req)
		if err != nil {
			return nil, err
		}
		reservationOrderAliasNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationOrderAliasName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreate(req.Context(), reservationOrderAliasNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		r.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		r.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		r.beginCreate.remove(req)
	}

	return resp, nil
}

func (r *ReservationOrderAliasResponsesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.BillingBenefits/reservationOrderAliases/(?P<reservationOrderAliasName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	reservationOrderAliasNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationOrderAliasName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), reservationOrderAliasNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ReservationOrderAliasResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ReservationOrderAliasResponsesServerTransport
var reservationOrderAliasResponsesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
