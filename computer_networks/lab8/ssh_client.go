package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pkg/sftp"
	"github.com/skorobogatov/input"
	"golang.org/x/crypto/ssh"
)

func connect(socket string, config *ssh.ClientConfig) (conn *ssh.Client, err error) {
	conn, err = ssh.Dial("tcp", socket, config)
	if err != nil {
		return
	}
	return
}

func main() {
	var conn *ssh.Client
	config := &ssh.ClientConfig{
		User:            "aleks",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{ssh.Password("12345")},
	}
	conn, err := connect("127.0.0.1:8081", config)

	// open an SFTP session over an existing ssh connection.
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// walk a directory
	w := client.Walk("./")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}

	fmt.Println("filename: ")
	fname := input.Gets()

	data, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
		return
	}

	// leave your mark
	f, err := client.Create("./files/" + fname)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(data); err != nil {
		log.Fatal(err)
	}
	f.Close()

}
