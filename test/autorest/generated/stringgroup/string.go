// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// StringOperations contains the methods for the String group.
type StringOperations interface {
	// GetBase64Encoded - Get value that is base64 encoded
	GetBase64Encoded(ctx context.Context) (*ByteArrayResponse, error)
	// GetBase64URLEncoded - Get value that is base64url encoded
	GetBase64URLEncoded(ctx context.Context) (*ByteArrayResponse, error)
	// GetEmpty - Get empty string value value ''
	GetEmpty(ctx context.Context) (*StringResponse, error)
	// GetMBCS - Get mbcs string value '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
	GetMBCS(ctx context.Context) (*StringResponse, error)
	// GetNotProvided - Get String value when no string value is sent in response payload
	GetNotProvided(ctx context.Context) (*StringResponse, error)
	// GetNull - Get null string value value
	GetNull(ctx context.Context) (*StringResponse, error)
	// GetNullBase64URLEncoded - Get null value that is expected to be base64url encoded
	GetNullBase64URLEncoded(ctx context.Context) (*ByteArrayResponse, error)
	// GetWhitespace - Get string value with leading and trailing whitespace '<tab><space><space>Now is the time for all good men to come to the aid of their country<tab><space><space>'
	GetWhitespace(ctx context.Context) (*StringResponse, error)
	// PutBase64URLEncoded - Put value that is base64url encoded
	PutBase64URLEncoded(ctx context.Context, stringBody []byte) (*http.Response, error)
	// PutEmpty - Set string value empty ''
	PutEmpty(ctx context.Context) (*http.Response, error)
	// PutMBCS - Set string value mbcs '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
	PutMBCS(ctx context.Context) (*http.Response, error)
	// PutNull - Set string value null
	PutNull(ctx context.Context, stringPutNullOptions *StringPutNullOptions) (*http.Response, error)
	// PutWhitespace - Set String value with leading and trailing whitespace '<tab><space><space>Now is the time for all good men to come to the aid of their country<tab><space><space>'
	PutWhitespace(ctx context.Context) (*http.Response, error)
}

// stringOperations implements the StringOperations interface.
type stringOperations struct {
	*Client
}

// GetBase64Encoded - Get value that is base64 encoded
func (client *stringOperations) GetBase64Encoded(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.getBase64EncodedCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBase64EncodedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBase64EncodedCreateRequest creates the GetBase64Encoded request.
func (client *stringOperations) getBase64EncodedCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/base64Encoding"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getBase64EncodedHandleResponse handles the GetBase64Encoded response.
func (client *stringOperations) getBase64EncodedHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getBase64EncodedHandleError(resp)
	}
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsByteArray(&result.Value, azcore.Base64StdFormat)
}

// getBase64EncodedHandleError handles the GetBase64Encoded error response.
func (client *stringOperations) getBase64EncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetBase64URLEncoded - Get value that is base64url encoded
func (client *stringOperations) GetBase64URLEncoded(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.getBase64UrlEncodedCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBase64UrlEncodedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBase64UrlEncodedCreateRequest creates the GetBase64URLEncoded request.
func (client *stringOperations) getBase64UrlEncodedCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/base64UrlEncoding"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getBase64UrlEncodedHandleResponse handles the GetBase64URLEncoded response.
func (client *stringOperations) getBase64UrlEncodedHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getBase64UrlEncodedHandleError(resp)
	}
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsByteArray(&result.Value, azcore.Base64URLFormat)
}

// getBase64UrlEncodedHandleError handles the GetBase64URLEncoded error response.
func (client *stringOperations) getBase64UrlEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetEmpty - Get empty string value value ''
func (client *stringOperations) GetEmpty(ctx context.Context) (*StringResponse, error) {
	req, err := client.getEmptyCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *stringOperations) getEmptyCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/empty"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *stringOperations) getEmptyHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getEmptyHandleError(resp)
	}
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getEmptyHandleError handles the GetEmpty error response.
func (client *stringOperations) getEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetMBCS - Get mbcs string value '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
func (client *stringOperations) GetMBCS(ctx context.Context) (*StringResponse, error) {
	req, err := client.getMbcsCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getMbcsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getMbcsCreateRequest creates the GetMBCS request.
func (client *stringOperations) getMbcsCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/mbcs"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getMbcsHandleResponse handles the GetMBCS response.
func (client *stringOperations) getMbcsHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getMbcsHandleError(resp)
	}
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getMbcsHandleError handles the GetMBCS error response.
func (client *stringOperations) getMbcsHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetNotProvided - Get String value when no string value is sent in response payload
func (client *stringOperations) GetNotProvided(ctx context.Context) (*StringResponse, error) {
	req, err := client.getNotProvidedCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNotProvidedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client *stringOperations) getNotProvidedCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/notProvided"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client *stringOperations) getNotProvidedHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getNotProvidedHandleError(resp)
	}
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getNotProvidedHandleError handles the GetNotProvided error response.
func (client *stringOperations) getNotProvidedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetNull - Get null string value value
func (client *stringOperations) GetNull(ctx context.Context) (*StringResponse, error) {
	req, err := client.getNullCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNullCreateRequest creates the GetNull request.
func (client *stringOperations) getNullCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/null"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *stringOperations) getNullHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getNullHandleError(resp)
	}
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getNullHandleError handles the GetNull error response.
func (client *stringOperations) getNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetNullBase64URLEncoded - Get null value that is expected to be base64url encoded
func (client *stringOperations) GetNullBase64URLEncoded(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.getNullBase64UrlEncodedCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNullBase64UrlEncodedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNullBase64UrlEncodedCreateRequest creates the GetNullBase64URLEncoded request.
func (client *stringOperations) getNullBase64UrlEncodedCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/nullBase64UrlEncoding"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNullBase64UrlEncodedHandleResponse handles the GetNullBase64URLEncoded response.
func (client *stringOperations) getNullBase64UrlEncodedHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getNullBase64UrlEncodedHandleError(resp)
	}
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsByteArray(&result.Value, azcore.Base64URLFormat)
}

