package main

import (
	"flag"
	"gopl.io/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {

	list,err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}


//Functionning of a web crawler is simple, go to a website extract all links in it and crawl also these websites, -> infinite loop
func main() {
	var depth = flag.Int("depth",3,"define links to crawl in each website")
	flag.Parse()
	worklist := make(chan []string,*depth)
	go func() {
		worklist <- os.Args[1:1]
	}()

	seen := make(map[string]bool)

	for list := range worklist {
		for _,link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}

}

//command-line go run web_crawler.go http://gopl.io/ -depth=3