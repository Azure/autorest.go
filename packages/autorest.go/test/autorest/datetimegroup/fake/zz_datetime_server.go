// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/datetimegroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"time"
)

// DatetimeServer is a fake server for instances of the datetimegroup.DatetimeClient type.
type DatetimeServer struct {
	// GetInvalid is the fake for method DatetimeClient.GetInvalid
	// HTTP status codes to indicate success: http.StatusOK
	GetInvalid func(ctx context.Context, options *datetimegroup.DatetimeClientGetInvalidOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetInvalidResponse], errResp azfake.ErrorResponder)

	// GetLocalNegativeOffsetLowercaseMaxDateTime is the fake for method DatetimeClient.GetLocalNegativeOffsetLowercaseMaxDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetLocalNegativeOffsetLowercaseMaxDateTime func(ctx context.Context, options *datetimegroup.DatetimeClientGetLocalNegativeOffsetLowercaseMaxDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetLocalNegativeOffsetLowercaseMaxDateTimeResponse], errResp azfake.ErrorResponder)

	// GetLocalNegativeOffsetMinDateTime is the fake for method DatetimeClient.GetLocalNegativeOffsetMinDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetLocalNegativeOffsetMinDateTime func(ctx context.Context, options *datetimegroup.DatetimeClientGetLocalNegativeOffsetMinDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetLocalNegativeOffsetMinDateTimeResponse], errResp azfake.ErrorResponder)

	// GetLocalNegativeOffsetUppercaseMaxDateTime is the fake for method DatetimeClient.GetLocalNegativeOffsetUppercaseMaxDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetLocalNegativeOffsetUppercaseMaxDateTime func(ctx context.Context, options *datetimegroup.DatetimeClientGetLocalNegativeOffsetUppercaseMaxDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetLocalNegativeOffsetUppercaseMaxDateTimeResponse], errResp azfake.ErrorResponder)

	// GetLocalNoOffsetMinDateTime is the fake for method DatetimeClient.GetLocalNoOffsetMinDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetLocalNoOffsetMinDateTime func(ctx context.Context, options *datetimegroup.DatetimeClientGetLocalNoOffsetMinDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetLocalNoOffsetMinDateTimeResponse], errResp azfake.ErrorResponder)

	// GetLocalPositiveOffsetLowercaseMaxDateTime is the fake for method DatetimeClient.GetLocalPositiveOffsetLowercaseMaxDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetLocalPositiveOffsetLowercaseMaxDateTime func(ctx context.Context, options *datetimegroup.DatetimeClientGetLocalPositiveOffsetLowercaseMaxDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetLocalPositiveOffsetLowercaseMaxDateTimeResponse], errResp azfake.ErrorResponder)

	// GetLocalPositiveOffsetMinDateTime is the fake for method DatetimeClient.GetLocalPositiveOffsetMinDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetLocalPositiveOffsetMinDateTime func(ctx context.Context, options *datetimegroup.DatetimeClientGetLocalPositiveOffsetMinDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetLocalPositiveOffsetMinDateTimeResponse], errResp azfake.ErrorResponder)

	// GetLocalPositiveOffsetUppercaseMaxDateTime is the fake for method DatetimeClient.GetLocalPositiveOffsetUppercaseMaxDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetLocalPositiveOffsetUppercaseMaxDateTime func(ctx context.Context, options *datetimegroup.DatetimeClientGetLocalPositiveOffsetUppercaseMaxDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetLocalPositiveOffsetUppercaseMaxDateTimeResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method DatetimeClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *datetimegroup.DatetimeClientGetNullOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetNullResponse], errResp azfake.ErrorResponder)

	// GetOverflow is the fake for method DatetimeClient.GetOverflow
	// HTTP status codes to indicate success: http.StatusOK
	GetOverflow func(ctx context.Context, options *datetimegroup.DatetimeClientGetOverflowOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetOverflowResponse], errResp azfake.ErrorResponder)

	// GetUTCLowercaseMaxDateTime is the fake for method DatetimeClient.GetUTCLowercaseMaxDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetUTCLowercaseMaxDateTime func(ctx context.Context, options *datetimegroup.DatetimeClientGetUTCLowercaseMaxDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetUTCLowercaseMaxDateTimeResponse], errResp azfake.ErrorResponder)

	// GetUTCMinDateTime is the fake for method DatetimeClient.GetUTCMinDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetUTCMinDateTime func(ctx context.Context, options *datetimegroup.DatetimeClientGetUTCMinDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetUTCMinDateTimeResponse], errResp azfake.ErrorResponder)

	// GetUTCUppercaseMaxDateTime is the fake for method DatetimeClient.GetUTCUppercaseMaxDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetUTCUppercaseMaxDateTime func(ctx context.Context, options *datetimegroup.DatetimeClientGetUTCUppercaseMaxDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetUTCUppercaseMaxDateTimeResponse], errResp azfake.ErrorResponder)

	// GetUTCUppercaseMaxDateTime7Digits is the fake for method DatetimeClient.GetUTCUppercaseMaxDateTime7Digits
	// HTTP status codes to indicate success: http.StatusOK
	GetUTCUppercaseMaxDateTime7Digits func(ctx context.Context, options *datetimegroup.DatetimeClientGetUTCUppercaseMaxDateTime7DigitsOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetUTCUppercaseMaxDateTime7DigitsResponse], errResp azfake.ErrorResponder)

	// GetUnderflow is the fake for method DatetimeClient.GetUnderflow
	// HTTP status codes to indicate success: http.StatusOK
	GetUnderflow func(ctx context.Context, options *datetimegroup.DatetimeClientGetUnderflowOptions) (resp azfake.Responder[datetimegroup.DatetimeClientGetUnderflowResponse], errResp azfake.ErrorResponder)

	// PutLocalNegativeOffsetMaxDateTime is the fake for method DatetimeClient.PutLocalNegativeOffsetMaxDateTime
	// HTTP status codes to indicate success: http.StatusOK
	PutLocalNegativeOffsetMaxDateTime func(ctx context.Context, datetimeBody time.Time, options *datetimegroup.DatetimeClientPutLocalNegativeOffsetMaxDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientPutLocalNegativeOffsetMaxDateTimeResponse], errResp azfake.ErrorResponder)

	// PutLocalNegativeOffsetMinDateTime is the fake for method DatetimeClient.PutLocalNegativeOffsetMinDateTime
	// HTTP status codes to indicate success: http.StatusOK
	PutLocalNegativeOffsetMinDateTime func(ctx context.Context, datetimeBody time.Time, options *datetimegroup.DatetimeClientPutLocalNegativeOffsetMinDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientPutLocalNegativeOffsetMinDateTimeResponse], errResp azfake.ErrorResponder)

	// PutLocalPositiveOffsetMaxDateTime is the fake for method DatetimeClient.PutLocalPositiveOffsetMaxDateTime
	// HTTP status codes to indicate success: http.StatusOK
	PutLocalPositiveOffsetMaxDateTime func(ctx context.Context, datetimeBody time.Time, options *datetimegroup.DatetimeClientPutLocalPositiveOffsetMaxDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientPutLocalPositiveOffsetMaxDateTimeResponse], errResp azfake.ErrorResponder)

	// PutLocalPositiveOffsetMinDateTime is the fake for method DatetimeClient.PutLocalPositiveOffsetMinDateTime
	// HTTP status codes to indicate success: http.StatusOK
	PutLocalPositiveOffsetMinDateTime func(ctx context.Context, datetimeBody time.Time, options *datetimegroup.DatetimeClientPutLocalPositiveOffsetMinDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientPutLocalPositiveOffsetMinDateTimeResponse], errResp azfake.ErrorResponder)

	// PutUTCMaxDateTime is the fake for method DatetimeClient.PutUTCMaxDateTime
	// HTTP status codes to indicate success: http.StatusOK
	PutUTCMaxDateTime func(ctx context.Context, datetimeBody time.Time, options *datetimegroup.DatetimeClientPutUTCMaxDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientPutUTCMaxDateTimeResponse], errResp azfake.ErrorResponder)

	// PutUTCMaxDateTime7Digits is the fake for method DatetimeClient.PutUTCMaxDateTime7Digits
	// HTTP status codes to indicate success: http.StatusOK
	PutUTCMaxDateTime7Digits func(ctx context.Context, datetimeBody time.Time, options *datetimegroup.DatetimeClientPutUTCMaxDateTime7DigitsOptions) (resp azfake.Responder[datetimegroup.DatetimeClientPutUTCMaxDateTime7DigitsResponse], errResp azfake.ErrorResponder)

	// PutUTCMinDateTime is the fake for method DatetimeClient.PutUTCMinDateTime
	// HTTP status codes to indicate success: http.StatusOK
	PutUTCMinDateTime func(ctx context.Context, datetimeBody time.Time, options *datetimegroup.DatetimeClientPutUTCMinDateTimeOptions) (resp azfake.Responder[datetimegroup.DatetimeClientPutUTCMinDateTimeResponse], errResp azfake.ErrorResponder)
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
	srv *DatetimeServer
}

