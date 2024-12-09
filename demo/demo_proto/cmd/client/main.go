package main

import (
	"log"

	"github.com/CaspianGao/mygomall/demo/demo_proto/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
    r, err := consul.NewConsulResolver("127.0.0.1:8500")
    if err != nil {
        log.Fatal(err)
    }
    client, err := echo.NewClient("greet.server", client.WithResolver(r))
    if err != nil {
        log.Fatal(err)
    }
    
}
