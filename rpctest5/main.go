package main

import (
	"log"
	api "rpc/kitex_gen/api/echo"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server/genericserver"
)

func main() {
	// Parse IDL with Local Files
	// YOUR_IDL_PATH thrift file path,eg: ./idl/example.thrift
	p, err := generic.NewThriftFileProvider("idl/hello.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	svr := genericserver.NewServer(new(EchoImpl), g)
	if err != nil {
		panic(err)
	}

	err = svr.Run()
	if err != nil {
		panic(err)
	}
	// resp is a JSON string
}


