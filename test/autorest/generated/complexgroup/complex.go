// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/complexgroup/internal/complexgroup"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// DefaultEndpoint is the default endpoint used for the ComplexClient service.
const DefaultEndpoint = "http://localhost:3000"

// ComplexClient is the test Infrastructure for AutoRest Swagger
type ComplexClient struct {
	s azinternal.Service
	u *url.URL
	p azcore.Pipeline
}

// ComplexClientOptions ...
type ComplexClientOptions struct {
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
	return azinternal.ColorValues()
}

// DefaultComplexClientOptions creates a ComplexClientOptions type initialized with default values.
func DefaultComplexClientOptions() ComplexClientOptions {
	return ComplexClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

// NewComplexClient creates an instance of the ComplexClient client.
func NewComplexClient(endpoint string, options *ComplexClientOptions) (*ComplexClient, error) {
	if options == nil {
		o := DefaultComplexClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewComplexClientWithPipeline(endpoint, p)
}

// NewComplexClientWithPipeline creates an instance of the ComplexClient client.
func NewComplexClientWithPipeline(endpoint string, p azcore.Pipeline) (*ComplexClient, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &ComplexClient{u: u, p: p}, nil
}

// GetValid Get complex type {id: 2, name: 'abc', color: 'YELLOW'}
func (client *ComplexClient) GetValid(ctx context.Context) (*GetValidResponse, error) {
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
func (client *ComplexClient) PutValid(ctx context.Context, complexBody Basic) (*PutValidResponse, error) {
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
