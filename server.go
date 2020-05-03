package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"strings"
)

const (
	server = "localhost"
	port   = "9000"
)

func main() {
	fmt.Println(server + ":" + port)
	listener, err := net.Listen("tcp", server+":"+port)

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		command, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		//time.Sleep(1 + time.Second)

		//t := time.Now()
		//myTime := t.Format(time.RFC3339) + "\n"

		if strings.TrimSpace(string(command)) == "ls" {
			conn.Write([]byte(listDirectory()))
		} else {
			conn.Write([]byte("Command not found\n"))
		}
	}
}

func listDirectory() string {
	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}
	result := strings.Join(files, " ")
	return result + "\n"
}
