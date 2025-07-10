package traitsgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewTraitsClient(options *azcore.ClientOptions) (*TraitsClient, error) {
	internal, err := azcore.NewClient("traitsgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &TraitsClient{
		internal: internal,
	}, nil
}
