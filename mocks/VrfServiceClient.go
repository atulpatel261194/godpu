/* SPDX-License-Identifier: Apache-2.0
   Copyright (c) 2023 Dell Inc, or its subsidiaries.
*/
// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	_go "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// VrfServiceClient is an autogenerated mock type for the VrfServiceClient type
type VrfServiceClient struct {
	mock.Mock
}

type VrfServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *VrfServiceClient) EXPECT() *VrfServiceClient_Expecter {
	return &VrfServiceClient_Expecter{mock: &_m.Mock}
}

// CreateVrf provides a mock function with given fields: ctx, in, opts
func (_m *VrfServiceClient) CreateVrf(ctx context.Context, in *_go.CreateVrfRequest, opts ...grpc.CallOption) (*_go.Vrf, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.Vrf
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.CreateVrfRequest, ...grpc.CallOption) (*_go.Vrf, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.CreateVrfRequest, ...grpc.CallOption) *_go.Vrf); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.Vrf)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.CreateVrfRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VrfServiceClient_CreateVrf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateVrf'
type VrfServiceClient_CreateVrf_Call struct {
	*mock.Call
}

// CreateVrf is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.CreateVrfRequest
//   - opts ...grpc.CallOption
func (_e *VrfServiceClient_Expecter) CreateVrf(ctx interface{}, in interface{}, opts ...interface{}) *VrfServiceClient_CreateVrf_Call {
	return &VrfServiceClient_CreateVrf_Call{Call: _e.mock.On("CreateVrf",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *VrfServiceClient_CreateVrf_Call) Run(run func(ctx context.Context, in *_go.CreateVrfRequest, opts ...grpc.CallOption)) *VrfServiceClient_CreateVrf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.CreateVrfRequest), variadicArgs...)
	})
	return _c
}

func (_c *VrfServiceClient_CreateVrf_Call) Return(_a0 *_go.Vrf, _a1 error) *VrfServiceClient_CreateVrf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VrfServiceClient_CreateVrf_Call) RunAndReturn(run func(context.Context, *_go.CreateVrfRequest, ...grpc.CallOption) (*_go.Vrf, error)) *VrfServiceClient_CreateVrf_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteVrf provides a mock function with given fields: ctx, in, opts
func (_m *VrfServiceClient) DeleteVrf(ctx context.Context, in *_go.DeleteVrfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *emptypb.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.DeleteVrfRequest, ...grpc.CallOption) (*emptypb.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.DeleteVrfRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.DeleteVrfRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VrfServiceClient_DeleteVrf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteVrf'
type VrfServiceClient_DeleteVrf_Call struct {
	*mock.Call
}

// DeleteVrf is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.DeleteVrfRequest
//   - opts ...grpc.CallOption
func (_e *VrfServiceClient_Expecter) DeleteVrf(ctx interface{}, in interface{}, opts ...interface{}) *VrfServiceClient_DeleteVrf_Call {
	return &VrfServiceClient_DeleteVrf_Call{Call: _e.mock.On("DeleteVrf",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *VrfServiceClient_DeleteVrf_Call) Run(run func(ctx context.Context, in *_go.DeleteVrfRequest, opts ...grpc.CallOption)) *VrfServiceClient_DeleteVrf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.DeleteVrfRequest), variadicArgs...)
	})
	return _c
}

func (_c *VrfServiceClient_DeleteVrf_Call) Return(_a0 *emptypb.Empty, _a1 error) *VrfServiceClient_DeleteVrf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VrfServiceClient_DeleteVrf_Call) RunAndReturn(run func(context.Context, *_go.DeleteVrfRequest, ...grpc.CallOption) (*emptypb.Empty, error)) *VrfServiceClient_DeleteVrf_Call {
	_c.Call.Return(run)
	return _c
}

// GetVrf provides a mock function with given fields: ctx, in, opts
func (_m *VrfServiceClient) GetVrf(ctx context.Context, in *_go.GetVrfRequest, opts ...grpc.CallOption) (*_go.Vrf, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.Vrf
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.GetVrfRequest, ...grpc.CallOption) (*_go.Vrf, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.GetVrfRequest, ...grpc.CallOption) *_go.Vrf); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.Vrf)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.GetVrfRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VrfServiceClient_GetVrf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVrf'
type VrfServiceClient_GetVrf_Call struct {
	*mock.Call
}

