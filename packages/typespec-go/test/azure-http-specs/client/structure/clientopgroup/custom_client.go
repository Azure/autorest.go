package clientopgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewFirstClient(endpoint string, clientType ClientType, options *azcore.ClientOptions) (*FirstClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FirstClient{
		internal: internal,
		endpoint: endpoint,
		client:   clientType,
	}, nil
}

func NewSecondClient(endpoint string, clientType ClientType, options *azcore.ClientOptions) (*SecondClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &SecondClient{
		internal: internal,
		endpoint: endpoint,
		client:   clientType,
	}, nil
}

func NewSecondGroup5Client(endpoint string, clientType ClientType, options *azcore.ClientOptions) (*SecondGroup5Client, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &SecondGroup5Client{
		internal: internal,
		endpoint: endpoint,
		client:   clientType,
	}, nil
}

func NewFirstGroup3Client(endpoint string, clientType ClientType, options *azcore.ClientOptions) (*FirstGroup3Client, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FirstGroup3Client{
		internal: internal,
		endpoint: endpoint,
		client:   clientType,
	}, nil
}

func NewFirstGroup4Client(endpoint string, clientType ClientType, options *azcore.ClientOptions) (*FirstGroup4Client, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FirstGroup4Client{
		internal: internal,
		endpoint: endpoint,
		client:   clientType,
	}, nil
}
