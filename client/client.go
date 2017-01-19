package main

import (
	"../" // TODO find out the proper way to do this
  "flag"
	"fmt"
	"log"
  "net"
	"net/rpc"
)
func login(client *rpc.Client, handle string) {
  // login handler
  loginRequest := RpcScheme.LoginRequest{handle}
  loginResponse := RpcScheme.LoginResponse{}
  err := client.Call("Server.Login", &loginRequest, &loginResponse)
  if err != nil {
    log.Fatal("Error: ", err)
  }
}

func listUsers(client *rpc.Client, handle string) {
  // list handler
  listRequest := RpcScheme.ListRequest{}
  listResponse := RpcScheme.ListResponse{}
  err := client.Call("Server.List", &listRequest, &listResponse)
  if err != nil {
    log.Fatal("Error: ", err)
  }
  
  // check if we are the only user online
  if len(listResponse.Users) == 1 {
    fmt.Println("There are no other users online.")
    return 
  }

  fmt.Println("List of users currently online:")
  for i := 0; i < len(listResponse.Users); i++ {
    if listResponse.Users[i] != handle {
      fmt.Println("\t", listResponse.Users[i])
    }
  }
}

func tellUser(client *rpc.Client, handle string, user string, message string) {
  tellRequest := RpcScheme.TellRequest{handle, user, message}
  tellResponse := RpcScheme.TellResponse{}
  err := client.Call("Server.Tell", &tellRequest, &tellResponse)
	if err != nil {
		log.Fatal("error: ", err)
	}

  fmt.Println(tellResponse.Result)
}

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
  fmt.Println("****************************************************")
  fmt.Printf("\tConnecting to %v:%v...\n", *hostPtr, *portPtr)
  fmt.Println("****************************************************")

  login(client, *handlePtr)
  listUsers(client, *handlePtr)
  tellUser(client, *handlePtr, "Stalin", "What's up dude?") 
}
