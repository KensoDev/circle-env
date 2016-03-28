package main

import (
	"github.com/codegangsta/cli"
	"github.com/kensodev/circleenv"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "circle"
	app.Usage = "Circle CI commands"
	app.Version = "0.0.3"
	app.Commands = append(app.Commands, circleenv.SyncCommands()...)

	app.Run(os.Args)
}
