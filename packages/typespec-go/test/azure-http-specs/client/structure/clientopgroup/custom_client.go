package clientopgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewFirstClient(options *azcore.ClientOptions) (*FirstClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FirstClient{
		internal: internal,
	}, nil
}

func NewSecondClient(options *azcore.ClientOptions) (*SecondClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &SecondClient{
		internal: internal,
	}, nil
}

func NewSecondGroup5Client(options *azcore.ClientOptions) (*SecondGroup5Client, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &SecondGroup5Client{
		internal: internal,
	}, nil
}

func NewFirstGroup3Client(options *azcore.ClientOptions) (*FirstGroup3Client, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FirstGroup3Client{
		internal: internal,
	}, nil
}

func NewFirstGroup4Client(options *azcore.ClientOptions) (*FirstGroup4Client, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FirstGroup4Client{
		internal: internal,
	}, nil
}
