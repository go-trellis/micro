package main

import (
	"github.com/iTrellis/trellis/cmd"
	"github.com/iTrellis/trellis/examples/components"
	"github.com/iTrellis/trellis/service"

	_ "github.com/iTrellis/trellis/server/api"
)

func main() {
	c := cmd.New()
	if err := c.Init(cmd.ConfigFile("config.yaml")); err != nil {
		panic(err)
	}

	// api.RegistUseFuncs()
	// Explicit to register component function
	cmd.DefaultCompManager.RegisterComponentFunc(
		&service.Service{Name: "component_ping", Version: "v1"},
		components.NewPing)

	if err := c.Start(); err != nil {
		panic(err)
	}

	c.BlockExit()
}