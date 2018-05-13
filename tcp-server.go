package main

import (
	"bufio"
	"fmt"
	"net"
	"strings" //Try to remove this
)

var UserDB map[string]string
var UserConn map[string]net.Conn

func main() {
	//Initialize data maps
	UserDB := make(map[string]string)
	UserConn := make(map[string]net.Conn)
	UserDB["Xan"] = "123"

	fmt.Println("Server On. Waiting for request.")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	User(conn,UserDB,UserConn)	//Need to improve.
}

func User(conn net.Conn, UserDB map[string]string, UserConn map[string]net.Conn) { //User Handler
	defer conn.Close();
	connflag, username := Login(conn,UserDB,UserConn) //!!!!!!flag, username
	for connflag == true {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {

			fmt.Println(username + " disconnected")
			break
		}
		fmt.Print("Person>", message)

		//modify how i/o works
		newmessage := message
		conn.Write([]byte(newmessage + "\n"))
	}
}

func GetClientString(conn net.Conn)string{
	str, _ := bufio.NewReader(conn).ReadString('\n')
	str = strings.TrimSpace(str) //Trim extra spaces
	return str
}

func MessageToClient(conn net.Conn,mess string){
	conn.Write([]byte(mess))
}
func Login(conn net.Conn,UserDB map[string]string, UserConn map[string]net.Conn) (bool, string) {
	//Initializing data maps

	//Read Username Request
	//conn.Write([]byte("Reading your username...\n"))
	MessageToClient(conn,"Reading your username...\n")
	username, _ := bufio.NewReader(conn).ReadString('\n')
	username = strings.TrimSpace(username) //Trim extra spaces
	fmt.Println("Username: " + username)

	//Register password / Login /Kick
	if _, ok := UserDB[username]; ok == false {

		//If username absent, enter a new password
		conn.Write([]byte("Enter a new password for "+username+"\n"))
		pass := GetClientString(conn); //GetPasswordValue

		UserDB[username] = pass   //Accept username and pass
		UserConn[username] = conn //Accept connection
		fmt.Println("Pass: " + pass)
		fmt.Println(username + " registered and connected.\n")
		return true, username

	} else {

		//USER exists, check password
		conn.Write([]byte("Enter password for " + username + "\n"))
		pass := GetClientString(conn); //GetPasswordValue

		if UserDB[username] != pass {
			conn.Close() //Close connection
			fmt.Printf("%s kicked.\nREASON: WRONG PASSWORD.\n", username)
			return false, username //Send failure
		} else {
			UserConn[username] = conn //Accept connection
			fmt.Printf("%s connected.\n", username)
			return true, username //Return true
		}
	}

}
