package main

import (
	"bufio"
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

func handleConn(src io.Writer) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(src, text)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
