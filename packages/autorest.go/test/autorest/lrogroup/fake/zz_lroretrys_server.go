// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/lrogroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// LRORetrysServer is a fake server for instances of the lrogroup.LRORetrysClient type.
type LRORetrysServer struct {
	// BeginDelete202Retry200 is the fake for method LRORetrysClient.BeginDelete202Retry200
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete202Retry200 func(ctx context.Context, options *lrogroup.LRORetrysClientBeginDelete202Retry200Options) (resp azfake.PollerResponder[lrogroup.LRORetrysClientDelete202Retry200Response], errResp azfake.ErrorResponder)

	// BeginDeleteAsyncRelativeRetrySucceeded is the fake for method LRORetrysClient.BeginDeleteAsyncRelativeRetrySucceeded
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeleteAsyncRelativeRetrySucceeded func(ctx context.Context, options *lrogroup.LRORetrysClientBeginDeleteAsyncRelativeRetrySucceededOptions) (resp azfake.PollerResponder[lrogroup.LRORetrysClientDeleteAsyncRelativeRetrySucceededResponse], errResp azfake.ErrorResponder)

	// BeginDeleteProvisioning202Accepted200Succeeded is the fake for method LRORetrysClient.BeginDeleteProvisioning202Accepted200Succeeded
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDeleteProvisioning202Accepted200Succeeded func(ctx context.Context, options *lrogroup.LRORetrysClientBeginDeleteProvisioning202Accepted200SucceededOptions) (resp azfake.PollerResponder[lrogroup.LRORetrysClientDeleteProvisioning202Accepted200SucceededResponse], errResp azfake.ErrorResponder)

	// BeginPost202Retry200 is the fake for method LRORetrysClient.BeginPost202Retry200
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginPost202Retry200 func(ctx context.Context, options *lrogroup.LRORetrysClientBeginPost202Retry200Options) (resp azfake.PollerResponder[lrogroup.LRORetrysClientPost202Retry200Response], errResp azfake.ErrorResponder)

	// BeginPostAsyncRelativeRetrySucceeded is the fake for method LRORetrysClient.BeginPostAsyncRelativeRetrySucceeded
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginPostAsyncRelativeRetrySucceeded func(ctx context.Context, options *lrogroup.LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions) (resp azfake.PollerResponder[lrogroup.LRORetrysClientPostAsyncRelativeRetrySucceededResponse], errResp azfake.ErrorResponder)

	// BeginPut201CreatingSucceeded200 is the fake for method LRORetrysClient.BeginPut201CreatingSucceeded200
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginPut201CreatingSucceeded200 func(ctx context.Context, product lrogroup.Product, options *lrogroup.LRORetrysClientBeginPut201CreatingSucceeded200Options) (resp azfake.PollerResponder[lrogroup.LRORetrysClientPut201CreatingSucceeded200Response], errResp azfake.ErrorResponder)

	// BeginPutAsyncRelativeRetrySucceeded is the fake for method LRORetrysClient.BeginPutAsyncRelativeRetrySucceeded
	// HTTP status codes to indicate success: http.StatusOK
	BeginPutAsyncRelativeRetrySucceeded func(ctx context.Context, product lrogroup.Product, options *lrogroup.LRORetrysClientBeginPutAsyncRelativeRetrySucceededOptions) (resp azfake.PollerResponder[lrogroup.LRORetrysClientPutAsyncRelativeRetrySucceededResponse], errResp azfake.ErrorResponder)
}

// NewLRORetrysServerTransport creates a new instance of LRORetrysServerTransport with the provided implementation.
// The returned LRORetrysServerTransport instance is connected to an instance of lrogroup.LRORetrysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLRORetrysServerTransport(srv *LRORetrysServer) *LRORetrysServerTransport {
	return &LRORetrysServerTransport{
		srv:                                    srv,
		beginDelete202Retry200:                 newTracker[azfake.PollerResponder[lrogroup.LRORetrysClientDelete202Retry200Response]](),
		beginDeleteAsyncRelativeRetrySucceeded: newTracker[azfake.PollerResponder[lrogroup.LRORetrysClientDeleteAsyncRelativeRetrySucceededResponse]](),
		beginDeleteProvisioning202Accepted200Succeeded: newTracker[azfake.PollerResponder[lrogroup.LRORetrysClientDeleteProvisioning202Accepted200SucceededResponse]](),
		beginPost202Retry200:                           newTracker[azfake.PollerResponder[lrogroup.LRORetrysClientPost202Retry200Response]](),
		beginPostAsyncRelativeRetrySucceeded:           newTracker[azfake.PollerResponder[lrogroup.LRORetrysClientPostAsyncRelativeRetrySucceededResponse]](),
		beginPut201CreatingSucceeded200:                newTracker[azfake.PollerResponder[lrogroup.LRORetrysClientPut201CreatingSucceeded200Response]](),
		beginPutAsyncRelativeRetrySucceeded:            newTracker[azfake.PollerResponder[lrogroup.LRORetrysClientPutAsyncRelativeRetrySucceededResponse]](),
	}
}

