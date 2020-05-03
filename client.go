package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	server = "localhost"
	port   = "9000"
)

func main() {
	fmt.Println("Connected with " + server + ":" + port)
	conn, err := net.Dial("tcp", server+":"+port)

	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	handleConn(conn)

}

func handleConn(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

func handleError(err error) {
}
