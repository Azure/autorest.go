package coreclientlocationgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewClientLocationClient(endpoint string, options *azcore.ClientOptions) (*ClientLocationClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientLocationClient{
		internal: internal,
		endpoint: endpoint,
	}, nil
}
