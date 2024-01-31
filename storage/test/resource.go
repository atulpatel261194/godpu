// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023-2024 Intel Corporation

// Package test implements compliance storage tests
package test

import "go.einride.tech/aip/resourcename"

func resourceIDToVolumeName(resourceID string) string {
	return resourcename.Join(
		"volumes", resourceID,
	)
}

func resourceIDToSubsystemName(resourceID string) string {
	return resourcename.Join(
		"subsystems", resourceID,
	)
}

func resourceIDToNamespaceName(subsysResourceID, ctrlrResourceID string) string {
	return resourcename.Join(
		"subsystems", subsysResourceID,
		"namespaces", ctrlrResourceID,
	)
}

func resourceIDToControllerName(subsysResourceID, ctrlrResourceID string) string {
	return resourcename.Join(
		"subsystems", subsysResourceID,
		"controllers", ctrlrResourceID,
	)
}

func resourceIDToRemoteControllerName(resourceID string) string {
	return resourcename.Join(
		"nvmeRemoteControllers", resourceID,
	)
}

func resourceIDToNvmePathName(ctrlrResourceID, pathResourceID string) string {
	return resourcename.Join(
		"nvmeRemoteControllers", ctrlrResourceID,
		"nvmePaths", pathResourceID,
	)
}
