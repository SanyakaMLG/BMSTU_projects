package main

import (
	"github.com/mgutz/logxi/v1"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func getAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func getChildren(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func isElem(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

func isText(node *html.Node) bool {
	return node != nil && node.Type == html.TextNode
}

func isDiv(node *html.Node, class string) bool {
	return isElem(node, "div") && getAttr(node, "class") == class
}

func isArticle(node *html.Node) bool {
	return isElem(node, "article") && strings.HasPrefix(getAttr(node, "class"), "article story")
}

type Item struct {
	Ref, Title string
}

func readItem(item *html.Node) *Item {
	children := getChildren(item)
	for _, child := range children {
		if isDiv(child, "info") {
			child2 := getChildren(getChildren(getChildren(child)[0])[1])[0]
			return &Item{
				Ref: getAttr(child2, "href"),
				Title: getChildren(child2)[0].Data,
			}
		}
	}
	return nil
}

func search(node *html.Node) []*Item {
	if isDiv(node, "content article-list") {
		var items []*Item
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isArticle(c) {
				if item := readItem(c); item != nil {
					items = append(items, item)
				}
			}
		}
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := search(c); items != nil {
			return items
		}
	}
	return nil
}

func downloadNews() []*Item {
	log.Info("sending request to foxnews.com")
	if response, err := http.Get("https://www.foxnews.com"); err != nil {
		log.Error("request to foxnews.com failed", "error", err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		log.Info("got response from foxnews.com", "status", status)
		if status == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				log.Error("invalid HTML from foxnews.com", "error", err)
			} else {
				log.Info("HTML from foxnews.com parsed successfully")
				return search(doc)
			}
		}
	}
	return nil
}
