// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package backend implements the go library for OPI backend storage
package backend

import (
	grpcOpi "github.com/opiproject/godpu/grpc"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/grpc"
)

// CreateNvmeClient defines the function type used to retrieve NvmeRemoteControllerServiceClient
type CreateNvmeClient func(cc grpc.ClientConnInterface) pb.NvmeRemoteControllerServiceClient

// Client is used for managing storage devices on OPI server
type Client struct {
	connector        grpcOpi.Connector
	createNvmeClient CreateNvmeClient
}

// New creates a new instance of Client
func New(addr string, tls string) (*Client, error) {
	connector, err := grpcOpi.New(addr, tls)
	if err != nil {
		return nil, err
	}

	return NewWithArgs(
		connector,
		pb.NewNvmeRemoteControllerServiceClient,
	)
}

// NewWithArgs creates a new instance of Client with non-default members
func NewWithArgs(
	connector grpcOpi.Connector,
	createNvmeClient CreateNvmeClient,
) (*Client, error) {
	return &Client{
		connector:        connector,
		createNvmeClient: createNvmeClient,
	}, nil
}
