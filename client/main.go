package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	server = "localhost"
	port   = "9000"
)

func main() {
	fmt.Println("Connected with " + server + ":" + port)
	conn, err := net.Dial("tcp", server+":"+port)

	handleError(err)
	defer conn.Close()

	handleConn(conn)

}

func handleConn(src io.Reader) {
	if _, err := io.Copy(os.Stdout, src); err != nil {
		log.Fatal(err)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
