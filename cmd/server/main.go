package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listening on port :6379")

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	resp := NewResp(conn)
	writer := NewWriter(conn)

	for {
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// Temporary smoke response until command handlers are wired up
		if err := writer.Write(Value{typ: "string", str: "OK"}); err != nil {
			fmt.Println(err)
			return
		}
	}
}