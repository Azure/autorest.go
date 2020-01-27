// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/complexgroup/internal/complexgroup"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// PrimitiveClient is the test Infrastructure for AutoRest Swagger
type PrimitiveClient struct {
	s azinternal.PrimitiveClient
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

// GetLong ...
func (client *PrimitiveClient) GetLong(ctx context.Context) (*GetLongResponse, error) {
	req, err := client.s.GetLongCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutLong ...
func (client *PrimitiveClient) PutLong(ctx context.Context, complexBody LongWrapper) (*PutLongResponse, error) {
	req, err := client.s.PutLongCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.PutLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetFloat ...
func (client *PrimitiveClient) GetFloat(ctx context.Context) (*GetFloatResponse, error) {
	req, err := client.s.GetFloatCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutFloat ...
func (client *PrimitiveClient) PutFloat(ctx context.Context, complexBody FloatWrapper) (*PutFloatResponse, error) {
	req, err := client.s.PutFloatCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.PutFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetDouble ...
func (client *PrimitiveClient) GetDouble(ctx context.Context) (*GetDoubleResponse, error) {
	req, err := client.s.GetDoubleCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutDouble ...
func (client *PrimitiveClient) PutDouble(ctx context.Context, complexBody DoubleWrapper) (*PutDoubleResponse, error) {
	req, err := client.s.PutDoubleCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.PutDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetBool ...
func (client *PrimitiveClient) GetBool(ctx context.Context) (*GetBoolResponse, error) {
	req, err := client.s.GetBoolCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutBool ...
func (client *PrimitiveClient) PutBool(ctx context.Context, complexBody BooleanWrapper) (*PutBoolResponse, error) {
	req, err := client.s.PutBoolCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.PutBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetString ...
func (client *PrimitiveClient) GetString(ctx context.Context) (*GetStringResponse, error) {
	req, err := client.s.GetStringCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.GetStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutString ...
func (client *PrimitiveClient) PutString(ctx context.Context, complexBody StringWrapper) (*PutStringResponse, error) {
	req, err := client.s.PutStringCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.s.PutStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// // GetDate ...
// func (client *PrimitiveClient) GetDate(ctx context.Context) (*GetDateResponse, error) {
// 	req, err := client.s.GetDateCreateRequest(*client.u)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp, err := client.p.Do(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	s, err := client.s.GetDateHandleResponse(resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return s, nil
// }

// // PutDate ...
// func (client *PrimitiveClient) PutDate(ctx context.Context, complexBody DateWrapper) (*PutDateResponse, error) {
// 	req, err := client.s.PutDateCreateRequest(*client.u, complexBody)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp, err := client.p.Do(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	s, err := client.s.PutDateHandleResponse(resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return s, nil
// }
