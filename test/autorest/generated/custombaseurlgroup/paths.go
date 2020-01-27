// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package custombaseurlgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/custombaseurlgroup/internal/custombaseurlgroup"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// CustomBaseURLClient is the test Infrastructure for AutoRest Swagger
type CustomBaseURLClient struct {
	s azinternal.PathsClient
	u *url.URL
	p azcore.Pipeline
}

// CustomBaseURLClientOptions ...
type CustomBaseURLClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport

	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions

	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// DefaultCustomBaseURLClientOptions creates a CustomBaseURLClientOptions type initialized with default values.
func DefaultCustomBaseURLClientOptions() CustomBaseURLClientOptions {
	return CustomBaseURLClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

// NewCustomBaseURLClient creates an instance of the CustomBaseURLClient.
func NewCustomBaseURLClient(endpoint string, options *CustomBaseURLClientOptions) (*CustomBaseURLClient, error) {
	if options == nil {
		o := DefaultCustomBaseURLClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewCustomBaseURLClientWithPipeline(endpoint, p)
}

// NewCustomBaseURLClientWithPipeline creates an instance of the CustomBaseURLClient with a custom pipeline.
func NewCustomBaseURLClientWithPipeline(endpoint string, p azcore.Pipeline) (*CustomBaseURLClient, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &CustomBaseURLClient{u: u, p: p}, nil
}

// GetEmpty ...
func (client *CustomBaseURLClient) GetEmpty(ctx context.Context) (*GetEmptyResponse, error) {
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
