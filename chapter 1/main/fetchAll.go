package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	start := time.Now()
	ch := make(chan string) // The main function creates a channel of strings using make.
	for _, url := range os.Args[1:] {
		go fetch(url,ch)
	}

	f,err := os.Create("Response.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v",err)
	}

	w := bufio.NewWriter(f)
	for range os.Args[1:] {
		w.WriteString(<- ch + "\n")
		w.Flush()
		fmt.Println(<- ch)
	}

	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string,ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) //Send ERROR to the channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard,response.Body)
	response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While Reading %s : %v",url,err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes,url)
}

// Command to execute the program on the command line go run fetchAll.go https://www.facebook.com/ https://github.com/ https://www.google.com/