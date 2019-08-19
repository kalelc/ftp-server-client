package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const (
	server = "localhost"
	port   = "9000"
)

func main() {
	fmt.Println(server + ":" + port)
	listener, err := net.Listen("tcp", server+":"+port)

	handleError(err)

	for {
		conn, err := listener.Accept()

		handleError(err)

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().String()+"\n")
		handleError(err)
		time.Sleep(1 + time.Second)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
