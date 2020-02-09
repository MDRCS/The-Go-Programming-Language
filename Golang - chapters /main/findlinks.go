package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
)


func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			} }
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() { //os.Stdin
	var r io.Reader
	var err error
	r, err = os.Open("web.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1) }
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	} }
