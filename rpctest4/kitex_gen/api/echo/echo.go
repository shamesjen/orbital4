// Code generated by Kitex v0.5.2. DO NOT EDIT.

package echo

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	api "hello/kitex_gen/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return echoServiceInfo
}

var echoServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Echo"
	handlerType := (*api.Echo)(nil)
	methods := map[string]kitex.MethodInfo{
		"call": kitex.NewMethodInfo(callHandler, newEchoCallArgs, newEchoCallResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func callHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.EchoCallArgs)
	realResult := result.(*api.EchoCallResult)
	success, err := handler.(api.Echo).Call(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEchoCallArgs() interface{} {
	return api.NewEchoCallArgs()
}

func newEchoCallResult() interface{} {
	return api.NewEchoCallResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Call(ctx context.Context, req *api.Request) (r *api.Response, err error) {
	var _args api.EchoCallArgs
	_args.Req = req
	var _result api.EchoCallResult
	if err = p.c.Call(ctx, "call", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
