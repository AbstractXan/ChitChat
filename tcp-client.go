package main

import (
	io "iofunc"
	"net"
)

type UserType struct {
	Name     string
	Serverip string
}

func NewUser() *UserType {
	var u UserType

	io.ToConsole("Enter your userame.")
	u.Name = io.FromConsole()

	io.ToConsole("Enter Server IP (Port defaults to 8081).")
	u.Serverip = io.FromConsole() + ":8081"
	return &u
}
func main() {
	//Setup New User
	User := NewUser()

	//Connect
	conn := User.GetConn()
	User.LoginHandler(conn)
	go User.Writer(conn)
	User.Listener(conn)
}

func (User *UserType) GetConn() net.Conn {
	var conn net.Conn
	var err error

	for {
		conn, err = net.Dial("tcp", User.Serverip)
		if err == nil {
			break
		}
	}
	return conn
}

//Listens to any input from connection
func (User *UserType) Listener(conn net.Conn) {
	defer conn.Close()
	for {
		str, err := io.FromConnErr(conn)
		if err != nil {
			io.ToConsole("Server Disconnected")
			return
		}
		io.ToConsole(str)
	}
}

//Keeps Writing to the connection
func (User *UserType) Writer(conn net.Conn) {
	defer conn.Close()
	for {
		text := io.FromConsole()
		if text == "quit" {
			io.ToConsole("Closing Connection")
			return
		}
		io.ToConn(conn, text)
	}
}

//Login Handler made wrt Server
func (User *UserType) LoginHandler(conn net.Conn) {
	//Accept Server Name
	io.ToConsole(io.FromConn(conn))

	//Sending Name
	io.ToConn(conn, User.Name)

	//Sending Password
	io.ToConsole(io.FromConn(conn))

	//Reading password from console
	pass := io.FromConsole()
	io.ToConn(conn, pass)

	//Affirmation from Server
	io.ToConsole(io.FromConn(conn))
}
