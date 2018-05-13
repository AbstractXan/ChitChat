package main

import(
	"net"
	"fmt"
	"bufio"
	"os"
)

//Set some 'anon' 
var username = "Xan" 


func GetServerMessage(conn net.Conn){
  serv_message, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Print(serv_message)
}

func main() {
  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  defer conn.Close();
  
  //Sending username
  GetServerMessage(conn)
  conn.Write([]byte(username + "\n"))
  //fmt.Fprintf(conn, username + "\n")

  //Sending Password
  GetServerMessage(conn)
  fmt.Print("Password: ")
  
  //Reading password from console
  reader := bufio.NewReader(os.Stdin)
  pass , _ := reader.ReadString('\n')
  conn.Write([]byte(pass + "\n"))

  for { 
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("You> ")
    txt, _ := reader.ReadString('\n')

    fmt.Fprintf(conn, txt + "\n")
    
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Println("Server> "+message)
  }
}


//Changes:
//  fmt.Fprintf(conn, username + "\n") -> Removed as it only prints.
//  conn.Write([]byte(username + "\n")) -> Actually passes variables