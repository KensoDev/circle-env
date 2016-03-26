package main

import (
	"github.com/kensodev/circleenv"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	filename    = kingpin.Flag("filename", ".env file location").Default(".env").String()
	projectName = kingpin.Flag("project-name", "Cirle project name").Required().String()
	token       = kingpin.Flag("token", "Circle API token").Required().String()
	username    = kingpin.Flag("username", "Circle User Name").Required().String()
)

func main() {
	kingpin.Parse()

	configuration := circleenv.Configuration{
		UserName:    *username,
		Token:       *token,
		ProjectName: *projectName,
		FileName:    *filename,
	}

	sender := circleenv.NewSender(configuration)
	sender.Send()
}
