package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var i int
	var method, uri string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			method = xs[0]
			uri = xs[1]
			fmt.Println("METHOD", method)
			fmt.Println("URI", uri)
		}
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	body += "\n"
	body += method
	body += "\n"
	body += uri
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}
