package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/mmcdole/gofeed"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var addr = flag.String("addr", "151.248.113.144:8021", "http service address")

func clear() {
	db, _ := sql.Open("mysql", login+":"+password+"@tcp("+host+")/"+dbname+"?parseTime=true")
	db.Query("TRUNCATE iu9networkslabs.iu9alexeev")
}

func save() {
	db, err := sql.Open("mysql", login+":"+password+"@tcp("+host+")/"+dbname)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://news.rambler.ru/rss/technology/")
	if err != nil {
		panic(err)
	}
	for _, item := range feed.Items {
		var categories string
		for _, category := range item.Categories {
			categories += category
		}
		var authors string
		for _, author := range item.Authors {
			authors += author.Name + " "
		}
		rows, err := db.Query("select * from iu9networkslabs.iu9alexeev where title = ? and link = ? and description = ?"+
			" and pubdate = ? and categories = ? and authors = ?",
			item.Title, item.Link, item.Description, item.Published, categories, authors)
		if err != nil {
			panic(err)
		}
		if !rows.Next() {
			_, err = db.Exec("insert into iu9networkslabs.iu9alexeev (title, link, pubdate, description, categories, authors)"+
				"values(?, ?, ?, ?, ?, ?)", item.Title, item.Link, item.Published, item.Description, categories, authors)
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/task2_1", task2_1)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func task2_1(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("recv: %s", message)
		msg := string(message)
		switch msg {
		case "saveNews":
			save()
		case "clear":
			clear()
		}
		log.Println(msg)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
