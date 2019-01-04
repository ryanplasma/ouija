package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

// SSHCommand runs a command over ssh
func SSHCommand(user string, password string, host string, port string, app string, command string) string {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	fmt.Println("ssh " + user + "@" + host + ":" + port)

	conn, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		fmt.Printf("Failed to connect to %v\n", host+":"+port)
		fmt.Println(err)
		os.Exit(2)
	}

	session, err := conn.NewSession()
	if err != nil {
		fmt.Printf("Cannot create SSH session to %v\n", host+":"+port)
		fmt.Println(err)
		os.Exit(2)
	}

	defer session.Close()

	// Create buffers for stdout and stderr
	var o, e bytes.Buffer

	session.Stdout = &o
	session.Stderr = &e

	// Run a command with Run and read stdout and stderr
	if err := session.Run("cd ~/" + app + "&& phenv python2.7 /opt/phantom/bin/compile_app.pyc -i"); err != nil {
		fmt.Println("Error running command", err)
	}

	if e.String() != "" {
		return fmt.Sprintf("Error: %s", e.String())
	}

	return fmt.Sprintf(o.String())
}
