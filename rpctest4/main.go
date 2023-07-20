package main

import (
	"context"
	"encoding/json"
	"fmt"
	"hello/kitex_gen/api"

	//echo "hello/kitex_gen/api/echo"

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
    m := request.(string)
    var jsonRequest map[string]interface{}
    json.Unmarshal([]byte(m), &jsonRequest)
    fmt.Println(m)
    req := api.NewRequest()
    message, ok := jsonRequest["message"].(string) // use "message" here
    if ok {
        req.Message = message
    }
    fmt.Println(req.Message)

    
    respMap := map[string]interface{}{
        "Msg":            req.Message,
        "AdditionalData": "Hongwei",
    }
    jsonResponse, err := json.Marshal(respMap)
    if err != nil {
        return nil, err
    }
    fmt.Println(string(jsonResponse))
    fmt.Println(respMap)

    // resp := map[string]interface{}{
    //     "Msg":            req.Message,
    //     "AdditionalData": "Hongwei",
    // }
    // return resp, nil

    return string(jsonResponse), nil
}