// GetVrf is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.GetVrfRequest
//   - opts ...grpc.CallOption
func (_e *VrfServiceClient_Expecter) GetVrf(ctx interface{}, in interface{}, opts ...interface{}) *VrfServiceClient_GetVrf_Call {
	return &VrfServiceClient_GetVrf_Call{Call: _e.mock.On("GetVrf",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *VrfServiceClient_GetVrf_Call) Run(run func(ctx context.Context, in *_go.GetVrfRequest, opts ...grpc.CallOption)) *VrfServiceClient_GetVrf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.GetVrfRequest), variadicArgs...)
	})
	return _c
}

func (_c *VrfServiceClient_GetVrf_Call) Return(_a0 *_go.Vrf, _a1 error) *VrfServiceClient_GetVrf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VrfServiceClient_GetVrf_Call) RunAndReturn(run func(context.Context, *_go.GetVrfRequest, ...grpc.CallOption) (*_go.Vrf, error)) *VrfServiceClient_GetVrf_Call {
	_c.Call.Return(run)
	return _c
}

// ListVrfs provides a mock function with given fields: ctx, in, opts
func (_m *VrfServiceClient) ListVrfs(ctx context.Context, in *_go.ListVrfsRequest, opts ...grpc.CallOption) (*_go.ListVrfsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.ListVrfsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.ListVrfsRequest, ...grpc.CallOption) (*_go.ListVrfsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.ListVrfsRequest, ...grpc.CallOption) *_go.ListVrfsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.ListVrfsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.ListVrfsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VrfServiceClient_ListVrfs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListVrfs'
type VrfServiceClient_ListVrfs_Call struct {
	*mock.Call
}

// ListVrfs is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.ListVrfsRequest
//   - opts ...grpc.CallOption
func (_e *VrfServiceClient_Expecter) ListVrfs(ctx interface{}, in interface{}, opts ...interface{}) *VrfServiceClient_ListVrfs_Call {
	return &VrfServiceClient_ListVrfs_Call{Call: _e.mock.On("ListVrfs",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *VrfServiceClient_ListVrfs_Call) Run(run func(ctx context.Context, in *_go.ListVrfsRequest, opts ...grpc.CallOption)) *VrfServiceClient_ListVrfs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.ListVrfsRequest), variadicArgs...)
	})
	return _c
}

func (_c *VrfServiceClient_ListVrfs_Call) Return(_a0 *_go.ListVrfsResponse, _a1 error) *VrfServiceClient_ListVrfs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VrfServiceClient_ListVrfs_Call) RunAndReturn(run func(context.Context, *_go.ListVrfsRequest, ...grpc.CallOption) (*_go.ListVrfsResponse, error)) *VrfServiceClient_ListVrfs_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateVrf provides a mock function with given fields: ctx, in, opts
func (_m *VrfServiceClient) UpdateVrf(ctx context.Context, in *_go.UpdateVrfRequest, opts ...grpc.CallOption) (*_go.Vrf, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.Vrf
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.UpdateVrfRequest, ...grpc.CallOption) (*_go.Vrf, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.UpdateVrfRequest, ...grpc.CallOption) *_go.Vrf); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.Vrf)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.UpdateVrfRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VrfServiceClient_UpdateVrf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateVrf'
type VrfServiceClient_UpdateVrf_Call struct {
	*mock.Call
}

// UpdateVrf is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.UpdateVrfRequest
//   - opts ...grpc.CallOption
func (_e *VrfServiceClient_Expecter) UpdateVrf(ctx interface{}, in interface{}, opts ...interface{}) *VrfServiceClient_UpdateVrf_Call {
	return &VrfServiceClient_UpdateVrf_Call{Call: _e.mock.On("UpdateVrf",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *VrfServiceClient_UpdateVrf_Call) Run(run func(ctx context.Context, in *_go.UpdateVrfRequest, opts ...grpc.CallOption)) *VrfServiceClient_UpdateVrf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.UpdateVrfRequest), variadicArgs...)
	})
	return _c
}

func (_c *VrfServiceClient_UpdateVrf_Call) Return(_a0 *_go.Vrf, _a1 error) *VrfServiceClient_UpdateVrf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VrfServiceClient_UpdateVrf_Call) RunAndReturn(run func(context.Context, *_go.UpdateVrfRequest, ...grpc.CallOption) (*_go.Vrf, error)) *VrfServiceClient_UpdateVrf_Call {
	_c.Call.Return(run)
	return _c
}

// NewVrfServiceClient creates a new instance of VrfServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVrfServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *VrfServiceClient {
	mock := &VrfServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}