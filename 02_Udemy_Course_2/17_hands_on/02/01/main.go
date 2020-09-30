//Create a basic server using TCP.
//The server should use net.Listen to listen on port 8080.
//Remember to close the listener using defer.
//Remember that from the "net" package you first need to LISTEN, then you need to ACCEPT an incoming connection.
//Now write a response back on the connection.
//Use io.WriteString to write the response: I see you connected.
package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		io.WriteString(conn, "i see you connected")
		conn.Close()
	}
}
