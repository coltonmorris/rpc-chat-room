package main

import (
	"../" // TODO find out the proper way to do this
	"errors"
	"fmt"
)

type Handler int

// A user tells another user a message
func (t *Handler) Tell(request *RpcScheme.TellRequest, response *RpcScheme.TellResponse) error {
  fmt.Println(request.Message)
  fmt.Println(request.User)
  if false {
    return errors.New("error happened")
  }
  return nil
}

