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

	io.Copy(os.Stdout, conn)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