// LRORetrysServerTransport connects instances of lrogroup.LRORetrysClient to instances of LRORetrysServer.
// Don't use this type directly, use NewLRORetrysServerTransport instead.
type LRORetrysServerTransport struct {
	srv                                            *LRORetrysServer
	beginDelete202Retry200                         *tracker[azfake.PollerResponder[lrogroup.LRORetrysClientDelete202Retry200Response]]
	beginDeleteAsyncRelativeRetrySucceeded         *tracker[azfake.PollerResponder[lrogroup.LRORetrysClientDeleteAsyncRelativeRetrySucceededResponse]]
	beginDeleteProvisioning202Accepted200Succeeded *tracker[azfake.PollerResponder[lrogroup.LRORetrysClientDeleteProvisioning202Accepted200SucceededResponse]]
	beginPost202Retry200                           *tracker[azfake.PollerResponder[lrogroup.LRORetrysClientPost202Retry200Response]]
	beginPostAsyncRelativeRetrySucceeded           *tracker[azfake.PollerResponder[lrogroup.LRORetrysClientPostAsyncRelativeRetrySucceededResponse]]
	beginPut201CreatingSucceeded200                *tracker[azfake.PollerResponder[lrogroup.LRORetrysClientPut201CreatingSucceeded200Response]]
	beginPutAsyncRelativeRetrySucceeded            *tracker[azfake.PollerResponder[lrogroup.LRORetrysClientPutAsyncRelativeRetrySucceededResponse]]
}

// Do implements the policy.Transporter interface for LRORetrysServerTransport.
func (l *LRORetrysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return l.dispatchToMethodFake(req, method)
}

