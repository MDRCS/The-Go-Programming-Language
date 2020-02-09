// Clock_server is a TCP server that periodically writes the time.
//sequantial program for tcp connections
// we cannot launch two netcat client in the same time and make them work

package main

import (
	"io"
	"log"
	"net"
	"time"
)


func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue }
		//handleConn(conn) // handle one connection at a time
		go handleConn(conn) // handle connections concurrently
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

// command-line go run concurrent_clock_server.go &
// nc localhost 8000

//nc : netcat is a client to the server that we created