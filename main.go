package main

import (
	"bufio"
	"log"
	"net"
)

func handle(conn net.Conn) {
	log.Printf("Now serving: %s \n", conn.RemoteAddr().String())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		log.Println(ln)
	}
	defer conn.Close()

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			// Print the error using a log.Fatal would exit the server
			log.Println(err)
		}
		// Using a go routine to handle the connection
		go handle(conn)
	}
}
