package main

import (
	"fmt"
	"log"
	"net"
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
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
