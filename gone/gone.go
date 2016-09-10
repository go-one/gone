package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/go-one/gone/gone/lib"
)


var commands []cli.Command

func RegisterCommands(cs ...cli.Command) {
	commands = append(commands, cs...)
}
func main() {
	app := cli.NewApp()
	app.Version = lib.Version
	app.Name = "gone"
	app.Usage = "Tool for build and run gone based applications"
	app.CommandNotFound = func(c *cli.Context, s string) {
		app.Command("help").Run(c)
	}
	app.Commands = append(app.Commands, commands...)
	app.Run(os.Args)
}
