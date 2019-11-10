package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"ReadRpc/service"
	"ReadRpc/srv/protoc"
	"go.etcd.io/etcd/clientv3"
	"context"
	"time"
	"fmt"
)

func main() {
	lis, err := net.Listen("tcp", ":9002")  //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	rpcServer := grpc.NewServer() //创建gRPC服务
	//注册书籍服务
	srv.RegisterBookServiceServer(rpcServer, &service.BookService{})
	srv.RegisterParserServiceServer(rpcServer, &service.ParserService{})
	//注册反射服务
	reflection.Register(rpcServer)
	//服务注册
	RegisterService()
	// 将监听交给gRPC服务处理
	err = rpcServer.Serve(lis)
	if  err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func RegisterService()  {
	client,err :=clientv3.New(clientv3.Config{
		Context:context.Background(),
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:time.Second * 5,
	})
	if err != nil {
		panic(err)
	}
	res, err := client.Put(client.Ctx(), "go-service/127.0.0.1:9001", "127.0.0.1:9002")
	if err != nil {
		panic(err)
	}
	fmt.Println("register Success:",res)
}