package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
		stream, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		command := strings.TrimSpace(string(stream))
		commands := strings.Split(command, " ")

		switch commands[0] {
		case "ls":
			conn.Write([]byte(listDirectory()))
		case "mkdir":
			conn.Write([]byte(makeDirectory(commands[1])))
		case "cd":
			conn.Write([]byte(changeDirectory(commands[1])))
		default:
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

func makeDirectory(value string) (response string) {
	err := os.Mkdir(value, 0755)
	if err != nil {
		response = err.Error() + "\n"
	} else {
		response = value + " created\n"
	}
	return
}

func changeDirectory(value string) (response string) {
	os.Chdir(value)

	mydir, err := os.Getwd()
	if err == nil {
		response = mydir + "\n"
	} else {
		response = err.Error() + "\n"
	}

	return
}
