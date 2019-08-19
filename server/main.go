package main

import (
	"bufio"
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

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().String()+"\n")
		handleError(err)
		time.Sleep(1 + time.Second)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(message)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
