package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var addr = flag.String("addr", "localhost:8030", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getOnlineUsers() int {
	config := &ssh.ClientConfig{
		User: "iu9lab",
		Auth: []ssh.AuthMethod{
			ssh.Password(""),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "151.248.113.144:443", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	session.Stderr = os.Stdout
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	in, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}
	err = session.Shell()
	if err != nil {
		panic(err)
	}
	line := "w -h -f -s"
	str := ""
	r := bytes.NewBufferString(str)
	session.Stdout = r
	fmt.Fprintf(in, "%s\n", line)
	time.Sleep(3 * time.Second)
	fmt.Fprintf(in, "%s\n", line)
	str = r.String()
	splited := strings.Split(str, "\n")
	fmt.Println(splited)
	firstword := make([]string, 0)
	for _, split := range splited {
		fmt.Println(split)
		words := strings.Split(split, " ")
		firstword = append(firstword, words[0])
	}
	n := 0
	for _, word := range firstword {
		if word == "iu9lab" {
			n++
		}
	}
	return n
}
func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/task3", task3)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func task3(w http.ResponseWriter, r *http.Request) {
	c, _ := upgrader.Upgrade(w, r, nil)
	//n := 0
	for {
		cmd := exec.Command("ps", "-ax")
		//n = getOnlineUsers()
		//switch {
		//case n == 0:
		//	c.WriteMessage(websocket.TextMessage, []byte("IU9 gone"))
		//case n == 1:
		//	c.WriteMessage(websocket.TextMessage, []byte("IU9 online"))
		//case n > 1:
		//	c.WriteMessage(websocket.TextMessage, []byte("IU9 online\n"+strconv.Itoa(n)))
		//}
		output, _ := cmd.Output()
		c.WriteMessage(websocket.TextMessage, output)
		time.Sleep(time.Second)
	}

}