// getNullBase64UrlEncodedHandleError handles the GetNullBase64URLEncoded error response.
func (client *stringOperations) getNullBase64UrlEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetWhitespace - Get string value with leading and trailing whitespace '<tab><space><space>Now is the time for all good men to come to the aid of their country<tab><space><space>'
func (client *stringOperations) GetWhitespace(ctx context.Context) (*StringResponse, error) {
	req, err := client.getWhitespaceCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getWhitespaceHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getWhitespaceCreateRequest creates the GetWhitespace request.
func (client *stringOperations) getWhitespaceCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/whitespace"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getWhitespaceHandleResponse handles the GetWhitespace response.
func (client *stringOperations) getWhitespaceHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getWhitespaceHandleError(resp)
	}
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getWhitespaceHandleError handles the GetWhitespace error response.
func (client *stringOperations) getWhitespaceHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutBase64URLEncoded - Put value that is base64url encoded
func (client *stringOperations) PutBase64URLEncoded(ctx context.Context, stringBody []byte) (*http.Response, error) {
	req, err := client.putBase64UrlEncodedCreateRequest(stringBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putBase64UrlEncodedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putBase64UrlEncodedCreateRequest creates the PutBase64URLEncoded request.
func (client *stringOperations) putBase64UrlEncodedCreateRequest(stringBody []byte) (*azcore.Request, error) {
	urlPath := "/string/base64UrlEncoding"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsByteArray(stringBody, azcore.Base64URLFormat)
}

// putBase64UrlEncodedHandleResponse handles the PutBase64URLEncoded response.
func (client *stringOperations) putBase64UrlEncodedHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putBase64UrlEncodedHandleError(resp)
	}
	return resp.Response, nil
}

// putBase64UrlEncodedHandleError handles the PutBase64URLEncoded error response.
func (client *stringOperations) putBase64UrlEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutEmpty - Set string value empty ''
func (client *stringOperations) PutEmpty(ctx context.Context) (*http.Response, error) {
	req, err := client.putEmptyCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putEmptyCreateRequest creates the PutEmpty request.
func (client *stringOperations) putEmptyCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/empty"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON("")
}

// putEmptyHandleResponse handles the PutEmpty response.
func (client *stringOperations) putEmptyHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putEmptyHandleError(resp)
	}
	return resp.Response, nil
}

// putEmptyHandleError handles the PutEmpty error response.
func (client *stringOperations) putEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutMBCS - Set string value mbcs '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
func (client *stringOperations) PutMBCS(ctx context.Context) (*http.Response, error) {
	req, err := client.putMbcsCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putMbcsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putMbcsCreateRequest creates the PutMBCS request.
func (client *stringOperations) putMbcsCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/mbcs"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON("啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€")
}

// putMbcsHandleResponse handles the PutMBCS response.
func (client *stringOperations) putMbcsHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putMbcsHandleError(resp)
	}
	return resp.Response, nil
}

// putMbcsHandleError handles the PutMBCS error response.
func (client *stringOperations) putMbcsHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutNull - Set string value null
func (client *stringOperations) PutNull(ctx context.Context, stringPutNullOptions *StringPutNullOptions) (*http.Response, error) {
	req, err := client.putNullCreateRequest(stringPutNullOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putNullCreateRequest creates the PutNull request.
func (client *stringOperations) putNullCreateRequest(stringPutNullOptions *StringPutNullOptions) (*azcore.Request, error) {
	urlPath := "/string/null"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	if stringPutNullOptions != nil {
		return req, req.MarshalAsJSON(stringPutNullOptions.StringBody)
	}
	return req, nil
}

// putNullHandleResponse handles the PutNull response.
func (client *stringOperations) putNullHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putNullHandleError(resp)
	}
	return resp.Response, nil
}

// putNullHandleError handles the PutNull error response.
func (client *stringOperations) putNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutWhitespace - Set String value with leading and trailing whitespace '<tab><space><space>Now is the time for all good men to come to the aid of their country<tab><space><space>'
func (client *stringOperations) PutWhitespace(ctx context.Context) (*http.Response, error) {
	req, err := client.putWhitespaceCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putWhitespaceHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putWhitespaceCreateRequest creates the PutWhitespace request.
func (client *stringOperations) putWhitespaceCreateRequest() (*azcore.Request, error) {
	urlPath := "/string/whitespace"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON("    Now is the time for all good men to come to the aid of their country    ")
}

// putWhitespaceHandleResponse handles the PutWhitespace response.
func (client *stringOperations) putWhitespaceHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putWhitespaceHandleError(resp)
	}
	return resp.Response, nil
}

// putWhitespaceHandleError handles the PutWhitespace error response.
func (client *stringOperations) putWhitespaceHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}