package main 
import 
("fmt"
 "net"
 "io"
 "os"
)

func main() {
	fmt.Println("Listening on port 6379...")

	// Create a new server
	l, err := net.Listen("tcp", ":6379") // 6379 - default Redis port 
	if err != nil{
		fmt.Println(err)
		return
	}

	// Listen for connections 
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	
	for { //infinite loop
		buf := make([]byte, 1024)

		//read message from client
		_, err := conn.Read(buf) // _ amount of read bytes

		if err != nil {
			if err == io.EOF { // End Of File
				break
			} 
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)

		}
		conn.Write([]byte("+OK\r\n"))
	} 
}
