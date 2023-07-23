package main

import (
	"context"
	"encoding/json"
	"fmt"
	api "rpc/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Call implements the EchoImpl interface.
func (s *EchoImpl) genericCall(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
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
        "Msg": req.Message,
        "AdditionalData": "world",
    }
    jsonResponse, err := json.Marshal(respMap)
    if err != nil {
        return nil, err
    }
    fmt.Println(string(jsonResponse))
    fmt.Println(respMap)

    return string(jsonResponse), nil
}
