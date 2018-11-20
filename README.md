# Simple TCP server written in Go.

I wrote this programm to get familiar with the net package in go. My final goal is to continue to learn about go and
code using go routines and concurrency patterns.

## Being concurrent
The implementation of the main() function tells our TCP server to start a new goroutine each time it
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

The net.Listen() call is used to accept network connections and thus act as a server. The return value of net.Listen()
is of the net.Conn type, which implements both io.Reader and io.Writer interfaces. The for loop allows our program to
keep accepting new TCP clients using Accept() that will be handled by instances of the handle() function, which are
executed as goroutines.