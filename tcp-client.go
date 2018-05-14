package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
  "strings"
)

//Set some 'anon' by default before distributing codes
//Keep a server check for users to type in new usernames
var username string
var serverip = "127.0.0.1:8081"

//========================================MAKE A LIBRARY======================================// 
//Read from connection.
func FromConn(conn net.Conn)string{
  str, _ := bufio.NewReader(conn).ReadString('\n')
  str = strings.TrimSpace(str) //Trim extra spaces
  return str
}

//Write to Connection
func ToConn(conn net.Conn,mess string){
  conn.Write([]byte(mess+"\n"))
}

//Read from input
func FromConsole()string{
  pass, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  return pass
}

//Print to Output
func ToConsole(str string){
  fmt.Println(str)
}
//================================================================================================//

func main() {

  ToConsole("Enter your username.")
  username=FromConsole()

  var conn net.Conn
  var err error
  //fmt.Println("Coudn't connect to server: %s\nPlease check your server ip",serverip)
  for{
  conn, err = net.Dial("tcp", serverip)
   if err == nil{break}
  }

	//Sending username
  ToConsole(FromConn(conn))
	ToConn(conn,username + "\n")
	//fmt.Fprintf(conn, username + "\n")

	//Sending Password
  ToConsole(FromConn(conn))
	ToConsole("Password: ")

	//Reading password from console
	pass:=FromConsole()
	ToConn(conn,pass + "\n")

  /*GetConfirmation
  Conf:=FromConn()
  if Conf==false{
    ToConsole("Confirmation failed")
    conn.Close();
    return
  }
*/


	for {
    fmt.Print("You> ")
    text := FromConsole()
    ToConn(conn,text)
    ToConsole("Server> "+ FromConn(conn))
	}


}