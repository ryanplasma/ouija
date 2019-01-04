package main

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"
)

func getBuildCommand() cli.Command {
	buildCommand := cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the app code on the Phantom server",
		Action: func(c *cli.Context) error {
			output := SSHCommand(
				c.String("user"),
				c.String("password"),
				c.String("host"),
				c.String("port"),
				c.String("app"),
				getPhantomBuildCommandString(c.String("app")),
			)

			fmt.Println(output)

			return nil
		},
	}

	buildFlags := []cli.Flag{
		altsrc.NewStringFlag(
			cli.StringFlag{
				Name:   "user",
				Usage:  "Set the SSH User",
				Value:  getUsername(),
				EnvVar: "OUIJA_USER",
			},
		),
		cli.StringFlag{
			Name:   "password",
			Usage:  "Set the SSH Password",
			EnvVar: "OUIJA_PASSWORD",
		},
		altsrc.NewStringFlag(
			cli.StringFlag{
				Name:   "host",
				Usage:  "Set the SSH host",
				Value:  "127.0.0.1",
				EnvVar: "OUIJA_HOST",
			},
		),
		altsrc.NewStringFlag(
			cli.StringFlag{
				Name:   "port",
				Usage:  "Set the SSH port",
				Value:  "22",
				EnvVar: "OUIJA_PORT",
			},
		),
		altsrc.NewStringFlag(
			cli.StringFlag{
				Name:   "app",
				Usage:  "Set the App Name",
				Value:  getCurrentFolder(),
				EnvVar: "OUIJA_APP",
			},
		),
		cli.StringFlag{
			Name:  "load",
			Value: "ouija.yml",
		},
	}

	buildCommand.Before = altsrc.InitInputSourceWithContext(buildFlags, altsrc.NewYamlSourceFromFlagFunc("load"))
	buildCommand.Flags = buildFlags

	return buildCommand
}
