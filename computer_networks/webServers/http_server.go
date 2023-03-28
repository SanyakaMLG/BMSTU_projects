package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
)

var addr1 = flag.String("addr", "localhost:8080", "http service address")
var messageType = [3]string{"register", "data", "token"}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func postHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		err := r.ParseForm()
		if err != nil {
			return
		}
		pkg, err := json.Marshal(r.Form.Get("arg"))
		u := url.URL{Scheme: "ws", Host: *addr1, Path: "/echo"}
		log.Printf("connecting to %s", u.String())
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Fatal("dial:", err)
		}
		defer c.Close()
		err = c.WriteMessage(websocket.TextMessage, pkg)
		if err != nil {
			log.Println("write:", err)
			return
		}
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		tml := template.Must(template.ParseFiles("form.html"))
		tml.Execute(w, string(message))
	}

}
func syncHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "form2.html")
}

func register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "register.html")
	case "POST":
		err := r.ParseForm()
		if err != nil {
			return
		}
		login := r.Form["login"][0]
		password := r.Form["password"][0]
		Hpass, _ := HashPassword(password)
		u := url.URL{Scheme: "ws", Host: *addr1, Path: "/register"}
		log.Printf("connecting to %s", u.String())
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Fatal("dial:", err)
		}
		pkg := map[string]string{"msgType": messageType[0], "login": login, "password": Hpass}
		msg, _ := json.Marshal(pkg)

		defer c.Close()
		err = c.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("write:", err)
			return
		}
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		tml := template.Must(template.ParseFiles("form.html"))
		tml.Execute(w, string(message))
	}
}

func main() {

	http.HandleFunc("/part1", postHandle)
	http.HandleFunc("/part2", syncHandle)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/console", cons)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func cons(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "console.html")
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}
