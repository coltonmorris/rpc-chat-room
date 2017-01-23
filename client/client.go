package main

import (
  // "bufio"
  "flag"
	"fmt"
	"log"
  "net"
	"net/rpc"
  "time"
  "bufio"
)

func main() {
  // handle command line arguments
  handlePtr := flag.String("handle", "Colton", "-handle=\"your handle/name\"")
  hostPtr := flag.String("host", "127.0.0.1", "-host=\"the host address\"")
	portPtr := flag.String("port", "3410", "-port=\"the port number\"")
  flag.Parse()
  
  // initialize RPC client
	client, err := rpc.DialHTTP("tcp", net.JoinHostPort(*hostPtr, *portPtr))
	if err != nil {
		log.Fatal("dialing:", err)
	}
  fmt.Printf("%T", client)
  fmt.Println("****************************************************")
  fmt.Printf("\tConnecting to %v:%v...\n", *hostPtr, *portPtr)
  fmt.Println("****************************************************")

  Login(client, "Trump")
  Login(client, "Stalin")
  Login(client, "Putin")
  Login(client, *handlePtr)
  ListUsers(client, *handlePtr)
  TellUser(client, "Putin", *handlePtr, "What's up dude?") 
  CheckMessages(client, *handlePtr)
  CheckMessages(client, *handlePtr)
  Help()
  Say(client, "Trump", "Grab her by the pussy")
  // Logout(client, *handlePtr)
  Shutdown(client, *handlePtr)
  // CheckMessages(client, *handlePtr)
  // go func() {
    for {
      CheckMessages(client, *handlePtr)
      time.Sleep(time.Second)
    }
  // }()

  scanner := bufio.NewScanner(strings.NewReader())

  //go func() {
  //  for {
  //    CheckMessages(client,*handlePtr)
  //    time.Sleep(time.Second)
  //  }
  //}()
}
