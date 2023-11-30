//go:build go1.18
// +build go1.18

package visibilitygroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewVisibilityClient(options *azcore.ClientOptions) (*VisibilityClient, error) {
	internal, err := azcore.NewClient("visibilitygroup", "v0.1.1", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &VisibilityClient{
		internal: internal,
	}, nil
}
