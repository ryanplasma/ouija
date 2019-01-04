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

func getDownloadCommand() cli.Command {
	downloadCommand := cli.Command{
		Name:    "download",
		Aliases: []string{"d"},
		Usage:   "Download the app tarball from the Phantom server",
		Action: func(c *cli.Context) error {
			destFolder, _ := os.Getwd()
			srcFolder := c.String("user") + "@" + c.String("host") + ":~/" + c.String("app") + ".tgz"

			cmd := exec.Command("scp", srcFolder, destFolder)

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

	downloadFlags := []cli.Flag{
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

	downloadCommand.Before = altsrc.InitInputSourceWithContext(downloadFlags, altsrc.NewYamlSourceFromFlagFunc("load"))
	downloadCommand.Flags = downloadFlags

	return downloadCommand
}
