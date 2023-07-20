package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server/genericserver"
)

func main() {
    // Parse IDL with Local Files
    // YOUR_IDL_PATH thrift file path,eg: ./idl/example.thrift
    p, err := generic.NewThriftFileProvider("./idl/hello.thrift")
    if err != nil {
        panic(err)
    }
    g, err := generic.JSONThriftGeneric(p)
    if err != nil {
        panic(err)
    }
    svr := genericserver.NewServer(new(GenericServiceImpl), g)
    if err != nil {
        panic(err)
    }
    err = svr.Run()
    if err != nil {
        panic(err)
    }
    // resp is a JSON string
}

type GenericServiceImpl struct {
}

func (g *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
        // use jsoniter or other json parse sdk to assert request
        m := request.(string)
        fmt.Printf("Recv: %v\n", m)
        return  "{\"Msg\": \"world\"}", nil
}


