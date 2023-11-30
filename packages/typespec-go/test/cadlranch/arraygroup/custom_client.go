//go:build go1.18
// +build go1.18

package arraygroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewBooleanValueClient(options *azcore.ClientOptions) (*BooleanValueClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &BooleanValueClient{
		internal: internal,
	}, nil
}

func NewDatetimeValueClient(options *azcore.ClientOptions) (*DatetimeValueClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &DatetimeValueClient{
		internal: internal,
	}, nil
}

func NewDurationValueClient(options *azcore.ClientOptions) (*DurationValueClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &DurationValueClient{
		internal: internal,
	}, nil
}

func NewFloat32ValueClient(options *azcore.ClientOptions) (*Float32ValueClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &Float32ValueClient{
		internal: internal,
	}, nil
}

func NewInt32ValueClient(options *azcore.ClientOptions) (*Int32ValueClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &Int32ValueClient{
		internal: internal,
	}, nil
}

func NewInt64ValueClient(options *azcore.ClientOptions) (*Int64ValueClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &Int64ValueClient{
		internal: internal,
	}, nil
}

func NewModelValueClient(options *azcore.ClientOptions) (*ModelValueClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ModelValueClient{
		internal: internal,
	}, nil
}

func NewNullableFloatValueClient(options *azcore.ClientOptions) (*NullableFloatValueClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &NullableFloatValueClient{
		internal: internal,
	}, nil
}

func NewStringValueClient(options *azcore.ClientOptions) (*StringValueClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &StringValueClient{
		internal: internal,
	}, nil
}

func NewUnknownValueClient(options *azcore.ClientOptions) (*UnknownValueClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &UnknownValueClient{
		internal: internal,
	}, nil
}

func newClient(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient("arraygroup", "v0.1.0", runtime.PipelineOptions{}, options)
}
