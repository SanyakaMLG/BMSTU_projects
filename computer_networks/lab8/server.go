package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
)

type Data struct {
	Files []File
}

type File struct {
	Name string
}

func handleMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		files, err := ioutil.ReadDir("files")
		if err != nil {
			fmt.Println(err)
		}
		var data Data
		for _, file := range files {
			if filepath.Ext(file.Name()) != ".exe" {
				data.Files = append(data.Files, File{Name: file.Name()})
			}
		}
		t, _ := template.ParseFiles("templates/index.html")
		if err := t.Execute(w, data); err != nil {
			fmt.Println(err)
		}
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		filename := r.FormValue("file")
		array := r.FormValue("array")
		ext := filepath.Ext(filename)
		if ext == ".c" {
			if len(array) > 0 {
				exec.Command("gcc", "-o", "files/"+filename+".exe", "files/"+filename).Run()
				out, err := exec.Command("files/"+filename+".exe", array).Output()
				if err != nil {
					fmt.Println(err)
				}
				w.Write(out)
			} else {
				w.Write([]byte("Enter array"))
			}
		} else {
			message, err := ioutil.ReadFile("files/" + filename)
			if err != nil {
				log.Fatalln(err)
			}
			w.WriteHeader(http.StatusOK)
			switch ext {
			case ".jpg":
				w.Header().Set("Content-Type", "image/jpg")
				w.Write(message)
			case ".jpeg":
				w.Header().Set("Content-Type", "image/jpeg")
				w.Write(message)
			case ".png":
				w.Header().Set("Content-Type", "image/png")
				w.Write(message)
			case ".gif":
				w.Header().Set("Content-Type", "image/gif")
				w.Write(message)
			case ".html":
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				fmt.Fprint(w, string(message))
			case ".txt":
				w.Header().Set("Content-Type", "text/txt; charset=utf-8")
				w.Write(message)
			}
		}
	}
}

func main() {
	st := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", st))
	http.HandleFunc("/", handleMethod)
	log.Println("starting listener")
	log.Println("listener failed", "error", http.ListenAndServe("127.0.0.1:8080", nil))
}
