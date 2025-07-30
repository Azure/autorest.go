package coreclientlocationgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewClientLocationMoveToExistingSubAdminOperationsClient(endpoint string, options *azcore.ClientOptions) (*ClientLocationMoveToExistingSubAdminOperationsClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientLocationMoveToExistingSubAdminOperationsClient{
		internal: internal,
		endpoint: endpoint,
	}, nil
}

func NewClientLocationMoveToExistingSubClient(endpoint string, options *azcore.ClientOptions) (*ClientLocationMoveToExistingSubClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientLocationMoveToExistingSubClient{
		internal: internal,
		endpoint: endpoint,
	}, nil

}

func NewClientLocationMoveToNewSubProductOperationsClient(endpoint string, options *azcore.ClientOptions) (*ClientLocationMoveToNewSubProductOperationsClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientLocationMoveToNewSubProductOperationsClient{
		internal: internal,
		endpoint: endpoint,
	}, nil
}

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

func NewClientLocationMoveToRootClient(endpoint string, options *azcore.ClientOptions) (*ClientLocationMoveToRootClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientLocationMoveToRootClient{
		internal: internal,
		endpoint: endpoint,
	}, nil
}
