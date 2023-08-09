// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (C) 2023 Intel Corporation

// Package cmd implements the CLI commands
package cmd

import (
	"context"
	"log"
	"time"

	"github.com/opiproject/godpu/grpc"

	"github.com/opiproject/godpu/storage"
	"github.com/spf13/cobra"
)

// NewStorageTestCommand returns the storage tests command
func NewStorageTestCommand() *cobra.Command {
	var (
		addr string
	)
	cmd := &cobra.Command{
		Use:     "test",
		Aliases: []string{"s"},
		Short:   "Test storage functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// Set up a connection to the server.
			client, err := grpc.New(addr)
			if err != nil {
				log.Fatalf("error creating new client: %v", err)
			}

			// Contact the server and print out its response.
			conn, closer, err := client.NewConn()
			if err != nil {
				log.Fatalf("error creating gRPC connection: %v", err)
			}
			defer closer()

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			log.Printf("==============================================================================")
			log.Printf("Test frontend")
			log.Printf("==============================================================================")
			err = storage.DoFrontend(ctx, conn)
			if err != nil {
				log.Panicf("DoFrontend tests failed with error: %v", err)
			}

			log.Printf("==============================================================================")
			log.Printf("Test backend")
			log.Printf("==============================================================================")
			err = storage.DoBackend(ctx, conn)
			if err != nil {
				log.Panicf("DoBackend tests failed with error: %v", err)
			}

			log.Printf("==============================================================================")
			log.Printf("Test middleend")
			log.Printf("==============================================================================")
			err = storage.DoMiddleend(ctx, conn)
			if err != nil {
				log.Panicf("DoMiddleend tests failed with error: %v", err)
			}
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&addr, "addr", "localhost:50151", "address or OPI gRPC server")
	return cmd
}

// NewStorageCommand tests the storage functionality
func NewStorageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "storage",
		Aliases: []string{"g"},
		Short:   "Tests storage functionality",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("[ERROR] %s", err.Error())
			}
		},
	}

	cmd.AddCommand(NewStorageTestCommand())
	return cmd
}
