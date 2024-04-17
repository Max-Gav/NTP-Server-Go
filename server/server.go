package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

var ntpServer = "il.pool.ntp.org:123"

func printResponse(ntpResponse []byte) {
	transmitTime := binary.BigEndian.Uint32(ntpResponse[40:44])
	fraction := binary.BigEndian.Uint32(ntpResponse[44:48])
	ntpTime := time.Unix(int64(transmitTime-2208988800), int64(uint64(fraction)*(1<<32/uint64(1e9))))

	fmt.Println("Current NTP time:", ntpTime)
}

func handleClient(client_conn net.Conn) {
	defer client_conn.Close()

	conn, err := net.Dial("udp", ntpServer)
	if err != nil {
		fmt.Println("Error connecting to NTP server: ", err)
		return
	}
	defer conn.Close()

	ntpRequest := make([]byte, 48)
	ntpRequest[0] = 0x1B // Set NTP version and mode (client)
	_, err = conn.Write(ntpRequest)
	if err != nil {
		fmt.Println("Error sending NTP request:", err)
		return
	}

	ntpResponse := make([]byte, 48)
	_, err = conn.Read(ntpResponse)
	if err != nil {
		fmt.Println("Error receiving NTP response:", err)
		return
	}

	_, err = client_conn.Write([]byte(ntpResponse))
	if err != nil {
		fmt.Println("Error sending NTP time to client:", err)
		return
	}

	printResponse(ntpResponse)
}

func main() {
	listener, err := net.Listen("tcp", ":123")
	if err != nil {
		fmt.Println("Error setting up server: ", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is running and listening on port 123")

	for {
		fmt.Println("Waiting for a client!")
		client_conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting client: ", err)
			continue
		}

		fmt.Println("Handling client")
		go handleClient(client_conn)

	}

}
