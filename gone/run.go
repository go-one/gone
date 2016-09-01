package main

import (
	"github.com/go-one/gone-framework/gone/lib"
	"github.com/urfave/cli"
)

var runFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "config,c",
		Usage: "Path of config",
		Value: "./build.conf",
	},
}

func init() {
	RegisterCommands(cli.Command{
		Name:    "run",
		Usage:   "Run application",
		Aliases: []string{"r"},
		Action:  RunCommandHandler,
		Flags:   runFlags,
		Before:  lib.ShowBanner,
	})
}

// Main docs
func RunCommandHandler(c *cli.Context) error {
	configPath := c.String("config")
	lib.InfoLog("\nRunning application:")
	lib.IncrLogOffset()
	lib.InfoLog("Using %s file as config", configPath)
	builder := lib.NewBuilder(configPath)
	err := builder.Build()
	return err
}
