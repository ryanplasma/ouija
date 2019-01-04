package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func getUsername() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return user.Name
}

func getCurrentFolder() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	baseFolder := filepath.Base(wd)
	return baseFolder
}

func getPhantomBuildCommandString(app string) string {
	return fmt.Sprintf("cd ~/%s && phenv python2.7 /opt/phantom/bin/compile_app.pyc -i", app)
}
