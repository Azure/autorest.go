// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/complexgroup/internal/complexgroup"
)

// BasicOperations contains the methods for the Basic group.
type BasicOperations interface {
	GetEmpty(ctx context.Context) (*GetEmptyResponse, error)
	GetInvalid(ctx context.Context) (*GetInvalidResponse, error)
	GetNotProvided(ctx context.Context) (*GetNotProvidedResponse, error)
	GetNull(ctx context.Context) (*GetNullResponse, error)
	GetValid(ctx context.Context) (*GetValidResponse, error)
	PutValid(ctx context.Context, complexBody Basic) (*PutValidResponse, error)
}

type basicOperations struct {
	*Client
	azinternal.BasicOperations
}

// GetValid Get complex type {id: 2, name: 'abc', color: 'YELLOW'}
func (client *basicOperations) GetValid(ctx context.Context) (*GetValidResponse, error) {
	req, err := client.GetValidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutValid Please put {id: 2, name: 'abc', color: 'Magenta'}
// Parameters:
// complexBody - Please put {id: 2, name: 'abc', color: 'Magenta'}
func (client *basicOperations) PutValid(ctx context.Context, complexBody Basic) (*PutValidResponse, error) {
	// TODO check validation requirements?
	req, err := client.PutValidCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.PutValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetInvalid ..
func (client *basicOperations) GetInvalid(ctx context.Context) (*GetInvalidResponse, error) {
	req, err := client.GetInvalidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetEmpty ..
func (client *basicOperations) GetEmpty(ctx context.Context) (*GetEmptyResponse, error) {
	req, err := client.GetEmptyCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetNull ..
func (client *basicOperations) GetNull(ctx context.Context) (*GetNullResponse, error) {
	req, err := client.GetNullCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetNotProvided ...
func (client *basicOperations) GetNotProvided(ctx context.Context) (*GetNotProvidedResponse, error) {
	req, err := client.GetNotProvidedCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetNotProvidedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

var _ BasicOperations = (*basicOperations)(nil)
