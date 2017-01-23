package main

import (
	"../" // TODO find out the proper way to do this
	"errors"
	"fmt"
)

func (server Server) Tell(request *RpcScheme.TellRequest, response *RpcScheme.TellResponse) error {
  state := <- server

  // dont message yourself silly
  if request.Sender == request.Target {
    // TODO refactor so that you call Tell on yourself. currently cant reuse these functions
    state.Users[request.Sender] = append(state.Users[request.Sender], "You can not send a message to yourself")

  // check if the Target exists
  } else if _, ok := state.Users[request.Target]; ok {
    message := request.Sender + " tells you: " + request.Message
    state.Users[request.Target] = append(state.Users[request.Target], message)

    fmt.Printf("%v just sent a message to %v.\n", request.Sender, request.Target)

  // user doesnt exist
  } else {
    state.Users[request.Sender] = append(state.Users[request.Sender], "User does not exist, you cannot tell something to nobody")
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

  // TODO
  // Now tell everyone this person logged in

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

func (server Server) CheckMessages(request *RpcScheme.CheckMessagesRequest, response *RpcScheme.CheckMessagesResponse) error {
  state := <- server
  
  // return the messages
  response.Messages = state.Users[request.Handle]
  
  // empty the messages
  state.Users[request.Handle] = nil

  server <- state
  return nil
}

func (server Server) Say(request *RpcScheme.SayRequest, response *RpcScheme.SayResponse) error {
  state := <- server

  fmt.Println(request.Sender, " just said to everybody: ", request.Message)
  for key, _ := range state.Users {
    // dont send the message to the sender
    if key != request.Sender {
      state.Users[key] = append(state.Users[key], request.Message)
    }
  }
  
  server <- state
  return nil
}

func (server Server) Logout(request *RpcScheme.LogoutRequest, response *RpcScheme.LogoutResponse) error {
  state := <- server
  // check if the handle exists, and delete it. Otherwise return an error
  if _, ok := state.Users[request.Handle]; ok {
    delete(state.Users, request.Handle)
    fmt.Println("************************* ", state.Users)
    server <- state
    return nil
  } else {
    server <- state
    err := "Could not logout user: " + request.Handle
    return errors.New(err)
  }
}

func (server Server) Shutdown(request *RpcScheme.ShutdownRequest, response *RpcScheme.ShutdownResponse) error {
  state := <- server
  
  // begin the shutdown process by setting the flag
  state.Shutdown = true

  // TODO this is not gauranteed to be logged out to the clients.
  // say to everybody taht the server is shutdown
  message := request.Handle + " just shutdown the server."
  for key, _ := range state.Users {
    // dont send the message to the sender
    if key != request.Handle {
      state.Users[key] = append(state.Users[key], message)
    }
  }


  server <- state
  return nil
}
