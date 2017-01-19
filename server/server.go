package main

import (
	"../" // TODO find out the proper way to do this
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Arith int

func (t *Arith) Multiply(args *RpcScheme.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// have the reciever be a pointer if it is changing state
func (t *Arith) Divide(args *RpcScheme.Args, quo *RpcScheme.Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func (t *Arith) Tell(args *RpcScheme.Args, tmp *RpcScheme.Temp) error {
	tmp.Message = "testing you bonobo"
	fmt.Println("hello ther colton")
	return nil
}

func main() {
	// parse command line arguments
	portPtr := flag.String("port", "3410", "the port number")
	hostPtr := flag.String("host", "127.0.0.1", "the host address")
	flag.Parse()

	// combine host and port
	address := net.JoinHostPort(*hostPtr, *portPtr)

	// begin rpc listener
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	fmt.Println("Listening on ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
