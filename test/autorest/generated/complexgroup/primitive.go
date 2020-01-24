package complexgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/complexgroup/internal/complexgroup"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// PrimitiveClient is the test Infrastructure for AutoRest Swagger
type PrimitiveClient struct {
	s azinternal.Service
	u *url.URL
	p azcore.Pipeline
}

// PrimitiveClientOptions ...
type PrimitiveClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport

	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions

	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// DefaultPrimitiveClientOptions creates a PrimitiveClientOptions type initialized with default values.
func DefaultPrimitiveClientOptions() PrimitiveClientOptions {
	return PrimitiveClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

// NewPrimitiveClient creates an instance of the PrimitiveClient client.
func NewPrimitiveClient(endpoint string, options *PrimitiveClientOptions) (*PrimitiveClient, error) {
	if options == nil {
		o := DefaultPrimitiveClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewPrimitiveClientWithPipeline(endpoint, p)
}

// NewPrimitiveClientWithPipeline creates an instance of the PrimitiveClient client.
func NewPrimitiveClientWithPipeline(endpoint string, p azcore.Pipeline) (*PrimitiveClient, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &PrimitiveClient{u: u, p: p}, nil
}

// GetInt ...
func (client *PrimitiveClient) GetInt(ctx context.Context) (*GetIntResponse, error) {
	req, err := client.s.GetIntCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetIntHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutInt ...
func (client *PrimitiveClient) PutInt(ctx context.Context, complexBody IntWrapper) (*PutIntResponse, error) {
	req, err := client.s.PutIntCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.PutIntHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}
