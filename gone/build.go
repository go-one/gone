package main

import (
	"path/filepath"

	"github.com/go-one/gone/gone/lib"
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
	appPath := c.Args().First()
	if appPath == "" {
		appPath, _ = filepath.Abs("./")
	} else {
		gopath := lib.GetGOPATH()
		if gopath == "" {
			lib.ErrorLog("GOPATH doesn't set")
			return nil
		}
		appPath = filepath.Join(gopath, "src", appPath)
	}
	app := lib.NewApplication(appPath)
	app.Build()
	return nil
}
