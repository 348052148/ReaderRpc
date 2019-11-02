package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"test/ReadRpc/service"
	"test/ReadRpc/srv/protoc"
)

func main() {
	lis, err := net.Listen("tcp", ":8088")  //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	rpcServer := grpc.NewServer() //创建gRPC服务
	//注册书籍服务
	srv.RegisterBookServiceServer(rpcServer, &service.BookService{})
	srv.RegisterParserServiceServer(rpcServer, &service.ParserService{})
	//注册反射服务
	reflection.Register(rpcServer)
	// 将监听交给gRPC服务处理
	err = rpcServer.Serve(lis)
	if  err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
