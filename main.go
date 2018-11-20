package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			// Print the error. Using a log.Fatal would exit the server
			log.Println(err)
		}
		io.WriteString(conn, "\nHello from TCP")
		fmt.Fprintln(conn, "\nHow is your day ?")
		fmt.Fprintf(conn, "%v", "Well, I hope!\n-----\n")

		conn.Close()
	}
}
