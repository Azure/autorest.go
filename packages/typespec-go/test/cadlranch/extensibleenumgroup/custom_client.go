//go:build go1.18
// +build go1.18

package extensibleenumgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewExtensibleClient(options *azcore.ClientOptions) (*ExtensibleClient, error) {
	internal, err := azcore.NewClient("extensibleenumgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ExtensibleClient{
		internal: internal,
	}, nil
}
