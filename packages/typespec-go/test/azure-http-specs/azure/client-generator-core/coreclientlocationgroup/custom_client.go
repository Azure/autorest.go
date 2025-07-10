package coreclientlocationgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewClientLocationMoveToExistingSubAdminOperationsClient(options *azcore.ClientOptions) (*ClientLocationMoveToExistingSubAdminOperationsClient, error) {
	internal, err := azcore.NewClient("clientnamespacegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientLocationMoveToExistingSubAdminOperationsClient{
		internal: internal,
	}, nil
}

func NewClientLocationMoveToExistingSubClient(options *azcore.ClientOptions) (*ClientLocationMoveToExistingSubClient, error) {
	internal, err := azcore.NewClient("clientnamespacegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientLocationMoveToExistingSubClient{
		internal: internal,
	}, nil
}

func NewClientLocationMoveToNewSubProductOperationsClient(options *azcore.ClientOptions) (*ClientLocationMoveToNewSubProductOperationsClient, error) {
	internal, err := azcore.NewClient("clientnamespacegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientLocationMoveToNewSubProductOperationsClient{
		internal: internal,
	}, nil
}

func NewClientLocationClient(options *azcore.ClientOptions) (*ClientLocationClient, error) {
	internal, err := azcore.NewClient("clientnamespacegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientLocationClient{
		internal: internal,
	}, nil
}
func NewClientLocationMoveToRootClient(options *azcore.ClientOptions) (*ClientLocationMoveToRootClient, error) {
	internal, err := azcore.NewClient("clientnamespacegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientLocationMoveToRootClient{
		internal: internal,
	}, nil
}
