package main

import (
	"github.com/jlaffaye/ftp"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"time"
)

const (
	RSS  = "https://news.rambler.ru/rss/technology/"
)

func main() {
	name := "Alexander_Alexeev_" + time.Now().Format("2006-01-02_15-04-05") + ".txt"
	log.Printf("Parsing %s...\n", RSS)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(RSS)
	if err != nil {
		log.Fatal("Failed to parse URL: ", err)
	}
	log.Printf("Writing to %s...\n", name)
	file, err := os.Create(name)
	if err != nil {
		log.Fatal("Failed to create a file: ", err)
	}
	defer file.Close()
	if _, err := file.WriteString("### " + feed.Description + "\n"); err != nil {
		log.Fatal("Failed to write the description: ", err)
	}
	for _, item := range feed.Items {
		if _, err := file.WriteString("- " + item.Title + ";\n" + item.Description + "\n" +
			item.Link + "\n" + item.Published + "\n\n\n"); err != nil {
			log.Fatal("Failed to write a title: ", err)
		}
	}
	log.Printf("Connecting to %s FTP-server...\n", ADDR)
	conn, err := ftp.Dial(ADDR, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	err = conn.Login(USR, PASS)
	if err != nil {
		log.Fatal("Failed to login: ", err)
	}
	readFile, err := os.Open(name)
	if err != nil {
		log.Fatal("Failed to open file: ", err)
	}
	defer readFile.Close()
	log.Printf("Storing %s...\n", name)
	if err = conn.Stor(name, readFile); err != nil {
		log.Fatal("Failed to issue STOR FTP command: ", err)
	}
	if err := conn.Quit(); err != nil {
		log.Fatal("Failed to quit: ", err)
	}
	log.Println("Completed successfully.")
}
