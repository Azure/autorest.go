// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegroup

import (
	"context"
	azinternal "generatortests/autorest/generated/bytegroup/internal/bytegroup"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// DefaultEndpoint is the default endpoint used for the byteclient service.
const DefaultEndpoint = "http://localhost:3000"

// ByteClient is the test Infrastructure for AutoRest Swagger BAT
type ByteClient struct {
	s azinternal.Service
	u *url.URL
	p azcore.Pipeline
}

// ByteClientOptions ...
type ByteClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport

	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions

	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// DefaultByteClientOptions creates a ByteClientOptions type initialized with default values.
func DefaultByteClientOptions() ByteClientOptions {
	return ByteClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

// NewByteClient creates an instance of the ByteClient client.
func NewByteClient(endpoint string, options *ByteClientOptions) (*ByteClient, error) {
	if options == nil {
		o := DefaultByteClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewByteClientWithPipeline(endpoint, p)
}

// NewByteClientWithPipeline creates an instance of the ByteClient client.
func NewByteClientWithPipeline(endpoint string, p azcore.Pipeline) (*ByteClient, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &ByteClient{u: u, p: p}, nil
}

// GetEmpty get empty byte value ''
func (client *ByteClient) GetEmpty(ctx context.Context) (*GetEmptyResponse, error) {
	req, err := client.s.GetEmptyCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	ba, err := client.s.GetEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return ba, nil
}

// GetInvalid get invalid byte value ':::SWAGGER::::'
func (client *ByteClient) GetInvalid(ctx context.Context) (*GetInvalidResponse, error) {
	req, err := client.s.GetInvalidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	ba, err := client.s.GetInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return ba, nil
}

// GetNonASCII get non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
func (client *ByteClient) GetNonASCII(ctx context.Context) (*GetNonASCIIResponse, error) {
	req, err := client.s.GetNonASCIICreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	ba, err := client.s.GetNonASCIIHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return ba, nil
}

// GetNull get null byte value
func (client *ByteClient) GetNull(ctx context.Context) (*GetNullResponse, error) {
	req, err := client.s.GetNullCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	ba, err := client.s.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return ba, nil
}

// PutNonASCII put non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
// Parameters:
// byteBody - base64-encoded non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
func (client *ByteClient) PutNonASCII(ctx context.Context, byteBody []byte) (*PutNonASCIIResponse, error) {
	// TODO check validation requirements?
	req, err := client.s.PutNonASCIICreateRequest(*client.u, byteBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	ba, err := client.s.PutNonASCIIHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return ba, nil
}
