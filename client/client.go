package main

import (
	"fmt"
	"net"
)

func main() {
	server_conn, err := net.Dial("tcp", "127.0.0.1:123")
	if err != nil {
		fmt.Println("Error connecting to the server: ", err)
		return
	}
	defer server_conn.Close()
	fmt.Println("Connected to the server.")

	_, err = server_conn.Write([]byte(""))
	if err != nil {
		fmt.Println("Error writing to the server: ", err)
		return
	}

	buffer := make([]byte, 1024)
	byteSize, err := server_conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from the server: ", err)
		return
	}
	fmt.Println("Read from buffer: ", string(buffer[:byteSize]))

}
