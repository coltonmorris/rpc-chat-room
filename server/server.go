package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	// parse command line arguments
	portPtr := flag.String("port", "3410", "the port number")
	hostPtr := flag.String("host", "127.0.0.1", "the host address")
	flag.Parse()

	// combine host and port
	address := net.JoinHostPort(*hostPtr, *portPtr)

	// begin rpc listener
  handlers := new(Handler)
	rpc.Register(handlers)
	rpc.HandleHTTP()
	fmt.Println("Listening on ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
