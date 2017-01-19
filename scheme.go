package RpcScheme

type TellRequest struct {
  Handle, User, Message string
}
type TellResponse struct {
  Result string
}

type LoginRequest struct {
  Handle string
}
type LoginResponse struct { }

type ListRequest struct { }
type ListResponse struct {
  Users []string
}

