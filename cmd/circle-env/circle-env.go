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
	app.Version = "1.0.0"
	app.Commands = append(app.Commands, circleenv.SyncCommands()...)
	app.Commands = append(app.Commands, circleenv.ReplaceCommands()...)

	app.Run(os.Args)
}
