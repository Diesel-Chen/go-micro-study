package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	hello "go-micro-study/example/helloworld/proto"
)

func main() {
	//注册客户端
	service:=micro.NewService(
		micro.Name("go.micro.service.client"),
		)

	//初始化
	service.Init()

	client:=hello.NewTestService("go.micro.service.hello",client.DefaultClient)
	resp,err:=client.SayHello(context.TODO(),&hello.HelloReq{Name: "朱晨"})
	if err!=nil{
		panic(err)
	}
	fmt.Println("res:",resp)

}
