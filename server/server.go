package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"net"
	"time"
)

func handleClient(client_conn net.Conn) {
	defer client_conn.Close()
	ntpTime, err := ntp.Time("pool.ntp.client")
	if err != nil {
		fmt.Println("Error getting ntp time: ", err)
	}
	unixTime := ntpTime.Unix()

	ntpTimeString := time.Unix(unixTime, 0).String()
	fmt.Println("Returning ntp time: ", ntpTimeString)
	_, err = client_conn.Write([]byte(ntpTimeString))
	if err != nil {
		fmt.Println("Error sending NTP time to client:", err)
		return
	}
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
