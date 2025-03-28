// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/complexgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PrimitiveServer is a fake server for instances of the complexgroup.PrimitiveClient type.
type PrimitiveServer struct {
	// GetBool is the fake for method PrimitiveClient.GetBool
	// HTTP status codes to indicate success: http.StatusOK
	GetBool func(ctx context.Context, options *complexgroup.PrimitiveClientGetBoolOptions) (resp azfake.Responder[complexgroup.PrimitiveClientGetBoolResponse], errResp azfake.ErrorResponder)

	// GetByte is the fake for method PrimitiveClient.GetByte
	// HTTP status codes to indicate success: http.StatusOK
	GetByte func(ctx context.Context, options *complexgroup.PrimitiveClientGetByteOptions) (resp azfake.Responder[complexgroup.PrimitiveClientGetByteResponse], errResp azfake.ErrorResponder)

	// GetDate is the fake for method PrimitiveClient.GetDate
	// HTTP status codes to indicate success: http.StatusOK
	GetDate func(ctx context.Context, options *complexgroup.PrimitiveClientGetDateOptions) (resp azfake.Responder[complexgroup.PrimitiveClientGetDateResponse], errResp azfake.ErrorResponder)

	// GetDateTime is the fake for method PrimitiveClient.GetDateTime
	// HTTP status codes to indicate success: http.StatusOK
	GetDateTime func(ctx context.Context, options *complexgroup.PrimitiveClientGetDateTimeOptions) (resp azfake.Responder[complexgroup.PrimitiveClientGetDateTimeResponse], errResp azfake.ErrorResponder)

	// GetDateTimeRFC1123 is the fake for method PrimitiveClient.GetDateTimeRFC1123
	// HTTP status codes to indicate success: http.StatusOK
	GetDateTimeRFC1123 func(ctx context.Context, options *complexgroup.PrimitiveClientGetDateTimeRFC1123Options) (resp azfake.Responder[complexgroup.PrimitiveClientGetDateTimeRFC1123Response], errResp azfake.ErrorResponder)

	// GetDouble is the fake for method PrimitiveClient.GetDouble
	// HTTP status codes to indicate success: http.StatusOK
	GetDouble func(ctx context.Context, options *complexgroup.PrimitiveClientGetDoubleOptions) (resp azfake.Responder[complexgroup.PrimitiveClientGetDoubleResponse], errResp azfake.ErrorResponder)

	// GetDuration is the fake for method PrimitiveClient.GetDuration
	// HTTP status codes to indicate success: http.StatusOK
	GetDuration func(ctx context.Context, options *complexgroup.PrimitiveClientGetDurationOptions) (resp azfake.Responder[complexgroup.PrimitiveClientGetDurationResponse], errResp azfake.ErrorResponder)

	// GetFloat is the fake for method PrimitiveClient.GetFloat
	// HTTP status codes to indicate success: http.StatusOK
	GetFloat func(ctx context.Context, options *complexgroup.PrimitiveClientGetFloatOptions) (resp azfake.Responder[complexgroup.PrimitiveClientGetFloatResponse], errResp azfake.ErrorResponder)

	// GetInt is the fake for method PrimitiveClient.GetInt
	// HTTP status codes to indicate success: http.StatusOK
	GetInt func(ctx context.Context, options *complexgroup.PrimitiveClientGetIntOptions) (resp azfake.Responder[complexgroup.PrimitiveClientGetIntResponse], errResp azfake.ErrorResponder)

	// GetLong is the fake for method PrimitiveClient.GetLong
	// HTTP status codes to indicate success: http.StatusOK
	GetLong func(ctx context.Context, options *complexgroup.PrimitiveClientGetLongOptions) (resp azfake.Responder[complexgroup.PrimitiveClientGetLongResponse], errResp azfake.ErrorResponder)

	// GetString is the fake for method PrimitiveClient.GetString
	// HTTP status codes to indicate success: http.StatusOK
	GetString func(ctx context.Context, options *complexgroup.PrimitiveClientGetStringOptions) (resp azfake.Responder[complexgroup.PrimitiveClientGetStringResponse], errResp azfake.ErrorResponder)

	// PutBool is the fake for method PrimitiveClient.PutBool
	// HTTP status codes to indicate success: http.StatusOK
	PutBool func(ctx context.Context, complexBody complexgroup.BooleanWrapper, options *complexgroup.PrimitiveClientPutBoolOptions) (resp azfake.Responder[complexgroup.PrimitiveClientPutBoolResponse], errResp azfake.ErrorResponder)

	// PutByte is the fake for method PrimitiveClient.PutByte
	// HTTP status codes to indicate success: http.StatusOK
	PutByte func(ctx context.Context, complexBody complexgroup.ByteWrapper, options *complexgroup.PrimitiveClientPutByteOptions) (resp azfake.Responder[complexgroup.PrimitiveClientPutByteResponse], errResp azfake.ErrorResponder)

	// PutDate is the fake for method PrimitiveClient.PutDate
	// HTTP status codes to indicate success: http.StatusOK
	PutDate func(ctx context.Context, complexBody complexgroup.DateWrapper, options *complexgroup.PrimitiveClientPutDateOptions) (resp azfake.Responder[complexgroup.PrimitiveClientPutDateResponse], errResp azfake.ErrorResponder)

	// PutDateTime is the fake for method PrimitiveClient.PutDateTime
	// HTTP status codes to indicate success: http.StatusOK
	PutDateTime func(ctx context.Context, complexBody complexgroup.DatetimeWrapper, options *complexgroup.PrimitiveClientPutDateTimeOptions) (resp azfake.Responder[complexgroup.PrimitiveClientPutDateTimeResponse], errResp azfake.ErrorResponder)

	// PutDateTimeRFC1123 is the fake for method PrimitiveClient.PutDateTimeRFC1123
	// HTTP status codes to indicate success: http.StatusOK
	PutDateTimeRFC1123 func(ctx context.Context, complexBody complexgroup.Datetimerfc1123Wrapper, options *complexgroup.PrimitiveClientPutDateTimeRFC1123Options) (resp azfake.Responder[complexgroup.PrimitiveClientPutDateTimeRFC1123Response], errResp azfake.ErrorResponder)

	// PutDouble is the fake for method PrimitiveClient.PutDouble
	// HTTP status codes to indicate success: http.StatusOK
	PutDouble func(ctx context.Context, complexBody complexgroup.DoubleWrapper, options *complexgroup.PrimitiveClientPutDoubleOptions) (resp azfake.Responder[complexgroup.PrimitiveClientPutDoubleResponse], errResp azfake.ErrorResponder)

	// PutDuration is the fake for method PrimitiveClient.PutDuration
	// HTTP status codes to indicate success: http.StatusOK
	PutDuration func(ctx context.Context, complexBody complexgroup.DurationWrapper, options *complexgroup.PrimitiveClientPutDurationOptions) (resp azfake.Responder[complexgroup.PrimitiveClientPutDurationResponse], errResp azfake.ErrorResponder)

	// PutFloat is the fake for method PrimitiveClient.PutFloat
	// HTTP status codes to indicate success: http.StatusOK
	PutFloat func(ctx context.Context, complexBody complexgroup.FloatWrapper, options *complexgroup.PrimitiveClientPutFloatOptions) (resp azfake.Responder[complexgroup.PrimitiveClientPutFloatResponse], errResp azfake.ErrorResponder)

	// PutInt is the fake for method PrimitiveClient.PutInt
	// HTTP status codes to indicate success: http.StatusOK
	PutInt func(ctx context.Context, complexBody complexgroup.IntWrapper, options *complexgroup.PrimitiveClientPutIntOptions) (resp azfake.Responder[complexgroup.PrimitiveClientPutIntResponse], errResp azfake.ErrorResponder)

	// PutLong is the fake for method PrimitiveClient.PutLong
	// HTTP status codes to indicate success: http.StatusOK
	PutLong func(ctx context.Context, complexBody complexgroup.LongWrapper, options *complexgroup.PrimitiveClientPutLongOptions) (resp azfake.Responder[complexgroup.PrimitiveClientPutLongResponse], errResp azfake.ErrorResponder)

	// PutString is the fake for method PrimitiveClient.PutString
	// HTTP status codes to indicate success: http.StatusOK
	PutString func(ctx context.Context, complexBody complexgroup.StringWrapper, options *complexgroup.PrimitiveClientPutStringOptions) (resp azfake.Responder[complexgroup.PrimitiveClientPutStringResponse], errResp azfake.ErrorResponder)
}

