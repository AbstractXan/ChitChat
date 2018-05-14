package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings" //Try to remove this
)

var UserDB map[string]string
var UserConn map[string]net.Conn

func main() {
	//Initialize data maps
	UserDB := make(map[string]string)
	UserConn := make(map[string]net.Conn)
	UserDB["Xan"] = "123"

	fmt.Println("Server On.\nWaiting for users.")

	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	go User(conn, UserDB, UserConn) //Need to improve.

	ln1, err := net.Listen("tcp", ":8081")
	for err != nil {
		ln1, err = net.Listen("tcp", ":8081")
	}
	conn1, _ := ln1.Accept()
	User(conn1, UserDB, UserConn) //Need to improve.

	/* LATER ADDING COMMANDS TO SERVER CONSOLE
	for command := FromConsole();command!="exit"{

	}*/
}

func User(conn net.Conn, UserDB map[string]string, UserConn map[string]net.Conn) { //User Handler
	defer conn.Close()

	//Login
	connflag, username := Login(conn, UserDB, UserConn)
	// If connection exists, read from user

	//Printing Online users
	ToConsole("Online users are:")
	for v, c := range UserConn {
		fmt.Printf(v + " ")
		fmt.Println(c)
		fmt.Print("\n")
	}

	//Temporary service. connflag==true means connection has been established
	for connflag == true {

		//Have to write explicitely to catch error
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil { //Disconnected user

			fmt.Println(username + " disconnected")
			break
		}
		ToConsole(username+"> " + message)
		//modify how i/o works
		//newmessage := message
		fmt.Print("You> ")
		text := FromConsole()
		ToConn(conn, text)
	}
}

func PrintOnline() {
	ToConsole("Online users are:")
	for v, c := range UserConn {
		fmt.Printf(v + " ")
		fmt.Println(c)
		fmt.Print("\n")
	}
}

func Login(conn net.Conn, UserDB map[string]string, UserConn map[string]net.Conn) (bool, string) {

	//Read Username Request
	ToConn(conn, "Reading your username...\n")
	username := FromConn(conn)
	ToConsole("Username:" + username)
	//fmt.Println("Username: " + username)

	//Register password / Login /Kick
	if _, ok := UserDB[username]; ok == false {

		//If username absent, enter a new password
		ToConn(conn, "Enter a new password for "+username+"\n")
		pass := FromConn(conn) //GetPasswordValue

		UserDB[username] = pass   //Accept username and pass
		UserConn[username] = conn //Accept connection
		ToConsole("Pass: " + pass)
		ToConsole(username + " registered and connected.\n")
		return true, username

	} else {

		//USER exists, check password
		ToConn(conn, "Enter password for "+username+"\n")
		pass := FromConn(conn) //GetPasswordValue

		if UserDB[username] != pass {

			conn.Close() //Close connection
			ToConsole(username + " kicked.\nREASON: WRONG PASSWORD.\n")
			return false, username //Send failure
		} else {

			UserConn[username] = conn //Accept connection
			ToConsole(username + " connected.\n")
			return true, username //Return true
		}
	}
}

//func request()

/*func chatroom(conn1 net.Conn,conn2 net.Conn)
{
	p1,p2=net.pipe()

	go func(){
		for connflag == true {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil { //Disconnected user

			fmt.Println(username + " disconnected")
			break
		}
		ToConsole("Person> "+message)
		//modify how i/o works
		newmessage := message
		ToConn(conn,newmessage)
	}
  	}
}
*/
//========================================MAKE A LIBRARY======================================//
//Read from connection.
func FromConn(conn net.Conn) string {
	str, _ := bufio.NewReader(conn).ReadString('\n')
	str = strings.TrimSpace(str) //Trim extra spaces
	return str
}

//Write to Connection
func ToConn(conn net.Conn, mess string) {
	conn.Write([]byte(mess + "\n"))
}

//Read from input
func FromConsole() string {
	pass, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return pass
}

//Print to Output
func ToConsole(str string) {
	fmt.Println(str)
}

//================================================================================================//
