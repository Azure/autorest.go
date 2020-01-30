// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/complexgroup/internal/complexgroup"
)

// PrimitiveOperations contains the methods for the Primitive group.
type PrimitiveOperations interface {
	GetBool(ctx context.Context) (*GetBoolResponse, error)
	GetDouble(ctx context.Context) (*GetDoubleResponse, error)
	GetFloat(ctx context.Context) (*GetFloatResponse, error)
	GetInt(ctx context.Context) (*GetIntResponse, error)
	GetLong(ctx context.Context) (*GetLongResponse, error)
	GetString(ctx context.Context) (*GetStringResponse, error)
	PutBool(ctx context.Context, complexBody BooleanWrapper) (*PutBoolResponse, error)
	PutDouble(ctx context.Context, complexBody DoubleWrapper) (*PutDoubleResponse, error)
	PutFloat(ctx context.Context, complexBody FloatWrapper) (*PutFloatResponse, error)
	PutInt(ctx context.Context, complexBody IntWrapper) (*PutIntResponse, error)
	PutLong(ctx context.Context, complexBody LongWrapper) (*PutLongResponse, error)
	PutString(ctx context.Context, complexBody StringWrapper) (*PutStringResponse, error)
}

type primitiveOperations struct {
	*Client
	azinternal.PrimitiveOperations
}

// GetInt ...
func (client *primitiveOperations) GetInt(ctx context.Context) (*GetIntResponse, error) {
	req, err := client.GetIntCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetIntHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutInt ...
func (client *primitiveOperations) PutInt(ctx context.Context, complexBody IntWrapper) (*PutIntResponse, error) {
	req, err := client.PutIntCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.PutIntHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetLong ...
func (client *primitiveOperations) GetLong(ctx context.Context) (*GetLongResponse, error) {
	req, err := client.GetLongCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutLong ...
func (client *primitiveOperations) PutLong(ctx context.Context, complexBody LongWrapper) (*PutLongResponse, error) {
	req, err := client.PutLongCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.PutLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetFloat ...
func (client *primitiveOperations) GetFloat(ctx context.Context) (*GetFloatResponse, error) {
	req, err := client.GetFloatCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutFloat ...
func (client *primitiveOperations) PutFloat(ctx context.Context, complexBody FloatWrapper) (*PutFloatResponse, error) {
	req, err := client.PutFloatCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.PutFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetDouble ...
func (client *primitiveOperations) GetDouble(ctx context.Context) (*GetDoubleResponse, error) {
	req, err := client.GetDoubleCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutDouble ...
func (client *primitiveOperations) PutDouble(ctx context.Context, complexBody DoubleWrapper) (*PutDoubleResponse, error) {
	req, err := client.PutDoubleCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.PutDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetBool ...
func (client *primitiveOperations) GetBool(ctx context.Context) (*GetBoolResponse, error) {
	req, err := client.GetBoolCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutBool ...
func (client *primitiveOperations) PutBool(ctx context.Context, complexBody BooleanWrapper) (*PutBoolResponse, error) {
	req, err := client.PutBoolCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.PutBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetString ...
func (client *primitiveOperations) GetString(ctx context.Context) (*GetStringResponse, error) {
	req, err := client.GetStringCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.GetStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// PutString ...
func (client *primitiveOperations) PutString(ctx context.Context, complexBody StringWrapper) (*PutStringResponse, error) {
	req, err := client.PutStringCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	s, err := client.PutStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// // GetDate ...
// func (client *primitiveOperations) GetDate(ctx context.Context) (*GetDateResponse, error) {
// 	req, err := client.GetDateCreateRequest(*client.u)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp, err := client.p.Do(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	s, err := client.GetDateHandleResponse(resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return s, nil
// }

// // PutDate ...
// func (client *primitiveOperations) PutDate(ctx context.Context, complexBody DateWrapper) (*PutDateResponse, error) {
// 	req, err := client.PutDateCreateRequest(*client.u, complexBody)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp, err := client.p.Do(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	s, err := client.PutDateHandleResponse(resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return s, nil
// }

var _ PrimitiveOperations = (*primitiveOperations)(nil)
