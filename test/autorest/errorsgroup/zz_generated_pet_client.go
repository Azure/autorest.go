// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// PetClient contains the methods for the Pet group.
// Don't use this type directly, use NewPetClient() instead.
type PetClient struct {
	con *Connection
}

// NewPetClient creates a new instance of PetClient with the specified values.
func NewPetClient(con *Connection) *PetClient {
	return &PetClient{con: con}
}

// DoSomething - Asks pet to do something
// If the operation fails it returns one of the following error types.
// - *PetActionError, *PetHungryOrThirstyError, *PetSadError
func (client *PetClient) DoSomething(ctx context.Context, whatAction string, options *PetDoSomethingOptions) (PetDoSomethingResponse, error) {
	req, err := client.doSomethingCreateRequest(ctx, whatAction, options)
	if err != nil {
		return PetDoSomethingResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PetDoSomethingResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PetDoSomethingResponse{}, client.doSomethingHandleError(resp)
	}
	return client.doSomethingHandleResponse(resp)
}

// doSomethingCreateRequest creates the DoSomething request.
func (client *PetClient) doSomethingCreateRequest(ctx context.Context, whatAction string, options *PetDoSomethingOptions) (*azcore.Request, error) {
	urlPath := "/errorStatusCodes/Pets/doSomething/{whatAction}"
	if whatAction == "" {
		return nil, errors.New("parameter whatAction cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{whatAction}", url.PathEscape(whatAction))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// doSomethingHandleResponse handles the DoSomething response.
func (client *PetClient) doSomethingHandleResponse(resp *azcore.Response) (PetDoSomethingResponse, error) {
	result := PetDoSomethingResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PetAction); err != nil {
		return PetDoSomethingResponse{}, err
	}
	return result, nil
}

// doSomethingHandleError handles the DoSomething error response.
func (client *PetClient) doSomethingHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	var errType petActionError
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(errType.wrapped, resp.Response)
}

// GetPetByID - Gets pets by id.
// If the operation fails it returns one of the following error types.
// - *AnimalNotFound, *LinkNotFound, *NotFoundErrorBase
func (client *PetClient) GetPetByID(ctx context.Context, petID string, options *PetGetPetByIDOptions) (PetGetPetByIDResponse, error) {
	req, err := client.getPetByIDCreateRequest(ctx, petID, options)
	if err != nil {
		return PetGetPetByIDResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PetGetPetByIDResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return PetGetPetByIDResponse{}, client.getPetByIDHandleError(resp)
	}
	return client.getPetByIDHandleResponse(resp)
}

// getPetByIDCreateRequest creates the GetPetByID request.
func (client *PetClient) getPetByIDCreateRequest(ctx context.Context, petID string, options *PetGetPetByIDOptions) (*azcore.Request, error) {
	urlPath := "/errorStatusCodes/Pets/{petId}/GetPet"
	if petID == "" {
		return nil, errors.New("parameter petID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{petId}", url.PathEscape(petID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getPetByIDHandleResponse handles the GetPetByID response.
func (client *PetClient) getPetByIDHandleResponse(resp *azcore.Response) (PetGetPetByIDResponse, error) {
	result := PetGetPetByIDResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Pet); err != nil {
		return PetGetPetByIDResponse{}, err
	}
	return result, nil
}

// getPetByIDHandleError handles the GetPetByID error response.
func (client *PetClient) getPetByIDHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	switch resp.StatusCode {
	case http.StatusBadRequest:
		var errType string
		if err := resp.UnmarshalAsJSON(&errType); err != nil {
			return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
		}
		return azcore.NewResponseError(fmt.Errorf("%v", errType), resp.Response)
	case http.StatusNotFound:
		var errType notFoundErrorBase
		if err := resp.UnmarshalAsJSON(&errType); err != nil {
			return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
		}
		return azcore.NewResponseError(errType.wrapped, resp.Response)
	case http.StatusNotImplemented:
		var errType int32
		if err := resp.UnmarshalAsJSON(&errType); err != nil {
			return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
		}
		return azcore.NewResponseError(fmt.Errorf("%v", errType), resp.Response)
	default:
		if len(body) == 0 {
			return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
		}
		return azcore.NewResponseError(errors.New(string(body)), resp.Response)
	}
}

// HasModelsParam - Ensure you can correctly deserialize the returned PetActionError and deserialization doesn't conflict with the input param name 'models'
// If the operation fails it returns one of the following error types.
// - *PetActionError, *PetHungryOrThirstyError, *PetSadError
func (client *PetClient) HasModelsParam(ctx context.Context, options *PetHasModelsParamOptions) (PetHasModelsParamResponse, error) {
	req, err := client.hasModelsParamCreateRequest(ctx, options)
	if err != nil {
		return PetHasModelsParamResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PetHasModelsParamResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PetHasModelsParamResponse{}, client.hasModelsParamHandleError(resp)
	}
	return PetHasModelsParamResponse{RawResponse: resp.Response}, nil
}

// hasModelsParamCreateRequest creates the HasModelsParam request.
func (client *PetClient) hasModelsParamCreateRequest(ctx context.Context, options *PetHasModelsParamOptions) (*azcore.Request, error) {
	urlPath := "/errorStatusCodes/Pets/hasModelsParam"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Models != nil {
		reqQP.Set("models", *options.Models)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// hasModelsParamHandleError handles the HasModelsParam error response.
func (client *PetClient) hasModelsParamHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	var errType petActionError
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(errType.wrapped, resp.Response)
}
