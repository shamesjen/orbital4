package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
	etcd "github.com/kitex-contrib/registry-etcd"
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
	fmt.Println("test1")
	r, err := etcd.NewEtcdRegistry([]string{"etcd:2379"})
	if err != nil {
		log.Fatalf("Failed to create etcd registry: %v", err)
	}
	fmt.Println("test")
	// svr := genericserver.NewServer(new(GenericServiceImpl), g, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Call"}), server.WithRegistry(r))
	// if err != nil {
	// 	panic(err)
	// }
    addr, err := net.ResolveTCPAddr("tcp", "server:8888")
    if err != nil {
        log.Fatalf("Failed to resolve server address: %v", err)
    }

    svr := genericserver.NewServer(
        new(GenericServiceImpl), 
        g, 
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Call"}),
        server.WithServiceAddr(addr), 
        server.WithRegistry(r),
    )

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
	m := request.(string)
	var jsonRequest map[string]interface{}

	err = json.Unmarshal([]byte(m), &jsonRequest)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(m)
	fmt.Println(jsonRequest)

	dataValue, ok := jsonRequest["message"].(string)
	if !ok {
		fmt.Println("data provided is not a string")
		return
	}
	fmt.Println(dataValue)

	jsonRequest["message"] = "Hello!, " + dataValue

	// var respMap map[string]interface{}

	jsonResponse, err := json.Marshal(jsonRequest)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsonResponse))
	// fmt.Println(respMap)

	return string(jsonResponse), nil
}
