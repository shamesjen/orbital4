// Code generated by hertz generator.

package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	api "hello/biz/model/api"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

// Call .
// @router /post [POST]
func Call(ctx context.Context, c *app.RequestContext) {
	//url to sendreq?
	var requestURL string = "http://example.com/life/client/11?vint64=1&items=item0,item1,item2"
	var IDLPATH string = "/Users/james/Documents/GitHub/apigateway4/idl/hello.thrift"
	var jsonData map[string]interface{}

	//return data in bytes
	response := c.GetRawData()

	err := json.Unmarshal(response, &jsonData)

	if err != nil {
		fmt.Println("Error", err)
		c.String(consts.StatusBadRequest, "bad post request")
		return
	}

	//wtv key value ned be consistet
	dataValue, ok := jsonData["text"].(string)

	fmt.Println("message is " + dataValue)

	if !ok {
		c.String(consts.StatusBadRequest, "data provided not a string")
		return
	}

	//working until here

	responseFromRPC, err := makeThriftCall(IDLPATH, response, requestURL, ctx, dataValue)

	if err != nil {
		fmt.Println(err)
		c.String(consts.StatusBadRequest, "error in thrift call ")
		return
	}

	fmt.Println("Post request successful")

	c.JSON(consts.StatusOK, responseFromRPC)
}

func makeThriftCall(IDLPath string, response []byte, requestURL string, ctx context.Context, dataValue string) (interface{}, error) {
	p, err := generic.NewThriftFileProvider(IDLPath)
	if err != nil {
		fmt.Println("error creating thrift file provider")
		return 0, err
	}
	fmt.Println(dataValue)
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		return 0, errors.New(("error creating thrift generic"))
	}

	cli, err := genericclient.NewClient("Call", g, client.WithHostPorts("0.0.0.0:8888"))

	if err != nil {
		return 0, errors.New(("invalid client name"))
	}

	// req, err := http.NewRequest(http.MethodGet, requestURL, bytes.NewBuffer(response))
	// req.Header.Set("token", "1")
	// if err != nil {
	// 	fmt.Println("error construting req")
	// 	return 0, err
	// }
	req := api.NewRequest()
	req.Message = dataValue
	// customReq, err := generic.FromHTTPRequest(req)

	// if err != nil {
	// 	fmt.Println("error constructing xcustom req")
	// 	return 0, err
	// }

	// fmt.Println(customReq)

	resp, err := cli.GenericCall(ctx, "call", dataValue)

	fmt.Println("generic call successful")
	fmt.Println(resp)

	if err != nil {
		fmt.Println("error making generic call")
		return 0, err
	}

	realResp := resp

	return realResp, nil
}
