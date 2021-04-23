package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	hello "go-micro-study/example/helloworld/proto"
)
type Hello struct {

}

func (h *Hello) SayHello(ctx context.Context, req *hello.HelloReq, resp *hello.HelloResp) error {
	resp.Msg = "Hello World "+req.Name
	return nil
}

func main() {
	//new service
	service:=micro.NewService(
		micro.Name("go.micro.service.hello"),
		)

	//初始化service
	service.Init()

	//注册
    hello.RegisterTestHandler(service.Server(),new(Hello))

	//启动服务
	if err:=service.Run();err!=nil{
		panic(err)
	}



}
