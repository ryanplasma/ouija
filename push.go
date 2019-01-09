package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"
)

func getPushCommand() cli.Command {
	pushCommand := cli.Command{
		Name:    "push",
		Aliases: []string{"p"},
		Usage:   "Push the app code to the Phantom server",
		Action: func(c *cli.Context) error {
			srcFolder, _ := os.Getwd()
			destFolder := c.String("user") + "@" + c.String("host") + ":~/" + c.String("app") + "/."

			fmt.Println("Running: rsync -av " + srcFolder + "/ " + destFolder)

			cmd := exec.Command("rsync", "-av", srcFolder+"/", destFolder)

			var o bytes.Buffer
			cmd.Stdout = &o

			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(o.String())
			return nil
		},
	}

	pushFlags := []cli.Flag{
		altsrc.NewStringFlag(
			cli.StringFlag{
				Name:   "user",
				Usage:  "Set the SSH User",
				Value:  getUsername(),
				EnvVar: "OUIJA_USER",
			},
		),
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
				Name:   "app",
				Usage:  "Set the App Name",
				Value:  getCurrentFolder(),
				EnvVar: "OUIJA_PORT",
			},
		),
		cli.StringFlag{
			Name:  "load",
			Value: "ouija.yml",
		},
	}

	pushCommand.Before = altsrc.InitInputSourceWithContext(pushFlags, altsrc.NewYamlSourceFromFlagFunc("load"))
	pushCommand.Flags = pushFlags

	return pushCommand
}
