# FTP Server and Client using Golang

This is a simple project using Golang Net Library with Goroutines to replicate FTP Server and Client Applications.

- `go run server/main.go`: Run FTP Server application, this receives and executes commands.

- `go run client/main.go`: Run FTP Client application, it sends commands to the server.

### Commands Allowed from Client

- [x] `ls`: List the names of the files in the current remote directory.

- [ ] `cd`: Change directory on the remote machine.

- [ ] `mkdir`: Make a new directory within the current remote directory.

- [ ] `get`: Copy one file from the remote machine to the local machine.

- [ ] `put`: Copy one file from the local machine to the remote machine.
