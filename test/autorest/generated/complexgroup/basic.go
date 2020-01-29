// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/complexgroup/internal/complexgroup"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// BasicClient is the test Infrastructure for AutoRest Swagger
type BasicClient struct {
	s azinternal.BasicClient
	u *url.URL
	p azcore.Pipeline
}

// BasicClientOptions ...
type BasicClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport

	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions

	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// PossibleColorValues ...
func PossibleColorValues() []ColorType {
	return azinternal.PossibleColorValues()
}

// DefaultBasicClientOptions creates a BasicClientOptions type initialized with default values.
func DefaultBasicClientOptions() BasicClientOptions {
	return BasicClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

// NewBasicClient creates an instance of the BasicClient client.
func NewBasicClient(options *BasicClientOptions) (*BasicClient, error) {
	return NewBasicClientWithEndpoint("http://localhost:3000", options)
}

// NewBasicClientWithEndpoint creates an instance of the BasicClient client.
func NewBasicClientWithEndpoint(endpoint string, options *BasicClientOptions) (*BasicClient, error) {
	if options == nil {
		o := DefaultBasicClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewBasicClientWithPipeline(endpoint, p)
}

// NewBasicClientWithPipeline creates an instance of the BasicClient client.
func NewBasicClientWithPipeline(endpoint string, p azcore.Pipeline) (*BasicClient, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &BasicClient{u: u, p: p}, nil
}

// GetValid Get complex type {id: 2, name: 'abc', color: 'YELLOW'}
func (client *BasicClient) GetValid(ctx context.Context) (*GetValidResponse, error) {
	req, err := client.s.GetValidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutValid Please put {id: 2, name: 'abc', color: 'Magenta'}
// Parameters:
// complexBody - Please put {id: 2, name: 'abc', color: 'Magenta'}
func (client *BasicClient) PutValid(ctx context.Context, complexBody Basic) (*PutValidResponse, error) {
	// TODO check validation requirements?
	req, err := client.s.PutValidCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.PutValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetInvalid ..
func (client *BasicClient) GetInvalid(ctx context.Context) (*GetInvalidResponse, error) {
	req, err := client.s.GetInvalidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetEmpty ..
func (client *BasicClient) GetEmpty(ctx context.Context) (*GetEmptyResponse, error) {
	req, err := client.s.GetEmptyCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetNull ..
func (client *BasicClient) GetNull(ctx context.Context) (*GetNullResponse, error) {
	req, err := client.s.GetNullCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetNotProvided ...
func (client *BasicClient) GetNotProvided(ctx context.Context) (*GetNotProvidedResponse, error) {
	req, err := client.s.GetNotProvidedCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetNotProvidedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}
