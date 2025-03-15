// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Package shared contains shared data between the host and plugins.
package core

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

// KV is the interface that we're exposing as a plugin.
type ICore interface {
	GetManifest() (Manifest, error)
	GetStatus() (Status, error)
	OnStage(stage string)
}
