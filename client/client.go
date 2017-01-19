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

  // list request
  // TODO first firgure out the parition of server.go and handlers.go
  // TODO, start here. modify scheme.go


	// Synchronous call
	args := RpcScheme.Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot RpcScheme.Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

	var tmp RpcScheme.Temp
	err = client.Call("Arith.Tell", args, &tmp)
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Printf("Hey: %v", tmp.Message)
}
