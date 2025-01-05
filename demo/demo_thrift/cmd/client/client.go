package main

import (
	"context"
	"fmt"

	"github.com/CaspianGao/mygomall/demo/demo_thrift/kitex_gen/api"
	"github.com/CaspianGao/mygomall/demo/demo_thrift/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	// "github.com/cloudwego/kitex/transport"
)

func main() {
	// 初始化服务用户 没有调用服务发现所以直接写死服务地址
	cli, err := echo.NewClient("demo_thrift",
		client.WithHostPorts("localhost:8888"),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),	// 设置传输协议
		// client.WithTransportProtocol(transport.TTHeader),		
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{	
			ServiceName: "demo_thrift_client",
		}),
	)
	if err != nil {
		panic(err)
	}

	res, err := cli.Echo(context.Background(), &api.Request{
		Message: "hello",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
