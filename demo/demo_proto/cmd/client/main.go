package main

import (
	"context"
	"fmt"
	"log"

	"github.com/CaspianGao/mygomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/CaspianGao/mygomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
    // 创建服务解析器
    r, err := consul.NewConsulResolver("127.0.0.1:8500")    
    if err != nil {
        log.Fatal(err)
    }  
    
    c, err := echoservice.NewClient("demo_proto", client.WithResolver(r))
    if err != nil {
        log.Fatal(err)
    }
    res, err := c.Echo(context.TODO(), &pbapi.Request{Message: "Hello World!"})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%v\n", res)
}
