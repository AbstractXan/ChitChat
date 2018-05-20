package main

import (
	io "iofunc"
	"net"
)

type ServerType struct {
	UserDB   map[string]string
	UserConn map[string]net.Conn
	Name     string
	//GoRoutines int
}

//Initializes and returns a new Server
func NewServerType() *ServerType {
	var s ServerType

	//Server name
	io.ToConsole("Enter a Server Name: ")
	s.Name = io.FromConsole()

	//Variables
	s.UserDB = make(map[string]string)
	s.UserConn = make(map[string]net.Conn)

	//Constants
	s.UserDB["Xan"] = "123"
	return &s
}

func main() {
	Server := NewServerType()
	io.ToConsole("Server On.\nWaiting for users.")

	//Loop to recieve and accept new connections
	for i := 0; ; i++ {
		ln1, err := net.Listen("tcp", ":8081")
		for err != nil {
			ln1, err = net.Listen("tcp", ":8081")
		}

		conn1, _ := ln1.Accept()

		go Server.UserHandler(conn1) //Need to improve.

	}
}

//User Handler
func (Server *ServerType) UserHandler(conn net.Conn) {
	defer conn.Close()

	//Login
	connflag, username := Server.LoginHandler(conn)

	//Printing Online users
	//Temporary service. connflag==true means connection has been established
	if connflag == true {
		Server.BroadcastToAll(Server.GetOnline())
		go Server.Writer(conn)          //Run Writer
		Server.Listener(conn, username) //UserHandler terminates when listener terminates
	}
}

//Handles Login for each user
func (Server *ServerType) LoginHandler(conn net.Conn) (bool, string) {

	//Send Server Name
	io.ToConn(conn, "[SERVER] Welcome to Chit-Chat server: "+Server.Name+ "Reading your username.")

	//Read Username Request
	username := io.FromConn(conn)
	
	//Register password / Login /Kick
	//Password DNE => Add new user
	//Password Exists => Check password
	if _, boolpass := Server.UserDB[username]; boolpass == false {

		//If username absent, enter a new password
		io.ToConn(conn, "[Server] >> Enter a new password for "+username+"\n")
		pass := io.FromConn(conn) //GetPasswordValue

		Server.UserDB[username] = pass   //Accept username and pass
		Server.UserConn[username] = conn //Accept connection
		io.ToConsole("[LOGIN] >>" + username + ": " + pass + " registered and connected.\n")
		io.ToConn(conn, "[SERVER] >> "+username+" registered and connected.\n")
		return true, username

		//Check if user already online
	} else if _, boolOnline := Server.UserConn[username]; boolOnline == false {
		//USER exists, check password
		io.ToConn(conn, "[SERVER] >> Enter password for "+username+"\n")
		pass := io.FromConn(conn) //GetPasswordValue

		//No pwd match => KICK
		if Server.UserDB[username] != pass {

			io.ToConsole("[LOGIN] >> " + username + " kicked.WRONG PASSWORD.\n")
			io.ToConn(conn, "[SERVER] >> "+username+" kicked. WRONG PASSWORD.\n")
			conn.Close()           //Close connection
			return false, username //Send failure

			//Pwd Match => Accept
		} else {

			Server.UserConn[username] = conn //Accept connection
			io.ToConsole("[LOGIN] >> " + username + " connected.")
			io.ToConn(conn, "[SERVER] >> Connected.")
			return true, username //Return true
		}

		//Kick if already online
	} else {
		io.ToConn(conn, "[SERVER] >> Enter password for "+username+"\n")
		_ = io.FromConn(conn) //GetPasswordValue
		io.ToConsole("[LOGIN HANDLER] >> " + username + " kicked. REASON: User is Already Online.")
		io.ToConn(conn, "[SERVER] >> "+username+" kicked. REASON: USER IS ALREADY ONLINE.")
		conn.Close()           //Close connection
		return false, username //Send failure
	}
}

//Other useful IO functions
//Listens until connection error
func (Server *ServerType) Listener(conn net.Conn, username string) {
	//Delete username from active connections
	defer delete(Server.UserConn, username)
	defer conn.Close()
	for {
		str, err := io.FromConnErr(conn)
		if err != nil {
			io.ToConsole(username + " Disconnected")
			return
		}
		io.ToConsole(username + "> " + str)
		Server.BroadcastFromOne(conn, username+"> "+str) //Broadcast from one user to all others
	}
}

//Broadcast(str) from one user(username,conn) to all others in UserConn map
func (Server *ServerType) BroadcastFromOne(conn net.Conn, str string) {
	for _, uconn := range Server.UserConn {
		if conn == uconn {
			continue
		}
		io.ToConn(uconn, str)
	}
}

//Keeps Writing stuff, tracks commands
func (Server *ServerType) Writer(conn net.Conn) {
	defer conn.Close()
	for {
		text := io.FromConsole()
		if text == "quit" {
			io.ToConsole("Closing Connection")
			return
		}
		Server.BroadcastToAll("SERVER> " + text)
	}
}

//Broadcast(str) to all users in UserConn map
func (Server *ServerType) BroadcastToAll(str string) {
	for _, uconn := range Server.UserConn {
		io.ToConn(uconn, str)
	}
}

//Command functions
//Prints Online Users
func (Server *ServerType) GetOnline() string {
	str := "Online users are: "
	for user, _ := range Server.UserConn {
		str = str + user + " | "
	}
	io.ToConsole(str)
	return str
}

/* Additional updates:
1. Server Commands
2. Client Request Handler
3. Chatrooms
4. Server Name
5. Server Discoverable to all
*/
