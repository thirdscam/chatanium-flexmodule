// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Package shared contains shared data between the host and plugins.
package core

type Interface interface {
	GetManifest() (Manifest, error)
	GetStatus() (Status, error)
	OnStage(stage string)
}

type Manifest struct {
	Name        string
	Version     string
	Author      string
	Repository  string
	Permissions Permissions
}

type Status struct {
	IsReady bool
}

type Permissions []string
