package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os/exec"
	"strings"
)

func handler(sess ssh.Session) {
	term := terminal.NewTerminal(sess, "> ")
	for {
		line, err := term.ReadLine()
		if err != nil {
			break
		}
		s := strings.Split(line, " ")
		cmd := exec.Command(s[0], s[1:]...)
		cmd.Stdout = sess
		cmd.Stderr = sess
		err = cmd.Run()
		if err != nil {
			fmt.Fprintln(sess, err)
		}
	}
	log.Println("terminal closed")
}

func main() {
	s := &ssh.Server{
		Addr:    ":2222",
		Handler: handler,
	}

	s.ListenAndServe()
}
