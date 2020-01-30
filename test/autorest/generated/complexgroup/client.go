// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ClientOptions ...
type ClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport

	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions

	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// DefaultClientOptions creates a ClientOptions type initialized with default values.
func DefaultClientOptions() ClientOptions {
	return ClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

// Client is the test Infrastructure for AutoRest Swagger
type Client struct {
	u                   *url.URL
	p                   azcore.Pipeline
	basicOperations     BasicOperations
	primitiveOperations PrimitiveOperations
}

// NewDefaultClient creates an instance of the Client type.
// It uses the default endpoint http://localhost:3000
func NewDefaultClient(options *ClientOptions) (*Client, error) {
	return NewClient("http://localhost:3000", options)
}

// NewClient creates an instance of the Client type with the specified endpoint.
func NewClient(endpoint string, options *ClientOptions) (*Client, error) {
	if options == nil {
		o := DefaultClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewClientWithPipeline(endpoint, p)
}

// NewClientWithPipeline creates an instance of the Client type with the specified endpoint and pipeline.
func NewClientWithPipeline(endpoint string, p azcore.Pipeline) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c := &Client{u: u, p: p}
	c.basicOperations = &basicOperations{Client: c}
	c.primitiveOperations = &primitiveOperations{Client: c}
	return c, nil
}

// BasicOperations returns the BasicOperations associated with this client.
func (client *Client) BasicOperations() BasicOperations {
	return client.basicOperations
}

// PrimitiveOperations returns the PrimitiveOperations associated with this client.
func (client *Client) PrimitiveOperations() PrimitiveOperations {
	return client.primitiveOperations
}
