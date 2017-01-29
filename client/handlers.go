package main

import (
	"../" // TODO find out the proper way to do this
  "fmt"
  "log"
  "net/rpc"
  "os"
)

func Login(client *rpc.Client, handle string) {
  // login handler
  request := RpcScheme.LoginRequest{handle}
  response := RpcScheme.LoginResponse{}
  err := client.Call("Server.Login", &request, &response)
  if err != nil {
    // log.Fatal("Error: ", err)
    fmt.Println("User already exists: ", handle)
  }
}

func ListUsers(client *rpc.Client, handle string) {
  // list handler
  request := RpcScheme.ListRequest{}
  response := RpcScheme.ListResponse{}
  err := client.Call("Server.List", &request, &response)
  if err != nil {
    log.Fatal("Error: ", err)
  }
  
  // check if we are the only user online
  if len(response.Users) == 1 {
    fmt.Println("There are no other users online.")
    return 
  }

  // TODO probably should just add this to the messages on the server side
  fmt.Println("List of users currently online:")
  for i := 0; i < len(response.Users); i++ {
    if response.Users[i] != handle {
      fmt.Println("\t", response.Users[i])
    }
  }
}

func TellUser(client *rpc.Client, sender string, target string, message string) {
  request := RpcScheme.TellRequest{sender, target, message}
  response := RpcScheme.TellResponse{}
  err := client.Call("Server.Tell", &request, &response)
	if err != nil {
		log.Fatal("error: ", err)
	}
}

func Say(client *rpc.Client, sender string, message string) {
  request := RpcScheme.SayRequest{sender, message}
  response := RpcScheme.SayResponse{}
  err := client.Call("Server.Say", &request, &response)
	if err != nil {
		log.Fatal("error: ", err)
	}
}

func Logout(client *rpc.Client, handle string) {
  request := RpcScheme.LogoutRequest{handle}
  response := RpcScheme.LogoutResponse{}
  err := client.Call("Server.Logout", &request, &response)
  if err != nil {
    log.Fatal("error: ", err)
  } else {
    os.Exit(0)
  }
}

func Shutdown(client *rpc.Client, handle string ) {
  request := RpcScheme.ShutdownRequest{handle}
  response := RpcScheme.ShutdownResponse{}
  err := client.Call("Server.Shutdown", &request, &response)
  if err != nil {
    log.Fatal("error: ", err)
  }
}

func CheckMessages(client *rpc.Client, handle string) {
  request := RpcScheme.CheckMessagesRequest{handle}
  response := RpcScheme.CheckMessagesResponse{}
  err := client.Call("Server.CheckMessages", &request, &response)
  if err != nil {
    log.Fatal("error: ", err)
  }
  for i := 0; i < len(response.Messages); i++ {
    fmt.Println(response.Messages[i])
  }
}

func Help() {
  fmt.Println("The 5 commands available are the following:")
  fmt.Println("\tquit  --This logs you out")
  fmt.Println("\ttell <user> <message> --This sends a message to a specific user")
  fmt.Println("\tsay <message>  --This sends a message to all users")
  fmt.Println("\tlist  --This shows all the users online")
  fmt.Println("\tshutdown --This shuts the server down")
}
