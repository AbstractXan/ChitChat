package iofunc
import (
	"bufio"
	"net"
	"os"
	"strings"
	//"fmt"
)

//Recieve from conn
func FromConn(conn net.Conn) string {
	str, _ := bufio.NewReader(conn).ReadString('\n')
	str = strings.TrimSpace(str) //Trim extra spaces
	return str
}

//Recieve with error
func FromConnErr(conn net.Conn) (string,error) {
	str, err := bufio.NewReader(conn).ReadString('\n')
	if err!=nil{
			return "error",err
		}
	str = strings.TrimSpace(str) //Trim extra spaces
	return str, err
}


//Write to Connection
func ToConn(conn net.Conn, mess string) {
	conn.Write([]byte(mess + "\n"))
}

//Read from input
func FromConsole() string {
	pass, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	pass = strings.TrimSpace(pass)
	return pass
}

//Print to Output
func ToConsole(str string) {
	fmt.Println(str)
}

