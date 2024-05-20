// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023-2024 Intel Corporation

// Package storage implements the storage related CLI commands
package storage

import (
	"time"

	"github.com/opiproject/godpu/cmd/common"
	"github.com/opiproject/godpu/cmd/storage/backend"
	"github.com/opiproject/godpu/cmd/storage/frontend"
	"github.com/spf13/cobra"
)

// NewStorageCommand tests the storage functionality
func NewStorageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "storage",
		Aliases: []string{"g"},
		Short:   "Tests storage functionality",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	flags := cmd.PersistentFlags()
	flags.Duration(common.TimeoutCmdLineArg, 10*time.Second, "timeout for a cmd")

	cmd.AddCommand(newStorageCreateCommand())
	cmd.AddCommand(newStorageDeleteCommand())
	cmd.AddCommand(newStorageGetCommand())
	cmd.AddCommand(newStorageTestCommand())

	return cmd
}

func newStorageCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Creates resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(frontend.NewCreateCommand())
	cmd.AddCommand(backend.NewCreateCommand())

	return cmd
}

func newStorageDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"d"},
		Short:   "Deletes resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(frontend.NewDeleteCommand())
	cmd.AddCommand(backend.NewDeleteCommand())

	return cmd
}

func newStorageGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "Gets resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(backend.NewGetCommand())

	return cmd
}
