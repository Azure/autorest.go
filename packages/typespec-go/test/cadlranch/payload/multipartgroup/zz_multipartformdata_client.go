// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package multipartgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
)

// MultiPartFormDataClient contains the methods for the MultiPartFormData group.
// Don't use this type directly, use [MultiPartClient.NewMultiPartFormDataClient] instead.
type MultiPartFormDataClient struct {
	internal *azcore.Client
}

// AnonymousModel - Test content-type: multipart/form-data
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientAnonymousModelOptions contains the optional parameters for the MultiPartFormDataClient.AnonymousModel
//     method.
func (client *MultiPartFormDataClient) AnonymousModel(ctx context.Context, profileImage io.ReadSeekCloser, options *MultiPartFormDataClientAnonymousModelOptions) (MultiPartFormDataClientAnonymousModelResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.AnonymousModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.anonymousModelCreateRequest(ctx, profileImage, options)
	if err != nil {
		return MultiPartFormDataClientAnonymousModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientAnonymousModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientAnonymousModelResponse{}, err
	}
	return MultiPartFormDataClientAnonymousModelResponse{}, nil
}

// anonymousModelCreateRequest creates the AnonymousModel request.
func (client *MultiPartFormDataClient) anonymousModelCreateRequest(ctx context.Context, profileImage io.ReadSeekCloser, _ *MultiPartFormDataClientAnonymousModelOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/anonymous-model"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData := map[string]any{}
	formData["profileImage"] = profileImage
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// Basic - Test content-type: multipart/form-data
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientBasicOptions contains the optional parameters for the MultiPartFormDataClient.Basic method.
func (client *MultiPartFormDataClient) Basic(ctx context.Context, body MultiPartRequest, options *MultiPartFormDataClientBasicOptions) (MultiPartFormDataClientBasicResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.Basic"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.basicCreateRequest(ctx, body, options)
	if err != nil {
		return MultiPartFormDataClientBasicResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientBasicResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientBasicResponse{}, err
	}
	return MultiPartFormDataClientBasicResponse{}, nil
}

// basicCreateRequest creates the Basic request.
func (client *MultiPartFormDataClient) basicCreateRequest(ctx context.Context, body MultiPartRequest, _ *MultiPartFormDataClientBasicOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/mixed-parts"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData, err := body.toMultipartFormData()
	if err != nil {
		return nil, err
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// BinaryArrayParts - Test content-type: multipart/form-data for scenario contains multi binary parts
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientBinaryArrayPartsOptions contains the optional parameters for the MultiPartFormDataClient.BinaryArrayParts
//     method.
func (client *MultiPartFormDataClient) BinaryArrayParts(ctx context.Context, body BinaryArrayPartsRequest, options *MultiPartFormDataClientBinaryArrayPartsOptions) (MultiPartFormDataClientBinaryArrayPartsResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.BinaryArrayParts"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.binaryArrayPartsCreateRequest(ctx, body, options)
	if err != nil {
		return MultiPartFormDataClientBinaryArrayPartsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientBinaryArrayPartsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientBinaryArrayPartsResponse{}, err
	}
	return MultiPartFormDataClientBinaryArrayPartsResponse{}, nil
}

// binaryArrayPartsCreateRequest creates the BinaryArrayParts request.
func (client *MultiPartFormDataClient) binaryArrayPartsCreateRequest(ctx context.Context, body BinaryArrayPartsRequest, _ *MultiPartFormDataClientBinaryArrayPartsOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/binary-array-parts"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData, err := body.toMultipartFormData()
	if err != nil {
		return nil, err
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// CheckFileNameAndContentType - Test content-type: multipart/form-data
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientCheckFileNameAndContentTypeOptions contains the optional parameters for the MultiPartFormDataClient.CheckFileNameAndContentType
//     method.
func (client *MultiPartFormDataClient) CheckFileNameAndContentType(ctx context.Context, body MultiPartRequest, options *MultiPartFormDataClientCheckFileNameAndContentTypeOptions) (MultiPartFormDataClientCheckFileNameAndContentTypeResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.CheckFileNameAndContentType"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkFileNameAndContentTypeCreateRequest(ctx, body, options)
	if err != nil {
		return MultiPartFormDataClientCheckFileNameAndContentTypeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientCheckFileNameAndContentTypeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientCheckFileNameAndContentTypeResponse{}, err
	}
	return MultiPartFormDataClientCheckFileNameAndContentTypeResponse{}, nil
}

// checkFileNameAndContentTypeCreateRequest creates the CheckFileNameAndContentType request.
func (client *MultiPartFormDataClient) checkFileNameAndContentTypeCreateRequest(ctx context.Context, body MultiPartRequest, _ *MultiPartFormDataClientCheckFileNameAndContentTypeOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/check-filename-and-content-type"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData, err := body.toMultipartFormData()
	if err != nil {
		return nil, err
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// Complex - Test content-type: multipart/form-data for mixed scenarios
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientComplexOptions contains the optional parameters for the MultiPartFormDataClient.Complex
//     method.
func (client *MultiPartFormDataClient) Complex(ctx context.Context, body ComplexPartsRequest, options *MultiPartFormDataClientComplexOptions) (MultiPartFormDataClientComplexResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.Complex"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.complexCreateRequest(ctx, body, options)
	if err != nil {
		return MultiPartFormDataClientComplexResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientComplexResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientComplexResponse{}, err
	}
	return MultiPartFormDataClientComplexResponse{}, nil
}

// complexCreateRequest creates the Complex request.
func (client *MultiPartFormDataClient) complexCreateRequest(ctx context.Context, body ComplexPartsRequest, _ *MultiPartFormDataClientComplexOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/complex-parts"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData, err := body.toMultipartFormData()
	if err != nil {
		return nil, err
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// ComplexWithHTTPPart - Test content-type: multipart/form-data for mixed scenarios
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientComplexWithHTTPPartOptions contains the optional parameters for the MultiPartFormDataClient.ComplexWithHTTPPart
//     method.
func (client *MultiPartFormDataClient) ComplexWithHTTPPart(ctx context.Context, body ComplexHTTPPartsModelRequest, options *MultiPartFormDataClientComplexWithHTTPPartOptions) (MultiPartFormDataClientComplexWithHTTPPartResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.ComplexWithHTTPPart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.complexWithHTTPPartCreateRequest(ctx, body, options)
	if err != nil {
		return MultiPartFormDataClientComplexWithHTTPPartResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientComplexWithHTTPPartResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientComplexWithHTTPPartResponse{}, err
	}
	return MultiPartFormDataClientComplexWithHTTPPartResponse{}, nil
}

// complexWithHTTPPartCreateRequest creates the ComplexWithHTTPPart request.
func (client *MultiPartFormDataClient) complexWithHTTPPartCreateRequest(ctx context.Context, body ComplexHTTPPartsModelRequest, _ *MultiPartFormDataClientComplexWithHTTPPartOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/complex-parts-with-httppart"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData, err := body.toMultipartFormData()
	if err != nil {
		return nil, err
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// FileWithHTTPPartOptionalContentType - Test content-type: multipart/form-data for optional content type
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientFileWithHTTPPartOptionalContentTypeOptions contains the optional parameters for the MultiPartFormDataClient.FileWithHTTPPartOptionalContentType
//     method.
func (client *MultiPartFormDataClient) FileWithHTTPPartOptionalContentType(ctx context.Context, body FileWithHTTPPartOptionalContentTypeRequest, options *MultiPartFormDataClientFileWithHTTPPartOptionalContentTypeOptions) (MultiPartFormDataClientFileWithHTTPPartOptionalContentTypeResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.FileWithHTTPPartOptionalContentType"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fileWithHTTPPartOptionalContentTypeCreateRequest(ctx, body, options)
	if err != nil {
		return MultiPartFormDataClientFileWithHTTPPartOptionalContentTypeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientFileWithHTTPPartOptionalContentTypeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientFileWithHTTPPartOptionalContentTypeResponse{}, err
	}
	return MultiPartFormDataClientFileWithHTTPPartOptionalContentTypeResponse{}, nil
}

// fileWithHTTPPartOptionalContentTypeCreateRequest creates the FileWithHTTPPartOptionalContentType request.
func (client *MultiPartFormDataClient) fileWithHTTPPartOptionalContentTypeCreateRequest(ctx context.Context, body FileWithHTTPPartOptionalContentTypeRequest, _ *MultiPartFormDataClientFileWithHTTPPartOptionalContentTypeOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/file-with-http-part-optional-content-type"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData, err := body.toMultipartFormData()
	if err != nil {
		return nil, err
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// FileWithHTTPPartRequiredContentType - Test content-type: multipart/form-data
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientFileWithHTTPPartRequiredContentTypeOptions contains the optional parameters for the MultiPartFormDataClient.FileWithHTTPPartRequiredContentType
//     method.
func (client *MultiPartFormDataClient) FileWithHTTPPartRequiredContentType(ctx context.Context, body FileWithHTTPPartRequiredContentTypeRequest, options *MultiPartFormDataClientFileWithHTTPPartRequiredContentTypeOptions) (MultiPartFormDataClientFileWithHTTPPartRequiredContentTypeResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.FileWithHTTPPartRequiredContentType"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fileWithHTTPPartRequiredContentTypeCreateRequest(ctx, body, options)
	if err != nil {
		return MultiPartFormDataClientFileWithHTTPPartRequiredContentTypeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientFileWithHTTPPartRequiredContentTypeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientFileWithHTTPPartRequiredContentTypeResponse{}, err
	}
	return MultiPartFormDataClientFileWithHTTPPartRequiredContentTypeResponse{}, nil
}

// fileWithHTTPPartRequiredContentTypeCreateRequest creates the FileWithHTTPPartRequiredContentType request.
func (client *MultiPartFormDataClient) fileWithHTTPPartRequiredContentTypeCreateRequest(ctx context.Context, body FileWithHTTPPartRequiredContentTypeRequest, _ *MultiPartFormDataClientFileWithHTTPPartRequiredContentTypeOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/check-filename-and-required-content-type-with-httppart"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData, err := body.toMultipartFormData()
	if err != nil {
		return nil, err
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// FileWithHTTPPartSpecificContentType - Test content-type: multipart/form-data
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientFileWithHTTPPartSpecificContentTypeOptions contains the optional parameters for the MultiPartFormDataClient.FileWithHTTPPartSpecificContentType
//     method.
func (client *MultiPartFormDataClient) FileWithHTTPPartSpecificContentType(ctx context.Context, body FileWithHTTPPartSpecificContentTypeRequest, options *MultiPartFormDataClientFileWithHTTPPartSpecificContentTypeOptions) (MultiPartFormDataClientFileWithHTTPPartSpecificContentTypeResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.FileWithHTTPPartSpecificContentType"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fileWithHTTPPartSpecificContentTypeCreateRequest(ctx, body, options)
	if err != nil {
		return MultiPartFormDataClientFileWithHTTPPartSpecificContentTypeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientFileWithHTTPPartSpecificContentTypeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientFileWithHTTPPartSpecificContentTypeResponse{}, err
	}
	return MultiPartFormDataClientFileWithHTTPPartSpecificContentTypeResponse{}, nil
}

// fileWithHTTPPartSpecificContentTypeCreateRequest creates the FileWithHTTPPartSpecificContentType request.
func (client *MultiPartFormDataClient) fileWithHTTPPartSpecificContentTypeCreateRequest(ctx context.Context, body FileWithHTTPPartSpecificContentTypeRequest, _ *MultiPartFormDataClientFileWithHTTPPartSpecificContentTypeOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/check-filename-and-specific-content-type-with-httppart"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData, err := body.toMultipartFormData()
	if err != nil {
		return nil, err
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// JSONPart - Test content-type: multipart/form-data for scenario contains json part and binary part
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientJSONPartOptions contains the optional parameters for the MultiPartFormDataClient.JSONPart
//     method.
func (client *MultiPartFormDataClient) JSONPart(ctx context.Context, body JSONPartRequest, options *MultiPartFormDataClientJSONPartOptions) (MultiPartFormDataClientJSONPartResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.JSONPart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.jsonPartCreateRequest(ctx, body, options)
	if err != nil {
		return MultiPartFormDataClientJSONPartResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientJSONPartResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientJSONPartResponse{}, err
	}
	return MultiPartFormDataClientJSONPartResponse{}, nil
}

// jsonPartCreateRequest creates the JSONPart request.
func (client *MultiPartFormDataClient) jsonPartCreateRequest(ctx context.Context, body JSONPartRequest, _ *MultiPartFormDataClientJSONPartOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/json-part"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData, err := body.toMultipartFormData()
	if err != nil {
		return nil, err
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// MultiBinaryParts - Test content-type: multipart/form-data for scenario contains multi binary parts
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultiPartFormDataClientMultiBinaryPartsOptions contains the optional parameters for the MultiPartFormDataClient.MultiBinaryParts
//     method.
func (client *MultiPartFormDataClient) MultiBinaryParts(ctx context.Context, body MultiBinaryPartsRequest, options *MultiPartFormDataClientMultiBinaryPartsOptions) (MultiPartFormDataClientMultiBinaryPartsResponse, error) {
	var err error
	const operationName = "MultiPartFormDataClient.MultiBinaryParts"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.multiBinaryPartsCreateRequest(ctx, body, options)
	if err != nil {
		return MultiPartFormDataClientMultiBinaryPartsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultiPartFormDataClientMultiBinaryPartsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultiPartFormDataClientMultiBinaryPartsResponse{}, err
	}
	return MultiPartFormDataClientMultiBinaryPartsResponse{}, nil
}

// multiBinaryPartsCreateRequest creates the MultiBinaryParts request.
func (client *MultiPartFormDataClient) multiBinaryPartsCreateRequest(ctx context.Context, body MultiBinaryPartsRequest, _ *MultiPartFormDataClientMultiBinaryPartsOptions) (*policy.Request, error) {
	urlPath := "/multipart/form-data/multi-binary-parts"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"multipart/form-data"}
	formData, err := body.toMultipartFormData()
	if err != nil {
		return nil, err
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}
