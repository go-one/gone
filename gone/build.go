package main

import (
	"github.com/go-one/gone-framework/gone/lib"
	"github.com/urfave/cli"
)

var buildFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "config,c",
		Usage: "Path of config",
		Value: "./build.conf",
	},
}

func init() {
	RegisterCommands(cli.Command{
		Name:    "build",
		Usage:   "Build application",
		Aliases: []string{"b"},
		Action:  BuildCommandHandler,
		Flags:   buildFlags,
		Before:  lib.ShowBanner,
	})
}

// Main docs
func BuildCommandHandler(c *cli.Context) error {
	configPath := c.String("config")
	lib.InfoLog("Using %s file as config", configPath)
	builder := lib.NewBuilder(configPath)
	err := builder.Build()
	return err
}
