package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type State struct {
  Status string
	Users map[string][]string
	End   chan bool
}

type Server chan *State

func main() {
	// parse command line arguments
	portPtr := flag.String("port", "3410", "the port number")
	hostPtr := flag.String("host", "127.0.0.1", "the host address")
	flag.Parse()

	// combine host and port
	address := net.JoinHostPort(*hostPtr, *portPtr)

  //prepare reciever with a buffer size of 1 to prevent race condiitons
	state := &State{
    Status: "Server is running",
		Users : make(map[string][]string),
		End : make(chan bool),
	}
  server := Server(make(chan *State, 1))
  server <- state

	// begin rpc listener
	rpc.Register(server)
	rpc.HandleHTTP()
	fmt.Println("Listening on ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
