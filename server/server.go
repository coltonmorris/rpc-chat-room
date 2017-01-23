package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
  "time"
)

type State struct {
	Users map[string][]string
	Shutdown bool
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
		Users : make(map[string][]string),
		Shutdown : false,
	}
  server := Server(make(chan *State, 1))
  server <- state

	// begin rpc listener
	rpc.Register(server)
	rpc.HandleHTTP()
	fmt.Println("Listening on ", address)

  go http.ListenAndServe(address, nil)
  for {
    state = <- server
    
    if state.Shutdown == true {
      fmt.Println("exit")
      return
    }
      server <- state
      time.Sleep(time.Second)
  }
//	err := http.ListenAndServe(address, nil)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
  //NOTES:
  /*
    maybe do go http.ListenAndServe(address, nil)

    <- server.Shutdown  this will be called when the server closes? or if it's full?
  */

  // return, aka exit this main loop, when the shutdown channel is true
  // maybe defer, which calls a function when main is returned
}