// NewPrimitiveServerTransport creates a new instance of PrimitiveServerTransport with the provided implementation.
// The returned PrimitiveServerTransport instance is connected to an instance of complexgroup.PrimitiveClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrimitiveServerTransport(srv *PrimitiveServer) *PrimitiveServerTransport {
	return &PrimitiveServerTransport{srv: srv}
}

// PrimitiveServerTransport connects instances of complexgroup.PrimitiveClient to instances of PrimitiveServer.
// Don't use this type directly, use NewPrimitiveServerTransport instead.
type PrimitiveServerTransport struct {
	srv *PrimitiveServer
}

// Do implements the policy.Transporter interface for PrimitiveServerTransport.
func (p *PrimitiveServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrimitiveServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if primitiveServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = primitiveServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrimitiveClient.GetBool":
				res.resp, res.err = p.dispatchGetBool(req)
			case "PrimitiveClient.GetByte":
				res.resp, res.err = p.dispatchGetByte(req)
			case "PrimitiveClient.GetDate":
				res.resp, res.err = p.dispatchGetDate(req)
			case "PrimitiveClient.GetDateTime":
				res.resp, res.err = p.dispatchGetDateTime(req)
			case "PrimitiveClient.GetDateTimeRFC1123":
				res.resp, res.err = p.dispatchGetDateTimeRFC1123(req)
			case "PrimitiveClient.GetDouble":
				res.resp, res.err = p.dispatchGetDouble(req)
			case "PrimitiveClient.GetDuration":
				res.resp, res.err = p.dispatchGetDuration(req)
			case "PrimitiveClient.GetFloat":
				res.resp, res.err = p.dispatchGetFloat(req)
			case "PrimitiveClient.GetInt":
				res.resp, res.err = p.dispatchGetInt(req)
			case "PrimitiveClient.GetLong":
				res.resp, res.err = p.dispatchGetLong(req)
			case "PrimitiveClient.GetString":
				res.resp, res.err = p.dispatchGetString(req)
			case "PrimitiveClient.PutBool":
				res.resp, res.err = p.dispatchPutBool(req)
			case "PrimitiveClient.PutByte":
				res.resp, res.err = p.dispatchPutByte(req)
			case "PrimitiveClient.PutDate":
				res.resp, res.err = p.dispatchPutDate(req)
			case "PrimitiveClient.PutDateTime":
				res.resp, res.err = p.dispatchPutDateTime(req)
			case "PrimitiveClient.PutDateTimeRFC1123":
				res.resp, res.err = p.dispatchPutDateTimeRFC1123(req)
			case "PrimitiveClient.PutDouble":
				res.resp, res.err = p.dispatchPutDouble(req)
			case "PrimitiveClient.PutDuration":
				res.resp, res.err = p.dispatchPutDuration(req)
			case "PrimitiveClient.PutFloat":
				res.resp, res.err = p.dispatchPutFloat(req)
			case "PrimitiveClient.PutInt":
				res.resp, res.err = p.dispatchPutInt(req)
			case "PrimitiveClient.PutLong":
				res.resp, res.err = p.dispatchPutLong(req)
			case "PrimitiveClient.PutString":
				res.resp, res.err = p.dispatchPutString(req)
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

func (p *PrimitiveServerTransport) dispatchGetBool(req *http.Request) (*http.Response, error) {
	if p.srv.GetBool == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBool not implemented")}
	}
	respr, errRespr := p.srv.GetBool(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BooleanWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchGetByte(req *http.Request) (*http.Response, error) {
	if p.srv.GetByte == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByte not implemented")}
	}
	respr, errRespr := p.srv.GetByte(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ByteWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchGetDate(req *http.Request) (*http.Response, error) {
	if p.srv.GetDate == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDate not implemented")}
	}
	respr, errRespr := p.srv.GetDate(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DateWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchGetDateTime(req *http.Request) (*http.Response, error) {
	if p.srv.GetDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDateTime not implemented")}
	}
	respr, errRespr := p.srv.GetDateTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DatetimeWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchGetDateTimeRFC1123(req *http.Request) (*http.Response, error) {
	if p.srv.GetDateTimeRFC1123 == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDateTimeRFC1123 not implemented")}
	}
	respr, errRespr := p.srv.GetDateTimeRFC1123(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Datetimerfc1123Wrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchGetDouble(req *http.Request) (*http.Response, error) {
	if p.srv.GetDouble == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDouble not implemented")}
	}
	respr, errRespr := p.srv.GetDouble(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DoubleWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchGetDuration(req *http.Request) (*http.Response, error) {
	if p.srv.GetDuration == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDuration not implemented")}
	}
	respr, errRespr := p.srv.GetDuration(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DurationWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchGetFloat(req *http.Request) (*http.Response, error) {
	if p.srv.GetFloat == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetFloat not implemented")}
	}
	respr, errRespr := p.srv.GetFloat(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FloatWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchGetInt(req *http.Request) (*http.Response, error) {
	if p.srv.GetInt == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInt not implemented")}
	}
	respr, errRespr := p.srv.GetInt(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IntWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchGetLong(req *http.Request) (*http.Response, error) {
	if p.srv.GetLong == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLong not implemented")}
	}
	respr, errRespr := p.srv.GetLong(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LongWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchGetString(req *http.Request) (*http.Response, error) {
	if p.srv.GetString == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetString not implemented")}
	}
	respr, errRespr := p.srv.GetString(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StringWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrimitiveServerTransport) dispatchPutBool(req *http.Request) (*http.Response, error) {
	if p.srv.PutBool == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutBool not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.BooleanWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutBool(req.Context(), body, nil)
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

func (p *PrimitiveServerTransport) dispatchPutByte(req *http.Request) (*http.Response, error) {
	if p.srv.PutByte == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutByte not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.ByteWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutByte(req.Context(), body, nil)
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

func (p *PrimitiveServerTransport) dispatchPutDate(req *http.Request) (*http.Response, error) {
	if p.srv.PutDate == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutDate not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.DateWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutDate(req.Context(), body, nil)
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

func (p *PrimitiveServerTransport) dispatchPutDateTime(req *http.Request) (*http.Response, error) {
	if p.srv.PutDateTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutDateTime not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.DatetimeWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutDateTime(req.Context(), body, nil)
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

func (p *PrimitiveServerTransport) dispatchPutDateTimeRFC1123(req *http.Request) (*http.Response, error) {
	if p.srv.PutDateTimeRFC1123 == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutDateTimeRFC1123 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.Datetimerfc1123Wrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutDateTimeRFC1123(req.Context(), body, nil)
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

func (p *PrimitiveServerTransport) dispatchPutDouble(req *http.Request) (*http.Response, error) {
	if p.srv.PutDouble == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutDouble not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.DoubleWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutDouble(req.Context(), body, nil)
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

func (p *PrimitiveServerTransport) dispatchPutDuration(req *http.Request) (*http.Response, error) {
	if p.srv.PutDuration == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutDuration not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.DurationWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutDuration(req.Context(), body, nil)
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

func (p *PrimitiveServerTransport) dispatchPutFloat(req *http.Request) (*http.Response, error) {
	if p.srv.PutFloat == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutFloat not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.FloatWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutFloat(req.Context(), body, nil)
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

func (p *PrimitiveServerTransport) dispatchPutInt(req *http.Request) (*http.Response, error) {
	if p.srv.PutInt == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutInt not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.IntWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutInt(req.Context(), body, nil)
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

func (p *PrimitiveServerTransport) dispatchPutLong(req *http.Request) (*http.Response, error) {
	if p.srv.PutLong == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutLong not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.LongWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutLong(req.Context(), body, nil)
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

func (p *PrimitiveServerTransport) dispatchPutString(req *http.Request) (*http.Response, error) {
	if p.srv.PutString == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutString not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.StringWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PutString(req.Context(), body, nil)
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

// set this to conditionally intercept incoming requests to PrimitiveServerTransport
var primitiveServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
