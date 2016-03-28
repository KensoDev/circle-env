package circleenv

import (
	"github.com/codegangsta/cli"
)

func SyncCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "sync",
			Aliases: []string{"p"},
			Usage:   "Sync Environment Variables to CircleCI",
			Action: func(c *cli.Context) {
				configuration := &Configuration{
					UserName:    c.String("username"),
					Token:       c.String("token"),
					ProjectName: c.String("project-name"),
					FileName:    c.String("filename"),
				}
				sender := NewSender(configuration)
				sender.Send()
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "project-name",
					Usage: "CircleCI project name",
				},
				cli.StringFlag{
					Name:  "username",
					Usage: "CircleCI username (Usually your username or Github org)",
				},
				cli.StringFlag{
					Name:  "token",
					Usage: "CircleCI token",
				},
				cli.StringFlag{
					Name:  "filename",
					Usage: "Environment variable filename",
					Value: ".env",
				},
			},
		},
	}
}
