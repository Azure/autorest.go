package clientnamespacegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewClientNamespaceClient(options *azcore.ClientOptions) (*ClientNamespaceClient, error) {
	internal, err := azcore.NewClient("clientnamespacegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ClientNamespaceClient{
		internal: internal,
	}, nil
}