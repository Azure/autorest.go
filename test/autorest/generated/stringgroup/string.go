// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/stringgroup/internal/stringgroup"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// DefaultEndpoint is the default endpoint used for the stringClient service.
const DefaultEndpoint = "http://localhost:3000"

// StringClient is the test Infrastructure for AutoRest Swagger BAT
type StringClient struct {
	s azinternal.Service
	u *url.URL
	p azcore.Pipeline
}

// StringClientOptions ...
type StringClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport

	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions

	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// DefaultStringClientOptions creates a StringClientOptions type initialized with default values.
func DefaultStringClientOptions() StringClientOptions {
	return StringClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

// NewStringClient creates an instance of the StringClient client.
func NewStringClient(endpoint string, options *StringClientOptions) (*StringClient, error) {
	if options == nil {
		o := DefaultStringClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewStringClientWithPipeline(endpoint, p)
}

// NewStringClientWithPipeline creates an instance of the StringClient client.
func NewStringClientWithPipeline(endpoint string, p azcore.Pipeline) (*StringClient, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &StringClient{u: u, p: p}, nil
}

// GetMBCS ..
func (client *StringClient) GetMBCS(ctx context.Context) (*GetMBCSResponse, error) {
	req, err := client.s.GetMBCSCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	ba, err := client.s.GetMBCSHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return ba, nil
}

// PutMBCS Set string value mbcs '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
// Parameters:
// stringBody - Set string value mbcs '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
func (client *StringClient) PutMBCS(ctx context.Context, stringBody string) (*PutMBCSResponse, error) {
	// TODO check validation requirements?
	req, err := client.s.PutMBCSCreateRequest(*client.u, stringBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	ba, err := client.s.PutMBCSHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return ba, nil
}
