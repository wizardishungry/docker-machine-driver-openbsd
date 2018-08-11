package ssh

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

// Dial utility function for getting an SSH client
func Dial(host string, port string, username string, password string) (*ssh.Client, error) {
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		Timeout: 1 * time.Second,
		User:    username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return ssh.Dial("tcp", host+":"+port, config)
}

// CopyID copy ssh key to host
func CopyID(client *ssh.Client, key []byte) error {
	username := client.User()
	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	cmd := `mkdir -p ~` + username + `/.ssh ; echo "` + string(key) + `" >>~` + username + `/.ssh/authorized_keys`
	fmt.Println(cmd)
	if err := session.Run(cmd); err != nil {
		return errors.New("Failed to run: " + err.Error())
	}
	return nil
}
