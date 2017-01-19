package main

import (
	"../" // TODO find out the proper way to do this
	"errors"
	"fmt"
)

func (server Server) Tell(request *RpcScheme.TellRequest, response *RpcScheme.TellResponse) error {
  state := <- server

  // dont message yourself silly
  if request.Handle == request.User {
    response.Result = "You can not send a message to yourself"

  // check if user exists
  } else if user, ok := state.Users[request.User]; ok {
    user = append(user, request.Message)
    response.Result = "Successfully sent message"
    fmt.Printf("%v just sent a message to %v", request.Handle, request.User)

  // user doesnt exist
  } else {
    response.Result = "User does not exist."
  }

  server <- state
  return nil
}

func (server Server) Login(request *RpcScheme.LoginRequest, response *RpcScheme.LoginResponse) error {
  state := <- server
  handle := request.Handle

  // check if user already exists
  if _, ok := state.Users[handle]; ok {
    server <- state
    return errors.New("User already exists")
  }
  
  state.Users[handle] = append(state.Users[handle], "You have successfully logged in.")

  server <- state
  return nil
}

func (server Server) List(request *RpcScheme.ListRequest, response *RpcScheme.ListResponse) error {
  state := <- server
  
  var users []string
  for key := range state.Users {
    users = append(users, key)
  }

  response.Users = users

  server <- state
  return nil
}

