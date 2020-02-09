package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
)

func main() {
	var r io.Reader
	var err error
	r, err = os.Open("web.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1) }
	outline(nil, doc)
}
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	} }
