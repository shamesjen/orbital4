package main

import (
	"context"
	api "rpc/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Call implements the EchoImpl interface.
func (s *EchoImpl) Call(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}