func (l *LRORetrysServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if lroRetrysServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = lroRetrysServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "LRORetrysClient.BeginDelete202Retry200":
				res.resp, res.err = l.dispatchBeginDelete202Retry200(req)
			case "LRORetrysClient.BeginDeleteAsyncRelativeRetrySucceeded":
				res.resp, res.err = l.dispatchBeginDeleteAsyncRelativeRetrySucceeded(req)
			case "LRORetrysClient.BeginDeleteProvisioning202Accepted200Succeeded":
				res.resp, res.err = l.dispatchBeginDeleteProvisioning202Accepted200Succeeded(req)
			case "LRORetrysClient.BeginPost202Retry200":
				res.resp, res.err = l.dispatchBeginPost202Retry200(req)
			case "LRORetrysClient.BeginPostAsyncRelativeRetrySucceeded":
				res.resp, res.err = l.dispatchBeginPostAsyncRelativeRetrySucceeded(req)
			case "LRORetrysClient.BeginPut201CreatingSucceeded200":
				res.resp, res.err = l.dispatchBeginPut201CreatingSucceeded200(req)
			case "LRORetrysClient.BeginPutAsyncRelativeRetrySucceeded":
				res.resp, res.err = l.dispatchBeginPutAsyncRelativeRetrySucceeded(req)
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

func (l *LRORetrysServerTransport) dispatchBeginDelete202Retry200(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete202Retry200 == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete202Retry200 not implemented")}
	}
	beginDelete202Retry200 := l.beginDelete202Retry200.get(req)
	if beginDelete202Retry200 == nil {
		respr, errRespr := l.srv.BeginDelete202Retry200(req.Context(), nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete202Retry200 = &respr
		l.beginDelete202Retry200.add(req, beginDelete202Retry200)
	}

	resp, err := server.PollerResponderNext(beginDelete202Retry200, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginDelete202Retry200.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete202Retry200) {
		l.beginDelete202Retry200.remove(req)
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginDeleteAsyncRelativeRetrySucceeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDeleteAsyncRelativeRetrySucceeded == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteAsyncRelativeRetrySucceeded not implemented")}
	}
	beginDeleteAsyncRelativeRetrySucceeded := l.beginDeleteAsyncRelativeRetrySucceeded.get(req)
	if beginDeleteAsyncRelativeRetrySucceeded == nil {
		respr, errRespr := l.srv.BeginDeleteAsyncRelativeRetrySucceeded(req.Context(), nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeleteAsyncRelativeRetrySucceeded = &respr
		l.beginDeleteAsyncRelativeRetrySucceeded.add(req, beginDeleteAsyncRelativeRetrySucceeded)
	}

	resp, err := server.PollerResponderNext(beginDeleteAsyncRelativeRetrySucceeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginDeleteAsyncRelativeRetrySucceeded.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeleteAsyncRelativeRetrySucceeded) {
		l.beginDeleteAsyncRelativeRetrySucceeded.remove(req)
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginDeleteProvisioning202Accepted200Succeeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDeleteProvisioning202Accepted200Succeeded == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteProvisioning202Accepted200Succeeded not implemented")}
	}
	beginDeleteProvisioning202Accepted200Succeeded := l.beginDeleteProvisioning202Accepted200Succeeded.get(req)
	if beginDeleteProvisioning202Accepted200Succeeded == nil {
		respr, errRespr := l.srv.BeginDeleteProvisioning202Accepted200Succeeded(req.Context(), nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeleteProvisioning202Accepted200Succeeded = &respr
		l.beginDeleteProvisioning202Accepted200Succeeded.add(req, beginDeleteProvisioning202Accepted200Succeeded)
	}

	resp, err := server.PollerResponderNext(beginDeleteProvisioning202Accepted200Succeeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginDeleteProvisioning202Accepted200Succeeded.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeleteProvisioning202Accepted200Succeeded) {
		l.beginDeleteProvisioning202Accepted200Succeeded.remove(req)
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginPost202Retry200(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPost202Retry200 == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPost202Retry200 not implemented")}
	}
	beginPost202Retry200 := l.beginPost202Retry200.get(req)
	if beginPost202Retry200 == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		var options *lrogroup.LRORetrysClientBeginPost202Retry200Options
		if !reflect.ValueOf(body).IsZero() {
			options = &lrogroup.LRORetrysClientBeginPost202Retry200Options{
				Product: &body,
			}
		}
		respr, errRespr := l.srv.BeginPost202Retry200(req.Context(), options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPost202Retry200 = &respr
		l.beginPost202Retry200.add(req, beginPost202Retry200)
	}

	resp, err := server.PollerResponderNext(beginPost202Retry200, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginPost202Retry200.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPost202Retry200) {
		l.beginPost202Retry200.remove(req)
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginPostAsyncRelativeRetrySucceeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPostAsyncRelativeRetrySucceeded == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPostAsyncRelativeRetrySucceeded not implemented")}
	}
	beginPostAsyncRelativeRetrySucceeded := l.beginPostAsyncRelativeRetrySucceeded.get(req)
	if beginPostAsyncRelativeRetrySucceeded == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		var options *lrogroup.LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &lrogroup.LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions{
				Product: &body,
			}
		}
		respr, errRespr := l.srv.BeginPostAsyncRelativeRetrySucceeded(req.Context(), options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPostAsyncRelativeRetrySucceeded = &respr
		l.beginPostAsyncRelativeRetrySucceeded.add(req, beginPostAsyncRelativeRetrySucceeded)
	}

	resp, err := server.PollerResponderNext(beginPostAsyncRelativeRetrySucceeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginPostAsyncRelativeRetrySucceeded.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPostAsyncRelativeRetrySucceeded) {
		l.beginPostAsyncRelativeRetrySucceeded.remove(req)
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginPut201CreatingSucceeded200(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPut201CreatingSucceeded200 == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPut201CreatingSucceeded200 not implemented")}
	}
	beginPut201CreatingSucceeded200 := l.beginPut201CreatingSucceeded200.get(req)
	if beginPut201CreatingSucceeded200 == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginPut201CreatingSucceeded200(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPut201CreatingSucceeded200 = &respr
		l.beginPut201CreatingSucceeded200.add(req, beginPut201CreatingSucceeded200)
	}

	resp, err := server.PollerResponderNext(beginPut201CreatingSucceeded200, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		l.beginPut201CreatingSucceeded200.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPut201CreatingSucceeded200) {
		l.beginPut201CreatingSucceeded200.remove(req)
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginPutAsyncRelativeRetrySucceeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPutAsyncRelativeRetrySucceeded == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPutAsyncRelativeRetrySucceeded not implemented")}
	}
	beginPutAsyncRelativeRetrySucceeded := l.beginPutAsyncRelativeRetrySucceeded.get(req)
	if beginPutAsyncRelativeRetrySucceeded == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginPutAsyncRelativeRetrySucceeded(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPutAsyncRelativeRetrySucceeded = &respr
		l.beginPutAsyncRelativeRetrySucceeded.add(req, beginPutAsyncRelativeRetrySucceeded)
	}

	resp, err := server.PollerResponderNext(beginPutAsyncRelativeRetrySucceeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.beginPutAsyncRelativeRetrySucceeded.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPutAsyncRelativeRetrySucceeded) {
		l.beginPutAsyncRelativeRetrySucceeded.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to LRORetrysServerTransport
var lroRetrysServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
