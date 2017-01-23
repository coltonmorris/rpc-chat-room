package RpcScheme

type TellRequest struct {
  Sender, Target, Message string
}
type TellResponse struct { }

type LoginRequest struct {
  Handle string
}
type LoginResponse struct { }

type ListRequest struct { }
type ListResponse struct {
  Users []string
}

type CheckMessagesRequest struct {
  Handle string
}
type CheckMessagesResponse struct {
  Messages []string
}

type SayRequest struct {
  Sender, Message string
}
type SayResponse struct { }

type LogoutRequest struct {
  Handle string
}
type LogoutResponse struct { }

type ShutdownRequest struct { 
  Handle string
}
type ShutdownResponse struct { }
