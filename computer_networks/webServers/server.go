package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var msgType = [4]string{"register", "data", "token", "error"}
var token = ""

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func findUser(login string) map[string]string {
	file, _ := ioutil.ReadFile("users.txt")
	users := make([]map[string]string, 1000)
	json.Unmarshal(file, &users)
	for _, user := range users {
		if user["login"] == login {
			return user
		}
	}
	return nil
}

func reg(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		data := make(map[string]string)
		json.Unmarshal(message, &data)
		fmt.Println(data)
		if data["msgType"] == msgType[0] {
			usr := map[string]string{"login": data["login"], "password": data["password"]}
			file, _ := ioutil.ReadFile("users.txt")
			users := make([]map[string]string, 0)
			json.Unmarshal(file, &users)
			log.Print("11111 ")
			log.Println(users)
			users = append(users, usr)
			users1, _ := json.Marshal(users)
			err = ioutil.WriteFile("users.txt", users1, 077)
			if err != nil {
				return
			}
			err = c.WriteMessage(mt, []byte("user "+data["login"]+" was successfully registered"))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}

}

func log_in(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		data := make(map[string]string)
		json.Unmarshal(message, &data)
		fmt.Println(data)
		if data["msgType"] == msgType[0] {
			user := findUser(data["login"])
			if user != nil && CheckPasswordHash(data["password"], user["password"]) {
				btoken, _ := bcrypt.GenerateFromPassword([]byte(time.Now().String()), 14)
				token = string(btoken)
				log.Println(token)
				msg := map[string]string{"msgType": msgType[2], "token": token}
				pkg, _ := json.Marshal(msg)
				err = c.WriteMessage(mt, pkg)
				if err != nil {
					log.Println("write:", err)
					break
				}
			} else {
				msg := map[string]string{"msgType": msgType[3], "message": "Wrong login or password!"}
				pkg, _ := json.Marshal(msg)
				err = c.WriteMessage(mt, pkg)
				if err != nil {
					log.Println("write:", err)
					break
				}
			}
		}
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(message)
		var msg string
		if len(message) == 0 || message[0] != 34 {
			message = []byte("\"" + string(message) + "\"")
		}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println(err)
			break
		}
		command := strings.Split(string(msg), " ")
		cmd := exec.Command(command[0], command[1:]...)
		log.Printf("recv: %s", message)
		output, err := cmd.Output()
		if err != nil {
			log.Println(err)
			c.WriteMessage(mt, []byte("Incorrect command"))
			continue
		}
		err = c.WriteMessage(mt, output)
		if err != nil {
			log.Println("write:", err)
			break
			//}
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/register", reg)
	http.HandleFunc("/login", log_in)
	http.HandleFunc("/console", console)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func console(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		msg := make(map[string]string)
		json.Unmarshal(message, &msg)
		if msg["msgType"] == msgType[2] && msg["token"] == token {
			command := msg["command"]
			comm := strings.Split(command, " ")
			cmd := exec.Command(comm[0], comm[1:]...)
			b := make([]byte, 0)
			buf := bytes.NewBuffer(b)
			cmd.Stdout = buf
			err := cmd.Run()
			if err != nil {
				c.WriteMessage(mt, []byte(err.Error()))
			} else {
				err := c.WriteMessage(mt, buf.Bytes())
				if err != nil {
					log.Println("write:", err)
					break
				}
			}
		} else {
			c.WriteMessage(mt, []byte("Tokens dont match!"))
		}
	}
}
