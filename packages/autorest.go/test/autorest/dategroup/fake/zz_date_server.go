// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/dategroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"time"
)

// DateServer is a fake server for instances of the dategroup.DateClient type.
type DateServer struct {
	// GetInvalidDate is the fake for method DateClient.GetInvalidDate
	// HTTP status codes to indicate success: http.StatusOK
	GetInvalidDate func(ctx context.Context, options *dategroup.DateClientGetInvalidDateOptions) (resp azfake.Responder[dategroup.DateClientGetInvalidDateResponse], errResp azfake.ErrorResponder)

	// GetMaxDate is the fake for method DateClient.GetMaxDate
	// HTTP status codes to indicate success: http.StatusOK
	GetMaxDate func(ctx context.Context, options *dategroup.DateClientGetMaxDateOptions) (resp azfake.Responder[dategroup.DateClientGetMaxDateResponse], errResp azfake.ErrorResponder)

	// GetMinDate is the fake for method DateClient.GetMinDate
	// HTTP status codes to indicate success: http.StatusOK
	GetMinDate func(ctx context.Context, options *dategroup.DateClientGetMinDateOptions) (resp azfake.Responder[dategroup.DateClientGetMinDateResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method DateClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *dategroup.DateClientGetNullOptions) (resp azfake.Responder[dategroup.DateClientGetNullResponse], errResp azfake.ErrorResponder)

	// GetOverflowDate is the fake for method DateClient.GetOverflowDate
	// HTTP status codes to indicate success: http.StatusOK
	GetOverflowDate func(ctx context.Context, options *dategroup.DateClientGetOverflowDateOptions) (resp azfake.Responder[dategroup.DateClientGetOverflowDateResponse], errResp azfake.ErrorResponder)

	// GetUnderflowDate is the fake for method DateClient.GetUnderflowDate
	// HTTP status codes to indicate success: http.StatusOK
	GetUnderflowDate func(ctx context.Context, options *dategroup.DateClientGetUnderflowDateOptions) (resp azfake.Responder[dategroup.DateClientGetUnderflowDateResponse], errResp azfake.ErrorResponder)

	// PutMaxDate is the fake for method DateClient.PutMaxDate
	// HTTP status codes to indicate success: http.StatusOK
	PutMaxDate func(ctx context.Context, dateBody time.Time, options *dategroup.DateClientPutMaxDateOptions) (resp azfake.Responder[dategroup.DateClientPutMaxDateResponse], errResp azfake.ErrorResponder)

	// PutMinDate is the fake for method DateClient.PutMinDate
	// HTTP status codes to indicate success: http.StatusOK
	PutMinDate func(ctx context.Context, dateBody time.Time, options *dategroup.DateClientPutMinDateOptions) (resp azfake.Responder[dategroup.DateClientPutMinDateResponse], errResp azfake.ErrorResponder)
}

// NewDateServerTransport creates a new instance of DateServerTransport with the provided implementation.
// The returned DateServerTransport instance is connected to an instance of dategroup.DateClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDateServerTransport(srv *DateServer) *DateServerTransport {
	return &DateServerTransport{srv: srv}
}

// DateServerTransport connects instances of dategroup.DateClient to instances of DateServer.
// Don't use this type directly, use NewDateServerTransport instead.
type DateServerTransport struct {
	srv *DateServer
}

// Do implements the policy.Transporter interface for DateServerTransport.
func (d *DateServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DateServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if dateServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = dateServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DateClient.GetInvalidDate":
				res.resp, res.err = d.dispatchGetInvalidDate(req)
			case "DateClient.GetMaxDate":
				res.resp, res.err = d.dispatchGetMaxDate(req)
			case "DateClient.GetMinDate":
				res.resp, res.err = d.dispatchGetMinDate(req)
			case "DateClient.GetNull":
				res.resp, res.err = d.dispatchGetNull(req)
			case "DateClient.GetOverflowDate":
				res.resp, res.err = d.dispatchGetOverflowDate(req)
			case "DateClient.GetUnderflowDate":
				res.resp, res.err = d.dispatchGetUnderflowDate(req)
			case "DateClient.PutMaxDate":
				res.resp, res.err = d.dispatchPutMaxDate(req)
			case "DateClient.PutMinDate":
				res.resp, res.err = d.dispatchPutMinDate(req)
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

func (d *DateServerTransport) dispatchGetInvalidDate(req *http.Request) (*http.Response, error) {
	if d.srv.GetInvalidDate == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInvalidDate not implemented")}
	}
	respr, errRespr := d.srv.GetInvalidDate(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateType)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DateServerTransport) dispatchGetMaxDate(req *http.Request) (*http.Response, error) {
	if d.srv.GetMaxDate == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetMaxDate not implemented")}
	}
	respr, errRespr := d.srv.GetMaxDate(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateType)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DateServerTransport) dispatchGetMinDate(req *http.Request) (*http.Response, error) {
	if d.srv.GetMinDate == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetMinDate not implemented")}
	}
	respr, errRespr := d.srv.GetMinDate(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateType)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DateServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if d.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNull not implemented")}
	}
	respr, errRespr := d.srv.GetNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateType)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DateServerTransport) dispatchGetOverflowDate(req *http.Request) (*http.Response, error) {
	if d.srv.GetOverflowDate == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetOverflowDate not implemented")}
	}
	respr, errRespr := d.srv.GetOverflowDate(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateType)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DateServerTransport) dispatchGetUnderflowDate(req *http.Request) (*http.Response, error) {
	if d.srv.GetUnderflowDate == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUnderflowDate not implemented")}
	}
	respr, errRespr := d.srv.GetUnderflowDate(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateType)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DateServerTransport) dispatchPutMaxDate(req *http.Request) (*http.Response, error) {
	if d.srv.PutMaxDate == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutMaxDate not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[dateType](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutMaxDate(req.Context(), time.Time(body), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DateServerTransport) dispatchPutMinDate(req *http.Request) (*http.Response, error) {
	if d.srv.PutMinDate == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutMinDate not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[dateType](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutMinDate(req.Context(), time.Time(body), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DateServerTransport
var dateServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
