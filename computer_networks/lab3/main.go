package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mmcdole/gofeed"
)

func main() {
	db, err := sql.Open("mysql", login+":"+password+"@tcp("+hostname+")/"+dbname)
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
