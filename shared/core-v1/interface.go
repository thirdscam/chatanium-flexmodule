// Package shared contains shared data between the host and plugins.
package core

type Hook interface {
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
	// Specifies whether the module is ready or not.
	//
	// The expected value is true, which tells the runtime
	// that the module is ready to be served.
	//
	// The runtime also gets the state of the module
	// five times in the first two-second interval.
	//
	// It then tries five more times at 10-second intervals,
	// and if the IsReady field is still not true,
	// it closes the module.
	//
	// So for an initial setup where you expect a long job (>1 minute),
	// just do the basics and set IsReady to true, and create another
	// boolean flag to handle it.
	IsReady bool
}

// Permissions is a list of permissions that a plugin requires.
//
// This is used to determine if a plugin is allowed to run.
// runtime will check if the host has the required permissions
// before starting the plugin.
type Permissions []string
