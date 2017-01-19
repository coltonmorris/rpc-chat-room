package main

import (
	"../" // TODO find out the proper way to do this
  "flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {
  // handle command line arguments
  handlePtr := flag.String("handle", "Colton", "your handle/name")
  hostPtr := flag.String("host", "127.0.0.1", "the host address")
	portPtr := flag.String("port", "3410", "the port number")
  flag.Parse()
  
  // initialize RPC client
	client, err := rpc.DialHTTP("tcp", net.JoinHostPort(*hostPtr, *portPtr))
	if err != nil {
		log.Fatal("dialing:", err)
	}

  // welcome user
  fmt.Printf("Hi %v, connecting to %v:%v...\n", *handlePtr, *hostPtr, *portPtr)

  // tell handler
  request := RpcScheme.TellRequest{"sup brotato", "Colton"}
  response := RpcScheme.TellResponse{}
	err = client.Call("Handler.Tell", &request, &response)
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Printf("Hey: %v\n", response)
}