// Do implements the policy.Transporter interface for DatetimeServerTransport.
func (d *DatetimeServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DatetimeServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if datetimeServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = datetimeServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DatetimeClient.GetInvalid":
				res.resp, res.err = d.dispatchGetInvalid(req)
			case "DatetimeClient.GetLocalNegativeOffsetLowercaseMaxDateTime":
				res.resp, res.err = d.dispatchGetLocalNegativeOffsetLowercaseMaxDateTime(req)
			case "DatetimeClient.GetLocalNegativeOffsetMinDateTime":
				res.resp, res.err = d.dispatchGetLocalNegativeOffsetMinDateTime(req)
			case "DatetimeClient.GetLocalNegativeOffsetUppercaseMaxDateTime":
				res.resp, res.err = d.dispatchGetLocalNegativeOffsetUppercaseMaxDateTime(req)
			case "DatetimeClient.GetLocalNoOffsetMinDateTime":
				res.resp, res.err = d.dispatchGetLocalNoOffsetMinDateTime(req)
			case "DatetimeClient.GetLocalPositiveOffsetLowercaseMaxDateTime":
				res.resp, res.err = d.dispatchGetLocalPositiveOffsetLowercaseMaxDateTime(req)
			case "DatetimeClient.GetLocalPositiveOffsetMinDateTime":
				res.resp, res.err = d.dispatchGetLocalPositiveOffsetMinDateTime(req)
			case "DatetimeClient.GetLocalPositiveOffsetUppercaseMaxDateTime":
				res.resp, res.err = d.dispatchGetLocalPositiveOffsetUppercaseMaxDateTime(req)
			case "DatetimeClient.GetNull":
				res.resp, res.err = d.dispatchGetNull(req)
			case "DatetimeClient.GetOverflow":
				res.resp, res.err = d.dispatchGetOverflow(req)
			case "DatetimeClient.GetUTCLowercaseMaxDateTime":
				res.resp, res.err = d.dispatchGetUTCLowercaseMaxDateTime(req)
			case "DatetimeClient.GetUTCMinDateTime":
				res.resp, res.err = d.dispatchGetUTCMinDateTime(req)
			case "DatetimeClient.GetUTCUppercaseMaxDateTime":
				res.resp, res.err = d.dispatchGetUTCUppercaseMaxDateTime(req)
			case "DatetimeClient.GetUTCUppercaseMaxDateTime7Digits":
				res.resp, res.err = d.dispatchGetUTCUppercaseMaxDateTime7Digits(req)
			case "DatetimeClient.GetUnderflow":
				res.resp, res.err = d.dispatchGetUnderflow(req)
			case "DatetimeClient.PutLocalNegativeOffsetMaxDateTime":
				res.resp, res.err = d.dispatchPutLocalNegativeOffsetMaxDateTime(req)
			case "DatetimeClient.PutLocalNegativeOffsetMinDateTime":
				res.resp, res.err = d.dispatchPutLocalNegativeOffsetMinDateTime(req)
			case "DatetimeClient.PutLocalPositiveOffsetMaxDateTime":
				res.resp, res.err = d.dispatchPutLocalPositiveOffsetMaxDateTime(req)
			case "DatetimeClient.PutLocalPositiveOffsetMinDateTime":
				res.resp, res.err = d.dispatchPutLocalPositiveOffsetMinDateTime(req)
			case "DatetimeClient.PutUTCMaxDateTime":
				res.resp, res.err = d.dispatchPutUTCMaxDateTime(req)
			case "DatetimeClient.PutUTCMaxDateTime7Digits":
				res.resp, res.err = d.dispatchPutUTCMaxDateTime7Digits(req)
			case "DatetimeClient.PutUTCMinDateTime":
				res.resp, res.err = d.dispatchPutUTCMinDateTime(req)
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

func (d *DatetimeServerTransport) dispatchGetInvalid(req *http.Request) (*http.Response, error) {
	if d.srv.GetInvalid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInvalid not implemented")}
	}
	respr, errRespr := d.srv.GetInvalid(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetLocalNegativeOffsetLowercaseMaxDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.GetLocalNegativeOffsetLowercaseMaxDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLocalNegativeOffsetLowercaseMaxDateTime not implemented")}
	}
	respr, errRespr := d.srv.GetLocalNegativeOffsetLowercaseMaxDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetLocalNegativeOffsetMinDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.GetLocalNegativeOffsetMinDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLocalNegativeOffsetMinDateTime not implemented")}
	}
	respr, errRespr := d.srv.GetLocalNegativeOffsetMinDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetLocalNegativeOffsetUppercaseMaxDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.GetLocalNegativeOffsetUppercaseMaxDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLocalNegativeOffsetUppercaseMaxDateTime not implemented")}
	}
	respr, errRespr := d.srv.GetLocalNegativeOffsetUppercaseMaxDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetLocalNoOffsetMinDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.GetLocalNoOffsetMinDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLocalNoOffsetMinDateTime not implemented")}
	}
	respr, errRespr := d.srv.GetLocalNoOffsetMinDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetLocalPositiveOffsetLowercaseMaxDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.GetLocalPositiveOffsetLowercaseMaxDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLocalPositiveOffsetLowercaseMaxDateTime not implemented")}
	}
	respr, errRespr := d.srv.GetLocalPositiveOffsetLowercaseMaxDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetLocalPositiveOffsetMinDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.GetLocalPositiveOffsetMinDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLocalPositiveOffsetMinDateTime not implemented")}
	}
	respr, errRespr := d.srv.GetLocalPositiveOffsetMinDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetLocalPositiveOffsetUppercaseMaxDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.GetLocalPositiveOffsetUppercaseMaxDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLocalPositiveOffsetUppercaseMaxDateTime not implemented")}
	}
	respr, errRespr := d.srv.GetLocalPositiveOffsetUppercaseMaxDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetOverflow(req *http.Request) (*http.Response, error) {
	if d.srv.GetOverflow == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetOverflow not implemented")}
	}
	respr, errRespr := d.srv.GetOverflow(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetUTCLowercaseMaxDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.GetUTCLowercaseMaxDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUTCLowercaseMaxDateTime not implemented")}
	}
	respr, errRespr := d.srv.GetUTCLowercaseMaxDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetUTCMinDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.GetUTCMinDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUTCMinDateTime not implemented")}
	}
	respr, errRespr := d.srv.GetUTCMinDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetUTCUppercaseMaxDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.GetUTCUppercaseMaxDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUTCUppercaseMaxDateTime not implemented")}
	}
	respr, errRespr := d.srv.GetUTCUppercaseMaxDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetUTCUppercaseMaxDateTime7Digits(req *http.Request) (*http.Response, error) {
	if d.srv.GetUTCUppercaseMaxDateTime7Digits == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUTCUppercaseMaxDateTime7Digits not implemented")}
	}
	respr, errRespr := d.srv.GetUTCUppercaseMaxDateTime7Digits(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchGetUnderflow(req *http.Request) (*http.Response, error) {
	if d.srv.GetUnderflow == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUnderflow not implemented")}
	}
	respr, errRespr := d.srv.GetUnderflow(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*dateTimeRFC3339)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimeServerTransport) dispatchPutLocalNegativeOffsetMaxDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.PutLocalNegativeOffsetMaxDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutLocalNegativeOffsetMaxDateTime not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[dateTimeRFC3339](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutLocalNegativeOffsetMaxDateTime(req.Context(), time.Time(body), nil)
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

func (d *DatetimeServerTransport) dispatchPutLocalNegativeOffsetMinDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.PutLocalNegativeOffsetMinDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutLocalNegativeOffsetMinDateTime not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[dateTimeRFC3339](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutLocalNegativeOffsetMinDateTime(req.Context(), time.Time(body), nil)
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

func (d *DatetimeServerTransport) dispatchPutLocalPositiveOffsetMaxDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.PutLocalPositiveOffsetMaxDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutLocalPositiveOffsetMaxDateTime not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[dateTimeRFC3339](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutLocalPositiveOffsetMaxDateTime(req.Context(), time.Time(body), nil)
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

func (d *DatetimeServerTransport) dispatchPutLocalPositiveOffsetMinDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.PutLocalPositiveOffsetMinDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutLocalPositiveOffsetMinDateTime not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[dateTimeRFC3339](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutLocalPositiveOffsetMinDateTime(req.Context(), time.Time(body), nil)
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

func (d *DatetimeServerTransport) dispatchPutUTCMaxDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.PutUTCMaxDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutUTCMaxDateTime not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[dateTimeRFC3339](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutUTCMaxDateTime(req.Context(), time.Time(body), nil)
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

func (d *DatetimeServerTransport) dispatchPutUTCMaxDateTime7Digits(req *http.Request) (*http.Response, error) {
	if d.srv.PutUTCMaxDateTime7Digits == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutUTCMaxDateTime7Digits not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[dateTimeRFC3339](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutUTCMaxDateTime7Digits(req.Context(), time.Time(body), nil)
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

func (d *DatetimeServerTransport) dispatchPutUTCMinDateTime(req *http.Request) (*http.Response, error) {
	if d.srv.PutUTCMinDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutUTCMinDateTime not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[dateTimeRFC3339](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutUTCMinDateTime(req.Context(), time.Time(body), nil)
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

// set this to conditionally intercept incoming requests to DatetimeServerTransport
var datetimeServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
