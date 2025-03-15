// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/thirdscam/chatanium-flexmodule/sample/shared"
	"github.com/thirdscam/chatanium-flexmodule/sample/shared/core-v1"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(io.Discard)

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("PLUGIN_PATH")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC,
		},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("core-v1")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// We should have a Counter store now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	core, ok := raw.(core.Interface)
	if !ok {
		fmt.Println("Plugin has no 'core-v1' plugin symbol")
		os.Exit(1)
	}

	manifest, err := core.GetManifest()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	fmt.Printf("MODULE_MANIFEST >> %+v\n", manifest)

	status, err := core.GetStatus()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	fmt.Printf("MODULE_STATUS >> %+v\n", status)

	core.OnStage("RUNTIME_STARTED")
	fmt.Printf("MODULE_STAGE_DISPATCHED >> RUNTIME_STARTED\n")

	status, err = core.GetStatus()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	fmt.Printf("MODULE_STATUS >> %+v\n", status)
}
