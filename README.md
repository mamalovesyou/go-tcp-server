This is a simple TCP server written in Go.

## Being concurrent
The implementation of the main() function tells the TCP server to start a new goroutine each time it
has to serve a TCP client:

```go
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
```