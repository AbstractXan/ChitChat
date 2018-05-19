package main

import (
	"fmt"
	io "iofunc" //Replace later on
	"net"
	"sync"
)

var UserDB map[string]string
var UserConn map[string]net.Conn
var wg = sync.WaitGroup{}

func main() {
	//Initialize data maps
	UserDB := make(map[string]string)
	UserConn := make(map[string]net.Conn)
	UserDB["Xan"] = "123"

	io.ToConsole("Server On.\nWaiting for users.")

	//ln, _ := net.Listen("tcp", ":8081")
	//conn, _ := ln.Accept()
	//go User(conn, UserDB, UserConn) //Need to improve.

	/* LATER ADDING COMMANDS TO SERVER CONSOLE
	for command := FromConsole();command!="exit"{

	}*/

	for i := 0; i < 2; i++ {
		ln1, err := net.Listen("tcp", ":8081")
		for err != nil {
			ln1, err = net.Listen("tcp", ":8081")
		}

		conn1, _ := ln1.Accept()
		go User(conn1, UserDB, UserConn) //Need to improve.

	}
}

func Broadcast(conn net.Conn, UserDB map[string]string, UserConn map[string]net.Conn) {

}

func User(conn net.Conn, UserDB map[string]string, UserConn map[string]net.Conn) { //User Handler
	defer conn.Close()

	//Login
	connflag, username := Login(conn, UserDB, UserConn)
	// If connection exists, read from user

	//Printing Online users
	PrintOnline(UserConn)

	//Temporary service. connflag==true means connection has been established
	for connflag == true {

		/*
			message, err := FromConnErr(conn)
			if err != nil { //Disconnected user

				fmt.Println(username + " disconnected")
				break
			}
			ToConsole(username + "> " + message)
			//modify how i/o works
			//newmessage := message
			fmt.Print("You> ")
			text := FromConsole()
			ToConn(conn, "SERVER> "+text)*/
		wg.Add(1)
		go Listener(conn, username)
		wg.Add(1)
		go Writer(conn)
		wg.Wait()
	}
}

func Listener(conn net.Conn, username string) {
	defer conn.Close()
	defer wg.Done()
	for {
		str, err := io.FromConnErr(conn)
		if err != nil {
			io.ToConsole(username + " Disconnected")
			return
		}
		io.ToConsole(str)
	}
}

func Writer(conn net.Conn) {
	defer conn.Close()
	defer wg.Done()
	for {
		//fmt.Print("You> ")
		text := io.FromConsole()
		if text == "quit" {
			io.ToConsole("Closing Connection")
			return
		}
		io.ToConn(conn, "SERVER> "+text)
	}
}

func PrintOnline(UserConn map[string]net.Conn) {
	io.ToConsole("Online users are:")
	for v, c := range UserConn {
		fmt.Printf(v + " ")
		fmt.Println(c)
		fmt.Print("\n")
	}
}

func Login(conn net.Conn, UserDB map[string]string, UserConn map[string]net.Conn) (bool, string) {

	//Read Username Request
	io.ToConn(conn, "Reading your username...\n")
	username := io.FromConn(conn)
	io.ToConsole("Username:" + username)
	//fmt.Println("Username: " + username)

	//Register password / Login /Kick
	if _, ok := UserDB[username]; ok == false {

		//If username absent, enter a new password
		io.ToConn(conn, "Enter a new password for "+username+"\n")
		pass := io.FromConn(conn) //GetPasswordValue

		UserDB[username] = pass   //Accept username and pass
		UserConn[username] = conn //Accept connection
		io.ToConsole("Pass: " + pass)
		io.ToConsole(username + " registered and connected.\n")
		return true, username

	} else {

		//USER exists, check password
		io.ToConn(conn, "Enter password for "+username+"\n")
		pass := io.FromConn(conn) //GetPasswordValue

		if UserDB[username] != pass {

			conn.Close() //Close connection
			io.ToConsole(username + " kicked.\nREASON: WRONG PASSWORD.\n")
			return false, username //Send failure
		} else {

			UserConn[username] = conn //Accept connection
			io.ToConsole(username + " connected.\n")
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
