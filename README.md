# Simple TCP server written in Go.

I wrote this program to get familiar with the net package and go routines. 
My main objective here is to continue to learn about Go.
Here is a list of goals I want to achieve:

- [x] Write a simple tcp server to handle connections
- [x] Use goroutines to handle connections 
- [x] Encoding the request with SHA1 algorithm
- [x] Sending response and close connection

## 1 - Accepting a connection

The net package provide a simple function `net.Listen`  to start to listen on a specific port.
It return a `Listener` which implement the `Accept ` function.

Here is what my first lines:

```go
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		// Use fatal to exit if the listener fails to start
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		break
	}
}



 
## 2 - Concurrency
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

## 3 - Encoding the request

This point is just to avoid to return the client the request itself. So I just wrote a simple function to encode the 
request with a sha1 algorithm.
```go
// SHA1 hashes using sha1 algorithm
func SHA1(text string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
```

## 4 - Sending response and close connection


