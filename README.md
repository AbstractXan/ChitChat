# ChitChat

Chit-Chat is a seamless local-network server-client based terminal texting program. 

## Setting up files

1. Copy all files (including /src) to your `GOPATH`
2. If you don't know your GOPATH, type `go env GOPATH` in your terminal.
  For Go 1.8 and onwards the default paths are `$HOME/go` on Unix-like systems and `%USERPROFILE%\go` on Windows.

## Running Server
```
$ go run tcp-server.go
```
- [ ] Typing "quit": should shutdown the server.

## Running Client
```
$ go run tcp-client.go
```
Typing "quit" disconnects the client(user) from the server.
