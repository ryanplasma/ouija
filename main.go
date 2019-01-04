package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "ouija"
	app.Usage = "Build Tools for Splunk Phantom App Developement"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ryan Plas",
			Email: "ryan@wordplas.com",
		},
	}

	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, getPushCommand())
	app.Commands = append(app.Commands, getBuildCommand())
	app.Commands = append(app.Commands, getDownloadCommand())

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
