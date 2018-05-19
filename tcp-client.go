package main

import (
	"net"
    io "iofunc" //Replace later on
    "sync"
)

//Set some 'anon' by default before distributing codes
//Keep a server check for users to type in new usernames
var username string
var serverip = "127.0.0.1:8081"
var wg = sync.WaitGroup{}

func main() {

  io.ToConsole("Enter your username.")
  username=io.FromConsole()

  var conn net.Conn
  var err error
  //fmt.Println("Coudn't connect to server: %s\nPlease check your server ip",serverip)
  for{
  conn, err = net.Dial("tcp", serverip)
   if err == nil{break}
  }

	//Sending username
  	io.ToConsole(io.FromConn(conn))
	io.ToConn(conn,username + "\n")
	//fmt.Fprintf(conn, username + "\n")

	//Sending Password
  	io.ToConsole(io.FromConn(conn))
	io.ToConsole("Password: ")

	//Reading password from console
	pass:=io.FromConsole()
	io.ToConn(conn,pass)

  /*GetConfirmation
  Conf:=FromConn()
  if Conf==false{
    ToConsole("Confirmation failed")
    conn.Close();
    return
  }
*/
  	wg.Add(1)
  	go Listener(conn)
  	wg.Add(1)
  	go Writer(conn)
  	wg.Wait()
}

func Listener(conn net.Conn){
	defer conn.Close();
	defer wg.Done()
	for{
		str, err := io.FromConnErr(conn)
		if err!=nil{
				io.ToConsole("Server Disconnected")
				return
			}
		io.ToConsole(str)
	}
}

func Writer(conn net.Conn){
	defer conn.Close();
	defer wg.Done()
	for {
    //fmt.Print("You> ")
    text := io.FromConsole()
    	if text=="quit" {
    		io.ToConsole("Closing Connection")
    		return
    	}
    io.ToConn(conn,text)
	}
}