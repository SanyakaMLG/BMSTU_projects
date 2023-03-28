package main

import (
	"bufio"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func stor(client *ftp.ServerConn, filename string) {
	f, err := os.Open("files/" + filename)
	if err != nil {
		log.Println("Can't open file " + filename)
		return
	}
	defer f.Close()

	err = client.Stor(filename, f)
	if err != nil {
		log.Println("Failed to store file on server")
	}
}

func retr(client *ftp.ServerConn, filename string) {
	r, err := client.Retr(filename)
	if err != nil {
		log.Println("Failed to issue RETR FTP command")
		return
	}
	defer r.Close()

	file, err := os.Create("store/" + filename)
	if err != nil {
		log.Println("Can't create file " + filename)
		return
	}
	if _, err = io.Copy(file, r); err != nil {
		log.Println("Failed to copy")
	}
}

func mkdir(client *ftp.ServerConn, path string) {
	err := client.MakeDir(path)
	if err != nil {
		log.Println("Failed to create directory " + path)
	}
}

func del(client *ftp.ServerConn, filename string) {
	err := client.Delete(filename)
	if err != nil {
		log.Println("Failed to delete file " + filename)
	}
}

func nlst(client *ftp.ServerConn, paths ...string) {
	var path string
	if len(paths) > 0 {
		path = paths[0]
	} else {
		path, _ = client.CurrentDir()
	}
	names, err := client.NameList(path)
	if err != nil {
		log.Println(err)
		return
	}
	for _, name := range names {
		log.Println(name)
	}
}

func quit(client *ftp.ServerConn) {
	if err := client.Quit(); err != nil {
		log.Println(err)
	}
}

func main() {
	c, err := ftp.Dial(host+":ftp", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(user, password)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("ftp>")
	for scanner.Scan() {
		cmd := scanner.Text()
		split := strings.Split(cmd, " ")
		switch split[0] {
		case "STOR":
			if len(split) == 2 {
				stor(c, split[1])
			} else {
				log.Println("wrong format, use: STOR <filename>")
			}
		case "RETR":
			if len(split) == 2 {
				retr(c, split[1])
			} else {
				log.Println("wrong format, use: RETR <filename>")
			}
		case "MKD":
			if len(split) == 2 {
				mkdir(c, split[1])
			} else {
				log.Println("wrong format, use: MKD <path>")
			}
		case "DEL":
			if len(split) == 2 {
				del(c, split[1])
			} else {
				log.Println("wrong format, use: DEL <filename>")
			}
		case "NLST":
			if len(split) == 1 {
				nlst(c)
			} else if len(split) == 2 {
				nlst(c, split[1])
			} else {
				log.Println("wrong format, use: NLST [path]")
			}
		case "QUIT":
			quit(c)
			return
		default:
			log.Println("unknown command")
		}
		fmt.Print("ftp>")
	}

}
