package circleenv

import (
	"github.com/codegangsta/cli"
	"os"
)

func ReplaceCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "replace",
			Aliases: []string{"p"},
			Usage:   "Replace ENV variables in files",
			Action: func(c *cli.Context) {
				cwd, err := os.Getwd()
				if err != nil {
					panic(err)
				}
				NewFileFinder(cwd).FindAndReplaceTemplateFiles()
			},
			Flags: []cli.Flag{},
		},
	}
}
