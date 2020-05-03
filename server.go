package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"strings"
	"time"
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

		time.Sleep(1 + time.Second)

		fmt.Println(string(command))

		if strings.TrimSpace(string(command)) == "ls" {
			fmt.Println("list directory")
		} else {
			fmt.Println("command invalid")
		}
		listDirectory()
	}
}

func listDirectory() {
	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}
