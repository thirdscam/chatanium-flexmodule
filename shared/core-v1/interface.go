// Package shared contains shared data between the host and plugins.
package core

type Interface interface {
	// GetManifest returns the manifest of the plugin.
	GetManifest() (Manifest, error)

	// GetStatus returns the status of the plugin.
	GetStatus() (Status, error)

	// OnStage is called when the plugin is in a certain stage.
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

// Permissions is a list of permissions that a plugin requires.
//
// This is used to determine if a plugin is allowed to run.
// runtime will check if the host has the required permissions
// before starting the plugin.
type Permissions []string
